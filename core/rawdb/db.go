package rawdb

import (
	"log"

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

func (d *DB) Get(key common.Address) []byte {
	val, closer, err := d.pebble.Get(key[:])

	if err != nil {
		log.Fatal(err, " key ", key)
	}
	defer closer.Close()
	return val
}

func (d *DB) CheckTX(key common.Hash) bool {
	_, closer, err := d.pebble.Get(key.Bytes())
	if err == pebble.ErrNotFound {
		return false
	}
	if err != nil {
		log.Fatal(err)
	}
	closer.Close()
	return true
}

func (t *TX) Set(key common.Address, val []byte) {
	err := t.batch.Set(key[:], val, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (t *TX) MarkTX(h common.Hash) {
	t.batch.Set(h[:], []byte{1}, nil)
}

func (t *TX) Commit() {
	err := t.batch.Commit(pebble.Sync)
	if err != nil {
		log.Fatal(err)

	}
}

func (d *DB) Beginx() *TX {
	return &TX{
		d.pebble.NewBatch(),
	}
}
