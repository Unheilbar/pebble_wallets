package types

import (
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

func Test__RlpEncode(t *testing.T) {
	expFrom := common.HexToAddress("0x5a443704dd4B594B382c22a083e2BD3090A6feF3")
	expTo := common.HexToAddress("0x5a443704dd4B594B382c22a083e2BD3090A6feF4")
	expId := common.HexToHash("0x121312312312312321312313213")
	expSig := common.Hex2Bytes("0x111")
	data := common.Hex2Bytes("0x3223")
	tx := NewTx(TxData{
		From:      expFrom,
		To:        &expTo,
		Id:        expId,
		Signature: expSig,
		Data:      data,
	})
	expParentHash := common.HexToHash("0x1")
	expTxHash := common.HexToHash("0x2")
	expRoot := common.HexToHash("0x3")

	header := &Header{
		ParentHash: expParentHash,
		TxHash:     expTxHash,
		Root:       expRoot,
	}

	block := Block{
		header:       header,
		Transactions: []*Transaction{tx},
	}

	var newBlock Block

	bblock, err := rlp.EncodeToBytes(&block)
	if err != nil {
		log.Fatal("err to encode", err)
	}

	err = rlp.DecodeBytes(bblock, &newBlock)
	if err != nil {
		log.Fatal("err to decode ", err)
	}
	var s common.Hash
	s.Hex()

	gotTx := newBlock.Transactions[0]
	if gotTx.Id().Hex() != expId.Hex() {
		log.Fatal("\nid doesnt match got ", gotTx.Id().Hex(), "\nexpected", expId.Hex())
	}

	gotHeaderRoot := newBlock.header.Root.Hex()
	if gotHeaderRoot != block.header.Root.Hex() {
		log.Fatal("\nheader root doesnt match got ", gotTx.Id().Hex(), "\nexpected", expId.Hex())
	}
}
