package types

import (
	"io"
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

type Header struct {
	ParentHash  common.Hash
	TxHash      common.Hash
	ReceiptHash common.Hash
	Number      *big.Int
	Root        common.Hash // state root
	Time        uint64
}

func (h *Header) Hash() common.Hash {
	return rlpHash(h)
}

func (h *Header) NumberU64() uint64 {
	return h.Number.Uint64()
}

type Block struct {
	header *Header

	Transactions []*Transaction
	Receipts     []*Receipt

	// caches
	hash atomic.Value
	size atomic.Value
}

func (block *Block) Hash() common.Hash {
	return block.header.Hash()
}

var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}

func NewBlock(header *Header, txs []*Transaction, receipts []*Receipt, hasher TrieHasher) *Block {
	b := &Block{
		Transactions: txs,
		Receipts:     receipts,
	}

	header.ReceiptHash = DeriveSha(Receipts(receipts), hasher)
	header.TxHash = DeriveSha(Transactions(txs), hasher)

	b.header = header
	return b
}

func NewBlockWithHeader(header *Header) *Block {
	return &Block{header: CopyHeader(header)}
}

func (b *Block) Number() *big.Int  { return new(big.Int).Set(b.header.Number) }
func (b *Block) NumberU64() uint64 { return b.Number().Uint64() }
func (b *Block) Root() common.Hash { return b.header.Root }
func (b *Block) Time() uint64      { return b.header.Time }

// Header returns the block header (as a copy).
func (b *Block) Header() *Header {
	return CopyHeader(b.header)
}

// CopyHeader creates a deep copy of a block header.
func CopyHeader(h *Header) *Header {
	cpy := *h
	if cpy.Number = new(big.Int); h.Number != nil {
		cpy.Number.Set(h.Number)
	}
	return &cpy
}

// Body is a simple (mutable, non-safe) data container for storing and moving
// a block's data contents (transactions and uncles) together.
type Body struct {
	Transactions []*Transaction
	Receipts     []*Receipt
}

// Body returns the non-header content of the block.
// Note the returned data is not an independent copy.
func (b *Block) Body() *Body {
	return &Body{
		Transactions: b.Transactions,
		Receipts:     b.Receipts,
	}
}

// WithBody returns a copy of the block with the given transaction and uncle contents.
func (b *Block) WithBody(transactions []*Transaction) *Block {
	block := &Block{
		header:       b.header,
		Transactions: make([]*Transaction, len(transactions)),
	}
	copy(block.Transactions, transactions)

	return block
}

// "external" block encoding. used for eth protocol, etc.
type extblock struct {
	Header *Header
	Txs    []*Transaction
}

// DecodeRLP decodes the Ethereum
func (b *Block) DecodeRLP(s *rlp.Stream) error {
	var eb extblock
	_, size, _ := s.Kind()
	if err := s.Decode(&eb); err != nil {
		return err
	}
	b.header, b.Transactions = eb.Header, eb.Txs
	b.size.Store(common.StorageSize(rlp.ListSize(size)))
	return nil
}

// EncodeRLP serializes b into the Ethereum RLP block format.
func (b *Block) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, extblock{
		Header: b.header,
		Txs:    b.Transactions,
	})
}
func (b *Block) ParentHash() common.Hash { return b.header.ParentHash }
