package tests

import (
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/Unheilbar/pebbke_wallets/core/types"
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

func (c *Chad) getContractDeployTX(contrCode []byte) *types.Transaction {
	return &types.Transaction{
		From:  common.Address{},
		To:    common.Address{},
		Input: contrCode,
		Value: new(big.Int),
	}
}
