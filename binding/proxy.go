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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"walletId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fromWalletId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toWalletId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234610057577398acac3b9c77c934c12780a2852a959e674970a360018060a01b0319731aea632c29d2978a5c6336a3b8bfe9d737eb8fe38160015416176001555f5416175f5561031f908161005c8239f35b5f80fdfe60406080815260049081361015610014575f80fd5b5f91823560e01c80634e3042441461011d57639b80b05014610034575f80fd5b829134610119576003196060368201126101145767ffffffffffffffff928035848111610110576100689036908301610245565b93602435908111610110576100809036908301610245565b9360018060a01b0386541690813b1561010c57866100c6936100d582968851998a97889687956341b3856960e01b875233908701526080602487015260848601906102ab565b918483030160448501526102ab565b604435606483015203925af190811561010357506100f05750f35b6100f99061021d565b6101005780f35b80fd5b513d84823e3d90fd5b8680fd5b8580fd5b505050fd5b5050fd5b50919034610219578060031936011261021957823567ffffffffffffffff81116102155761014e9036908501610245565b6001546001600160a01b039081169491929190853b156101f35784835180976336555b8560e01b8252858583015281838161018c604482018b6102ab565b602435602483015203925af1801561020b576101f7575b8495506001541692833b156101f3576101e0938592838551809781958294633b8c713360e21b8452339084015288602484015260448301906102ab565b03925af190811561010357506100f05750f35b8480fd5b9390946102039061021d565b9284906101a3565b83513d87823e3d90fd5b8280fd5b5080fd5b67ffffffffffffffff811161023157604052565b634e487b7160e01b5f52604160045260245ffd5b81601f820112156102a75780359067ffffffffffffffff928383116102315760405193601f8401601f19908116603f011685019081118582101761023157604052828452602083830101116102a757815f926020809301838601378301015290565b5f80fd5b91908251928382525f5b8481106102d5575050825f602080949584010152601f8019910116010190565b6020818301810151848301820152016102b556fea2646970667358221220209eb2c923c17223a5a5d91067e88a014dbb2bfb04b207dee247490d23e3644364736f6c63430008150033",
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

// Emission is a paid mutator transaction binding the contract method 0x4e304244.
//
// Solidity: function emission(string walletId, uint256 amount) returns()
func (_Proxy *ProxyTransactor) Emission(opts *bind.TransactOpts, walletId string, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "emission", walletId, amount)
}

// Emission is a paid mutator transaction binding the contract method 0x4e304244.
//
// Solidity: function emission(string walletId, uint256 amount) returns()
func (_Proxy *ProxySession) Emission(walletId string, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Emission(&_Proxy.TransactOpts, walletId, amount)
}

// Emission is a paid mutator transaction binding the contract method 0x4e304244.
//
// Solidity: function emission(string walletId, uint256 amount) returns()
func (_Proxy *ProxyTransactorSession) Emission(walletId string, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Emission(&_Proxy.TransactOpts, walletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string fromWalletId, string toWalletId, uint256 amount) returns()
func (_Proxy *ProxyTransactor) Transfer(opts *bind.TransactOpts, fromWalletId string, toWalletId string, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "transfer", fromWalletId, toWalletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string fromWalletId, string toWalletId, uint256 amount) returns()
func (_Proxy *ProxySession) Transfer(fromWalletId string, toWalletId string, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Transfer(&_Proxy.TransactOpts, fromWalletId, toWalletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string fromWalletId, string toWalletId, uint256 amount) returns()
func (_Proxy *ProxyTransactorSession) Transfer(fromWalletId string, toWalletId string, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Transfer(&_Proxy.TransactOpts, fromWalletId, toWalletId, amount)
}
