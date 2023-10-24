// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LsmMetaData contains all meta data concerning the Lsm contract.
var LsmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"emission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657610185908161001b8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f803560e01c80636a240f12146100c85763e60900d814610030575f80fd5b346100c55760603660031901126100c5576001600160801b036044358181168082036100c1576004358452836020526040842090815490848216038481116100ad57908492916001600160801b031993849116911617905560243584528360205260408420926100a4845493828516610120565b16911617905580f35b634e487b7160e01b86526011600452602486fd5b8380fd5b80fd5b50346100c55760403660031901126100c5576001600160801b03602435818116810361011c57600435835282602052604083209161010a835492828416610120565b16906001600160801b03191617905580f35b8280fd5b9190916001600160801b038080941691160191821161013b57565b634e487b7160e01b5f52601160045260245ffdfea26469706673582212205f90bbd6c124f2ece3c880932398e29b1806a695da37098f4c24224ccd3bb09c64736f6c63430008150033",
}

// LsmABI is the input ABI used to generate the binding from.
// Deprecated: Use LsmMetaData.ABI instead.
var LsmABI = LsmMetaData.ABI

// LsmBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LsmMetaData.Bin instead.
var LsmBin = LsmMetaData.Bin

// DeployLsm deploys a new Ethereum contract, binding an instance of Lsm to it.
func DeployLsm(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lsm, error) {
	parsed, err := LsmMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LsmBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lsm{LsmCaller: LsmCaller{contract: contract}, LsmTransactor: LsmTransactor{contract: contract}, LsmFilterer: LsmFilterer{contract: contract}}, nil
}

// Lsm is an auto generated Go binding around an Ethereum contract.
type Lsm struct {
	LsmCaller     // Read-only binding to the contract
	LsmTransactor // Write-only binding to the contract
	LsmFilterer   // Log filterer for contract events
}

// LsmCaller is an auto generated read-only Go binding around an Ethereum contract.
type LsmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LsmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LsmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LsmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LsmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LsmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LsmSession struct {
	Contract     *Lsm              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LsmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LsmCallerSession struct {
	Contract *LsmCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LsmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LsmTransactorSession struct {
	Contract     *LsmTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LsmRaw is an auto generated low-level Go binding around an Ethereum contract.
type LsmRaw struct {
	Contract *Lsm // Generic contract binding to access the raw methods on
}

// LsmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LsmCallerRaw struct {
	Contract *LsmCaller // Generic read-only contract binding to access the raw methods on
}

// LsmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LsmTransactorRaw struct {
	Contract *LsmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLsm creates a new instance of Lsm, bound to a specific deployed contract.
func NewLsm(address common.Address, backend bind.ContractBackend) (*Lsm, error) {
	contract, err := bindLsm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lsm{LsmCaller: LsmCaller{contract: contract}, LsmTransactor: LsmTransactor{contract: contract}, LsmFilterer: LsmFilterer{contract: contract}}, nil
}

// NewLsmCaller creates a new read-only instance of Lsm, bound to a specific deployed contract.
func NewLsmCaller(address common.Address, caller bind.ContractCaller) (*LsmCaller, error) {
	contract, err := bindLsm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LsmCaller{contract: contract}, nil
}

// NewLsmTransactor creates a new write-only instance of Lsm, bound to a specific deployed contract.
func NewLsmTransactor(address common.Address, transactor bind.ContractTransactor) (*LsmTransactor, error) {
	contract, err := bindLsm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LsmTransactor{contract: contract}, nil
}

// NewLsmFilterer creates a new log filterer instance of Lsm, bound to a specific deployed contract.
func NewLsmFilterer(address common.Address, filterer bind.ContractFilterer) (*LsmFilterer, error) {
	contract, err := bindLsm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LsmFilterer{contract: contract}, nil
}

// bindLsm binds a generic wrapper to an already deployed contract.
func bindLsm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LsmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lsm *LsmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lsm.Contract.LsmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lsm *LsmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lsm.Contract.LsmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lsm *LsmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lsm.Contract.LsmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lsm *LsmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lsm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lsm *LsmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lsm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lsm *LsmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lsm.Contract.contract.Transact(opts, method, params...)
}

// Emission is a paid mutator transaction binding the contract method 0x6a240f12.
//
// Solidity: function emission(bytes32 id, uint128 amount) returns()
func (_Lsm *LsmTransactor) Emission(opts *bind.TransactOpts, id [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Lsm.contract.Transact(opts, "emission", id, amount)
}

// Emission is a paid mutator transaction binding the contract method 0x6a240f12.
//
// Solidity: function emission(bytes32 id, uint128 amount) returns()
func (_Lsm *LsmSession) Emission(id [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Lsm.Contract.Emission(&_Lsm.TransactOpts, id, amount)
}

// Emission is a paid mutator transaction binding the contract method 0x6a240f12.
//
// Solidity: function emission(bytes32 id, uint128 amount) returns()
func (_Lsm *LsmTransactorSession) Emission(id [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Lsm.Contract.Emission(&_Lsm.TransactOpts, id, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xe60900d8.
//
// Solidity: function transfer(bytes32 from, bytes32 to, uint128 amount) returns()
func (_Lsm *LsmTransactor) Transfer(opts *bind.TransactOpts, from [32]byte, to [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Lsm.contract.Transact(opts, "transfer", from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xe60900d8.
//
// Solidity: function transfer(bytes32 from, bytes32 to, uint128 amount) returns()
func (_Lsm *LsmSession) Transfer(from [32]byte, to [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Lsm.Contract.Transfer(&_Lsm.TransactOpts, from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xe60900d8.
//
// Solidity: function transfer(bytes32 from, bytes32 to, uint128 amount) returns()
func (_Lsm *LsmTransactorSession) Transfer(from [32]byte, to [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Lsm.Contract.Transfer(&_Lsm.TransactOpts, from, to, amount)
}
