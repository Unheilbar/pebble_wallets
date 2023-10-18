package core

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"math/big"

	// Add VM and runtime to core
	"github.com/Unheilbar/pebbke_wallets/core/state"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/Unheilbar/pebbke_wallets/core/vm"

	// "github.com/Unheilbar/pebbke_wallets/core/vm/runtime"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/crypto"
)

type StateProcessor struct {
}

func NewStateProcessor() *StateProcessor {
	return &StateProcessor{}
}

func (p *StateProcessor) Process(block *types.Block, statedb *state.StateDB, id int) []*types.Receipt {
	var (
		receipts    []*types.Receipt
		blockHash   = block.Hash()
		blockNumber = block.Number()
	)
	evm := vm.NewEVM(NewEVMBlockContext(block.Header()), vm.TxContext{}, statedb, vm.DefaultCancunConfig())

	for i, tx := range block.Transactions {
		fmt.Println("process: tx ", tx.Hash(), "block", block.Hash(), "node id", id, "block number", blockNumber)
		statedb.SetTxContext(tx.Hash(), i)
		receipt, err := p.applyTransaction(tx, statedb, blockNumber, blockHash, evm)
		if err != nil {
			log.Fatal("couldn't aply transaction") //TODO if precheck tx fail (Signature, tx uniqueness etc..)
		}
		statedb.Finalise(true)
		receipts = append(receipts, receipt)
	}

	fmt.Println("process final root", statedb.IntermediateRoot(true), "block", blockHash, "node id", id)
	return receipts
}

func (p *StateProcessor) applyTransaction(tx *types.Transaction, statedb *state.StateDB, blockNumber *big.Int, blockHash common.Hash, evm *vm.EVM) (*types.Receipt, error) {
	evm.Reset(newTxContext(tx.From()), statedb)
	var (
		sender           = vm.AccountRef(tx.From())
		contractCreation = tx.To() == nil
		contrAddr        common.Address
		vmerr            error
	)

	// TODO CHECK UNIQUNESS of tx HASH signature etc..
	preCheck(tx)

	if contractCreation {
		_, contrAddr, _, vmerr = evm.Create(sender, tx.Data(), math.MaxUint64, new(big.Int))
	} else {
		//TODO save unique tx hash?
		_, _, vmerr = evm.Call(sender, common.Address{}, tx.Data(), math.MaxUint64, new(big.Int))
	}

	receipt := &types.Receipt{}

	if vmerr != nil {
		receipt.Status = types.ReceiptStatusFailed
	} else {
		receipt.Status = types.ReceiptStatusSuccessful
	}
	receipt.TxHash = tx.Hash()

	if contractCreation {
		receipt.ContractAddress = contrAddr
	}

	receipt.BlockHash = blockHash
	receipt.BlockNumber = blockNumber
	receipt.TransactionIndex = uint(statedb.TxIndex())
	receipt.Logs = getLogs(tx.Hash(), blockNumber.Uint64(), blockHash, statedb)
	return receipt, nil
}

func ApplyTransactions(chain *Blockchain, statedb *state.StateDB, header *types.Header, txs []*types.Transaction) ([]*types.Transaction, []*types.Receipt, error) {
	if len(txs) == 0 {
		return nil, nil, nil
	}
	// defer func(t time.Time) {
	// 	fmt.Println("apply txes tx/s ", float64(len(txs))/time.Since(t).Seconds())
	// }(time.Now())
	// blockCtx := NewEVMBlockContext(header)
	// // PEBBLE we use cancun by default
	// evm := vm.NewEVM(blockCtx, vm.TxContext{}, statedb, vm.DefaultCancunConfig())
	var receipts []*types.Receipt
	var appliedTxs []*types.Transaction
	blockHash := header.Hash()
	blockNumber := header.Number
	txCount := 0
	for _, tx := range txs {
		fmt.Println("transition tx hash ", tx.Hash(), "blockHash", blockHash)
		blockCtx := NewEVMBlockContext(header)
		// PEBBLE we use cancun by default
		evm := vm.NewEVM(blockCtx, vm.TxContext{}, statedb, vm.DefaultCancunConfig())
		evm.Reset(newTxContext(tx.From()), statedb)
		snap := statedb.Snapshot()
		statedb.SetTxContext(tx.Hash(), txCount)
		result, err := ApplyMessage(evm, tx)
		if err != nil {
			statedb.RevertToSnapshot(snap)
			log.Println("exec reverted", err)
			continue
		}
		statedb.Finalise(true)
		txCount++
		appliedTxs = append(appliedTxs, tx)
		receipt := &types.Receipt{}
		if result.Failed() {
			receipt.Status = types.ReceiptStatusFailed
		} else {
			receipt.Status = types.ReceiptStatusSuccessful
		}

		receipt.TxHash = tx.Hash()

		created := tx.To() == nil
		if created {
			contractAddress := crypto.CreateAddress(evm.TxContext.Origin, 1) //TODO think of something //nonce has been increased
			receipt.ContractAddress = contractAddress
			log.Println("successed deploy", contractAddress, "sender ", tx.From().Hex(), "nonce", 1)
		}

		receipt.Logs = getLogs(tx.Hash(), header.Number.Uint64(), header.Hash(), statedb)
		receipt.BlockHash = blockHash
		receipt.BlockNumber = blockNumber
		receipt.TransactionIndex = uint(statedb.TxIndex())
		revertReason := result.Revert()
		if revertReason != nil {
			r, _ := abi.UnpackRevert(revertReason)
			fmt.Println("revert reason", r)
			receipt.RevertReason = revertReason
		}

		receipts = append(receipts, receipt)
	}
	fmt.Println("transition final root ", statedb.IntermediateRoot(true), "blockHash", blockHash)
	return appliedTxs, receipts, nil
}

func getLogs(txHash common.Hash, blockNumber uint64, blockHash common.Hash, statedb *state.StateDB) []*types.Log {
	logs := statedb.GetLogs(txHash, blockNumber, blockHash)
	ret := make([]*types.Log, 0)
	for _, log := range logs {
		ret = append(ret, &types.Log{
			Data:        log.Data,
			TxHash:      log.TxHash,
			BlockNumber: blockNumber,
			TxIndex:     log.TxIndex,
			BlockHash:   blockHash,
			Index:       log.Index,
		})
	}

	return ret
}

func newTxContext(from common.Address) vm.TxContext {
	return vm.TxContext{
		Origin:     from,
		GasPrice:   new(big.Int),
		BlobHashes: []common.Hash{},
	}
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

// ChainContext supports retrieving headers and consensus parameters from the
// current blockchain to be used during transaction processing.
type ChainContext interface {
	// Engine retrieves the chain's consensus engine.
	Engine() consensus.Engine

	// GetHeader returns the header corresponding to the hash/number argument pair.
	GetHeader(common.Hash, uint64) *types.Header
}

// NewEVMBlockContext creates a new context for use in the EVM.
func NewEVMBlockContext(header *types.Header) vm.BlockContext {
	var (
		beneficiary common.Address
		random      *common.Hash
	)

	return vm.BlockContext{
		CanTransfer: CanTransfer,
		Transfer:    Transfer, //
		GetHash: func(n uint64) common.Hash {
			return common.BytesToHash(crypto.Keccak256([]byte(new(big.Int).SetUint64(n).String())))
		},
		Coinbase:    beneficiary,
		BlockNumber: new(big.Int).Set(header.Number),
		Time:        header.Time,
		Random:      random,
	}
}
