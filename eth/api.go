package eth

import (
	"context"
	"time"

	"github.com/Unheilbar/pebbke_wallets/core/types"
)

type EthAPIBackend struct {
	eth *Ethereum
}

func NewApi(eth *Ethereum) *EthAPIBackend {
	return &EthAPIBackend{eth}
}

func (b *EthAPIBackend) SendTxs(ctx context.Context, signedTx *types.Transaction) {
	// Quourum
	// validation for node need to happen here and cannot be done as a part of
	// validateTx in tx_pool.go as tx_pool validation will happen in every node
	// if !pcore.ValidateNodeForTxn(b.nodeId, signedTx.From()) {
	// 	return errors.New("cannot send transaction from this node")
	// }
	// End Quorum
	b.eth.txPool.AddTx(signedTx)
}

func (b *EthAPIBackend) SendTx(ctx context.Context, signedTx *types.Transaction) {
	// Quourum
	// validation for node need to happen here and cannot be done as a part of
	// validateTx in tx_pool.go as tx_pool validation will happen in every node
	// if !pcore.ValidateNodeForTxn(b.nodeId, signedTx.From()) {
	// 	return errors.New("cannot send transaction from this node")
	// }
	// End Quorum
	signedTx.WithTime(time.Now()) //
	b.eth.txPool.AddTx(signedTx)
}

// // For test purposes only
// func (b *EthAPIBackend) GetBalance(ctx context.Context, wallet string) {
// 	// Quourum
// 	// validation for node need to happen here and cannot be done as a part of
// 	// validateTx in tx_pool.go as tx_pool validation will happen in every node
// 	// if !pcore.ValidateNodeForTxn(b.nodeId, signedTx.From()) {
// 	// 	return errors.New("cannot send transaction from this node")
// 	// }
// 	// End Quorum
// 	b.eth.txPool.AddTx(signedTx)
// }
