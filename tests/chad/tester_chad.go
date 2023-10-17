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
	From    common.Address
	Private *ecdsa.PrivateKey
}

type Chad struct {
	accounts     map[common.Address]chadAcc
	orderedAccs  []chadAcc
	fakeBalances map[string]uint
}

const genesisPrivate = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"

func (c *Chad) GenerateAccEmissionsTx(contrAddr common.Address) []*types.Transaction {
	var ret []*types.Transaction

	for _, acc := range c.accounts {
		tx := getContractEmissionTX(acc.From.Hash().Bytes()[:32], contrAddr)
		c.fakeBalances[acc.From.Hex()] += emissionVal
		tx.From = acc.From
		sign, err := crypto.Sign(tx.Hash().Bytes(), acc.Private)
		if err != nil {
			log.Fatal("sign err", err)
		}
		tx.Signature = sign

		ret = append(ret, tx)
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
		tx := getContractTransferTX(accFrom.From.Hash().Bytes()[:32], accTo.From.Hex(), contrAddr)
		tx.From = accFrom.From
		tx.Id = crypto.Keccak256Hash([]byte(fmt.Sprint(genesisId)))
		genesisId++
		sign, err := crypto.Sign(tx.Hash().Bytes(), accFrom.Private)
		if err != nil {
			log.Fatal("sign err", err)
		}
		tx.Signature = sign
		ret = append(ret, tx)
		c.fakeBalances[accFrom.From.Hex()] -= transferVal
		c.fakeBalances[accTo.From.Hex()] += transferVal
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

		from := common.BytesToAddress(crypto.Keccak256(crypto.FromECDSAPub(&privateKey.PublicKey)[1:])[12:])
		c.accounts[from] = chadAcc{From: from, Private: privateKey}
		c.orderedAccs = append(c.orderedAccs, c.accounts[from])
		prev = string(crypto.Keccak256Hash([]byte(prev)).Hex()[2:])
	}
}

func (c *Chad) GetContractDeployTX(from common.Address, contrCode []byte) *types.Transaction {
	return &types.Transaction{
		From:  from,
		To:    common.Address{},
		Input: contrCode,
	}
}

const emissionVal = 100000000

func getContractEmissionTX(wallet []byte, contrAddr common.Address) *types.Transaction {
	var arr [32]byte
	copy(arr[:], wallet)
	input, err := packTX("emission", arr, big.NewInt(emissionVal))
	if err != nil {
		log.Fatal("err generate emission input", err, input)
	}

	return &types.Transaction{
		From:  common.Address{},
		To:    contrAddr,
		Input: input,
	}
}

const transferVal = 1

func getContractTransferTX(fromWallet []byte, toWallet string, contrAddr common.Address) *types.Transaction {
	var arrFrom [32]byte
	copy(arrFrom[:], fromWallet)
	var arrTo [32]byte
	copy(arrTo[:], toWallet)
	input, err := packTX("transfer", arrFrom, arrTo, big.NewInt(transferVal))
	if err != nil {
		log.Fatal("err generate transfer input", err, input)
	}

	return &types.Transaction{
		From:  common.Address{},
		To:    contrAddr,
		Input: input,
	}
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
