package tests

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/Unheilbar/pebbke_wallets/binding"
	"github.com/Unheilbar/pebbke_wallets/core"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core/vm/runtime"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

const walletsAmount = 10000
const transfersAmount = 10000

type Blockchain struct {
	db ethdb.Database
}

func NewBlockchain() (*Blockchain, error) {
	rdb, err := rawdb.NewPebbleDBDatabase("../chaindb", 1024, 16, "some", false, false)
	if err != nil {
		return nil, err
	}

	return &Blockchain{
		db: rdb,
	}, nil
}

func (b *Blockchain) InsertBlock(block *types.Block) error {
	batch := b.db.NewBatch()

	rlpBlock, err := rlp.EncodeToBytes(block)
	if err != nil {
		return err
	}
	fmt.Println("insert", block.Hash(), block.Number().Uint64())
	rawdb.WriteBodyRLP(batch, block.Hash(), block.Number().Uint64(), rlpBlock)
	return batch.Write()
}

// hasherPool holds LegacyKeccak256 hashers for rlpHash.

func Test__StateProcessor(t *testing.T) {
	blockchain, err := NewBlockchain()
	if err != nil {
		log.Fatal("create blockchain ", err)
	}

	var tester Chad

	tester.generateAccs(walletsAmount)
	stProcessor := newProcessor()

	rdb, err := rawdb.NewPebbleDBDatabase("../testdb", 1024, 16, "some", false, false)
	if err != nil {
		log.Fatal(err)
	}

	triedb := trie.NewDatabase(rdb, trie.HashDefaults)

	sb := state.NewDatabaseWithNodeDB(rdb, triedb)
	if err != nil {
		log.Fatal("err node db open", err)
	}

	statedb, err := state.New(common.Hash{}, sb, nil)
	if err != nil {
		log.Fatal("err state db open", err)
	}
	var blockID int64 = 1
	deployStartTime := time.Now()

	block := types.Block{
		Header: &types.Header{
			Number: big.NewInt(blockID),
		},
		Transactions: []*types.Transaction{GetContractDeployTX(common.Hex2Bytes(binding.StorageMetaData.Bin[2:]))},
	}

	receipts := stProcessor.Process(block, statedb)
	receipt := receipts[0]
	contrAddr := receipt.ContractAddress

	newRoot, err := statedb.Commit(uint64(blockID), true)
	if err != nil {
		log.Fatal("err commit sb deploy", err)
	}

	err = sb.TrieDB().Commit(newRoot, false)
	if err != nil {
		log.Fatal("err commit tdb deploy", err)
	}
	fmt.Println("deploy receipt addr", contrAddr, receipt.Status, "root", newRoot, "time", time.Since(deployStartTime))
	blockID++

	statedb, err = state.New(newRoot, sb, nil)
	if err != nil {
		log.Fatal("err state db open", err)
	}
	emissions := tester.generateAccEmissionsTx(receipts[0].ContractAddress)
	emissionsStartTime := time.Now()

	block = types.Block{
		Header: &types.Header{
			Number: big.NewInt(blockID),
		},
		Transactions: emissions,
	}
	receipts = stProcessor.Process(block, statedb)

	newRoot, err = statedb.Commit(uint64(blockID), true)
	if err != nil {
		log.Fatal("err state db commit", err)
	}
	err = sb.TrieDB().Commit(newRoot, false)
	if err != nil {
		log.Fatal("err commit emissions deploy", err)
	}
	emisRoot := newRoot
	evaltime := time.Since(emissionsStartTime)
	fmt.Println("transaction count", len(emissions))
	fmt.Println("emission done receipts len", len(receipts), "status", receipts[0].Status, "root", newRoot, "time", evaltime, float64(walletsAmount)/evaltime.Seconds(), "tx/s")

	blockID++
	transfers := tester.generateTransfers(transfersAmount, contrAddr)
	statedb, err = state.New(newRoot, sb, nil)
	if err != nil {
		log.Fatal("err state db open", err)
	}
	transfersStartTime := time.Now()
	receipts = stProcessor.Process(
		types.Block{
			Header: &types.Header{
				Number: big.NewInt(blockID),
			},
			Transactions: transfers,
		},
		statedb)

	newRoot, err = statedb.Commit(uint64(blockID), true)
	if err != nil {
		log.Fatal("err state db commit", err)
	}
	err = sb.TrieDB().Commit(newRoot, false)
	if err != nil {
		log.Fatal("err commit tb transfers deploy", err)
	}
	evaltime = time.Since(transfersStartTime)

	controlWallet := tester.orderedAccs[1].from.Hex()
	fmt.Println("transfers done receipts len", len(receipts), "status", receipts[0].Status, "root", newRoot, "time", evaltime, float64(transfersAmount)/evaltime.Seconds(), "tx/s")
	fmt.Println("control fake balance: ", tester.fakeBalances[controlWallet])
	fmt.Println("contr evm balance:    ", getWalletBalanceForRoot(newRoot, controlWallet, sb, contrAddr))
	fmt.Println("Close databases...root", newRoot)

	err = rdb.Close()
	err = sb.DiskDB().Close()
	err = triedb.Close()
	if err != nil {
		log.Fatal(err)
	}
	rdb, err = rawdb.NewPebbleDBDatabase("../testdb", 1024, 16, "some", false, false)
	if err != nil {
		log.Fatal(err)
	}

	triedb = trie.NewDatabase(rdb, trie.HashDefaults)
	sb = state.NewDatabaseWithNodeDB(rdb, triedb)
	statedb, err = state.New(newRoot, sb, nil)
	fmt.Println("reopen DB.. control Balance", getWalletBalanceForRoot(newRoot, controlWallet, sb, contrAddr))
	fmt.Println("emis moment Balance", getWalletBalanceForRoot(emisRoot, controlWallet, sb, contrAddr))

	h := &types.Header{Number: big.NewInt(blockID), Root: newRoot}
	bl := types.NewBlock(h, transfers, receipts)
	insertStart := time.Now()
	err = blockchain.InsertBlock(bl)
	fmt.Println("insert block time ", time.Since(insertStart))
	if err != nil {
		log.Fatal(err)
	}
	readBlock := ReadBlock(blockchain.db, bl.Hash(), bl.Number().Uint64())
	fmt.Println(readBlock.Transactions[0].Id)
	walletBalance1 := getWalletBalanceForRoot(readBlock.Root(), readBlock.Transactions[1].From.Hex(), sb, contrAddr)
	expected := tester.fakeBalances[readBlock.Transactions[1].From.Hex()]
	assert.Equal(t, expected, uint(walletBalance1.Uint64()))
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

func getWalletBalanceForRoot(root common.Hash, s string, sb state.Database, contrAddr common.Address) *big.Int {
	statedb, err := state.New(root, sb, nil)
	if err != nil {
		log.Fatal("err state db open", err)
	}

	cfg := new(runtime.Config)
	cfg.State = statedb
	setDefaults(cfg)
	nenv := runtime.NewEnv(cfg)
	input, err := packTX("getBalance", s)
	if err != nil {
		log.Fatal("err state db open", err)
	}

	ret, _, err := nenv.Call(vm.AccountRef(cfg.Origin),
		contrAddr,
		input,
		cfg.GasLimit,
		cfg.Value)
	if err != nil {
		log.Fatal(err)
	}
	sabi, _ := binding.StorageMetaData.GetAbi()
	res, err := sabi.Unpack("getBalance", ret)
	if err != nil {
		log.Fatal(err)
	}
	out0 := *abi.ConvertType(res[0], new(*big.Int)).(**big.Int)
	return out0
}
func ReadBlock(db ethdb.Reader, hash common.Hash, number uint64) *types.Block {
	fmt.Println("read ", hash, number)
	data := rawdb.ReadBodyRLP(db, hash, number)
	if len(data) == 0 {
		log.Fatal("read block nil")
		return nil
	}
	block := new(types.Block)
	if err := rlp.Decode(bytes.NewReader(data), block); err != nil {
		fmt.Println("error decode block", err)
		return nil
	}
	return block

}
