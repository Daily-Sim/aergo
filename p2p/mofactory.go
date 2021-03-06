/*
 * @file
 * @copyright defined in aergo/LICENSE.txt
 */

package p2p

import (
	"github.com/aergoio/aergo/p2p/p2pcommon"
	"github.com/aergoio/aergo/p2p/p2putil"
	"github.com/aergoio/aergo/p2p/subproto"
	"time"

	"github.com/aergoio/aergo/types"
	"github.com/gofrs/uuid"
)

type baseMOFactory struct {
}

func (mf *baseMOFactory) NewMsgRequestOrder(expectResponse bool, protocolID p2pcommon.SubProtocol, message p2pcommon.MessageBody) p2pcommon.MsgOrder {
	rmo := &pbRequestOrder{}
	msgID := uuid.Must(uuid.NewV4())
	if mf.newV030MsgOrder(&rmo.pbMessageOrder, msgID, uuid.Nil, protocolID, message) {
		return rmo
	}
	return nil
}

func (mf *baseMOFactory) NewMsgBlockRequestOrder(respReceiver p2pcommon.ResponseReceiver, protocolID p2pcommon.SubProtocol, message p2pcommon.MessageBody) p2pcommon.MsgOrder {
	rmo := &pbRequestOrder{}
	msgID := uuid.Must(uuid.NewV4())
	if mf.newV030MsgOrder(&rmo.pbMessageOrder, msgID, uuid.Nil, protocolID, message) {
		rmo.respReceiver = respReceiver
		return rmo
	}
	return nil
}

func (mf *baseMOFactory) NewMsgResponseOrder(reqID p2pcommon.MsgID, protocolID p2pcommon.SubProtocol, message p2pcommon.MessageBody) p2pcommon.MsgOrder {
	rmo := &pbResponseOrder{}
	msgID := uuid.Must(uuid.NewV4())
	if mf.newV030MsgOrder(&rmo.pbMessageOrder, msgID, uuid.FromBytesOrNil(reqID[:]), protocolID, message) {
		return rmo
	}
	return nil
}

func (mf *baseMOFactory) NewMsgBlkBroadcastOrder(noticeMsg *types.NewBlockNotice) p2pcommon.MsgOrder {
	rmo := &pbBlkNoticeOrder{}
	msgID := uuid.Must(uuid.NewV4())
	if mf.newV030MsgOrder(&rmo.pbMessageOrder, msgID, uuid.Nil, subproto.NewBlockNotice, noticeMsg) {
		rmo.blkHash = noticeMsg.BlockHash
		rmo.blkNo = noticeMsg.BlockNo
		return rmo
	}
	return nil
}

func (mf *baseMOFactory) NewMsgTxBroadcastOrder(message *types.NewTransactionsNotice) p2pcommon.MsgOrder {
	rmo := &pbTxNoticeOrder{}
	reqID := uuid.Must(uuid.NewV4())
	if mf.newV030MsgOrder(&rmo.pbMessageOrder, reqID, uuid.Nil, subproto.NewTxNotice, message) {
		rmo.txHashes = message.TxHashes
		return rmo
	}
	return nil
}

func (mf *baseMOFactory) NewMsgBPBroadcastOrder(noticeMsg *types.BlockProducedNotice) p2pcommon.MsgOrder {
	rmo := &pbBpNoticeOrder{}
	msgID := uuid.Must(uuid.NewV4())
	if mf.newV030MsgOrder(&rmo.pbMessageOrder, msgID, uuid.Nil, subproto.BlockProducedNotice, noticeMsg) {
		rmo.block = noticeMsg.Block
		return rmo
	}
	return nil
}

func (mf *baseMOFactory) newHandshakeMessage(protocolID p2pcommon.SubProtocol, message p2pcommon.MessageBody) p2pcommon.Message {
	// TODO define handshake specific datatype
	rmo := &pbRequestOrder{}
	msgID := uuid.Must(uuid.NewV4())
	if mf.newV030MsgOrder(&rmo.pbMessageOrder, msgID, uuid.Nil, protocolID, message) {
		return rmo.message
	}
	return nil
}

// newPbMsgOrder is base form of making sendrequest struct
func (mf *baseMOFactory)newV030MsgOrder(mo *pbMessageOrder, msgID, orgID uuid.UUID, protocolID p2pcommon.SubProtocol, messageBody p2pcommon.MessageBody) bool {
	id :=p2pcommon.MsgID(msgID)
	originalid := p2pcommon.MsgID(orgID)
	bytes, err := p2putil.MarshalMessageBody(messageBody)
	if err != nil {
		return false
	}
	msg := p2pcommon.NewMessageValue(protocolID, id, originalid, time.Now().UnixNano(), bytes)
	mo.protocolID = protocolID
	mo.needSign = true
	mo.message = msg

	return true
}
