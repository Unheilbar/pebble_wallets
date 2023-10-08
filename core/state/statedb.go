package state

import (
	"log"
	"sync"

	"github.com/Unheilbar/pebbke_wallets/core/rawdb"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
)

type Tx interface {
	Set(key common.Address, value uint)
	Commit()
}

type Database interface {
	Get(key common.Address) uint
	Beginx() *rawdb.TX
}

type StateDB struct {
	db Database

	Cached map[common.Address]uint
	mx     sync.Mutex
}

func NewState(filename string) *StateDB {
	db, err := pebble.Open(filename, &pebble.Options{})
	if err != nil {
		log.Fatal(err)
	}

	raw := rawdb.New(db)

	return &StateDB{
		db: raw,
	}
}

func (s *StateDB) ApplyTransactions(txs []*types.Transaction) {
	s.prefetch(txs)
	for _, tx := range txs {
		s.Cached[tx.From] -= tx.Value
		s.Cached[tx.To] += tx.Value
	}
}

func (s *StateDB) Commit() {
	tx := s.db.Beginx()
	for k, v := range s.Cached {
		tx.Set(k, v)
	}
	tx.Commit()
	s.Cached = nil
}

func (s *StateDB) InitWallets(wallets []common.Address) {
	s.Cached = make(map[common.Address]uint)
	for _, w := range wallets {
		s.Cached[w] = 10000000
	}
	s.Commit()
}

func (s *StateDB) prefetch(txs []*types.Transaction) {
	s.Cached = make(map[common.Address]uint, len(txs)*2)
	for _, tx := range txs {
		s.Cached[tx.From] = s.db.Get(tx.From)
		s.Cached[tx.To] = s.db.Get(tx.To)
	}
}
