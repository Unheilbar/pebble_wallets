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
	Bin: "0x6080806040523461005e577394a562ef266f41d4ac4b125c1c2a5aaf7e95246760018060a01b0319731aea632c29d2978a5c6336a3b8bfe9d737eb8fe3815f5416175f5560015416176001556298968060025561087e90816100638239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f803560e01c6394bf8d6114610026575f80fd5b34610373576080366003190112610373576004356001600160a01b03811692908390036105435760018060a01b03825416926301fdd06760e11b82526004820152602081602481865afa9081156103765782916107ba575b5060243503610775576064351561073f576040516389eba42160e01b81526024356004820152602081602481865afa90811561037657829161070d575b5060643510156106d35760405163110bee0960e21b81526024356004820152602081602481865afa9081156103765782916106a1575b50606435810180911161068d5760025411156106575760405163bfb5dcf960e01b908181526024356004820152602081602481875afa9081156104c6578391610618575b5060ff166002146105d3576040519081526044356004820152602081602481865afa908115610376578291610594575b5060ff1660021461054f5760405191630fa5acbd60e01b83526024356004840152602083602481845afa92831561037657829361050f575b50604051630fa5acbd60e01b81526044356004820152602081602481855afa9081156104c65783916104d1575b5060ff604051946390ce3e5f60e01b86521660048501528284602481855afa9384156104c657839461040a575b508293835b81518110156104005760ff60208260051b840101511660ff84161461023f575f19811461022b576001016101fc565b634e487b7160e01b85526011600452602485fd5b50505090915060015b156103bb578082913b15610381578190604460405180948193632ffd25e760e11b83528335600484015260643560248401525af18015610376576103ac575b50805481906001600160a01b0316803b1561038157818091604460405180948193631774ec6160e11b83528335600484015260643560248401525af1801561037657610398575b5080546001600160a01b0316803b156103815781809160446040518094819363e00e3c6f60e01b83528335600484015260643560248401525af1801561037657610384575b506001546001600160a01b0316803b156103815781809160646040518094819363df8af47560e01b835260243560048401526044356024840152833560448401525af18015610376576103635750f35b61036c906107ec565b6103735780f35b80fd5b6040513d84823e3d90fd5b50fd5b61038d906107ec565b61037357805f610313565b6103a1906107ec565b61037357805f6102ce565b6103b5906107ec565b5f610287565b60405162461bcd60e51b815260206004820152601960248201527f756e617661696c61626c65207472616e736665722074797065000000000000006044820152606490fd5b5050509091610248565b9093503d8084833e61041c8183610814565b81016020828203126104c257815167ffffffffffffffff928382116104a6570181601f820112156104be5780519283116104aa578260051b90604051936104666020840186610814565b84526020808501928201019283116104a657602001905b82821061048e57505050925f6101f7565b6020809161049b84610836565b81520191019061047d565b8580fd5b634e487b7160e01b85526041600452602485fd5b8480fd5b8380fd5b6040513d85823e3d90fd5b90506020813d602011610507575b816104ec60209383610814565b81010312610503576104fd90610836565b5f6101ca565b8280fd5b3d91506104df565b9092506020813d602011610547575b8161052b60209383610814565b810103126105435761053c90610836565b915f61019d565b5080fd5b3d915061051e565b60405162461bcd60e51b815260206004820152601760248201527f746f2077616c6c657420737461747573206661696c65640000000000000000006044820152606490fd5b90506020813d6020116105cb575b816105af60209383610814565b810103126105435760ff6105c4600292610836565b9150610165565b3d91506105a2565b60405162461bcd60e51b815260206004820152601960248201527f66726f6d2077616c6c657420737461747573206661696c6564000000000000006044820152606490fd5b90506020813d60201161064f575b8161063360209383610814565b810103126105035760ff610648600292610836565b9150610135565b3d9150610626565b60405162461bcd60e51b815260206004820152600e60248201526d1b1a5b5a5d08195e18d95959195960921b6044820152606490fd5b634e487b7160e01b82526011600452602482fd5b90506020813d6020116106cb575b816106bc60209383610814565b8101031261054357515f6100f1565b3d91506106af565b60405162461bcd60e51b81526020600482015260126024820152716e6f7420656e6f7567682062616c616e636560701b6044820152606490fd5b90506020813d602011610737575b8161072860209383610814565b8101031261054357515f6100bb565b3d915061071b565b60405162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a5908185b5bdd5b9d60921b6044820152606490fd5b60405162461bcd60e51b815260206004820152601a60248201527f6f726967696e20646f65736e74206d617463682077616c6c65740000000000006044820152606490fd5b90506020813d6020116107e4575b816107d560209383610814565b8101031261054357515f61007e565b3d91506107c8565b67ffffffffffffffff811161080057604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff82111761080057604052565b519060ff8216820361084457565b5f80fdfea2646970667358221220f58e36f8546e1852dce6147ff4e2defc87e9260b99f8fb4d726e0d00cf11bbc364736f6c63430008150033",
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
