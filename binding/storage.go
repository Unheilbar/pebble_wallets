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
	Bin: "0x608060405234801561000f575f80fd5b50620f424060018190555061090b806100275f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80632e64cec11461006457806338639f7b146100825780633a51d2461461009e5780636057361d146100ce5780639b80b050146100ea578063a5e2d81014610106575b5f80fd5b61006c610122565b6040516100799190610417565b60405180910390f35b61009c600480360381019061009791906105a7565b61012a565b005b6100b860048036038101906100b39190610601565b61017a565b6040516100c59190610417565b60405180910390f35b6100e860048036038101906100e39190610648565b6101a1565b005b61010460048036038101906100ff9190610673565b6101aa565b005b610120600480360381019061011b9190610755565b6103b9565b005b5f8054905090565b8060048360405161013b91906107ff565b9081526020016040518091039020546101549190610842565b60048360405161016491906107ff565b9081526020016040518091039020819055505050565b5f60048260405161018b91906107ff565b9081526020016040518091039020549050919050565b805f8190555050565b5f816004856040516101bc91906107ff565b9081526020016040518091039020546101d59190610875565b10156101df575f80fd5b600154816005856040516101f391906107ff565b90815260200160405180910390205461020c9190610842565b10610215575f80fd5b600280811115610228576102276108a8565b5b60068460405161023891906107ff565b90815260200160405180910390205f9054906101000a900460ff166002811115610265576102646108a8565b5b141580156102c2575060016002811115610282576102816108a8565b5b60068360405161029291906107ff565b90815260200160405180910390205f9054906101000a900460ff1660028111156102bf576102be6108a8565b5b14155b6102ca575f80fd5b806004846040516102db91906107ff565b9081526020016040518091039020546102f49190610875565b60048460405161030491906107ff565b9081526020016040518091039020819055508060048360405161032791906107ff565b9081526020016040518091039020546103409190610842565b60048360405161035091906107ff565b9081526020016040518091039020819055508060058460405161037391906107ff565b90815260200160405180910390205461038c9190610842565b60058460405161039c91906107ff565b908152602001604051809103902081905550805f81905550505050565b8060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505050565b5f819050919050565b610411816103ff565b82525050565b5f60208201905061042a5f830184610408565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61048f82610449565b810181811067ffffffffffffffff821117156104ae576104ad610459565b5b80604052505050565b5f6104c0610430565b90506104cc8282610486565b919050565b5f67ffffffffffffffff8211156104eb576104ea610459565b5b6104f482610449565b9050602081019050919050565b828183375f83830152505050565b5f61052161051c846104d1565b6104b7565b90508281526020810184848401111561053d5761053c610445565b5b610548848285610501565b509392505050565b5f82601f83011261056457610563610441565b5b813561057484826020860161050f565b91505092915050565b610586816103ff565b8114610590575f80fd5b50565b5f813590506105a18161057d565b92915050565b5f80604083850312156105bd576105bc610439565b5b5f83013567ffffffffffffffff8111156105da576105d961043d565b5b6105e685828601610550565b92505060206105f785828601610593565b9150509250929050565b5f6020828403121561061657610615610439565b5b5f82013567ffffffffffffffff8111156106335761063261043d565b5b61063f84828501610550565b91505092915050565b5f6020828403121561065d5761065c610439565b5b5f61066a84828501610593565b91505092915050565b5f805f6060848603121561068a57610689610439565b5b5f84013567ffffffffffffffff8111156106a7576106a661043d565b5b6106b386828701610550565b935050602084013567ffffffffffffffff8111156106d4576106d361043d565b5b6106e086828701610550565b92505060406106f186828701610593565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610724826106fb565b9050919050565b6107348161071a565b811461073e575f80fd5b50565b5f8135905061074f8161072b565b92915050565b5f806040838503121561076b5761076a610439565b5b5f61077885828601610741565b925050602061078985828601610593565b9150509250929050565b5f81519050919050565b5f81905092915050565b5f5b838110156107c45780820151818401526020810190506107a9565b5f8484015250505050565b5f6107d982610793565b6107e3818561079d565b93506107f38185602086016107a7565b80840191505092915050565b5f61080a82846107cf565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61084c826103ff565b9150610857836103ff565b925082820190508082111561086f5761086e610815565b5b92915050565b5f61087f826103ff565b915061088a836103ff565b92508282039050818111156108a2576108a1610815565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffdfea2646970667358221220ab91a745824e2da829b87e4a87676bfc344bfcb2f3d4fcd2249472232be47e7a64736f6c63430008150033",
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
