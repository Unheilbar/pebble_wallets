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

// ProxyMetaData contains all meta data concerning the Proxy contract.
var ProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fromWalletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"toWalletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234610057577398acac3b9c77c934c12780a2852a959e674970a360018060a01b0319731aea632c29d2978a5c6336a3b8bfe9d737eb8fe38160015416176001555f5416175f556101e3908161005c8239f35b5f80fdfe60806040818152600480361015610014575f80fd5b5f92833560e01c9081636ff6cdad146100c4575063ae9f75e314610036575f80fd5b8291346100c05760603660031901126100c05782546001600160a01b031691823b156100bb57839260848492845195869384926394bf8d6160e01b84523381850152356024840152602435604484015260443560648401525af19081156100b2575061009f5750f35b6100a890610185565b6100af5780f35b80fd5b513d84823e3d90fd5b505050fd5b5050fd5b84838534610181578060031936011261018157600154823592906001600160a01b03908116803b1561017d5786604481888094632ffd25e760e11b8352898884015260243560248401525af180156101735761015f575b8495506001541692833b1561015b576044859283855196879485936371bae41d60e01b8552339085015260248401525af19081156100b2575061009f5750f35b8480fd5b93909461016b90610185565b92849061011b565b83513d87823e3d90fd5b8580fd5b8280fd5b67ffffffffffffffff811161019957604052565b634e487b7160e01b5f52604160045260245ffdfea2646970667358221220df5d5e7a41a7d7c3be71b1c1090a18aa3feb4b80ab62d7bd619be0b150a3ebcf64736f6c63430008150033",
}

// ProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyMetaData.ABI instead.
var ProxyABI = ProxyMetaData.ABI

// ProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProxyMetaData.Bin instead.
var ProxyBin = ProxyMetaData.Bin

// DeployProxy deploys a new Ethereum contract, binding an instance of Proxy to it.
func DeployProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Proxy, error) {
	parsed, err := ProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// Emission is a paid mutator transaction binding the contract method 0x6ff6cdad.
//
// Solidity: function emission(bytes32 walletId, uint256 amount) returns()
func (_Proxy *ProxyTransactor) Emission(opts *bind.TransactOpts, walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "emission", walletId, amount)
}

// Emission is a paid mutator transaction binding the contract method 0x6ff6cdad.
//
// Solidity: function emission(bytes32 walletId, uint256 amount) returns()
func (_Proxy *ProxySession) Emission(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Emission(&_Proxy.TransactOpts, walletId, amount)
}

// Emission is a paid mutator transaction binding the contract method 0x6ff6cdad.
//
// Solidity: function emission(bytes32 walletId, uint256 amount) returns()
func (_Proxy *ProxyTransactorSession) Emission(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Emission(&_Proxy.TransactOpts, walletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xae9f75e3.
//
// Solidity: function transfer(bytes32 fromWalletId, bytes32 toWalletId, uint256 amount) returns()
func (_Proxy *ProxyTransactor) Transfer(opts *bind.TransactOpts, fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "transfer", fromWalletId, toWalletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xae9f75e3.
//
// Solidity: function transfer(bytes32 fromWalletId, bytes32 toWalletId, uint256 amount) returns()
func (_Proxy *ProxySession) Transfer(fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Transfer(&_Proxy.TransactOpts, fromWalletId, toWalletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xae9f75e3.
//
// Solidity: function transfer(bytes32 fromWalletId, bytes32 toWalletId, uint256 amount) returns()
func (_Proxy *ProxyTransactorSession) Transfer(fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Transfer(&_Proxy.TransactOpts, fromWalletId, toWalletId, amount)
}
