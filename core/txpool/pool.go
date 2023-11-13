package txpool

import (
	"errors"
	"sync"

	"github.com/Unheilbar/pebble_wallets/core/types"
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
	queuedOrderedTx *timestampTxList     // prioritized by tx timestamp queued transactions //clients request go here
	pendingTx       []*types.Transaction // currently processing transactions //minter requests only
	mx              sync.Mutex           // held when tx pool lists get modified

	// for readers
	all     map[common.Hash]*types.Transaction
	pending map[common.Hash]*types.Transaction
	queued  map[common.Hash]*types.Transaction
}

func (tp *TxPool) AddTx(tx *types.Transaction) error {
	// prevalidate before mutex and return error in case of validation error?
	tp.mx.Lock()
	defer tp.mx.Unlock()
	if tp.Get(tx.Hash()) != nil {
		return ErrAlreadyKnown
	}

	tp.queued[tx.Hash()] = tx
	tp.all[tx.Hash()] = tx
	tp.queuedOrderedTx.Push(txToTimestampedItem(tx))

	return nil
}

func (tp *TxPool) Get(h common.Hash) *types.Transaction {
	return tp.all[h]
}

func New() *TxPool {
	tp := &TxPool{
		all:     make(map[common.Hash]*types.Transaction),
		pending: make(map[common.Hash]*types.Transaction),
		queued:  make(map[common.Hash]*types.Transaction),
	}

	tp.queuedOrderedTx = new(timestampTxList)
	return tp
}

// Pendings returns transactions for processing
// clears currentPending list
// shpuld be executed only by minter
func (tp *TxPool) Pending(blockSize int) []*types.Transaction {
	tp.mx.Lock()
	defer tp.mx.Unlock()
	var counter int
	tp.pendingTx = make([]*types.Transaction, 0, blockSize)
	for tp.queuedOrderedTx.Len() > 0 {
		tx := tp.queuedOrderedTx.Pop().(*timestampTxItem).tx
		tp.pendingTx = append(tp.pendingTx, tx)
		if counter == blockSize {
			break
		}
		delete(tp.queued, tx.Hash())
		counter++
	}

	return tp.pendingTx
}

func txToTimestampedItem(t *types.Transaction) *timestampTxItem {
	return &timestampTxItem{tx: t, priority: t.Time().Unix()}
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
