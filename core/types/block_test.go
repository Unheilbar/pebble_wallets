package types

import (
	"bytes"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

var (
	expFrom = common.HexToAddress("0x5a443704dd4B594B382c22a083e2BD3090A6feF3")
	expTo   = common.HexToAddress("0x5a443704dd4B594B382c22a083e2BD3090A6feF4")
	expId   = common.HexToHash("0xf5f3ef46c9494b0381a2882498b12c4b2c86decaffc2554705838e3c80d3f1a3")
	expSig  = []byte{7, 8, 9}
	expData = []byte{1, 2, 3}

	expFrom2 = common.HexToAddress("0xbb443704dd4B594B382c22a083e2BD3090A6feD5")
	expId2   = common.HexToHash("0x88f3ef46c9494b0381a2882498b12c4b2c86decaffc2554705838e3c80d3f199")
	expData2 = []byte{3, 2, 1}
	expTo2   = common.HexToAddress("0x1433ef46c9494b0381a2882498b12c4b2c86decaffc2554705838e3c80d3f1a9")
	expSig2  = []byte{12, 13, 14}

	expParentHash = common.HexToHash("0x1")
	expTxHash     = common.HexToHash("0x2")
	expRoot       = common.HexToHash("0x3")

	expParentHash2 = common.HexToHash("0x4")
	expTxHash2     = common.HexToHash("0x5")
	expRoot2       = common.HexToHash("0x6")

	tx1 = NewTx(TxData{
		From:      expFrom,
		To:        &expTo,
		Id:        expId,
		Signature: expSig,
		Data:      expData,
	})

	tx2 = NewTx(TxData{
		From:      expFrom2,
		To:        &expTo2,
		Id:        expId2,
		Signature: expSig2,
		Data:      expData2,
	})

	blockNumber = big.NewInt(1)
)

func Test__RlpEncode(t *testing.T) {
	// pre check test data
	if bytes.Equal(tx1.Data(), tx2.Data()) {
		log.Fatal("test data wrong tx1 data and tx2 data")
	}

	header := &Header{
		ParentHash: expParentHash,
		TxHash:     expTxHash,
		Root:       expRoot,
		Number:     blockNumber,
	}

	expBlock := NewBlockWithHeader(header).WithBody([]*Transaction{tx1, tx2})

	if expBlock.Transactions[0].Id().Hex() != expId.Hex() ||
		expBlock.Transactions[1].Id().Hex() != expId2.Hex() ||
		expBlock.Transactions[0].To().Hex() != expTo.Hex() ||
		expBlock.Transactions[1].To().Hex() != expTo2.Hex() {
		log.Fatal("block build failed ")
	}

	bytesBlock, err := rlp.EncodeToBytes(&expBlock)
	if err != nil {
		log.Fatal("err to encode", err)
	}

	var gotBlock Block

	err = rlp.DecodeBytes(bytesBlock, &gotBlock)
	if err != nil {
		log.Fatal("err to decode ", err)
	}

	// assert encoded data
	var s common.Hash
	s.Hex()

	gotTx := gotBlock.Transactions[0]

	gotTxId := gotTx.Id()
	if !bytes.Equal(gotTxId.Bytes(), expId.Bytes()) {
		log.Fatal("\nfirst id doesnt match \ngot ", gotTx.Id().Hex(), "\nexpected ", expId.Hex())
	}

	gotTx2 := gotBlock.Transactions[1]
	gotTxId2 := gotTx2.Id()
	if !bytes.Equal(gotTxId2.Bytes(), expId2.Bytes()) {
		log.Fatal("\nsecond id doesnt match \ngot ", gotTxId2.Hex(), "\nexpected ", expTo2.Hex())
	}

	gotHeaderRoot := gotBlock.header.Root.Hex()
	if gotHeaderRoot != expBlock.header.Root.Hex() {
		log.Fatal("\nheader root doesnt match got ", gotTx.Id().Hex(), "\nexpected", expId.Hex())
	}

	gotData1 := gotTx.Data()
	gotData2 := gotTx2.Data()
	if !bytes.Equal(gotData1, expData) || !bytes.Equal(gotData2, expData2) {
		log.Fatal("\ndata doesnt match")
	}
}
