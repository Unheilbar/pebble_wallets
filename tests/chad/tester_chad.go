package chad

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/Unheilbar/pebbke_wallets/binding"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type chadAcc struct {
	Address common.Address
	Private *ecdsa.PrivateKey
}

// TODO fake balances swap to bytes32
type Chad struct {
	accounts     map[common.Address]chadAcc
	orderedAccs  []chadAcc
	fakeBalances map[string]uint
}

const genesisPrivate = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"

func (c *Chad) GenerateAccEmissionsTx(contrAddr common.Address) []*types.Transaction {
	var ret []*types.Transaction

	for _, acc := range c.accounts {
		payload := getContractEmissionPayload(acc.Address, acc.Address.Hash().Bytes()[:32], contrAddr)
		c.fakeBalances[acc.Address.Hex()] += emissionVal
		genTx := types.NewTx(types.TxData{
			From: acc.Address,
			To:   &contrAddr,
			Id:   crypto.Keccak256Hash([]byte{3, 2, 1}),
			Data: payload,
		})
		sign, err := crypto.Sign(genTx.Hash().Bytes(), acc.Private)
		if err != nil {
			log.Fatal("sign err", err)
		}
		genTx.WithSignature(sign)
		ret = append(ret, genTx)
	}
	return ret
}

func (c *Chad) GetTestAccByID(id uint) chadAcc {
	return c.orderedAccs[id]
}

func (c *Chad) GenerateTransfers(size int, contrAddr common.Address) []*types.Transaction {
	var ret []*types.Transaction
	var ptrFrom int
	var ptrTo int
	genesisId := 1
	for i := 0; i < size; i++ {
		accFrom := c.orderedAccs[ptrFrom%len(c.orderedAccs)]
		ptrFrom += i
		accTo := c.orderedAccs[ptrTo%len(c.orderedAccs)]
		ptrTo += i * 2
		data := getContractTransferPayload(accFrom.Address.Hash().Bytes()[:32], accTo.Address.Hash().Bytes()[:32], contrAddr)
		id := crypto.Keccak256Hash([]byte(fmt.Sprint(genesisId)))
		genesisId++

		genTx := types.NewTx(types.TxData{
			From: accFrom.Address,
			To:   &contrAddr,
			Id:   id,
			Data: data,
		})
		sign, err := crypto.Sign(genTx.Hash().Bytes(), accFrom.Private)
		if err != nil {
			log.Fatal("sign err", err)
		}
		genTx.WithSignature(sign)
		ret = append(ret, genTx)
		c.fakeBalances[accFrom.Address.Hex()] -= transferVal
		c.fakeBalances[accTo.Address.Hex()] += transferVal
	}

	return ret
}

func (c *Chad) GenerateAccs(size int) {
	c.accounts = make(map[common.Address]chadAcc)
	c.fakeBalances = make(map[string]uint)

	prev := genesisPrivate
	for i := 0; i < size; i++ {
		privateKey, err := crypto.HexToECDSA(prev)
		if err != nil {
			log.Fatal("err gen private", err)
		}

		newAddr := common.BytesToAddress(crypto.Keccak256(crypto.FromECDSAPub(&privateKey.PublicKey)[1:])[12:])
		c.accounts[newAddr] = chadAcc{Address: newAddr, Private: privateKey}
		c.orderedAccs = append(c.orderedAccs, c.accounts[newAddr])
		prev = string(crypto.Keccak256Hash([]byte(prev)).Hex()[2:])
	}
}

func (c *Chad) GetContractDeployTX(from common.Address, contrCode []byte, id common.Hash) *types.Transaction {
	genTx := types.NewTx(types.TxData{
		From: from,
		To:   nil,
		Id:   id,
		Data: contrCode,
	})
	return genTx
}

const emissionVal = 1000000000

func getContractEmissionPayload(pubKey common.Address, wallet []byte, contrAddr common.Address) []byte {
	var arr [32]byte
	copy(arr[:], wallet)
	data, err := packTX("emission", pubKey, arr, big.NewInt(emissionVal))
	if err != nil {
		log.Fatal("err generate emission input", err, data)
	}

	return data
}

const transferVal = 1

func getContractTransferPayload(fromWallet []byte, toWallet []byte, contrAddr common.Address) []byte {
	var arrFrom [32]byte
	copy(arrFrom[:], fromWallet)
	var arrTo [32]byte
	copy(arrTo[:], toWallet)
	data, err := packTX("transfer", arrFrom, arrTo, big.NewInt(transferVal))
	if err != nil {
		log.Fatal("err generate transfer input", err, data)
	}

	return data
}

func packTX(method string, params ...interface{}) ([]byte, error) {
	bind := binding.ProxyABI

	pabi, err := abi.JSON(strings.NewReader(bind))

	if err != nil {
		return nil, err
	}

	input, err := pabi.Pack(method, params...)
	if err != nil {
		return nil, err
	}

	return input, nil
}
