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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setBalanceA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50620f424060018190555061085c806100275f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80632e64cec11461005957806338639f7b146100775780636057361d146100935780639b80b050146100af578063a5e2d810146100cb575b5f80fd5b6100616100e7565b60405161006e91906103af565b60405180910390f35b610091600480360381019061008c919061053f565b6100ef565b005b6100ad60048036038101906100a89190610599565b61013f565b005b6100c960048036038101906100c491906105c4565b610148565b005b6100e560048036038101906100e091906106a6565b610351565b005b5f8054905090565b806004836040516101009190610750565b9081526020016040518091039020546101199190610793565b6004836040516101299190610750565b9081526020016040518091039020819055505050565b805f8190555050565b5f8160048560405161015a9190610750565b90815260200160405180910390205461017391906107c6565b101561017d575f80fd5b600154816005856040516101919190610750565b9081526020016040518091039020546101aa9190610793565b106101b3575f80fd5b6002808111156101c6576101c56107f9565b5b6006846040516101d69190610750565b90815260200160405180910390205f9054906101000a900460ff166002811115610203576102026107f9565b5b141580156102605750600160028111156102205761021f6107f9565b5b6006836040516102309190610750565b90815260200160405180910390205f9054906101000a900460ff16600281111561025d5761025c6107f9565b5b14155b610268575f80fd5b806004846040516102799190610750565b90815260200160405180910390205461029291906107c6565b6004846040516102a29190610750565b908152602001604051809103902081905550806004836040516102c59190610750565b9081526020016040518091039020546102de9190610793565b6004836040516102ee9190610750565b908152602001604051809103902081905550806005846040516103119190610750565b90815260200160405180910390205461032a9190610793565b60058460405161033a9190610750565b908152602001604051809103902081905550505050565b8060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505050565b5f819050919050565b6103a981610397565b82525050565b5f6020820190506103c25f8301846103a0565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610427826103e1565b810181811067ffffffffffffffff82111715610446576104456103f1565b5b80604052505050565b5f6104586103c8565b9050610464828261041e565b919050565b5f67ffffffffffffffff821115610483576104826103f1565b5b61048c826103e1565b9050602081019050919050565b828183375f83830152505050565b5f6104b96104b484610469565b61044f565b9050828152602081018484840111156104d5576104d46103dd565b5b6104e0848285610499565b509392505050565b5f82601f8301126104fc576104fb6103d9565b5b813561050c8482602086016104a7565b91505092915050565b61051e81610397565b8114610528575f80fd5b50565b5f8135905061053981610515565b92915050565b5f8060408385031215610555576105546103d1565b5b5f83013567ffffffffffffffff811115610572576105716103d5565b5b61057e858286016104e8565b925050602061058f8582860161052b565b9150509250929050565b5f602082840312156105ae576105ad6103d1565b5b5f6105bb8482850161052b565b91505092915050565b5f805f606084860312156105db576105da6103d1565b5b5f84013567ffffffffffffffff8111156105f8576105f76103d5565b5b610604868287016104e8565b935050602084013567ffffffffffffffff811115610625576106246103d5565b5b610631868287016104e8565b92505060406106428682870161052b565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6106758261064c565b9050919050565b6106858161066b565b811461068f575f80fd5b50565b5f813590506106a08161067c565b92915050565b5f80604083850312156106bc576106bb6103d1565b5b5f6106c985828601610692565b92505060206106da8582860161052b565b9150509250929050565b5f81519050919050565b5f81905092915050565b5f5b838110156107155780820151818401526020810190506106fa565b5f8484015250505050565b5f61072a826106e4565b61073481856106ee565b93506107448185602086016106f8565b80840191505092915050565b5f61075b8284610720565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61079d82610397565b91506107a883610397565b92508282019050808211156107c0576107bf610766565b5b92915050565b5f6107d082610397565b91506107db83610397565b92508282039050818111156107f3576107f2610766565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffdfea2646970667358221220e334960094035d77e91d69d1a01286bbfe14245951d30cb6f07a526c2c12132f64736f6c63430008150033",
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
