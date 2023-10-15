package core

import (
	"math/big"

	"github.com/Unheilbar/pebbke_wallets/core/vm"
	"github.com/ethereum/go-ethereum/common"
)

// CanTransfer checks whether there are enough funds in the address' account to make a transfer.
// This does not take the necessary gas in to account to make the transfer valid.
func CanTransfer(db vm.StateDB, addr common.Address, amount *big.Int) bool {
	return true // We always can transfer, cause we don't send anything
}

// Transfer subtracts amount from sender and adds amount to recipient using the given Db
func Transfer(db vm.StateDB, sender, recipient common.Address, amount *big.Int) {
	return // Mock since we never transfer any ethers
}
