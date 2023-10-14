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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"fromWalletId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toWalletId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461005e577394a562ef266f41d4ac4b125c1c2a5aaf7e95246760018060a01b0319731aea632c29d2978a5c6336a3b8bfe9d737eb8fe3815f5416175f55600154161760015562989680600255610a9d90816100638239f35b5f80fdfe60806040526004361015610011575f80fd5b5f803560e01c6341b3856914610025575f80fd5b34610458576080366003190112610458576004356001600160a01b03811691908290036104585760243567ffffffffffffffff811161095f5761006c9036906004016109c9565b9160443567ffffffffffffffff81116104c15761008d9036906004016109c9565b82546040516301fdd06760e11b815260048101939093526001600160a01b0316918381602481865afa9081156104835784916108e4575b5060208151910120845160208601200361089f576064351561086957604051631ec05a9960e11b815260206004820181905281806101056024820189610a34565b0381865afa908115610483578491610837575b5060643510156107fd5760405163ce29615d60e01b815260206004820181905281806101476024820189610a34565b0381865afa9081156104835784916107cb575b5060643581018091116107b75760025411156107815760405162b2610960e31b908181526020600482015260208180610196602482018a610a34565b0381875afa9081156105f0578591610742575b5060ff166002146106fd5760405190815260206004820152602081806101d26024820186610a34565b0381865afa9081156104835784916106be575b5060ff1660021461067957604051936325c562f560e11b855260206004860152602085806102166024820185610a34565b0381865afa948515610483578495610639575b506040516325c562f560e11b8152602060048201819052818061024f6024820187610a34565b0381875afa9081156105f05785916105fb575b5060ff604051966390ce3e5f60e01b88521660048701528486602481875afa9586156105f0578596610515575b508495855b815181101561050a5760ff60208260051b840101511660ff8416146102d7575f1981146102c357600101610294565b634e487b7160e01b87526011600452602487fd5b5050509091935060015b156104c557833b156104c1578260405180956336555b8560e01b8252604060048301528183816103146044820189610a34565b606435602483015203925af180156104b6576104a2575b825492935083926001600160a01b0316803b156104665783604051809263478bc4c960e11b825260406004830152818381610369604482018a610a34565b606435602483015203925af190811561048357849161048e575b505082546001600160a01b0316803b156104665783604051809263cc4ef18360e01b8252604060048301528183816103be604482018a610a34565b606435602483015203925af190811561048357849161046b575b50506001546001600160a01b031691823b156104665761042f9261041d85809460405196879586948593631a52239d60e31b8552606060048601526064850190610a34565b83810360031901602485015290610a34565b606435604483015203925af1801561045b576104485750f35b61045190610963565b6104585780f35b80fd5b6040513d84823e3d90fd5b505050fd5b61047490610963565b61047f57825f6103d8565b5050fd5b6040513d86823e3d90fd5b61049790610963565b61047f57825f610383565b9190926104ae90610963565b90829061032b565b6040513d85823e3d90fd5b8280fd5b60405162461bcd60e51b815260206004820152601960248201527f756e617661696c61626c65207472616e736665722074797065000000000000006044820152606490fd5b5050509091936102e1565b9095503d8086833e610527818361098b565b8101906020818303126105ec5780519167ffffffffffffffff83116105e85780601f8484010112156105e857828201519267ffffffffffffffff84116105d4578360051b906040519461057d602084018761098b565b855260208501926020838387010101116105d0579190602083850101925b602082828701010184106105b5575050505050945f61028f565b60208080946105c387610a59565b815201940193925061059b565b8880fd5b634e487b7160e01b88526041600452602488fd5b8680fd5b8580fd5b6040513d87823e3d90fd5b90506020813d602011610631575b816106166020938361098b565b8101031261062d5761062790610a59565b5f610262565b8480fd5b3d9150610609565b9094506020813d602011610671575b816106556020938361098b565b8101031261066d5761066690610a59565b935f610229565b8380fd5b3d9150610648565b60405162461bcd60e51b815260206004820152601760248201527f746f2077616c6c657420737461747573206661696c65640000000000000000006044820152606490fd5b90506020813d6020116106f5575b816106d96020938361098b565b8101031261066d5760ff6106ee600292610a59565b91506101e5565b3d91506106cc565b60405162461bcd60e51b815260206004820152601960248201527f66726f6d2077616c6c657420737461747573206661696c6564000000000000006044820152606490fd5b90506020813d602011610779575b8161075d6020938361098b565b8101031261062d5760ff610772600292610a59565b91506101a9565b3d9150610750565b60405162461bcd60e51b815260206004820152600e60248201526d1b1a5b5a5d08195e18d95959195960921b6044820152606490fd5b634e487b7160e01b84526011600452602484fd5b90506020813d6020116107f5575b816107e66020938361098b565b8101031261066d57515f61015a565b3d91506107d9565b60405162461bcd60e51b81526020600482015260126024820152716e6f7420656e6f7567682062616c616e636560701b6044820152606490fd5b90506020813d602011610861575b816108526020938361098b565b8101031261066d57515f610118565b3d9150610845565b60405162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a5908185b5bdd5b9d60921b6044820152606490fd5b60405162461bcd60e51b815260206004820152601a60248201527f6f726967696e20646f65736e74206d617463682077616c6c65740000000000006044820152606490fd5b90503d8085833e6108f5818361098b565b81019060208183031261062d5780519067ffffffffffffffff82116105ec570181601f8201121561062d5780519061092c826109ad565b9261093a604051948561098b565b828452602083830101116105ec57906109599160208085019101610a13565b5f6100c4565b5080fd5b67ffffffffffffffff811161097757604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff82111761097757604052565b67ffffffffffffffff811161097757601f01601f191660200190565b81601f82011215610a0f578035906109e0826109ad565b926109ee604051948561098b565b82845260208383010111610a0f57815f926020809301838601378301015290565b5f80fd5b5f5b838110610a245750505f910152565b8181015183820152602001610a15565b90602091610a4d81518092818552858086019101610a13565b601f01601f1916010190565b519060ff82168203610a0f5756fea2646970667358221220f1a9b6179758852856b3ea50f3a101dd6da601300c7cb3b20ef944b231722df564736f6c63430008150033",
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

// Transfer is a paid mutator transaction binding the contract method 0x41b38569.
//
// Solidity: function transfer(address origin, string fromWalletId, string toWalletId, uint256 amount) returns()
func (_Transfer *TransferTransactor) Transfer(opts *bind.TransactOpts, origin common.Address, fromWalletId string, toWalletId string, amount *big.Int) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "transfer", origin, fromWalletId, toWalletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x41b38569.
//
// Solidity: function transfer(address origin, string fromWalletId, string toWalletId, uint256 amount) returns()
func (_Transfer *TransferSession) Transfer(origin common.Address, fromWalletId string, toWalletId string, amount *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Transfer(&_Transfer.TransactOpts, origin, fromWalletId, toWalletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x41b38569.
//
// Solidity: function transfer(address origin, string fromWalletId, string toWalletId, uint256 amount) returns()
func (_Transfer *TransferTransactorSession) Transfer(origin common.Address, fromWalletId string, toWalletId string, amount *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Transfer(&_Transfer.TransactOpts, origin, fromWalletId, toWalletId, amount)
}
