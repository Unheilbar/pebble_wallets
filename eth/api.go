package eth

import (
	"bytes"
	"context"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/Unheilbar/pebbke_wallets/binding"
	"github.com/Unheilbar/pebbke_wallets/core"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/Unheilbar/pebbke_wallets/core/vm"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
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
	signedTx.WithTime(recieveTime)
	err := b.eth.txPool.AddTx(signedTx)
	return recieveTime, err
}

func preCheck(tx *types.Transaction) (bool, error) {
	sigPublicKey, err := crypto.Ecrecover(tx.Hash().Bytes(), tx.Signature())
	if err != nil {
		return false, err
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(sigPublicKey[1:])[12:])
	return bytes.Equal(addr.Bytes(), tx.From().Bytes()), nil
}

func (b *EthAPIBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.eth.blockchain.SubscribeChainHeadEvent(ch)
}

// For test purposes only
func (b *EthAPIBackend) GetBalance(wallet []byte, contrAddr common.Address) *big.Int {
	state, err := b.eth.blockchain.StateAt(b.eth.blockchain.CurrentBlock().Root())
	if err != nil {
		log.Fatal("cant get balance")
	}

	venv := vm.NewEVM(core.NewEVMBlockContext(b.eth.blockchain.CurrentBlock().Header()), vm.TxContext{}, state, vm.DefaultCancunConfig())

	var id [32]byte
	copy(id[:], wallet)

	input := packTX("getBalance", id)

	ret, _, err := venv.StaticCall(vm.AccountRef(common.Address{}), contrAddr, input, math.MaxBig256.Uint64())
	if err != nil {
		log.Fatal(err)
	}

	pabi, _ := abi.JSON(strings.NewReader(binding.ProxyABI))

	res, err := pabi.Unpack("getBalance", ret)
	if err != nil {
		log.Fatal(err)
	}

	out0 := *abi.ConvertType(res[0], new(*big.Int)).(**big.Int)

	return out0
}

func packTX(method string, params ...interface{}) []byte {
	bind := binding.ProxyABI

	pabi, err := abi.JSON(strings.NewReader(bind))

	if err != nil {
		log.Fatal("failed to read abi", err)
	}

	input, err := pabi.Pack(method, params...)
	if err != nil {
		log.Fatal("failed to pack tx", err)
	}

	return input
}
