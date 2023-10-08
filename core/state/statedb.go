package state

import (
	"log"
	"math/big"

	"github.com/Unheilbar/pebbke_wallets/core/rawdb"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type Tx interface {
	Set(key common.Address, data []byte)
	Commit()
}

type Database interface {
	Beginx() *rawdb.TX
	Get(key common.Address) []byte
	CheckTX(key common.Hash) bool
}

type StateDB struct {
	db Database

	dirtyObjects map[common.Address]*stateObject
	usedTxs      []common.Hash
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

func (s *StateDB) ApplyTransactions(txs []*types.Transaction) common.Hash {
	s.prefetch(txs)
	for _, tx := range txs {
		s.dirtyObjects[tx.From].data.Balance.Sub(s.dirtyObjects[tx.From].data.Balance, tx.Value)
		s.dirtyObjects[tx.From].data.MonthAmount.Add(s.dirtyObjects[tx.From].data.MonthAmount, tx.Value)
		s.dirtyObjects[tx.To].data.Balance.Add(s.dirtyObjects[tx.To].data.Balance, tx.Value)
	}
	return s.Commit()
}

func (s *StateDB) Commit() common.Hash {
	tx := s.db.Beginx()
	var stateChanges []byte
	for k, v := range s.dirtyObjects {
		obj, err := rlp.EncodeToBytes(v.data)
		stateChanges = append(stateChanges, obj...)
		if err != nil {
			log.Fatal(err)
		}
		tx.Set(k, obj)
	}
	for _, h := range s.usedTxs {
		tx.MarkTX(h)
	}
	hash := crypto.Keccak256(stateChanges)
	tx.Commit()
	s.dirtyObjects = nil
	return common.BytesToHash(hash)
}

func (s *StateDB) InitWallets(wallets []common.Address) {
	s.dirtyObjects = make(map[common.Address]*stateObject)
	for _, w := range wallets {
		s.dirtyObjects[w] = &stateObject{
			data: Account{Address: w, Balance: big.NewInt(1000000)},
		}
	}
	s.Commit()
}

func (s *StateDB) prefetch(txs []*types.Transaction) {
	s.dirtyObjects = make(map[common.Address]*stateObject, len(txs)*2)
	s.usedTxs = make([]common.Hash, 0, len(txs))
	for _, tx := range txs {
		if s.exists(tx) {
			log.Printf("tx exists with hash %s", tx.Id)
			continue
		}
		s.dirtyObjects[tx.From] = s.getStateObject(tx.From)
		s.dirtyObjects[tx.To] = s.getStateObject(tx.To)

	}
}

func (s *StateDB) exists(tx *types.Transaction) bool {
	k := append(tx.From[:], tx.Id.Bytes()...)
	h := crypto.Keccak256Hash(k)
	if s.db.CheckTX(h) {
		return true
	}
	s.usedTxs = append(s.usedTxs, crypto.Keccak256Hash(k))
	return false
}

func (s *StateDB) getStateObject(k common.Address) *stateObject {
	obj := new(Account)
	b := s.db.Get(k)
	err := rlp.DecodeBytes(b, &obj)
	if err != nil {
		log.Fatal("get ", err, b)
	}
	return &stateObject{data: *obj}
}
