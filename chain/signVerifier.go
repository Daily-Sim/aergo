package chain

import (
	"errors"
	"github.com/aergoio/aergo/account/key"
	"github.com/aergoio/aergo/internal/enc"
	"github.com/aergoio/aergo/message"
	"github.com/aergoio/aergo/pkg/component"
	"github.com/aergoio/aergo/types"
	"time"
)

type SignVerifier struct {
	comm component.IComponentRequester

	workerCnt int
	workCh    chan verifyWork
	doneCh    chan verifyWorkRes
	resultCh  chan *VerifyResult

	useMempool bool
	totalHit   int
}

type verifyWork struct {
	idx int
	tx  *types.Tx
}

type verifyWorkRes struct {
	work *verifyWork
	err  error
	hit  bool
}

type VerifyResult struct {
	failed bool
	errs   []error
}

var (
	ErrTxFormatInvalid = errors.New("tx invalid format")
	dfltUseMempool     = true
	//logger = log.NewLogger("signverifier")
)

func NewSignVerifier(comm component.IComponentRequester, workerCnt int, useMempool bool) *SignVerifier {
	sv := &SignVerifier{
		comm:       comm,
		workerCnt:  workerCnt,
		workCh:     make(chan verifyWork, workerCnt),
		doneCh:     make(chan verifyWorkRes, workerCnt),
		resultCh:   make(chan *VerifyResult, 1),
		useMempool: useMempool,
	}

	for i := 0; i < workerCnt; i++ {
		go sv.verifyTxLoop(i)
	}

	return sv
}

func (sv *SignVerifier) Stop() {
	close(sv.workCh)
	close(sv.doneCh)
}

func (sv *SignVerifier) verifyTxLoop(workerNo int) {
	logger.Debug().Int("worker", workerNo).Msg("verify worker run")

	for txWork := range sv.workCh {
		//logger.Debug().Int("worker", workerNo).Int("idx", txWork.idx).Msg("get work to verify tx")
		hit, err := sv.verifyTx(sv.comm, txWork.tx)

		if err != nil {
			logger.Error().Int("worker", workerNo).Bool("hit", hit).Str("hash", enc.ToString(txWork.tx.GetHash())).
				Err(err).Msg("error verify tx")
		}

		sv.doneCh <- verifyWorkRes{work: &txWork, err: err, hit: hit}
	}

	logger.Debug().Int("worker", workerNo).Msg("verify worker stop")
}

func (sv *SignVerifier) isExistInMempool(comm component.IComponentRequester, tx *types.Tx) (bool, error) {
	if !sv.useMempool {
		return false, nil
	}

	result, err := comm.RequestToFutureResult(message.MemPoolSvc, &message.MemPoolExist{Hash: tx.GetHash()}, time.Second,
		"chain/signverifier/verifytx")
	if err != nil {
		logger.Error().Err(err).Msg("failed to get verify from mempool")
		return false, err
	}

	msg := result.(*message.MemPoolExistRsp)
	if msg.Tx != nil {
		return true, nil
	}

	return false, nil
}

func (sv *SignVerifier) verifyTx(comm component.IComponentRequester, tx *types.Tx) (hit bool, err error) {
	account := tx.GetBody().GetAccount()
	if account == nil {
		return false, ErrTxFormatInvalid
	}

	if hit, err = sv.isExistInMempool(comm, tx); err != nil {
		return false, err
	}
	if hit {
		return hit, nil
	}

	err = key.VerifyTx(tx)

	if err != nil {
		return false, err
	}

	return false, nil
}

func (sv *SignVerifier) RequestVerifyTxs(txlist *types.TxList) {
	txs := txlist.GetTxs()
	txLen := len(txs)

	if txLen == 0 {
		sv.resultCh <- &VerifyResult{failed: false, errs: nil}
		return
	}

	errs := make([]error, txLen, txLen)

	//logger.Debug().Int("txlen", txLen).Msg("verify tx start")

	go func() {
		for i, tx := range txs {
			//logger.Debug().Int("idx", i).Msg("push tx start")
			sv.workCh <- verifyWork{idx: i, tx: tx}
		}
	}()

	go func() {
		var doneCnt = 0
		failed := false
		sv.totalHit = 0

	LOOP:
		for {
			select {
			case result := <-sv.doneCh:
				doneCnt++
				//logger.Debug().Int("donecnt", doneCnt).Msg("verify tx done")

				if result.work.idx < 0 || result.work.idx >= txLen {
					logger.Error().Int("idx", result.work.idx).Msg("Invalid Verify Result Index")
					continue
				}

				errs[result.work.idx] = result.err

				if result.err != nil {
					logger.Error().Err(result.err).Int("txno", result.work.idx).
						Msg("verifing tx failed")
					failed = true
				}

				if result.hit {
					sv.totalHit++
				}

				if doneCnt == txLen {
					break LOOP
				}
			}
		}

		sv.resultCh <- &VerifyResult{failed: failed, errs: errs}
		logger.Debug().Int("hit", sv.totalHit).Msg("verify tx done")
	}()
	return
}

func (sv *SignVerifier) WaitDone() (bool, []error) {
	select {
	case res := <-sv.resultCh:
		logger.Debug().Msg("wait verify tx")
		return res.failed, res.errs
	}
}

func (bv *SignVerifier) verifyTxsInplace(txlist *types.TxList) (bool, []error) {
	txs := txlist.GetTxs()
	txLen := len(txs)
	errs := make([]error, txLen, txLen)
	failed := false
	var hit bool

	logger.Debug().Int("txlen", txLen).Msg("verify tx inplace start")

	for i, tx := range txs {
		hit, errs[i] = bv.verifyTx(bv.comm, tx)
		failed = true

		if hit {
			bv.totalHit++
		}
	}

	logger.Debug().Int("totalhit", bv.totalHit).Msg("verify tx inplace done")
	return failed, errs
}
