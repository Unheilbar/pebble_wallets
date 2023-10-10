package vm

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core/vm/runtime"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
)

var size = 100000

func Test_Run(t *testing.T) {
	start := time.Now()
	rdb, err := rawdb.NewPebbleDBDatabase("../../test", 1024, 16, "some", false, false)
	cfg := new(runtime.Config)
	setDefaults(cfg)
	if err != nil {
		log.Fatal(err)
	}

	triedb := trie.NewDatabase(rdb, trie.HashDefaults)
	sb := state.NewDatabaseWithNodeDB(rdb, triedb)
	if err != nil {
		log.Fatal(err)
	}

	st, _ := state.New(common.Hash{}, sb, nil)
	h, _ := st.Commit(1, true)
	st, _ = state.New(h, sb, nil)
	h2, _ := st.Commit(2, true)
	fmt.Println(h, h2)

	evm := getVM(common.Hex2Bytes(erc20), cfg, sb)
	var blockID uint64 = 1
	for i := 0; i < size; i++ {
		r, _, err := evm.Call(
			vm.AccountRef(cfg.Origin),
			common.BytesToAddress([]byte("contract")),
			common.Hex2Bytes(calldata),
			cfg.GasLimit,
			cfg.Value,
		)
		if err != nil {
			log.Fatal(err)
		}
		if len(r) == 0 {
			log.Fatal("nil res")
		}
		h, err := cfg.State.Commit(blockID, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(h)
		blockID++
		cfg.State, err = state.New(h, sb, nil)
		if err != nil {
			log.Fatal(err)
		}

		evm = getVM(common.Hex2Bytes(erc20), cfg, nil)
	}
	res := time.Since(start)
	fmt.Println(res, "tx/s", float64(size)/(res.Seconds()))
}

func Execute(code, input []byte, cfg *runtime.Config, statedb state.Database) ([]byte, *state.StateDB, error) {
	if cfg == nil {
		cfg = new(runtime.Config)
	}

	if cfg.State == nil {
		cfg.State, _ = state.New(types.EmptyRootHash, statedb, nil)
	}
	var (
		address = common.BytesToAddress([]byte("contract"))
		vmenv   = runtime.NewEnv(cfg)
		sender  = vm.AccountRef(cfg.Origin)
		rules   = cfg.ChainConfig.Rules(vmenv.Context.BlockNumber, vmenv.Context.Random != nil, vmenv.Context.Time)
	)
	// Execute the preparatory steps for state transition which includes:
	// - prepare accessList(post-berlin)
	// - reset transient storage(eip 1153)
	cfg.State.Prepare(rules, cfg.Origin, cfg.Coinbase, &address, vm.ActivePrecompiles(rules), nil)
	cfg.State.CreateAccount(address)
	// set the receiver's (the executing contract) code for execution.
	cfg.State.SetCode(address, code)
	// Call the code with the given configuration.
	ret, _, err := vmenv.Call(
		sender,
		common.BytesToAddress([]byte("contract")),
		input,
		cfg.GasLimit,
		cfg.Value,
	)

	return ret, cfg.State, err
}

func getVM(code []byte, cfg *runtime.Config, statedb state.Database) *vm.EVM {
	if cfg == nil {
		cfg = new(runtime.Config)
	}

	if cfg.State == nil {
		cfg.State, _ = state.New(types.EmptyRootHash, statedb, nil)
	}
	var (
		address = common.BytesToAddress([]byte("contract"))
		vmenv   = runtime.NewEnv(cfg)
		//sender  = vm.AccountRef(cfg.Origin)
		rules = cfg.ChainConfig.Rules(vmenv.Context.BlockNumber, vmenv.Context.Random != nil, vmenv.Context.Time)
	)
	// Execute the preparatory steps for state transition which includes:
	// - prepare accessList(post-berlin)
	// - reset transient storage(eip 1153)
	cfg.State.Prepare(rules, cfg.Origin, cfg.Coinbase, &address, vm.ActivePrecompiles(rules), nil)
	cfg.State.CreateAccount(address)
	// set the receiver's (the executing contract) code for execution.
	cfg.State.SetCode(address, code)

	return vmenv
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
