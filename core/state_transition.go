package core

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/Unheilbar/pebble_wallets/core/types"
	"github.com/Unheilbar/pebble_wallets/core/vm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var ErrSenderNoEOA = errors.New("sender not an eoa") // only externally owned accounts can call state transition

// ExecutionResult includes all output after executing given evm
// message no matter the execution itself is successful or not.
type ExecutionResult struct {
	Err        error  // Any error encountered during the execution(listed in core/vm/errors.go)
	ReturnData []byte // Returned data from evm(function result or data supplied with revert opcode)
}

// Unwrap returns the internal evm error which allows us for further
// analysis outside.
func (result *ExecutionResult) Unwrap() error {
	return result.Err
}

// Failed returns the indicator whether the execution is successful or not
func (result *ExecutionResult) Failed() bool { return result.Err != nil }

// Return is a helper function to help caller distinguish between revert reason
// and function return. Return returns the data after execution if no error occurs.
func (result *ExecutionResult) Return() []byte {
	if result.Err != nil {
		return nil
	}
	return common.CopyBytes(result.ReturnData)
}

// Revert returns the concrete revert reason if the execution is aborted by `REVERT`
// opcode. Note the reason can be nil if no data supplied with revert opcode.
func (result *ExecutionResult) Revert() []byte {
	if result.Err != vm.ErrExecutionReverted {
		return nil
	}
	return common.CopyBytes(result.ReturnData)
}

func ApplyMessage(evm *vm.EVM, msg *types.Transaction) (*ExecutionResult, error) {
	return NewStateTransition(evm, msg).TransitionDb()
}

type StateTransition struct {
	tx    *types.Transaction
	state vm.StateDB
	evm   *vm.EVM
}

// NewStateTransition initialises and returns a new state transition object.
func NewStateTransition(evm *vm.EVM, tx *types.Transaction) *StateTransition {
	return &StateTransition{
		tx:    tx,
		evm:   evm,
		state: evm.StateDB,
	}
}

func (st *StateTransition) TransitionDb() (*ExecutionResult, error) {

	var (
		msg              = st.tx
		sender           = vm.AccountRef(msg.From())
		rules            = st.evm.ChainConfig().Rules(st.evm.Context.BlockNumber, st.evm.Context.Random != nil, st.evm.Context.Time)
		contractCreation = msg.To() == nil
	)

	if !contractCreation {
		if err := st.preCheck(); err != nil {
			log.Fatal("signature failed")
			return nil, err
		}
	}

	//PEBBLE TODO Later add access list to Transaction
	st.state.Prepare(rules, msg.From(), st.evm.Context.Coinbase, msg.To(), vm.ActivePrecompiles(rules), nil)

	var (
		ret       []byte
		vmerr     error                   // vm errors do not effect consensus and are therefore not assigned to err
		value            = new(big.Int)   // transfer 0 always
		remainGas uint64 = math.MaxUint64 // gas never ends
	)
	if contractCreation {
		ret, _, _, vmerr = st.evm.Create(sender, msg.Data(), math.MaxUint64, value)

		contractAddr := crypto.CreateAddress(sender.Address(), 1) // hack because we form contract address not with nonce
		check := st.evm.StateDB.GetCode(contractAddr)
		if check == nil {
			panic("evm create code err")
		}

		log.Println("successfull deploy", contractAddr.Hex())
	} else {
		// Increment the nonce for the next transaction
		// st.state.SetNonce(msg.From, st.state.GetNonce(sender.Address())+1) add unique tx ID to states
		ret, _, vmerr = st.evm.Call(sender, *st.tx.To(), msg.Data(), remainGas, value)
	}

	return &ExecutionResult{
		Err:        vmerr,
		ReturnData: ret,
	}, nil
}

func (st *StateTransition) preCheck() error {
	msg := st.tx
	sigPublicKey, err := crypto.Ecrecover(st.tx.Hash().Bytes(), msg.Signature())
	if err != nil {
		return err
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(sigPublicKey[1:])[12:])
	if !bytes.Equal(addr.Bytes(), st.tx.From().Bytes()) {
		return fmt.Errorf("invalid signature")
	}

	// codeHash := st.state.GetCodeHash(msg.From)
	// if codeHash != (common.Hash{}) && codeHash != types.EmptyCodeHash {
	// 	return fmt.Errorf("%w: address %v, codehash: %s", ErrSenderNoEOA,
	// 		msg.From.Hex(), codeHash)
	// }
	// TODO check tx unique?
	return nil
}
