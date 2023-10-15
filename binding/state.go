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

// StateMetaData contains all meta data concerning the State contract.
var StateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addMonth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fromWalletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"toWalletId\",\"type\":\"bytes32\"}],\"name\":\"getAvailableTransferTypes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSender\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"monthAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"setSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234610059576e1000000000000000000000000000015f5260056020527f36ba48b89d0af79da670bf69bab4afcdddb976f9908bcd5012434e39a0cf8a41805460ff1916600117905561038f908161005e8239f35b5f80fdfe6080604081815260049182361015610015575f80fd5b5f92833560e01c91826303fba0ce146102d4575081630fa5acbd146102aa5781632ee9d8c2146102585781633b1ce984146101c7578163442fb8241461019f5781635ffa4bce1461016457816371bae41d1461012957816389eba42114610102578163bfb5dcf9146100d3575063e00e3c6f14610090575f80fd5b346100cf5761009e36610322565b9190913233146100cb576100bd90838552600260205282852054610338565b918352600260205282205580f35b8380fd5b5080fd5b9050346100fe5760203660031901126100fe578160209360ff92358152600385522054169051908152f35b8280fd5b9050346100fe5760203660031901126100fe57602092829135815280845220549051908152f35b5050346100cf57806003193601126100cf57610143610308565b3233146100fe576001600160a01b0316825260016020528120602435905580f35b5050346100cf5761017436610322565b9190913233146100cb57610192908385528460205282852054610338565b9183528260205282205580f35b9050346100fe5760203660031901126100fe5760209282913581526002845220549051908152f35b8383346100cf57806003193601126100cf5782358252826020528082205491602435815281812054825193602085019160ff60f81b809260f81b16835260f81b166021850152600284528284019380851067ffffffffffffffff861117610245578484525190208152600560209081529190205460ff161515825290f35b634e487b7160e01b835260418652602483fd5b919050346100fe5761026936610322565b923233146102a6578185528460205282852054938403938411610293575083528260205282205580f35b634e487b7160e01b855260119052602484fd5b8480fd5b9050346100fe5760203660031901126100fe57816020938260ff9335825285522054169051908152f35b849084346100fe5760203660031901126100fe576020926001600160a01b036102fb610308565b1681526001845220548152f35b600435906001600160a01b038216820361031e57565b5f80fd5b604090600319011261031e576004359060243590565b9190820180921161034557565b634e487b7160e01b5f52601160045260245ffdfea2646970667358221220aa1c60e1c899378bd017fcc605312fda4cd60132bb151624703043e2a32d100e64736f6c63430008150033",
}

// StateABI is the input ABI used to generate the binding from.
// Deprecated: Use StateMetaData.ABI instead.
var StateABI = StateMetaData.ABI

// StateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateMetaData.Bin instead.
var StateBin = StateMetaData.Bin

// DeployState deploys a new Ethereum contract, binding an instance of State to it.
func DeployState(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *State, error) {
	parsed, err := StateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &State{StateCaller: StateCaller{contract: contract}, StateTransactor: StateTransactor{contract: contract}, StateFilterer: StateFilterer{contract: contract}}, nil
}

// State is an auto generated Go binding around an Ethereum contract.
type State struct {
	StateCaller     // Read-only binding to the contract
	StateTransactor // Write-only binding to the contract
	StateFilterer   // Log filterer for contract events
}

// StateCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateSession struct {
	Contract     *State            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateCallerSession struct {
	Contract *StateCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateTransactorSession struct {
	Contract     *StateTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateRaw struct {
	Contract *State // Generic contract binding to access the raw methods on
}

// StateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateCallerRaw struct {
	Contract *StateCaller // Generic read-only contract binding to access the raw methods on
}

// StateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateTransactorRaw struct {
	Contract *StateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewState creates a new instance of State, bound to a specific deployed contract.
func NewState(address common.Address, backend bind.ContractBackend) (*State, error) {
	contract, err := bindState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &State{StateCaller: StateCaller{contract: contract}, StateTransactor: StateTransactor{contract: contract}, StateFilterer: StateFilterer{contract: contract}}, nil
}

// NewStateCaller creates a new read-only instance of State, bound to a specific deployed contract.
func NewStateCaller(address common.Address, caller bind.ContractCaller) (*StateCaller, error) {
	contract, err := bindState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateCaller{contract: contract}, nil
}

// NewStateTransactor creates a new write-only instance of State, bound to a specific deployed contract.
func NewStateTransactor(address common.Address, transactor bind.ContractTransactor) (*StateTransactor, error) {
	contract, err := bindState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateTransactor{contract: contract}, nil
}

// NewStateFilterer creates a new log filterer instance of State, bound to a specific deployed contract.
func NewStateFilterer(address common.Address, filterer bind.ContractFilterer) (*StateFilterer, error) {
	contract, err := bindState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateFilterer{contract: contract}, nil
}

// bindState binds a generic wrapper to an already deployed contract.
func bindState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_State *StateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _State.Contract.StateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_State *StateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.Contract.StateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_State *StateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _State.Contract.StateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_State *StateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _State.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_State *StateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_State *StateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _State.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0x89eba421.
//
// Solidity: function balance(bytes32 walletId) view returns(uint256)
func (_State *StateCaller) Balance(opts *bind.CallOpts, walletId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "balance", walletId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x89eba421.
//
// Solidity: function balance(bytes32 walletId) view returns(uint256)
func (_State *StateSession) Balance(walletId [32]byte) (*big.Int, error) {
	return _State.Contract.Balance(&_State.CallOpts, walletId)
}

// Balance is a free data retrieval call binding the contract method 0x89eba421.
//
// Solidity: function balance(bytes32 walletId) view returns(uint256)
func (_State *StateCallerSession) Balance(walletId [32]byte) (*big.Int, error) {
	return _State.Contract.Balance(&_State.CallOpts, walletId)
}

// GetAvailableTransferTypes is a free data retrieval call binding the contract method 0x3b1ce984.
//
// Solidity: function getAvailableTransferTypes(bytes32 fromWalletId, bytes32 toWalletId) view returns(bool)
func (_State *StateCaller) GetAvailableTransferTypes(opts *bind.CallOpts, fromWalletId [32]byte, toWalletId [32]byte) (bool, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getAvailableTransferTypes", fromWalletId, toWalletId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAvailableTransferTypes is a free data retrieval call binding the contract method 0x3b1ce984.
//
// Solidity: function getAvailableTransferTypes(bytes32 fromWalletId, bytes32 toWalletId) view returns(bool)
func (_State *StateSession) GetAvailableTransferTypes(fromWalletId [32]byte, toWalletId [32]byte) (bool, error) {
	return _State.Contract.GetAvailableTransferTypes(&_State.CallOpts, fromWalletId, toWalletId)
}

// GetAvailableTransferTypes is a free data retrieval call binding the contract method 0x3b1ce984.
//
// Solidity: function getAvailableTransferTypes(bytes32 fromWalletId, bytes32 toWalletId) view returns(bool)
func (_State *StateCallerSession) GetAvailableTransferTypes(fromWalletId [32]byte, toWalletId [32]byte) (bool, error) {
	return _State.Contract.GetAvailableTransferTypes(&_State.CallOpts, fromWalletId, toWalletId)
}

// GetSender is a free data retrieval call binding the contract method 0x03fba0ce.
//
// Solidity: function getSender(address sender) view returns(bytes32)
func (_State *StateCaller) GetSender(opts *bind.CallOpts, sender common.Address) ([32]byte, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getSender", sender)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSender is a free data retrieval call binding the contract method 0x03fba0ce.
//
// Solidity: function getSender(address sender) view returns(bytes32)
func (_State *StateSession) GetSender(sender common.Address) ([32]byte, error) {
	return _State.Contract.GetSender(&_State.CallOpts, sender)
}

// GetSender is a free data retrieval call binding the contract method 0x03fba0ce.
//
// Solidity: function getSender(address sender) view returns(bytes32)
func (_State *StateCallerSession) GetSender(sender common.Address) ([32]byte, error) {
	return _State.Contract.GetSender(&_State.CallOpts, sender)
}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 walletId) view returns(uint8)
func (_State *StateCaller) GetWalletStatus(opts *bind.CallOpts, walletId [32]byte) (uint8, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getWalletStatus", walletId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 walletId) view returns(uint8)
func (_State *StateSession) GetWalletStatus(walletId [32]byte) (uint8, error) {
	return _State.Contract.GetWalletStatus(&_State.CallOpts, walletId)
}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 walletId) view returns(uint8)
func (_State *StateCallerSession) GetWalletStatus(walletId [32]byte) (uint8, error) {
	return _State.Contract.GetWalletStatus(&_State.CallOpts, walletId)
}

// GetWalletType is a free data retrieval call binding the contract method 0x0fa5acbd.
//
// Solidity: function getWalletType(bytes32 walletId) view returns(uint8)
func (_State *StateCaller) GetWalletType(opts *bind.CallOpts, walletId [32]byte) (uint8, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getWalletType", walletId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetWalletType is a free data retrieval call binding the contract method 0x0fa5acbd.
//
// Solidity: function getWalletType(bytes32 walletId) view returns(uint8)
func (_State *StateSession) GetWalletType(walletId [32]byte) (uint8, error) {
	return _State.Contract.GetWalletType(&_State.CallOpts, walletId)
}

// GetWalletType is a free data retrieval call binding the contract method 0x0fa5acbd.
//
// Solidity: function getWalletType(bytes32 walletId) view returns(uint8)
func (_State *StateCallerSession) GetWalletType(walletId [32]byte) (uint8, error) {
	return _State.Contract.GetWalletType(&_State.CallOpts, walletId)
}

// MonthAmount is a free data retrieval call binding the contract method 0x442fb824.
//
// Solidity: function monthAmount(bytes32 walletId) view returns(uint256)
func (_State *StateCaller) MonthAmount(opts *bind.CallOpts, walletId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "monthAmount", walletId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MonthAmount is a free data retrieval call binding the contract method 0x442fb824.
//
// Solidity: function monthAmount(bytes32 walletId) view returns(uint256)
func (_State *StateSession) MonthAmount(walletId [32]byte) (*big.Int, error) {
	return _State.Contract.MonthAmount(&_State.CallOpts, walletId)
}

// MonthAmount is a free data retrieval call binding the contract method 0x442fb824.
//
// Solidity: function monthAmount(bytes32 walletId) view returns(uint256)
func (_State *StateCallerSession) MonthAmount(walletId [32]byte) (*big.Int, error) {
	return _State.Contract.MonthAmount(&_State.CallOpts, walletId)
}

// Add is a paid mutator transaction binding the contract method 0x5ffa4bce.
//
// Solidity: function add(bytes32 walletId, uint256 amount) returns()
func (_State *StateTransactor) Add(opts *bind.TransactOpts, walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "add", walletId, amount)
}

// Add is a paid mutator transaction binding the contract method 0x5ffa4bce.
//
// Solidity: function add(bytes32 walletId, uint256 amount) returns()
func (_State *StateSession) Add(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.Add(&_State.TransactOpts, walletId, amount)
}

// Add is a paid mutator transaction binding the contract method 0x5ffa4bce.
//
// Solidity: function add(bytes32 walletId, uint256 amount) returns()
func (_State *StateTransactorSession) Add(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.Add(&_State.TransactOpts, walletId, amount)
}

// AddMonth is a paid mutator transaction binding the contract method 0xe00e3c6f.
//
// Solidity: function addMonth(bytes32 walletId, uint256 amount) returns()
func (_State *StateTransactor) AddMonth(opts *bind.TransactOpts, walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "addMonth", walletId, amount)
}

// AddMonth is a paid mutator transaction binding the contract method 0xe00e3c6f.
//
// Solidity: function addMonth(bytes32 walletId, uint256 amount) returns()
func (_State *StateSession) AddMonth(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.AddMonth(&_State.TransactOpts, walletId, amount)
}

// AddMonth is a paid mutator transaction binding the contract method 0xe00e3c6f.
//
// Solidity: function addMonth(bytes32 walletId, uint256 amount) returns()
func (_State *StateTransactorSession) AddMonth(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.AddMonth(&_State.TransactOpts, walletId, amount)
}

// SetSender is a paid mutator transaction binding the contract method 0x71bae41d.
//
// Solidity: function setSender(address sender, bytes32 walletId) returns()
func (_State *StateTransactor) SetSender(opts *bind.TransactOpts, sender common.Address, walletId [32]byte) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "setSender", sender, walletId)
}

// SetSender is a paid mutator transaction binding the contract method 0x71bae41d.
//
// Solidity: function setSender(address sender, bytes32 walletId) returns()
func (_State *StateSession) SetSender(sender common.Address, walletId [32]byte) (*types.Transaction, error) {
	return _State.Contract.SetSender(&_State.TransactOpts, sender, walletId)
}

// SetSender is a paid mutator transaction binding the contract method 0x71bae41d.
//
// Solidity: function setSender(address sender, bytes32 walletId) returns()
func (_State *StateTransactorSession) SetSender(sender common.Address, walletId [32]byte) (*types.Transaction, error) {
	return _State.Contract.SetSender(&_State.TransactOpts, sender, walletId)
}

// Sub is a paid mutator transaction binding the contract method 0x2ee9d8c2.
//
// Solidity: function sub(bytes32 walletId, uint256 amount) returns()
func (_State *StateTransactor) Sub(opts *bind.TransactOpts, walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "sub", walletId, amount)
}

// Sub is a paid mutator transaction binding the contract method 0x2ee9d8c2.
//
// Solidity: function sub(bytes32 walletId, uint256 amount) returns()
func (_State *StateSession) Sub(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.Sub(&_State.TransactOpts, walletId, amount)
}

// Sub is a paid mutator transaction binding the contract method 0x2ee9d8c2.
//
// Solidity: function sub(bytes32 walletId, uint256 amount) returns()
func (_State *StateTransactorSession) Sub(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.Sub(&_State.TransactOpts, walletId, amount)
}
