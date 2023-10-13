package eth

import (
	"log"

	"github.com/Unheilbar/pebbke_wallets/core"
	"github.com/Unheilbar/pebbke_wallets/core/txpool"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
)

type Ethereum struct {
	txPool     *txpool.TxPool
	blockchain *core.Blockchain
	chainDb    ethdb.Database // Block chain database
}

func New() *Ethereum {
	chainDb, err := rawdb.NewPebbleDBDatabase("../chaindb", 1024, 16, "some", false, false)
	if err != nil {
		log.Fatal("open pebble db when create ethereum ", err)
	}
	blockchain := core.NewBlockchain(chainDb)
	txPool := txpool.New()
	return &Ethereum{
		txPool:     txPool,
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (s *Ethereum) Blockchain() *core.Blockchain { return s.blockchain }
func (s *Ethereum) TxPool() *txpool.TxPool       { return s.txPool }
func (s *Ethereum) ChainDb() ethdb.Database      { return s.chainDb }
