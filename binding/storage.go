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

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setBalanceA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50620f4240600181905550610d05806100275f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80632e64cec11461006457806338639f7b146100825780633a51d2461461009e5780636057361d146100ce5780639b80b050146100ea578063a5e2d81014610106575b5f80fd5b61006c610122565b6040516100799190610496565b60405180910390f35b61009c60048036038101906100979190610626565b61012a565b005b6100b860048036038101906100b39190610680565b61019c565b6040516100c59190610496565b60405180910390f35b6100e860048036038101906100e391906106c7565b6101c3565b005b61010460048036038101906100ff91906106f2565b6101cc565b005b610120600480360381019061011b91906107d4565b610438565b005b5f8054905090565b8160075f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090816101749190610a16565b50806004836040516101869190610b47565b9081526020016040518091039020819055505050565b5f6004826040516101ad9190610b47565b9081526020016040518091039020549050919050565b805f8190555050565b5f816004856040516101de9190610b47565b9081526020016040518091039020546101f79190610b8a565b1015610201575f80fd5b600154816005856040516102159190610b47565b90815260200160405180910390205461022e9190610bbd565b10610237575f80fd5b60028081111561024a57610249610bf0565b5b60068460405161025a9190610b47565b90815260200160405180910390205f9054906101000a900460ff16600281111561028757610286610bf0565b5b141580156102e45750600160028111156102a4576102a3610bf0565b5b6006836040516102b49190610b47565b90815260200160405180910390205f9054906101000a900460ff1660028111156102e1576102e0610bf0565b5b14155b6102ec575f80fd5b828051906020012060075f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060405161033e9190610cb9565b60405180910390201461034f575f80fd5b806004846040516103609190610b47565b9081526020016040518091039020546103799190610b8a565b6004846040516103899190610b47565b908152602001604051809103902081905550806004836040516103ac9190610b47565b9081526020016040518091039020546103c59190610bbd565b6004836040516103d59190610b47565b908152602001604051809103902081905550806005846040516103f89190610b47565b9081526020016040518091039020546104119190610bbd565b6005846040516104219190610b47565b908152602001604051809103902081905550505050565b8060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505050565b5f819050919050565b6104908161047e565b82525050565b5f6020820190506104a95f830184610487565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61050e826104c8565b810181811067ffffffffffffffff8211171561052d5761052c6104d8565b5b80604052505050565b5f61053f6104af565b905061054b8282610505565b919050565b5f67ffffffffffffffff82111561056a576105696104d8565b5b610573826104c8565b9050602081019050919050565b828183375f83830152505050565b5f6105a061059b84610550565b610536565b9050828152602081018484840111156105bc576105bb6104c4565b5b6105c7848285610580565b509392505050565b5f82601f8301126105e3576105e26104c0565b5b81356105f384826020860161058e565b91505092915050565b6106058161047e565b811461060f575f80fd5b50565b5f81359050610620816105fc565b92915050565b5f806040838503121561063c5761063b6104b8565b5b5f83013567ffffffffffffffff811115610659576106586104bc565b5b610665858286016105cf565b925050602061067685828601610612565b9150509250929050565b5f60208284031215610695576106946104b8565b5b5f82013567ffffffffffffffff8111156106b2576106b16104bc565b5b6106be848285016105cf565b91505092915050565b5f602082840312156106dc576106db6104b8565b5b5f6106e984828501610612565b91505092915050565b5f805f60608486031215610709576107086104b8565b5b5f84013567ffffffffffffffff811115610726576107256104bc565b5b610732868287016105cf565b935050602084013567ffffffffffffffff811115610753576107526104bc565b5b61075f868287016105cf565b925050604061077086828701610612565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6107a38261077a565b9050919050565b6107b381610799565b81146107bd575f80fd5b50565b5f813590506107ce816107aa565b92915050565b5f80604083850312156107ea576107e96104b8565b5b5f6107f7858286016107c0565b925050602061080885828601610612565b9150509250929050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061086057607f821691505b6020821081036108735761087261081c565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026108d57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261089a565b6108df868361089a565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61091a6109156109108461047e565b6108f7565b61047e565b9050919050565b5f819050919050565b61093383610900565b61094761093f82610921565b8484546108a6565b825550505050565b5f90565b61095b61094f565b61096681848461092a565b505050565b5b818110156109895761097e5f82610953565b60018101905061096c565b5050565b601f8211156109ce5761099f81610879565b6109a88461088b565b810160208510156109b7578190505b6109cb6109c38561088b565b83018261096b565b50505b505050565b5f82821c905092915050565b5f6109ee5f19846008026109d3565b1980831691505092915050565b5f610a0683836109df565b9150826002028217905092915050565b610a1f82610812565b67ffffffffffffffff811115610a3857610a376104d8565b5b610a428254610849565b610a4d82828561098d565b5f60209050601f831160018114610a7e575f8415610a6c578287015190505b610a7685826109fb565b865550610add565b601f198416610a8c86610879565b5f5b82811015610ab357848901518255600182019150602085019450602081019050610a8e565b86831015610ad05784890151610acc601f8916826109df565b8355505b6001600288020188555050505b505050505050565b5f81905092915050565b5f5b83811015610b0c578082015181840152602081019050610af1565b5f8484015250505050565b5f610b2182610812565b610b2b8185610ae5565b9350610b3b818560208601610aef565b80840191505092915050565b5f610b528284610b17565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610b948261047e565b9150610b9f8361047e565b9250828203905081811115610bb757610bb6610b5d565b5b92915050565b5f610bc78261047e565b9150610bd28361047e565b9250828201905080821115610bea57610be9610b5d565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f81905092915050565b5f819050815f5260205f209050919050565b5f8154610c4581610849565b610c4f8186610c1d565b9450600182165f8114610c695760018114610c7e57610cb0565b60ff1983168652811515820286019350610cb0565b610c8785610c27565b5f5b83811015610ca857815481890152600182019150602081019050610c89565b838801955050505b50505092915050565b5f610cc48284610c39565b91508190509291505056fea264697066735822122015ee5366b646623a2ba5461816c2a424ee8a3ea21e33933dd66d69069c14e0c164736f6c63430008150033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string s) view returns(uint256)
func (_Storage *StorageCaller) GetBalance(opts *bind.CallOpts, s string) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getBalance", s)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string s) view returns(uint256)
func (_Storage *StorageSession) GetBalance(s string) (*big.Int, error) {
	return _Storage.Contract.GetBalance(&_Storage.CallOpts, s)
}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string s) view returns(uint256)
func (_Storage *StorageCallerSession) GetBalance(s string) (*big.Int, error) {
	return _Storage.Contract.GetBalance(&_Storage.CallOpts, s)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Storage *StorageCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Storage *StorageSession) Retrieve() (*big.Int, error) {
	return _Storage.Contract.Retrieve(&_Storage.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Storage *StorageCallerSession) Retrieve() (*big.Int, error) {
	return _Storage.Contract.Retrieve(&_Storage.CallOpts)
}

// SetBalance is a paid mutator transaction binding the contract method 0x38639f7b.
//
// Solidity: function setBalance(string s, uint256 n) returns()
func (_Storage *StorageTransactor) SetBalance(opts *bind.TransactOpts, s string, n *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setBalance", s, n)
}

// SetBalance is a paid mutator transaction binding the contract method 0x38639f7b.
//
// Solidity: function setBalance(string s, uint256 n) returns()
func (_Storage *StorageSession) SetBalance(s string, n *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetBalance(&_Storage.TransactOpts, s, n)
}

// SetBalance is a paid mutator transaction binding the contract method 0x38639f7b.
//
// Solidity: function setBalance(string s, uint256 n) returns()
func (_Storage *StorageTransactorSession) SetBalance(s string, n *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetBalance(&_Storage.TransactOpts, s, n)
}

// SetBalanceA is a paid mutator transaction binding the contract method 0xa5e2d810.
//
// Solidity: function setBalanceA(address a, uint256 n) returns()
func (_Storage *StorageTransactor) SetBalanceA(opts *bind.TransactOpts, a common.Address, n *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setBalanceA", a, n)
}

// SetBalanceA is a paid mutator transaction binding the contract method 0xa5e2d810.
//
// Solidity: function setBalanceA(address a, uint256 n) returns()
func (_Storage *StorageSession) SetBalanceA(a common.Address, n *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetBalanceA(&_Storage.TransactOpts, a, n)
}

// SetBalanceA is a paid mutator transaction binding the contract method 0xa5e2d810.
//
// Solidity: function setBalanceA(address a, uint256 n) returns()
func (_Storage *StorageTransactorSession) SetBalanceA(a common.Address, n *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetBalanceA(&_Storage.TransactOpts, a, n)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_Storage *StorageTransactor) Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "store", num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_Storage *StorageSession) Store(num *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Store(&_Storage.TransactOpts, num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_Storage *StorageTransactorSession) Store(num *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Store(&_Storage.TransactOpts, num)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string from, string to, uint256 amount) returns()
func (_Storage *StorageTransactor) Transfer(opts *bind.TransactOpts, from string, to string, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transfer", from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string from, string to, uint256 amount) returns()
func (_Storage *StorageSession) Transfer(from string, to string, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Transfer(&_Storage.TransactOpts, from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string from, string to, uint256 amount) returns()
func (_Storage *StorageTransactorSession) Transfer(from string, to string, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Transfer(&_Storage.TransactOpts, from, to, amount)
}
