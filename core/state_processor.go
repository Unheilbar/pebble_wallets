package core

import (
	"bytes"
	"log"
	"math"
	"math/big"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core/vm/runtime"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
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
		blockNumber = block.Number()
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
	evm.Reset(newTxContext(tx.From), statedb)
	var (
		sender           = vm.AccountRef(tx.From)
		contractCreation = tx.To == common.Address{}
		contrAddr        common.Address
		vmerr            error
	)

	// TODO CHECK UNIQUNESS of tx HASH signature etc..
	preCheck(tx)

	if contractCreation {
		_, contrAddr, _, vmerr = evm.Create(sender, tx.Input, p.cfg.GasLimit, new(big.Int))
	} else {
		//TODO save unique tx hash?
		_, _, vmerr = evm.Call(sender, tx.To, tx.Input, p.cfg.GasLimit, new(big.Int))
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

func ApplyTransactions(chain *Blockchain, statedb *state.StateDB, header *types.Header, txs []*types.Transaction) ([]*types.Transaction, []*types.Receipt, error) {
	if len(txs) == 0 {
		return nil, nil, nil
	}

	cfg := getDefaultCfg()
	evm := runtime.NewEnv(cfg)
	blockCtx := NewEVMBlockContext(header)
	evm.SetBlockContext(blockCtx)
	var receipts []*types.Receipt
	var appliedTxs []*types.Transaction
	blockHash := header.Hash()
	blockNumber := header.Number
	txCount := 0
	for _, tx := range txs {
		evm.Reset(newTxContext(tx.From), statedb)
		snap := statedb.Snapshot()
		statedb.SetTxContext(tx.Hash(), txCount)
		result, err := ApplyMessage(evm, tx)
		if err != nil {
			statedb.RevertToSnapshot(snap)
			log.Println("tx failed, skipped", err)
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

		created := tx.To == common.Address{}
		if created {
			contractAddress := crypto.CreateAddress(evm.TxContext.Origin, statedb.GetNonce(tx.From)-1) //TODO think of something //nonce has been increased
			receipt.ContractAddress = contractAddress
			log.Println("successed deploy", contractAddress, "sender ", tx.From.Hex(), "nonce", statedb.GetNonce(tx.From)-1)
		}

		receipt.Logs = getLogs(tx.Hash(), header.Number.Uint64(), header.Hash(), statedb)
		receipt.BlockHash = blockHash
		receipt.BlockNumber = blockNumber
		receipt.TransactionIndex = uint(statedb.TxIndex())
		revertReason := result.Revert()
		if revertReason != nil {
			receipt.RevertReason = revertReason
		}

		receipts = append(receipts, receipt)
	}

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

func getDefaultCfg() *runtime.Config {
	cfg := new(runtime.Config)
	cfg.ChainConfig = &params.ChainConfig{
		ChainID:             big.NewInt(1),
		HomesteadBlock:      new(big.Int),
		DAOForkBlock:        new(big.Int),
		DAOForkSupport:      false,
		EIP150Block:         new(big.Int),
		EIP155Block:         new(big.Int),
		EIP158Block:         new(big.Int),
		ByzantiumBlock:      new(big.Int),
		ConstantinopleBlock: new(big.Int),
		PetersburgBlock:     new(big.Int),
		IstanbulBlock:       new(big.Int),
		MuirGlacierBlock:    new(big.Int),
		BerlinBlock:         new(big.Int),
		LondonBlock:         new(big.Int),
	}

	cfg.ChainConfig.CancunTime = &[]uint64{0}[0] // default cancun mode

	if cfg.Difficulty == nil {
		cfg.Difficulty = new(big.Int)
	}
	if cfg.GasLimit == 0 {
		cfg.GasLimit = math.MaxUint64
	}
	if cfg.GasPrice == nil {
		cfg.GasPrice = new(big.Int)
	}

	if cfg.Value == nil {
		cfg.Value = new(big.Int)
	}
	if cfg.BlockNumber == nil {
		cfg.BlockNumber = big.NewInt(1)
	}
	if cfg.GetHashFn == nil { //TODO probably replace since can be used inside EVM
		cfg.GetHashFn = func(n uint64) common.Hash {
			return common.BytesToHash(crypto.Keccak256([]byte(new(big.Int).SetUint64(n).String())))
		}
	}

	if cfg.BaseFee == nil {
		cfg.BaseFee = big.NewInt(params.InitialBaseFee)
	}

	return cfg
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
		baseFee     *big.Int
		random      *common.Hash
	)

	return vm.BlockContext{
		CanTransfer: core.CanTransfer,
		Transfer:    core.Transfer,
		GetHash: func(n uint64) common.Hash {
			return common.BytesToHash(crypto.Keccak256([]byte(new(big.Int).SetUint64(n).String())))
		},
		Coinbase:    beneficiary,
		BlockNumber: new(big.Int).Set(header.Number),
		Time:        header.Time,
		Difficulty:  new(big.Int),
		BaseFee:     baseFee,
		GasLimit:    math.MaxUint64,
		Random:      random,
	}
}
