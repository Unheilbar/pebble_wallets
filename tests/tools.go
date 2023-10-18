package tests

import (
	"bytes"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func verifyTx(tx *types.Transaction) (bool, error) {
	sigPublicKey, err := crypto.Ecrecover(tx.Hash().Bytes(), tx.Signature())
	if err != nil {
		return false, err
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(sigPublicKey[1:])[12:])
	return bytes.Equal(addr.Bytes(), tx.From().Bytes()), nil
}
