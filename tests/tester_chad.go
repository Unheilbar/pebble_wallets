package tests

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
	accounts map[common.Address]chadAcc
}

const genesisPrivate = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"

func (c *Chad) generateAccEmissionsTx(contrAddr common.Address) []*types.Transaction {
	var ret []*types.Transaction
	for _, acc := range c.accounts {
		ret = append(ret, getContractEmissionTX(acc.from.Hex(), contrAddr))
	}
	return ret
}

func (c *Chad) generateAccs(size int) {
	c.accounts = make(map[common.Address]chadAcc)
	prev := genesisPrivate
	for i := 0; i < size; i++ {
		privateKey, err := crypto.HexToECDSA(prev)
		prev = string(crypto.Keccak256Hash([]byte(prev)).Hex()[2:])
		if err != nil {
			log.Fatal("err gen private", err)
		}

		from := common.BytesToAddress(crypto.Keccak256(crypto.FromECDSAPub(&privateKey.PublicKey)[1:])[12:])
		c.accounts[from] = chadAcc{from: from, private: privateKey}
	}
}

func getContractDeployTX(contrCode []byte) *types.Transaction {
	return &types.Transaction{
		From:  common.Address{},
		To:    common.Address{},
		Input: contrCode,
		Value: new(big.Int),
	}
}

func getContractEmissionTX(wallet string, contrAddr common.Address) *types.Transaction {
	input, err := packTX("setBalance", wallet, big.NewInt(int64(10000000)))
	if err != nil {
		log.Fatal("err generate emission input", err, input)
	}

	return &types.Transaction{
		From:  common.Address{},
		To:    contrAddr,
		Input: input,
		Value: new(big.Int),
	}
}
func packTX(method string, params ...interface{}) ([]byte, error) {
	bind := binding.StorageABI

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
