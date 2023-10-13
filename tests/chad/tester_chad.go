package chad

import (
	"crypto/ecdsa"
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
	from    common.Address
	private *ecdsa.PrivateKey
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
		tx := getContractEmissionTX(acc.from.Hex(), contrAddr)
		c.fakeBalances[acc.from.Hex()] += emissionVal
		tx.From = acc.from
		sign, err := crypto.Sign(tx.Hash().Bytes(), acc.private)
		if err != nil {
			log.Fatal("sign err", err)
		}
		tx.Signature = sign

		ret = append(ret, tx)
	}
	return ret
}

func (c *Chad) GenerateTransfers(size int, contrAddr common.Address) []*types.Transaction {
	var ret []*types.Transaction
	var ptrFrom int
	var ptrTo int
	for i := 0; i < size; i++ {
		accFrom := c.orderedAccs[ptrFrom%len(c.orderedAccs)]
		ptrFrom += i
		accTo := c.orderedAccs[ptrTo%len(c.orderedAccs)]
		ptrTo += i * 2
		tx := getContractTransferTX(accFrom.from.Hex(), accTo.from.Hex(), contrAddr)
		tx.From = accFrom.from
		tx.Id = accFrom.from.Hash()
		sign, err := crypto.Sign(tx.Hash().Bytes(), accFrom.private)
		if err != nil {
			log.Fatal("sign err", err)
		}
		tx.Signature = sign
		ret = append(ret, tx)
		c.fakeBalances[accFrom.from.Hex()] -= transferVal
		c.fakeBalances[accTo.from.Hex()] += transferVal
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
		c.accounts[from] = chadAcc{from: from, private: privateKey}
		c.orderedAccs = append(c.orderedAccs, c.accounts[from])
		prev = string(crypto.Keccak256Hash([]byte(prev)).Hex()[2:])
	}
}

func (c *Chad) GetContractDeployTX(senderId int, contrCode []byte) *types.Transaction {
	return &types.Transaction{
		From:  c.orderedAccs[senderId].from,
		To:    common.Address{},
		Input: contrCode,
	}
}

const emissionVal = 1000

func getContractEmissionTX(wallet string, contrAddr common.Address) *types.Transaction {
	input, err := packTX("emission", wallet, big.NewInt(emissionVal))
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

func getContractTransferTX(fromWallet string, toWallet string, contrAddr common.Address) *types.Transaction {
	input, err := packTX("transfer", fromWallet, toWallet, big.NewInt(transferVal))
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
