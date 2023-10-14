package txpool

import (
	"sync"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
)

type TxPool struct {
	queuedTx  map[common.Hash]*types.Transaction
	pendingTx map[common.Hash]*types.Transaction
	mx        sync.Mutex
}

func New() *TxPool {
	return &TxPool{
		queuedTx:  make(map[common.Hash]*types.Transaction),
		pendingTx: make(map[common.Hash]*types.Transaction),
	}
}

func (tp *TxPool) AddTxs(txs []*types.Transaction) {
	tp.mx.Lock()
	for _, tx := range txs {
		tp.queuedTx[tx.Hash()] = tx
	}
	tp.mx.Unlock()
}

func (tp *TxPool) Pending(size uint) []*types.Transaction {
	tp.mx.Lock()
	var counter uint
	var txs []*types.Transaction
	for hash, tx := range tp.queuedTx {
		if counter == size {
			break
		}
		counter++
		tp.pendingTx[hash] = tx
		txs = append(txs, tx)
		delete(tp.queuedTx, hash)
	}
	tp.mx.Unlock()
	return txs
}

func (tp *TxPool) RemovePendingTxs(hash []common.Hash) {
	tp.mx.Lock()
	for _, h := range hash {
		delete(tp.pendingTx, h)
	}
	tp.mx.Unlock()
}

func (tp *TxPool) PendingSize() int {
	return len(tp.pendingTx)
}

func (tp *TxPool) QueuedSize() int {
	return len(tp.queuedTx)
}
