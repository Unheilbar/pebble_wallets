package pool

import (
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
)

type TxPool struct {
	queuedTx  map[common.Hash]types.Transaction
	pendingTx map[common.Hash]types.Transaction
}

func New() *TxPool {
	return &TxPool{
		queuedTx:  make(map[common.Hash]types.Transaction),
		pendingTx: make(map[common.Hash]types.Transaction),
	}
}
