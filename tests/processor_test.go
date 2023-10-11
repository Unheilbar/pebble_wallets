package tests

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"testing"

	"github.com/Unheilbar/pebbke_wallets/binding"
	"github.com/Unheilbar/pebbke_wallets/core"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm/runtime"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
)

func Test__StateProcessor(t *testing.T) {
	var tester Chad

	tester.generateAccs(1000)
	stProcessor := newProcessor()

	rdb, err := rawdb.NewPebbleDBDatabase("../../testdb", 1024, 16, "some", false, false)

	triedb := trie.NewDatabase(rdb, trie.HashDefaults)
	sb := state.NewDatabaseWithNodeDB(rdb, triedb)
	if err != nil {
		log.Fatal("err node db open", err)
	}

	statedb, err := state.New(common.Hash{}, sb, nil)
	if err != nil {
		log.Fatal("err state db open", err)
	}

	receipts := stProcessor.Process(
		types.Block{
			Transactions: []*types.Transaction{tester.getContractDeployTX(common.Hex2Bytes(binding.StorageMetaData.Bin[2:]))},
			Number:       big.NewInt(1),
		},
		statedb)
	receipt := receipts[0]
	fmt.Println(receipt.ContractAddress, receipt.Status)
}

func newProcessor() *core.StateProcessor {
	cfg := new(runtime.Config)
	setDefaults(cfg)
	return core.NewStateProcessor(cfg)
}

// sets defaults on the config
func setDefaults(cfg *runtime.Config) {
	if cfg.ChainConfig == nil {
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
	}
	cfg.ChainConfig.CancunTime = &[]uint64{0}[0]

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
	if cfg.GetHashFn == nil {
		cfg.GetHashFn = func(n uint64) common.Hash {
			return common.BytesToHash(crypto.Keccak256([]byte(new(big.Int).SetUint64(n).String())))
		}
	}

	cfg.Origin = common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4")
	if cfg.BaseFee == nil {
		cfg.BaseFee = big.NewInt(params.InitialBaseFee)
	}
}
