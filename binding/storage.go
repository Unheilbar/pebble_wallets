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
	ABI: "[{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setBalanceA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506107328061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80632e64cec11461005957806338639f7b146100775780636057361d14610093578063a5e2d810146100af578063beabacc8146100cb575b5f80fd5b6100616100e7565b60405161006e91906102ea565b60405180910390f35b610091600480360381019061008c919061047a565b6100ef565b005b6100ad60048036038101906100a891906104d4565b610116565b005b6100c960048036038101906100c49190610559565b61011f565b005b6100e560048036038101906100e09190610597565b610165565b005b5f8054905090565b806001836040516101009190610653565b9081526020016040518091039020819055505050565b805f8190555050565b8060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505050565b5f8160035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546101af9190610696565b10156101b9575f80fd5b8060035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546102029190610696565b60035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205461028c91906106c9565b60035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505050565b5f819050919050565b6102e4816102d2565b82525050565b5f6020820190506102fd5f8301846102db565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6103628261031c565b810181811067ffffffffffffffff821117156103815761038061032c565b5b80604052505050565b5f610393610303565b905061039f8282610359565b919050565b5f67ffffffffffffffff8211156103be576103bd61032c565b5b6103c78261031c565b9050602081019050919050565b828183375f83830152505050565b5f6103f46103ef846103a4565b61038a565b9050828152602081018484840111156104105761040f610318565b5b61041b8482856103d4565b509392505050565b5f82601f83011261043757610436610314565b5b81356104478482602086016103e2565b91505092915050565b610459816102d2565b8114610463575f80fd5b50565b5f8135905061047481610450565b92915050565b5f80604083850312156104905761048f61030c565b5b5f83013567ffffffffffffffff8111156104ad576104ac610310565b5b6104b985828601610423565b92505060206104ca85828601610466565b9150509250929050565b5f602082840312156104e9576104e861030c565b5b5f6104f684828501610466565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610528826104ff565b9050919050565b6105388161051e565b8114610542575f80fd5b50565b5f813590506105538161052f565b92915050565b5f806040838503121561056f5761056e61030c565b5b5f61057c85828601610545565b925050602061058d85828601610466565b9150509250929050565b5f805f606084860312156105ae576105ad61030c565b5b5f6105bb86828701610545565b93505060206105cc86828701610545565b92505060406105dd86828701610466565b9150509250925092565b5f81519050919050565b5f81905092915050565b5f5b838110156106185780820151818401526020810190506105fd565b5f8484015250505050565b5f61062d826105e7565b61063781856105f1565b93506106478185602086016105fb565b80840191505092915050565b5f61065e8284610623565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6106a0826102d2565b91506106ab836102d2565b92508282039050818111156106c3576106c2610669565b5b92915050565b5f6106d3826102d2565b91506106de836102d2565b92508282019050808211156106f6576106f5610669565b5b9291505056fea264697066735822122086697f7812cd013bb279e91e4c2ea05d4098301695c22e81d4515e05fe534d5c64736f6c63430008150033",
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

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns()
func (_Storage *StorageTransactor) Transfer(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transfer", from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns()
func (_Storage *StorageSession) Transfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Transfer(&_Storage.TransactOpts, from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns()
func (_Storage *StorageTransactorSession) Transfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Transfer(&_Storage.TransactOpts, from, to, amount)
}
