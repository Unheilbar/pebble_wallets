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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"fromWalletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"toWalletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461005e57736c02e060d0e1cad7c039a9ae3abc29a40b3dff1f60018060a01b0319736027946b05e7ab6ef245093622ab18ead5453877815f5416175f5560015416176001556298968060025561078e90816100638239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f803560e01c639364c8ec14610026575f80fd5b34610331576080366003190112610331576004356001600160a01b038116929083900361068657606435906001600160801b038216820361067a5760018060a01b03835416936301fdd06760e11b82526004820152602081602481875afa90811561035a578391610650575b506024350361060b576001600160801b038116156105d5576040519263c052bd5360e01b908185526024356004860152604085602481845afa9485156105ca57849085966105a6575b506040519283526044356004840152604083602481855afa90811561037d5785938692610572575b506001600160801b0385166001600160801b0382161115610538576101288588610710565b6001600160801b03600254911610156105025760405163bfb5dcf960e01b81526024356004820152602081602481875afa90811561041f5760029160ff9189916104e3575b50161461049e5760405163bfb5dcf960e01b81526044356004820152602081602481875afa90811561041f5760029160ff91899161046f575b50161461042a57604051630ec73a6160e21b815260243560048201526044356024820152602081604481875afa90811561041f5787916103e4575b5080156103dc575b156103a0576001600160801b03808616911603916001600160801b03831161038c57610216858798610710565b813b156103885760405163b23d46ab60e01b8082526024803560048401526001600160801b03968716908301529190941660448501529296869188916064918391905af1801561037d57610369575b84955061027c8460018060a01b0387541694610710565b92803b15610365576040519283526044803560048501526001600160801b03948516602485015291909316908201529083908290606490829084905af190811561035a578391610343575b50506001546001600160a01b031690813b1561033f576001600160801b0360648492836040519586948593637ee3830f60e01b8552602435600486015260443560248601521660448401525af18015610334576103215750f35b61032a9061068a565b6103315780f35b80fd5b6040513d84823e3d90fd5b5050fd5b61034c9061068a565b61035757815f6102c7565b50fd5b6040513d85823e3d90fd5b8580fd5b9390946103759061068a565b928490610265565b6040513d87823e3d90fd5b8680fd5b634e487b7160e01b86526011600452602486fd5b60405162461bcd60e51b81526020600482015260146024820152733ab730bb30b4b630b13632903a3930b739b332b960611b6044820152606490fd5b5060016101e9565b90506020813d602011610417575b816103ff602093836106b2565b8101031261038857518015158103610388575f6101e1565b3d91506103f2565b6040513d89823e3d90fd5b60405162461bcd60e51b815260206004820152601760248201527f746f2077616c6c657420737461747573206661696c65640000000000000000006044820152606490fd5b610491915060203d602011610497575b61048981836106b2565b81019061073f565b5f6101a6565b503d61047f565b60405162461bcd60e51b815260206004820152601960248201527f66726f6d2077616c6c657420737461747573206661696c6564000000000000006044820152606490fd5b6104fc915060203d6020116104975761048981836106b2565b5f61016d565b60405162461bcd60e51b815260206004820152600e60248201526d1b1a5b5a5d08195e18d95959195960921b6044820152606490fd5b60405162461bcd60e51b81526020600482015260126024820152716e6f7420656e6f7567682062616c616e636560701b6044820152606490fd5b909350610597915060403d60401161059f575b61058f81836106b2565b8101906106ec565b90925f610103565b503d610585565b90506105c291955060403d60401161059f5761058f81836106b2565b94905f6100db565b6040513d86823e3d90fd5b60405162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a5908185b5bdd5b9d60921b6044820152606490fd5b60405162461bcd60e51b815260206004820152601a60248201527f6f726967696e20646f65736e74206d617463682077616c6c65740000000000006044820152606490fd5b90506020813d60201161067e575b8161066b602093836106b2565b8101031261067a57515f610092565b8280fd5b3d915061065e565b5080fd5b67ffffffffffffffff811161069e57604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff82111761069e57604052565b51906001600160801b03821682036106e857565b5f80fd5b91908260409103126106e85761070d6020610706846106d4565b93016106d4565b90565b9190916001600160801b038080941691160191821161072b57565b634e487b7160e01b5f52601160045260245ffd5b908160209103126106e8575160ff811681036106e8579056fea2646970667358221220e7735cd5551740c7a8685abdfa6c52ea073fed429b95c82d726a695f04b7b4f064736f6c63430008150033",
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

// Transfer is a paid mutator transaction binding the contract method 0x9364c8ec.
//
// Solidity: function transfer(address origin, bytes32 fromWalletId, bytes32 toWalletId, uint128 amount) returns()
func (_Transfer *TransferTransactor) Transfer(opts *bind.TransactOpts, origin common.Address, fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "transfer", origin, fromWalletId, toWalletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x9364c8ec.
//
// Solidity: function transfer(address origin, bytes32 fromWalletId, bytes32 toWalletId, uint128 amount) returns()
func (_Transfer *TransferSession) Transfer(origin common.Address, fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Transfer(&_Transfer.TransactOpts, origin, fromWalletId, toWalletId, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x9364c8ec.
//
// Solidity: function transfer(address origin, bytes32 fromWalletId, bytes32 toWalletId, uint128 amount) returns()
func (_Transfer *TransferTransactorSession) Transfer(origin common.Address, fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Transfer(&_Transfer.TransactOpts, origin, fromWalletId, toWalletId, amount)
}
