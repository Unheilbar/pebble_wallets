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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fromWalletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"toWalletId\",\"type\":\"bytes32\"}],\"name\":\"getAvailableTransferTypes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSender\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletState\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"setSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"balance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"monthAmount\",\"type\":\"uint128\"}],\"name\":\"setWalletState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234610059576e1000000000000000000000000000015f5260046020527fa47a91b734a2de8e1b077162bffc55ec14da8b73828e7814ac1f12dfc28187c8805460ff1916600117905561037b908161005e8239f35b5f80fdfe608060409080825260049182361015610016575f80fd5b5f91823560e01c90816303fba0ce146102c7575080630fa5acbd1461029c5780633b1ce9841461023157806371bae41d146101f7578063b23d46ab14610122578063bfb5dcf9146100f75763c052bd531461006f575f80fd5b346100f357602092836003193601126100ef57358252818352808220815191829193915b600190600282850110156100c757908260029287546001600160801b038116825260801c8982015201950192019193610093565b8285886100d382610315565b6001600160801b03818184511693015116908351928352820152f35b8280fd5b5080fd5b5082346100ef5760203660031901126100ef578160209360ff92358152600285522054169051908152f35b509190346101f45760603660031901126101f457602435926001600160801b03908185168095036100ef57604435938285168095036101f0573233146101f05781519582870187811067ffffffffffffffff8211176101dd57835286526020948587015235835282845282209390825b6001808210156101d957849085905b87600283106101b95750505086820155600101610192565b869386839498511690878960071b92831b921b19161793019501906101a1565b8480f35b634e487b7160e01b865260418352602486fd5b8380fd5b80fd5b50346100f357806003193601126100f3576102106102fb565b3233146100ef576001600160a01b0316825260016020528120602435905580f35b5082346100ef57816003193601126100ef57816020938260ff93358252600386528282205460243583528383205490845190888201928760f81b809260f81b16845260f81b1660218201526002815261028981610315565b5190208252855220541690519015158152f35b5082346100ef5760203660031901126100ef578160209360ff92358152600385522054169051908152f35b919050346100ef5760203660031901126100ef576020926001600160a01b036102ee6102fb565b1681526001845220548152f35b600435906001600160a01b038216820361031157565b5f80fd5b6040810190811067ffffffffffffffff82111761033157604052565b634e487b7160e01b5f52604160045260245ffdfea2646970667358221220bf71d2471c0daa7f869a185b647d6edf48568fd735ba66d0494f42b2cb9c692d64736f6c63430008150033",
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

// GetWalletState is a free data retrieval call binding the contract method 0xc052bd53.
//
// Solidity: function getWalletState(bytes32 walletId) view returns(uint128, uint128)
func (_State *StateCaller) GetWalletState(opts *bind.CallOpts, walletId [32]byte) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getWalletState", walletId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetWalletState is a free data retrieval call binding the contract method 0xc052bd53.
//
// Solidity: function getWalletState(bytes32 walletId) view returns(uint128, uint128)
func (_State *StateSession) GetWalletState(walletId [32]byte) (*big.Int, *big.Int, error) {
	return _State.Contract.GetWalletState(&_State.CallOpts, walletId)
}

// GetWalletState is a free data retrieval call binding the contract method 0xc052bd53.
//
// Solidity: function getWalletState(bytes32 walletId) view returns(uint128, uint128)
func (_State *StateCallerSession) GetWalletState(walletId [32]byte) (*big.Int, *big.Int, error) {
	return _State.Contract.GetWalletState(&_State.CallOpts, walletId)
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

// SetWalletState is a paid mutator transaction binding the contract method 0xb23d46ab.
//
// Solidity: function setWalletState(bytes32 walletId, uint128 balance, uint128 monthAmount) returns()
func (_State *StateTransactor) SetWalletState(opts *bind.TransactOpts, walletId [32]byte, balance *big.Int, monthAmount *big.Int) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "setWalletState", walletId, balance, monthAmount)
}

// SetWalletState is a paid mutator transaction binding the contract method 0xb23d46ab.
//
// Solidity: function setWalletState(bytes32 walletId, uint128 balance, uint128 monthAmount) returns()
func (_State *StateSession) SetWalletState(walletId [32]byte, balance *big.Int, monthAmount *big.Int) (*types.Transaction, error) {
	return _State.Contract.SetWalletState(&_State.TransactOpts, walletId, balance, monthAmount)
}

// SetWalletState is a paid mutator transaction binding the contract method 0xb23d46ab.
//
// Solidity: function setWalletState(bytes32 walletId, uint128 balance, uint128 monthAmount) returns()
func (_State *StateTransactorSession) SetWalletState(walletId [32]byte, balance *big.Int, monthAmount *big.Int) (*types.Transaction, error) {
	return _State.Contract.SetWalletState(&_State.TransactOpts, walletId, balance, monthAmount)
}
