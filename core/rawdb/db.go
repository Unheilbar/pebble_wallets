package rawdb

import (
	"log"
	"math/big"

	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
)

type DB struct {
	pebble *pebble.DB
}

func New(db *pebble.DB) *DB {
	return &DB{
		db,
	}
}

type TX struct {
	batch *pebble.Batch
}

func (d *DB) Get(key common.Address) uint {
	val, closer, err := d.pebble.Get(key[:])
	defer closer.Close()
	if err != nil {
		log.Fatal(err)
	}
	var v big.Int
	return uint(v.SetBytes(val).Uint64())
}

func (t *TX) Set(key common.Address, value uint) {
	err := t.batch.Set(key[:], big.NewInt(int64(value)).Bytes(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (t *TX) Commit() {
	err := t.batch.Commit(nil)
	if err != nil {
		log.Fatal(err)

	}
}

func (d *DB) Beginx() *TX {
	return &TX{
		d.pebble.NewBatch(),
	}
}

type tx interface {
	Set(key common.Address, value uint)
	Commit()
}

type Database interface {
	Get(key common.Address) uint
	Beginx() tx
}
