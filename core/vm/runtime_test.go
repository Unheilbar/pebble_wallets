package vm

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/Unheilbar/pebbke_wallets/binding"
	"github.com/Unheilbar/pebbke_wallets/core/rawdb"
	"github.com/Unheilbar/pebbke_wallets/core/state"
	"github.com/Unheilbar/pebbke_wallets/core/vm"
	"github.com/Unheilbar/pebbke_wallets/core/vm/runtime"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
)

var size = 100000

func Test_ExecuteVM(t *testing.T) {
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

	contrBin := common.Hex2Bytes(binding.StorageMetaData.Bin[2:])
	statedb, err := state.New(common.Hash{}, sb, nil)
	cfg.State = statedb

	contrAddr, err := deploy(cfg, contrBin)
	if err != nil {
		log.Fatal(err)
	}

	var blockID uint64 = 1
	h, err := statedb.Commit(blockID, true)
	if err != nil {
		log.Fatal(err)
	}

	blockID++
	start := time.Now()
	cfg.State, err = state.New(h, sb, nil)
	if err != nil {
		log.Fatal(err)
	}
	inputs := generateInputsEmissions(size)
	root := executeInputs(cfg, blockID, contrAddr, inputs)
	res := time.Since(start)
	fmt.Println(res, "tx/s", float64(size)/(res.Seconds()), "emissions storage root", root)

	blockID++
	cfg.State, err = state.New(root, sb, nil)
	inputs = generateInputsTransfers(size)
	start = time.Now()
	root = executeInputs(cfg, blockID, contrAddr, inputs)
	res = time.Since(start)
	fmt.Println(res, "tx/s", float64(size)/(res.Seconds()), "transfers storage root", root)
}

func deploy(cfg *runtime.Config, code []byte) (common.Address, error) {
	vevm := runtime.NewEnv(cfg)
	_, addr, _, err := vevm.Create(vm.AccountRef(cfg.Origin), code, cfg.GasLimit, cfg.Value)
	if err != nil {
		log.Fatal(err)
	}
	return addr, err
}

func executeInputs(cfg *runtime.Config, blockID uint64, contrAddr common.Address, inputs [][]byte) common.Hash {
	cfg.BlockNumber = big.NewInt(int64(blockID))
	for i := 0; i < len(inputs); i++ {
		evm := runtime.NewEnv(cfg)
		_, _, err := evm.Call(
			vm.AccountRef(cfg.Origin),
			contrAddr,
			inputs[i],
			cfg.GasLimit,
			cfg.Value,
		)

		if err != nil {
			log.Fatal(err)
		}
	}

	h, err := cfg.State.Commit(blockID, true)
	if err != nil {
		log.Fatal(err)
	}
	return h
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
			fmt.Println("call hash function")
			return common.BytesToHash(crypto.Keccak256([]byte(new(big.Int).SetUint64(n).String())))
		}
	}

	cfg.Origin = common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4")
	if cfg.BaseFee == nil {
		cfg.BaseFee = big.NewInt(params.InitialBaseFee)
	}
}

func PackTX(method string, params ...interface{}) ([]byte, error) {
	bind := binding.StorageABI

	pabi, err := abi.JSON(strings.NewReader(bind))

	if err != nil {
		return nil, err
	}

	input, err := pabi.Pack(method, params...)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func generateInputsEmissions(size int) [][]byte {
	users := generateAccounts(size)
	r := make([][]byte, 0)
	for i := 0; i < size; i++ {
		input, err := PackTX("setBalance", users[i].Hex(), big.NewInt(int64(10000000)))
		if err != nil {
			log.Fatal(err, input)
		}
		r = append(r, input)
	}
	return r
}

func generateInputsTransfers(size int) [][]byte {
	users := generateAccounts(size)
	r := make([][]byte, 0)
	for i := 0; i < size; i++ {
		fromIdx := i % len(users)
		toIdx := (i + 1) % len(users)
		input, err := PackTX("transfer", users[fromIdx].Hex(), users[toIdx].Hex(), big.NewInt(1))
		if err != nil {
			log.Fatal(err)
		}
		r = append(r, input)
	}

	return r
}

func generateAccounts(size int) []common.Address {
	var s []common.Address
	for i := 0; i < size; i++ {
		k := common.BytesToAddress(crypto.Keccak256(big.NewInt(int64(i)).Bytes())[12:])

		s = append(s, k)
	}
	return s
}
