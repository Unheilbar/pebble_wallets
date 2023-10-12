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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"

	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm/runtime"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

const walletsAmount = 10000

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

	rawdb.WriteBodyRLP(batch, block.Hash(), block.Number.Uint64(), rlpBlock)
	return nil
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
		Transactions: []*types.Transaction{getContractDeployTX(common.Hex2Bytes(binding.StorageMetaData.Bin[2:]))},
		Number:       big.NewInt(blockID),
	}

	receipts := stProcessor.Process(block, statedb)
	receipt := receipts[0]
	contrAddr := receipt.ContractAddress

	newRoot, err := statedb.Commit(uint64(blockID), true)
	fmt.Println("deploy receipt addr", contrAddr, receipt.Status, "root", newRoot, "time", time.Since(deployStartTime))
	blockID++
	if err != nil {
		log.Fatal("err commit deploy", err)
	}

	//Insert block in rawdb
	blockchain.InsertBlock(&block)

	statedb, err = state.New(newRoot, sb, nil)
	if err != nil {
		log.Fatal("err state db open", err)
	}
	emissions := tester.generateAccEmissionsTx(receipts[0].ContractAddress)
	emissionsStartTime := time.Now()

	block = types.Block{
		Transactions: emissions,
		Number:       big.NewInt(blockID),
	}
	receipts = stProcessor.Process(block, statedb)

	newRoot, err = statedb.Commit(uint64(blockID), true)
	if err != nil {
		log.Fatal("err state db open", err)
	}
	//Insert block in rawdb
	blockchain.InsertBlock(&block)

	evaltime := time.Since(emissionsStartTime)
	fmt.Println("transaction count", len(emissions))
	fmt.Println("emission done receipts len", len(receipts), "status", receipts[0].Status, "root", newRoot, "time", evaltime, float64(walletsAmount)/evaltime.Seconds(), "tx/s")

	_ = ReadBlock(blockchain.db, block.Hash(), block.Number.Uint64())
	//	for _, tx := range chainBlock.Transactions {
	//		fmt.Println("chain transactino id", tx.Id)
	//	}
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

func ReadBlock(db ethdb.Reader, hash common.Hash, number uint64) *types.Block {
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
