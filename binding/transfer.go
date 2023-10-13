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
	Bin: "0x6080806040523461005e57736027946b05e7ab6ef245093622ab18ead545387760018060a01b031973e11f8d55a93bf877a091a3c54c071aac5cc0b01d815f5416175f55600154161760015562989680600255610a3a90816100638239f35b5f80fdfe60806040526004361015610011575f80fd5b5f803560e01c6341b3856914610025575f80fd5b34610451576080366003190112610451576004356001600160a01b03811691908290036104515760243567ffffffffffffffff81116108fc5761006c903690600401610966565b9160443567ffffffffffffffff81116104ba5761008d903690600401610966565b82546040516301fdd06760e11b815260048101939093526001600160a01b0316918381602481865afa90811561047c578491610881575b5060208151910120845160208601200361083c57604051631ec05a9960e11b815260206004820181905281806100fd60248201896109d1565b0381865afa90811561047c57849161080a575b506064358103116107f65760405163ce29615d60e01b8152602060048201819052818061014060248201896109d1565b0381865afa90811561047c5784916107c4575b5060643581018091116107b057600254111561077a5760405162b2610960e31b90818152602060048201526020818061018f602482018a6109d1565b0381875afa9081156105e957859161073b575b5060ff166002146106f65760405190815260206004820152602081806101cb60248201866109d1565b0381865afa90811561047c5784916106b7575b5060ff1660021461067257604051936325c562f560e11b8552602060048601526020858061020f60248201856109d1565b0381865afa94851561047c578495610632575b506040516325c562f560e11b8152602060048201819052818061024860248201876109d1565b0381875afa9081156105e95785916105f4575b5060ff604051966390ce3e5f60e01b88521660048701528486602481875afa9586156105e957859661050e575b508495855b81518110156105035760ff60208260051b840101511660ff8416146102d0575f1981146102bc5760010161028d565b634e487b7160e01b87526011600452602487fd5b5050509091935060015b156104be57833b156104ba578260405180956336555b8560e01b82526040600483015281838161030d60448201896109d1565b606435602483015203925af180156104af5761049b575b825492935083926001600160a01b0316803b1561045f5783604051809263478bc4c960e11b825260406004830152818381610362604482018a6109d1565b606435602483015203925af190811561047c578491610487575b505082546001600160a01b0316803b1561045f5783604051809263cc4ef18360e01b8252604060048301528183816103b7604482018a6109d1565b606435602483015203925af190811561047c578491610464575b50506001546001600160a01b031691823b1561045f576104289261041685809460405196879586948593631a52239d60e31b85526060600486015260648501906109d1565b838103600319016024850152906109d1565b606435604483015203925af18015610454576104415750f35b61044a90610900565b6104515780f35b80fd5b6040513d84823e3d90fd5b505050fd5b61046d90610900565b61047857825f6103d1565b5050fd5b6040513d86823e3d90fd5b61049090610900565b61047857825f61037c565b9190926104a790610900565b908290610324565b6040513d85823e3d90fd5b8280fd5b60405162461bcd60e51b815260206004820152601960248201527f756e617661696c61626c65207472616e736665722074797065000000000000006044820152606490fd5b5050509091936102da565b9095503d8086833e6105208183610928565b8101906020818303126105e55780519167ffffffffffffffff83116105e15780601f8484010112156105e157828201519267ffffffffffffffff84116105cd578360051b90604051946105766020840187610928565b855260208501926020838387010101116105c9579190602083850101925b602082828701010184106105ae575050505050945f610288565b60208080946105bc876109f6565b8152019401939250610594565b8880fd5b634e487b7160e01b88526041600452602488fd5b8680fd5b8580fd5b6040513d87823e3d90fd5b90506020813d60201161062a575b8161060f60209383610928565b8101031261062657610620906109f6565b5f61025b565b8480fd5b3d9150610602565b9094506020813d60201161066a575b8161064e60209383610928565b810103126106665761065f906109f6565b935f610222565b8380fd5b3d9150610641565b60405162461bcd60e51b815260206004820152601760248201527f746f2077616c6c657420737461747573206661696c65640000000000000000006044820152606490fd5b90506020813d6020116106ee575b816106d260209383610928565b810103126106665760ff6106e76002926109f6565b91506101de565b3d91506106c5565b60405162461bcd60e51b815260206004820152601960248201527f66726f6d2077616c6c657420737461747573206661696c6564000000000000006044820152606490fd5b90506020813d602011610772575b8161075660209383610928565b810103126106265760ff61076b6002926109f6565b91506101a2565b3d9150610749565b60405162461bcd60e51b815260206004820152600e60248201526d1b1a5b5a5d08195e18d95959195960921b6044820152606490fd5b634e487b7160e01b84526011600452602484fd5b90506020813d6020116107ee575b816107df60209383610928565b8101031261066657515f610153565b3d91506107d2565b634e487b7160e01b83526011600452602483fd5b90506020813d602011610834575b8161082560209383610928565b8101031261066657515f610110565b3d9150610818565b60405162461bcd60e51b815260206004820152601a60248201527f6f726967696e20646f65736e74206d617463682077616c6c65740000000000006044820152606490fd5b90503d8085833e6108928183610928565b8101906020818303126106265780519067ffffffffffffffff82116105e5570181601f82011215610626578051906108c98261094a565b926108d76040519485610928565b828452602083830101116105e557906108f691602080850191016109b0565b5f6100c4565b5080fd5b67ffffffffffffffff811161091457604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff82111761091457604052565b67ffffffffffffffff811161091457601f01601f191660200190565b81601f820112156109ac5780359061097d8261094a565b9261098b6040519485610928565b828452602083830101116109ac57815f926020809301838601378301015290565b5f80fd5b5f5b8381106109c15750505f910152565b81810151838201526020016109b2565b906020916109ea815180928185528580860191016109b0565b601f01601f1916010190565b519060ff821682036109ac5756fea26469706673582212202a10a32caaf0e97307c11ab77cffc207b9dec8fe819d9d126414815c524bd3d464736f6c63430008150033",
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
