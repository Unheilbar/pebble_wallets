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
	Bin: "0x6080806040523461005e577394a562ef266f41d4ac4b125c1c2a5aaf7e95246760018060a01b0319731aea632c29d2978a5c6336a3b8bfe9d737eb8fe3815f5416175f55600154161760015562989680600255610a5f90816100638239f35b5f80fdfe60806040526004361015610011575f80fd5b5f803560e01c6341b3856914610025575f80fd5b34610450576080366003190112610450576004356001600160a01b03811691908290036104505760243567ffffffffffffffff81116109215761006c90369060040161098b565b9160443567ffffffffffffffff81116104b95761008d90369060040161098b565b82546040516301fdd06760e11b815260048101939093526001600160a01b0316918381602481865afa90811561047b5784916108a6575b5060208151910120845160208601200361086157604051631ec05a9960e11b815260206004820181905281806100fd60248201896109f6565b0381865afa90811561047b57849161082f575b5060643510156107f55760405163ce29615d60e01b8152602060048201819052818061013f60248201896109f6565b0381865afa90811561047b5784916107c3575b5060643581018091116107af5760025411156107795760405162b2610960e31b90818152602060048201526020818061018e602482018a6109f6565b0381875afa9081156105e857859161073a575b5060ff166002146106f55760405190815260206004820152602081806101ca60248201866109f6565b0381865afa90811561047b5784916106b6575b5060ff1660021461067157604051936325c562f560e11b8552602060048601526020858061020e60248201856109f6565b0381865afa94851561047b578495610631575b506040516325c562f560e11b8152602060048201819052818061024760248201876109f6565b0381875afa9081156105e85785916105f3575b5060ff604051966390ce3e5f60e01b88521660048701528486602481875afa9586156105e857859661050d575b508495855b81518110156105025760ff60208260051b840101511660ff8416146102cf575f1981146102bb5760010161028c565b634e487b7160e01b87526011600452602487fd5b5050509091935060015b156104bd57833b156104b9578260405180956336555b8560e01b82526040600483015281838161030c60448201896109f6565b606435602483015203925af180156104ae5761049a575b825492935083926001600160a01b0316803b1561045e5783604051809263478bc4c960e11b825260406004830152818381610361604482018a6109f6565b606435602483015203925af190811561047b578491610486575b505082546001600160a01b0316803b1561045e5783604051809263cc4ef18360e01b8252604060048301528183816103b6604482018a6109f6565b606435602483015203925af190811561047b578491610463575b50506001546001600160a01b031691823b1561045e576104279261041585809460405196879586948593631a52239d60e31b85526060600486015260648501906109f6565b838103600319016024850152906109f6565b606435604483015203925af18015610453576104405750f35b61044990610925565b6104505780f35b80fd5b6040513d84823e3d90fd5b505050fd5b61046c90610925565b61047757825f6103d0565b5050fd5b6040513d86823e3d90fd5b61048f90610925565b61047757825f61037b565b9190926104a690610925565b908290610323565b6040513d85823e3d90fd5b8280fd5b60405162461bcd60e51b815260206004820152601960248201527f756e617661696c61626c65207472616e736665722074797065000000000000006044820152606490fd5b5050509091936102d9565b9095503d8086833e61051f818361094d565b8101906020818303126105e45780519167ffffffffffffffff83116105e05780601f8484010112156105e057828201519267ffffffffffffffff84116105cc578360051b9060405194610575602084018761094d565b855260208501926020838387010101116105c8579190602083850101925b602082828701010184106105ad575050505050945f610287565b60208080946105bb87610a1b565b8152019401939250610593565b8880fd5b634e487b7160e01b88526041600452602488fd5b8680fd5b8580fd5b6040513d87823e3d90fd5b90506020813d602011610629575b8161060e6020938361094d565b810103126106255761061f90610a1b565b5f61025a565b8480fd5b3d9150610601565b9094506020813d602011610669575b8161064d6020938361094d565b810103126106655761065e90610a1b565b935f610221565b8380fd5b3d9150610640565b60405162461bcd60e51b815260206004820152601760248201527f746f2077616c6c657420737461747573206661696c65640000000000000000006044820152606490fd5b90506020813d6020116106ed575b816106d16020938361094d565b810103126106655760ff6106e6600292610a1b565b91506101dd565b3d91506106c4565b60405162461bcd60e51b815260206004820152601960248201527f66726f6d2077616c6c657420737461747573206661696c6564000000000000006044820152606490fd5b90506020813d602011610771575b816107556020938361094d565b810103126106255760ff61076a600292610a1b565b91506101a1565b3d9150610748565b60405162461bcd60e51b815260206004820152600e60248201526d1b1a5b5a5d08195e18d95959195960921b6044820152606490fd5b634e487b7160e01b84526011600452602484fd5b90506020813d6020116107ed575b816107de6020938361094d565b8101031261066557515f610152565b3d91506107d1565b60405162461bcd60e51b81526020600482015260126024820152716e6f7420656e6f7567682062616c616e636560701b6044820152606490fd5b90506020813d602011610859575b8161084a6020938361094d565b8101031261066557515f610110565b3d915061083d565b60405162461bcd60e51b815260206004820152601a60248201527f6f726967696e20646f65736e74206d617463682077616c6c65740000000000006044820152606490fd5b90503d8085833e6108b7818361094d565b8101906020818303126106255780519067ffffffffffffffff82116105e4570181601f82011215610625578051906108ee8261096f565b926108fc604051948561094d565b828452602083830101116105e4579061091b91602080850191016109d5565b5f6100c4565b5080fd5b67ffffffffffffffff811161093957604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff82111761093957604052565b67ffffffffffffffff811161093957601f01601f191660200190565b81601f820112156109d1578035906109a28261096f565b926109b0604051948561094d565b828452602083830101116109d157815f926020809301838601378301015290565b5f80fd5b5f5b8381106109e65750505f910152565b81810151838201526020016109d7565b90602091610a0f815180928185528580860191016109d5565b601f01601f1916010190565b519060ff821682036109d15756fea264697066735822122095e1b0f145cc075b925d0e573ec08862a821221afee3a8cd1cfe1ba1831251da64736f6c63430008150033",
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
