package eth

import (
	"context"
	"time"

	"github.com/Unheilbar/pebbke_wallets/core"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/event"
)

type EthAPIBackend struct {
	eth *Ethereum
}

func NewApi(eth *Ethereum) *EthAPIBackend {
	return &EthAPIBackend{eth}
}

func (b *EthAPIBackend) SendTxs(ctx context.Context, signedTx *types.Transaction) {
	// PEBBLE BATCH TRANSACTIONS AREN'T IMPLEMENTED YET
	// Quourum
	// validation for node need to happen here and cannot be done as a part of
	// validateTx in tx_pool.go as tx_pool validation will happen in every node
	// if !pcore.ValidateNodeForTxn(b.nodeId, signedTx.From()) {
	// 	return errors.New("cannot send transaction from this node")
	// }
	// End Quorum
	b.eth.txPool.AddTx(signedTx)
}

func (b *EthAPIBackend) SendTx(ctx context.Context, signedTx *types.Transaction) (time.Time, error) {
	recieveTime := time.Now()
	signedTx.WithTime(recieveTime) //
	err := b.eth.txPool.AddTx(signedTx)
	return recieveTime, err
}

func (b *EthAPIBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.eth.blockchain.SubscribeChainHeadEvent(ch)
}

// For test purposes only
func (b *EthAPIBackend) GetBalance(ctx context.Context, wallet []byte) {
	// // Quourum
	// // validation for node need to happen here and cannot be done as a part of
	// // validateTx in tx_pool.go as tx_pool validation will happen in every node
	// // if !pcore.ValidateNodeForTxn(b.nodeId, signedTx.From()) {
	// // 	return errors.New("cannot send transaction from this node")
	// // }
	// // End Quorum
	// b.eth.txPool.AddTx(signedTx)
}
