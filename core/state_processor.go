package core

import (
	"bytes"
	"log"
	"math/big"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core/vm/runtime"
	"github.com/ethereum/go-ethereum/crypto"
)

type StateProcessor struct {
	cfg *runtime.Config
}

func NewStateProcessor(cfg *runtime.Config) *StateProcessor {
	return &StateProcessor{cfg}
}

func (p *StateProcessor) Process(block types.Block, statedb *state.StateDB) []*types.Receipt {
	var (
		receipts    []*types.Receipt
		blockHash   = block.Hash()
		blockNumber = block.Number
	)
	evm := runtime.NewEnv(p.cfg)
	evm.SetBlockContext(newBlockContext(p.cfg, blockNumber))

	for i, tx := range block.Transactions {
		statedb.SetTxContext(tx.Hash(), i)
		receipt, err := p.applyTransaction(tx, statedb, blockNumber, blockHash, evm)
		if err != nil {
			log.Fatal("couldn't aply transaction") //TODO if precheck tx fail (Signature, tx uniqueness etc..)
		}
		receipts = append(receipts, receipt)
	}

	return receipts
}

func (p *StateProcessor) applyTransaction(tx *types.Transaction, statedb *state.StateDB, blockNumber *big.Int, blockHash common.Hash, evm *vm.EVM) (*types.Receipt, error) {
	evm.Reset(newTxContext(p.cfg, tx.From), statedb)
	var (
		sender           = vm.AccountRef(tx.From)
		contractCreation = tx.To == common.Address{}
		contrAddr        common.Address
		vmerr            error
	)

	// TODO CHECK UNIQUNESS of tx HASH signature etc..
	preCheck(tx)

	if contractCreation {
		_, contrAddr, _, vmerr = evm.Create(sender, tx.Input, p.cfg.GasLimit, tx.Value)
	} else {
		//TODO save unique tx hash?
		_, _, vmerr = evm.Call(sender, tx.To, tx.Input, p.cfg.GasLimit, tx.Value)
	}

	statedb.Finalise(true)
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

func getLogs(txHash common.Hash, blockNumber uint64, blockHash common.Hash, statedb *state.StateDB) []*types.Log {
	logs := statedb.GetLogs(txHash, blockNumber, blockHash)
	ret := make([]*types.Log, len(logs))
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

func newTxContext(cfg *runtime.Config, from common.Address) vm.TxContext {
	return vm.TxContext{
		Origin:     from,
		GasPrice:   cfg.GasPrice,
		BlobHashes: cfg.BlobHashes,
	}
}

func newBlockContext(cfg *runtime.Config, blockNumber *big.Int) vm.BlockContext {
	return vm.BlockContext{
		CanTransfer: core.CanTransfer,
		Transfer:    core.Transfer,
		GetHash:     cfg.GetHashFn,
		Coinbase:    cfg.Coinbase,
		BlockNumber: cfg.BlockNumber,
		Time:        cfg.Time,
		Difficulty:  cfg.Difficulty,
		GasLimit:    cfg.GasLimit,
		BaseFee:     cfg.BaseFee,
		Random:      cfg.Random,
	}
}

func preCheck(tx *types.Transaction) (bool, error) {
	sigPublicKey, err := crypto.Ecrecover(tx.Hash().Bytes(), tx.Signature)
	if err != nil {
		return false, err
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(sigPublicKey[1:])[12:])
	return bytes.Equal(addr.Bytes(), tx.From.Bytes()), nil
}
