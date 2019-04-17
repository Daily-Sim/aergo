package raftv2

import (
	"errors"
	"github.com/aergoio/aergo/consensus"
	"github.com/aergoio/aergo/types"
	"github.com/aergoio/etcd/raft"
	"github.com/aergoio/etcd/raft/raftpb"
)

var (
	ErrInvalidEntry       = errors.New("Invalid raftpb.entry")
	ErrWalEntryTooLowTerm = errors.New("term of wal entry is too low")
)

type WalDB struct {
	consensus.ChainWAL
}

func NewWalDB(chainWal consensus.ChainWAL) *WalDB {
	return &WalDB{chainWal}
}

func (wal *WalDB) SaveEntry(state raftpb.HardState, entries []raftpb.Entry) error {
	if !raft.IsEmptyHardState(state) {
		// save hardstate
		if err := wal.WriteHardState(&state); err != nil {
			return err
		}
	}

	if len(entries) == 0 {
		return nil
	}

	walEnts, blocks := wal.convertFromRaft(entries)

	if err := wal.WriteRaftEntry(walEnts, blocks); err != nil {
		return err
	}

	return nil
}

func (wal *WalDB) convertFromRaft(entries []raftpb.Entry) ([]*consensus.WalEntry, []*types.Block) {
	lenEnts := len(entries)
	if lenEnts == 0 {
		return nil, nil
	}

	getWalEntryType := func(entry *raftpb.Entry) consensus.EntryType {
		switch entry.Type {
		case raftpb.EntryNormal:
			if entry.Data != nil {
				return consensus.EntryBlock
			} else {
				return consensus.EntryEmpty
			}
		case raftpb.EntryConfChange:
			return consensus.EntryConfChange
		default:
			panic("not support raftpb entrytype")
		}
	}

	getWalData := func(entry *raftpb.Entry) (*types.Block, []byte, error) {
		if entry.Type == raftpb.EntryNormal && entry.Data != nil {
			block, err := unmarshalEntryData(entry.Data)
			if err != nil {
				return nil, nil, ErrInvalidEntry
			}

			return block, block.BlockHash(), nil
		} else {
			return nil, entry.Data, nil
		}
	}

	blocks := make([]*types.Block, lenEnts)
	walents := make([]*consensus.WalEntry, lenEnts)

	var (
		data  []byte
		block *types.Block
		err   error
	)
	for i, entry := range entries {
		if block, data, err = getWalData(&entry); err != nil {
			panic("entry unmarshalEntryData error")
		}

		blocks[i] = block

		walents[i] = &consensus.WalEntry{
			Type:  getWalEntryType(&entry),
			Term:  entry.Term,
			Index: entry.Index,
			Data:  data,
		}
	}

	return walents, blocks
}

var ErrInvalidWalEntry = errors.New("invalid wal entry")
var ErrWalConvBlock = errors.New("failed to convert bytes of block from wal entry")

func (wal *WalDB) convertWalToRaft(walEntry *consensus.WalEntry) (*raftpb.Entry, error) {
	var raftEntry = &raftpb.Entry{Term: walEntry.Term, Index: walEntry.Index}

	getDataFromWalEntry := func(walEntry *consensus.WalEntry) ([]byte, error) {
		if walEntry.Type != consensus.EntryBlock {
			return nil, ErrWalConvBlock
		}
		block, err := wal.GetBlock(walEntry.Data)
		if err != nil {
			return nil, err
		}
		data, err := marshalEntryData(block)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	switch walEntry.Type {
	case consensus.EntryConfChange:
		raftEntry.Type = raftpb.EntryConfChange
		raftEntry.Data = walEntry.Data

	case consensus.EntryEmpty:
		raftEntry.Type = raftpb.EntryNormal
		raftEntry.Data = nil

	case consensus.EntryBlock:
		data, err := getDataFromWalEntry(walEntry)
		if err != nil {
			return nil, err
		}
		raftEntry.Data = data
	default:
		return nil, ErrInvalidWalEntry
	}

	return raftEntry, nil
}

var (
	ErrWalGetHardState = errors.New("failed to read hard state")
	ErrWalGetLastIdx   = errors.New("failed to read last Idx")
)

// ReadAll returns hard state, all uncommitted entries
// - read last hard state
// - read  all uncommited entries after snapshot index
func (wal *WalDB) ReadAll(snapshot *raftpb.Snapshot) (state *raftpb.HardState, ents []raftpb.Entry, err error) {
	state, err = wal.GetHardState()
	if err != nil {
		return state, ents, ErrWalGetHardState
	}

	commitIdx := state.Commit
	lastIdx, err := wal.GetRaftEntryLastIdx()
	if err != nil {
		return state, ents, ErrWalGetLastIdx
	}

	var snapIdx, snapTerm uint64
	if snapshot != nil {
		snapIdx = snapshot.Metadata.Index
		snapTerm = snapshot.Metadata.Term
	}

	logger.Info().Uint64("snapidx", snapIdx).Uint64("snapterm", snapTerm).Uint64("commit", commitIdx).Uint64("last", lastIdx).Msg("read all entries of wal")

	for i := snapIdx + 1; i <= lastIdx; i++ {
		walEntry, err := wal.GetRaftEntry(i)
		if err != nil {
			logger.Error().Err(err).Uint64("idx", i).Msg("failed to get raft entry")
			return state, nil, err
		}

		if walEntry.Term < snapTerm {
			logger.Error().Str("wal", walEntry.ToString()).Err(ErrWalEntryTooLowTerm).Msg("invalid wal entry")
			return state, nil, ErrWalEntryTooLowTerm
		}

		raftEntry, err := wal.convertWalToRaft(walEntry)
		if err != nil {
			return state, nil, err
		}

		logger.Debug().Str("walentry", walEntry.ToString()).Msg("read wal entry")
		ents = append(ents, *raftEntry)
	}

	return state, ents, nil
}
