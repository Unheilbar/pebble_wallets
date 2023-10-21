package eth

import (
	"log"

	"github.com/Unheilbar/pebbke_wallets/core"
	"github.com/Unheilbar/pebbke_wallets/core/rawdb"
	"github.com/Unheilbar/pebbke_wallets/core/txpool"
	"github.com/ethereum/go-ethereum/ethdb"
)

type Ethereum struct {
	txPool     *txpool.TxPool
	blockchain *core.Blockchain
	chainDb    ethdb.Database // Block chain database
}

func New(dbPath string, apiServerHost string, apiServerPort string) *Ethereum {
	chainDb, err := rawdb.NewPebbleDBDatabase(dbPath, 2048, 16, "some", false, false)
	if err != nil {
		log.Fatal("open pebble db when create ethereum ", err)
	}
	blockchain := core.NewBlockchain(chainDb)

	txPool := txpool.New()

	ethereum := &Ethereum{
		txPool:     txPool,
		blockchain: blockchain,
		chainDb:    chainDb,
	}

	ethApi := NewApi(ethereum)

	NewServer(ethApi, apiServerHost, apiServerPort)

	return ethereum
}

func (s *Ethereum) Blockchain() *core.Blockchain { return s.blockchain }
func (s *Ethereum) TxPool() *txpool.TxPool       { return s.txPool }
func (s *Ethereum) ChainDb() ethdb.Database      { return s.chainDb }
