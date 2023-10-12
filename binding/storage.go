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
	Bin: "0x6080806040523461001d57620f424060015561064790816100228239f35b5f80fdfe60806040818152600480361015610014575f80fd5b5f92833560e01c9081632e64cec1146104595750806338639f7b146102ff5780633a51d246146102b75780636057361d1461029a5780639b80b0501461009a5763a5e2d81014610062575f80fd5b34610096578160031936011261009657356001600160a01b0381169081900361009657825260036020526024359082205580f35b8280fd5b50346100965760603660031901126100965767ffffffffffffffff918135838111610296576100cc90369084016104a9565b92602435908111610296576100e490369084016104a9565b90604435926100fc846100f687610503565b546105ab565b506101108461010a8761053b565b546105cc565b90600191825411156102925760ff61012787610573565b5416600381101561027f57600214159081610249575b501561024557338652602091600783528087209051908188825494610161866105d9565b93848452878401968883821691825f146102255750506001146101e9575b505061018d92500382610473565b51902090845190850120036101e557816101db916101b16101e1946100f687610503565b6101ba86610503565b556101d16101cb8361010a84610503565b91610503565b5561010a8461053b565b9161053b565b5580f35b8380fd5b8792508b52818b20908b915b85831061020d57505061018d93508201015f8061017f565b805483880185015286945088939092019181016101f5565b925093505061018d94915060ff19168652151560051b8201015f8061017f565b8580fd5b905060ff61025685610573565b541690600382101561026c57508114155f61013d565b634e487b7160e01b885260219052602487fd5b634e487b7160e01b885260218252602488fd5b8680fd5b8480fd5b8382346102b35760203660031901126102b35735815580f35b5080fd5b5091346102fc5760203660031901126102fc5782359067ffffffffffffffff82116102fc57506102ef6020936102f4923691016104a9565b610503565b549051908152f35b80fd5b508290346102b357826003193601126102b35767ffffffffffffffff9080358281116101e55761033290369083016104a9565b91338452602094600786528420918351918211610446575061035482546105d9565b601f8111610403575b5084601f82116001146103a1575083948192939491610396575b508160011b915f199060031b1c19161790555b6101e160243591610503565b905083015185610377565b601f198216958386528186209186905b8882106103eb57505082600194959697106103d3575b5050811b01905561038a565b8501515f1960f88460031b161c1916905585806103c7565b80600184958294958a015181550194019201906103b1565b828552858520601f830160051c81019187841061043c575b601f0160051c01905b818110610431575061035d565b858155600101610424565b909150819061041b565b634e487b7160e01b855260419052602484fd5b8490346102b357816003193601126102b357602091548152f35b90601f8019910116810190811067ffffffffffffffff82111761049557604052565b634e487b7160e01b5f52604160045260245ffd5b81601f820112156104ff5780359067ffffffffffffffff821161049557604051926104de601f8401601f191660200185610473565b828452602083830101116104ff57815f926020809301838601378301015290565b5f80fd5b60405190818151915f5b8381106105265750506020918101600481520301902090565b6020828201810151868301528593500161050d565b60405190818151915f5b83811061055e5750506020918101600581520301902090565b60208282018101518683015285935001610545565b60405190818151915f5b8381106105965750506020918101600681520301902090565b6020828201810151868301528593500161057d565b919082039182116105b857565b634e487b7160e01b5f52601160045260245ffd5b919082018092116105b857565b90600182811c92168015610607575b60208310146105f357565b634e487b7160e01b5f52602260045260245ffd5b91607f16916105e856fea2646970667358221220d410dd20b120a01d5523471ddef68c23f0185d0a7f249e60b5a8d5d2ee8baeda64736f6c63430008150033",
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
