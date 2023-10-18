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

// TransferMetaData contains all meta data concerning the Transfer contract.
var TransferMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"fromWalletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"toWalletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461005e57736c02e060d0e1cad7c039a9ae3abc29a40b3dff1f60018060a01b0319736027946b05e7ab6ef245093622ab18ead5453877815f5416175f5560015416176001556298968060025561067590816100638239f35b5f80fdfe60806040818152600480361015610014575f80fd5b5f91823560e01c6394bf8d6114610029575f80fd5b346105d45760803660031901126105d4576001600160a01b038235818116959193908690036105d05760249384359160449182359060643595818a54169a6301fdd06760e11b825284820152808b818b60209485935afa908115610379579087918c9161059f575b500361055f57861561052e5787516389eba42160e01b815284810187905281818b818f5afa908115610379579088918c916104fd575b5011156104c857875163110bee0960e21b815284810187905281818b818f5afa908115610379578b9161049b575b50878101809111610489576002541115610458578a885182818c8163bfb5dcf960e01b958682528c8b8301525afa90811561044e5760029160ff918e91610431575b5016146103f057885190815284810184905281818b818f5afa9081156103795760029160ff918d916103c3575b501614610383578751630ec73a6160e21b8152848101879052808a01849052818187818f5afa908115610379578b9161033f575b508015610337575b156103005750893b156102a8578886858a838b519e8f948593632ffd25e760e11b8552898b8601528401525af180156102f6576102e2575b88995080895416803b156102ca578987868b838c519586948593631774ec6160e11b85528a8c8601528401525af180156102c057908a916102ce575b50548116803b156102ca578987868b838c51958694859363e00e3c6f60e01b85528a8c8601528401525af180156102c057908a916102ac575b505060015416803b156102a8578895606494879389519a8b98899763df8af47560e01b89528801528601528401525af190811561029f575061028c5750f35b610295906105d8565b61029c5780f35b80fd5b513d84823e3d90fd5b8880fd5b6102b5906105d8565b6102a857885f61024d565b88513d8c823e3d90fd5b8980fd5b6102d7906105d8565b6102a857885f610214565b9790986102ee906105d8565b9688906101d8565b87513d8b823e3d90fd5b83733ab730bb30b4b630b13632903a3930b739b332b960611b8660148c6064958d519562461bcd60e51b8752860152840152820152fd5b5060016101a0565b90508181813d8311610372575b6103568183610600565b8101031261036e5751801515810361036e575f610198565b8a80fd5b503d61034c565b89513d8d823e3d90fd5b837f746f2077616c6c657420737461747573206661696c65640000000000000000008660178c6064958d519562461bcd60e51b8752860152840152820152fd5b6103e39150843d86116103e9575b6103db8183610600565b810190610622565b5f610164565b503d6103d1565b885162461bcd60e51b81528086018390526019818c01527f66726f6d2077616c6c657420737461747573206661696c65640000000000000081880152606490fd5b6104489150853d87116103e9576103db8183610600565b5f610137565b8a513d8e823e3d90fd5b836d1b1a5b5a5d08195e18d95959195960921b86600e8c6064958d519562461bcd60e51b8752860152840152820152fd5b634e487b7160e01b8b5260118552898bfd5b90508181813d83116104c1575b6104b28183610600565b8101031261036e57515f6100f5565b503d6104a8565b83716e6f7420656e6f7567682062616c616e636560701b8660128c6064958d519562461bcd60e51b8752860152840152820152fd5b809250838092503d8311610527575b6105168183610600565b8101031261036e578790515f6100c7565b503d61050c565b836d1a5b9d985b1a5908185b5bdd5b9d60921b86600e8c6064958d519562461bcd60e51b8752860152840152820152fd5b837f6f726967696e20646f65736e74206d617463682077616c6c657400000000000086601a8c6064958d519562461bcd60e51b8752860152840152820152fd5b809250838092503d83116105c9575b6105b88183610600565b8101031261036e578690515f610091565b503d6105ae565b8480fd5b8280fd5b67ffffffffffffffff81116105ec57604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff8211176105ec57604052565b9081602091031261063b575160ff8116810361063b5790565b5f80fdfea2646970667358221220e463b70f4caa29fafab585616776e353eba3eb986e07185987cdf504ac5e657064736f6c63430008150033",
}

// TransferABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferMetaData.ABI instead.
var TransferABI = TransferMetaData.ABI

// TransferBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransferMetaData.Bin instead.
var TransferBin = TransferMetaData.Bin

// DeployTransfer deploys a new Ethereum contract, binding an instance of Transfer to it.
func DeployTransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Transfer, error) {
	parsed, err := TransferMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Transfer{TransferCaller: TransferCaller{contract: contract}, TransferTransactor: TransferTransactor{contract: contract}, TransferFilterer: TransferFilterer{contract: contract}}, nil
}

// Transfer is an auto generated Go binding around an Ethereum contract.
type Transfer struct {
	TransferCaller     // Read-only binding to the contract
	TransferTransactor // Write-only binding to the contract
	TransferFilterer   // Log filterer for contract events
}

// TransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferSession struct {
	Contract     *Transfer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferCallerSession struct {
	Contract *TransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferTransactorSession struct {
	Contract     *TransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferRaw struct {
	Contract *Transfer // Generic contract binding to access the raw methods on
}

// TransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferCallerRaw struct {
	Contract *TransferCaller // Generic read-only contract binding to access the raw methods on
}

// TransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferTransactorRaw struct {
	Contract *TransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransfer creates a new instance of Transfer, bound to a specific deployed contract.
func NewTransfer(address common.Address, backend bind.ContractBackend) (*Transfer, error) {
	contract, err := bindTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transfer{TransferCaller: TransferCaller{contract: contract}, TransferTransactor: TransferTransactor{contract: contract}, TransferFilterer: TransferFilterer{contract: contract}}, nil
}

// NewTransferCaller creates a new read-only instance of Transfer, bound to a specific deployed contract.
func NewTransferCaller(address common.Address, caller bind.ContractCaller) (*TransferCaller, error) {
	contract, err := bindTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferCaller{contract: contract}, nil
}

// NewTransferTransactor creates a new write-only instance of Transfer, bound to a specific deployed contract.
func NewTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferTransactor, error) {
	contract, err := bindTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferTransactor{contract: contract}, nil
}

// NewTransferFilterer creates a new log filterer instance of Transfer, bound to a specific deployed contract.
func NewTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferFilterer, error) {
	contract, err := bindTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferFilterer{contract: contract}, nil
}

// bindTransfer binds a generic wrapper to an already deployed contract.
func bindTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransferMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.TransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0x94bf8d61.
//
// Solidity: function transfer(address origin, bytes32 fromWalletId, bytes32 toWalletId, uint256 amount) returns()
func (_Transfer *TransferTransactor) Transfer(opts *bind.TransactOpts, origin common.Address, fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "transfer", origin, fromWalletId, toWalletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x94bf8d61.
//
// Solidity: function transfer(address origin, bytes32 fromWalletId, bytes32 toWalletId, uint256 amount) returns()
func (_Transfer *TransferSession) Transfer(origin common.Address, fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Transfer(&_Transfer.TransactOpts, origin, fromWalletId, toWalletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x94bf8d61.
//
// Solidity: function transfer(address origin, bytes32 fromWalletId, bytes32 toWalletId, uint256 amount) returns()
func (_Transfer *TransferTransactorSession) Transfer(origin common.Address, fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Transfer(&_Transfer.TransactOpts, origin, fromWalletId, toWalletId, amount)
}
