package txpool

import (
	"errors"
	"sync"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
)

var ErrAlreadyKnown = errors.New("tx already in pool")

// TxStatus is the current status of a transaction as seen by the pool.
type TxStatus uint

const (
	TxStatusUnknown TxStatus = iota
	TxStatusQueued
	TxStatusPending
	TxStatusIncluded
)

// PEBBLE
// TxPool hash been simplified
// We add transaction  in ordered by time heap in case we are minter
// If we aren't minter return error
// If tx is already in pool return error
type TxPool struct {
	queuedTx  timestampTxList      // prioritized by tx timestamp queued transactions
	pendingTx []*types.Transaction // currently processing transactions
	mx        sync.Mutex           // held when tx pool lists get modified

	// for readers
	all     *lookup
	pending map[common.Hash]*types.Transaction
	queued  map[common.Hash]*types.Transaction
}

type lookup struct {
	slots  int
	lock   sync.RWMutex
	locals map[common.Hash]*types.Transaction
}

func (l *lookup) get(h common.Hash) *types.Transaction {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.locals[h]
}

func (tp *TxPool) AddTxs(txs []*types.Transaction) {
	//
}

func (tp *TxPool) add(tx *types.Transaction) {
	tp.all.get(tx.Hash())
}

func New() *TxPool {
	return &TxPool{}
}

// Pendings returns transactions for processing
// clears currentPending list
func (tp *TxPool) Pending(blockSize int) []*types.Transaction {
	tp.mx.Lock()
	defer tp.mx.Unlock()
	var counter int
	tp.pendingTx = make([]*types.Transaction, 0, blockSize)
	for tp.queuedTx.Len() > 0 {
		tx := tp.queuedTx.Pop().(*timestampTxItem).tx
		tp.pendingTx = append(tp.pendingTx, tx)
		if counter == blockSize {
			break
		}
		delete(tp.queued, tx.Hash())
		counter++
	}

	return tp.pendingTx
}

func (tp *TxPool) Has(hash common.Hash) TxStatus {
	_, ok := tp.queued[hash]
	if ok {
		return TxStatusQueued
	}
	_, ok = tp.pending[hash]
	if ok {
		return TxStatusPending
	}
	return TxStatusUnknown
}
