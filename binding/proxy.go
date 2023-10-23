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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"emission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fromWalletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"toWalletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461005757735dbc355b93dd7a0c0d759fd6a7859d261021922160018060a01b0319736027946b05e7ab6ef245093622ab18ead54538778160015416176001555f5416175f5561033e908161005c8239f35b5f80fdfe60806040818152600480361015610014575f80fd5b5f92833560e01c9081638e739461146101cb57508063b2923fe1146100dd5763e60900d814610041575f80fd5b8291346100d95760603660031901126100d95761005c6102b2565b83546001600160a01b03169290833b156100d55760848592836001600160801b0396865197889586946324d9323b60e21b8652338187015235602486015260243560448601521660648401525af19081156100cc57506100b95750f35b6100c2906102cc565b6100c95780f35b80fd5b513d84823e3d90fd5b8480fd5b5050fd5b509190346101c75760603660031901126101c7576001600160a01b039280358481169291908390036101c357602435926101156102b2565b95806001541696873b156101bf576001600160801b03606488928388519b8c94859363b23d46ab60e01b85528c8b8601521660248401528160448401525af180156101b5576101a1575b85965060015416803b1561019d57859283604492865197889586946371bae41d60e01b865285015260248401525af19081156100cc57506100b95750f35b8580fd5b9490956101ad906102cc565b93859061015f565b84513d88823e3d90fd5b8680fd5b8380fd5b5080fd5b91939050346100c95760203660031901126100c95760015463c052bd5360e01b83528435858401528390839060249082906001600160a01b03165afa9384156102a6578194610228575b6020856001600160801b03865191168152f35b9091935082903d841161029e575b601f8201601f191685019067ffffffffffffffff82118683101761028b5750839185918352810103126100c957506001600160801b0382610283602061027c81966102f4565b92016102f4565b50925f610215565b634e487b7160e01b845260419052602483fd5b3d9150610236565b508251903d90823e3d90fd5b604435906001600160801b03821682036102c857565b5f80fd5b67ffffffffffffffff81116102e057604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160801b03821682036102c85756fea26469706673582212203f1b81bb954eca3794e59ae0ed96c479084cf431e851b47398d879ae7ed045d264736f6c63430008150033",
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

// GetBalance is a free data retrieval call binding the contract method 0x8e739461.
//
// Solidity: function getBalance(bytes32 walletId) view returns(uint128)
func (_Proxy *ProxyCaller) GetBalance(opts *bind.CallOpts, walletId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getBalance", walletId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x8e739461.
//
// Solidity: function getBalance(bytes32 walletId) view returns(uint128)
func (_Proxy *ProxySession) GetBalance(walletId [32]byte) (*big.Int, error) {
	return _Proxy.Contract.GetBalance(&_Proxy.CallOpts, walletId)
}

// GetBalance is a free data retrieval call binding the contract method 0x8e739461.
//
// Solidity: function getBalance(bytes32 walletId) view returns(uint128)
func (_Proxy *ProxyCallerSession) GetBalance(walletId [32]byte) (*big.Int, error) {
	return _Proxy.Contract.GetBalance(&_Proxy.CallOpts, walletId)
}

// Emission is a paid mutator transaction binding the contract method 0xb2923fe1.
//
// Solidity: function emission(address senderId, bytes32 walletId, uint128 amount) returns()
func (_Proxy *ProxyTransactor) Emission(opts *bind.TransactOpts, senderId common.Address, walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "emission", senderId, walletId, amount)
}

// Emission is a paid mutator transaction binding the contract method 0xb2923fe1.
//
// Solidity: function emission(address senderId, bytes32 walletId, uint128 amount) returns()
func (_Proxy *ProxySession) Emission(senderId common.Address, walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Emission(&_Proxy.TransactOpts, senderId, walletId, amount)
}

// Emission is a paid mutator transaction binding the contract method 0xb2923fe1.
//
// Solidity: function emission(address senderId, bytes32 walletId, uint128 amount) returns()
func (_Proxy *ProxyTransactorSession) Emission(senderId common.Address, walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Emission(&_Proxy.TransactOpts, senderId, walletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xe60900d8.
//
// Solidity: function transfer(bytes32 fromWalletId, bytes32 toWalletId, uint128 amount) returns()
func (_Proxy *ProxyTransactor) Transfer(opts *bind.TransactOpts, fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "transfer", fromWalletId, toWalletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xe60900d8.
//
// Solidity: function transfer(bytes32 fromWalletId, bytes32 toWalletId, uint128 amount) returns()
func (_Proxy *ProxySession) Transfer(fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Transfer(&_Proxy.TransactOpts, fromWalletId, toWalletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xe60900d8.
//
// Solidity: function transfer(bytes32 fromWalletId, bytes32 toWalletId, uint128 amount) returns()
func (_Proxy *ProxyTransactorSession) Transfer(fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Transfer(&_Proxy.TransactOpts, fromWalletId, toWalletId, amount)
}
