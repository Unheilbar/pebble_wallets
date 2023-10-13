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

// StateMetaData contains all meta data concerning the State contract.
var StateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"walletId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"walletId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addMonth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"walletId\",\"type\":\"string\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"walletType\",\"type\":\"uint8\"}],\"name\":\"getAvailableTransferTypes\",\"outputs\":[{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSender\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"walletId\",\"type\":\"string\"}],\"name\":\"getWalletStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"walletId\",\"type\":\"string\"}],\"name\":\"getWalletType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"walletId\",\"type\":\"string\"}],\"name\":\"monthAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"walletId\",\"type\":\"string\"}],\"name\":\"setSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"walletId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608080604052346100a8575f80805260056020527f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc8054919068010000000000000000831015610094576001830180825583101561008057906020918152208160051c019060ff60f883549260031b161b19169055610d3c90816100ad8239f35b634e487b7160e01b82526032600452602482fd5b634e487b7160e01b82526041600452602482fd5b5f80fdfe6080604081815260049081361015610015575f80fd5b5f9260e08435811c91826303fba0ce14610abd575081630593084814610a5c57816336555b8514610a145781633d80b532146109b75781634b8ac5ea146109555781638f178992146108d457816390ce3e5f146102cf57508063cc4ef18314610260578063ce29615d146102005763ee31c4cc14610091575f80fd5b346101fc57806003193601126101fc576100a9610bb4565b9067ffffffffffffffff906024358281116101f8576100cb9036908601610c25565b923233146101f85760018060a01b0316855260019260209184835286209481519384116101e557506100fd8554610cce565b601f811161019f575b5081601f841160011461013f57509282939183928794610134575b50501b915f199060031b1c191617905580f35b015192505f80610121565b919083601f1981168789528489209489905b88838310610185575050501061016d575b505050811b01905580f35b01515f1960f88460031b161c191690555f8080610162565b858701518855909601959485019487935090810190610151565b858752828720601f850160051c8101918486106101db575b601f0160051c019085905b8281106101d0575050610106565b8881550185906101c2565b90915081906101b7565b634e487b7160e01b875260419052602486fd5b8580fd5b8280fd5b50823461025d57602036600319011261025d5782359067ffffffffffffffff821161025d57506020610239819461024b93369101610c25565b81845193828580945193849201610bce565b81016002815203019020549051908152f35b80fd5b8382346102cb5761027036610c7b565b903233146102c7576102b9916020916102a985519282519385818186019661029981838a610bce565b8101600281520301902054610cad565b9451938492839251928391610bce565b810160028152030190205580f35b8380fd5b5080fd5b849150346102cb57602093846003193601126101fc57359360ff948581168091036102c75792819293815260058252848120938551809485918882895492838152019182998752838720958c88915b83601f84011061074657908691610441995495858510610732575b85851061071b575b858510610704575b8585106106ed575b8585106106d7575b8585106106c0575b8585106106a9575b858510610692575b85851061067b575b50848410610664575b84841061064d575b848410610636575b84841061061f575b848410610608575b8484106105f1575b8484106105da575b8484106105c3575b8484106105ac575b848410610595575b84841061057e575b848410610567575b848410610550575b848410610539575b848410610522575b84841061050b575b8484106104f4575b8484106104dd575b8484106104c6575b8484106104b1575b5083831061049b575b838310610485575b505010610477575b509050949392940383610bef565b8451948186019282875251809352850193925b8281106104615785850386f35b8351871685529381019392810192600101610454565b60f81c81520184908a610433565b94600192958560f01c168152019301848d61042b565b94600192958560e81c168152019301848d610423565b85901c16855290930192600101848d8f61041a565b929582600191838860d81c16815201960192610412565b929582600191838860d01c1681520196019261040a565b929582600191838860c81c16815201960192610402565b929582600191838860c01c168152019601926103fa565b929582600191838860b81c168152019601926103f2565b929582600191838860b01c168152019601926103ea565b929582600191838860a81c168152019601926103e2565b929582600191838860a01c168152019601926103da565b929582600191838860981c168152019601926103d2565b929582600191838860901c168152019601926103ca565b929582600191838860881c168152019601926103c2565b929582600191838860801c168152019601926103ba565b929582600191838860781c168152019601926103b2565b929582600191838860701c168152019601926103aa565b929582600191838860681c168152019601926103a2565b929582600191838860601c1681520196019261039a565b929582600191838860581c16815201960192610392565b929582600191838860501c1681520196019261038a565b929582600191838860481c16815201960192610382565b6001919497838886931c168152019601928e610379565b939683600191848960381c16815201970193610371565b939683600191848960301c16815201970193610369565b939683600191848960281c16815201970193610361565b9396836001918489831c16815201970193610359565b939683600191848960181c16815201970193610351565b939683600191848960101c16815201970193610349565b939683600191848960081c16815201970193610341565b939683600191848916815201970193610339565b949060019397508080610400949896888c54948580938282168952828260081c168c8a0152828260101c16818a015282828d82828d606090838360181c168282015283836080961c1685820152838360a099828260281c168b85015260c09d8e848460301c1690860152838360381c16908501521c16610100820152838360481c16610120820152838360501c16610140820152610160848460581c169101521c166101808d0152828260681c166101a08d0152828260701c166101c08d0152828260781c166101e08d01521c166102008a0152828260881c166102208a0152828260901c166102408a0152828260981c166102608a01521c16610280870152828260a81c166102a0870152828260b01c166102c0870152828260b81c166102e08701521c16610300840152808260c81c16610320840152808260d01c16610340840152808260d81c166103608401528082881c16610380840152808260e81c166103a08401528160f01c166103c083015260f81c6103e0820152019501920187938a918c8995979461031e565b505090346101fc576108e536610c7b565b92323314610951578051928251936020818186019661090581838a610bce565b8101898152030190205494850394851161093e5750916020916109319351938492839251928391610bce565b8101858152030190205580f35b634e487b7160e01b865260119052602485fd5b8480fd5b5050823461025d57602036600319011261025d5782359067ffffffffffffffff821161025d57506109a56020809461099260ff9436908301610c25565b9082865194838680955193849201610bce565b82019081520301902054169051908152f35b505090346101fc5760203660031901126101fc57803567ffffffffffffffff81116102c7576020936109f0610a03938693369101610c25565b9082855194838680955193849201610bce565b820190815203019020549051908152f35b8483346102cb57610a2436610c7b565b903233146102c757610931916020916102a9855192825193858181860196610a4d81838a610bce565b81018b81520301902054610cad565b5050823461025d57602036600319011261025d5782359067ffffffffffffffff821161025d5750610aaa6020610a98819560ff94369101610c25565b81855193828580945193849201610bce565b8101600381520301902054169051908152f35b838691346101fc576020928360031936011261025d576001600160a01b03610ae3610bb4565b168152838260018083528584208491815495610afe87610cce565b928386528686019783811690815f14610b915750600114610b54575b50505050610b2a92500383610bef565b610b468351948593818552519283809286015285850190610bce565b601f01601f19168101030190f35b815285812095935091905b818310610b79575088945050820101610b2a888080610b1a565b85548884018501529485019487945091830191610b5f565b9350505050610b2a94925060ff19168552151560051b8201018692888080610b1a565b600435906001600160a01b0382168203610bca57565b5f80fd5b5f5b838110610bdf5750505f910152565b8181015183820152602001610bd0565b90601f8019910116810190811067ffffffffffffffff821117610c1157604052565b634e487b7160e01b5f52604160045260245ffd5b81601f82011215610bca5780359067ffffffffffffffff8211610c115760405192610c5a601f8401601f191660200185610bef565b82845260208383010111610bca57815f926020809301838601378301015290565b6040600319820112610bca576004359067ffffffffffffffff8211610bca57610ca691600401610c25565b9060243590565b91908201809211610cba57565b634e487b7160e01b5f52601160045260245ffd5b90600182811c92168015610cfc575b6020831014610ce857565b634e487b7160e01b5f52602260045260245ffd5b91607f1691610cdd56fea26469706673582212201cb2a6c92b87a77612c8bcae8315d1e548fde5a357ef1d60e2d2b2237727c49e64736f6c63430008150033",
}

// StateABI is the input ABI used to generate the binding from.
// Deprecated: Use StateMetaData.ABI instead.
var StateABI = StateMetaData.ABI

// StateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateMetaData.Bin instead.
var StateBin = StateMetaData.Bin

// DeployState deploys a new Ethereum contract, binding an instance of State to it.
func DeployState(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *State, error) {
	parsed, err := StateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &State{StateCaller: StateCaller{contract: contract}, StateTransactor: StateTransactor{contract: contract}, StateFilterer: StateFilterer{contract: contract}}, nil
}

// State is an auto generated Go binding around an Ethereum contract.
type State struct {
	StateCaller     // Read-only binding to the contract
	StateTransactor // Write-only binding to the contract
	StateFilterer   // Log filterer for contract events
}

// StateCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateSession struct {
	Contract     *State            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateCallerSession struct {
	Contract *StateCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateTransactorSession struct {
	Contract     *StateTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateRaw struct {
	Contract *State // Generic contract binding to access the raw methods on
}

// StateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateCallerRaw struct {
	Contract *StateCaller // Generic read-only contract binding to access the raw methods on
}

// StateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateTransactorRaw struct {
	Contract *StateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewState creates a new instance of State, bound to a specific deployed contract.
func NewState(address common.Address, backend bind.ContractBackend) (*State, error) {
	contract, err := bindState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &State{StateCaller: StateCaller{contract: contract}, StateTransactor: StateTransactor{contract: contract}, StateFilterer: StateFilterer{contract: contract}}, nil
}

// NewStateCaller creates a new read-only instance of State, bound to a specific deployed contract.
func NewStateCaller(address common.Address, caller bind.ContractCaller) (*StateCaller, error) {
	contract, err := bindState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateCaller{contract: contract}, nil
}

// NewStateTransactor creates a new write-only instance of State, bound to a specific deployed contract.
func NewStateTransactor(address common.Address, transactor bind.ContractTransactor) (*StateTransactor, error) {
	contract, err := bindState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateTransactor{contract: contract}, nil
}

// NewStateFilterer creates a new log filterer instance of State, bound to a specific deployed contract.
func NewStateFilterer(address common.Address, filterer bind.ContractFilterer) (*StateFilterer, error) {
	contract, err := bindState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateFilterer{contract: contract}, nil
}

// bindState binds a generic wrapper to an already deployed contract.
func bindState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_State *StateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _State.Contract.StateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_State *StateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.Contract.StateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_State *StateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _State.Contract.StateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_State *StateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _State.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_State *StateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_State *StateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _State.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0x3d80b532.
//
// Solidity: function balance(string walletId) view returns(uint256)
func (_State *StateCaller) Balance(opts *bind.CallOpts, walletId string) (*big.Int, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "balance", walletId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x3d80b532.
//
// Solidity: function balance(string walletId) view returns(uint256)
func (_State *StateSession) Balance(walletId string) (*big.Int, error) {
	return _State.Contract.Balance(&_State.CallOpts, walletId)
}

// Balance is a free data retrieval call binding the contract method 0x3d80b532.
//
// Solidity: function balance(string walletId) view returns(uint256)
func (_State *StateCallerSession) Balance(walletId string) (*big.Int, error) {
	return _State.Contract.Balance(&_State.CallOpts, walletId)
}

// GetAvailableTransferTypes is a free data retrieval call binding the contract method 0x90ce3e5f.
//
// Solidity: function getAvailableTransferTypes(uint8 walletType) view returns(uint8[])
func (_State *StateCaller) GetAvailableTransferTypes(opts *bind.CallOpts, walletType uint8) ([]uint8, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getAvailableTransferTypes", walletType)

	if err != nil {
		return *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)

	return out0, err

}

// GetAvailableTransferTypes is a free data retrieval call binding the contract method 0x90ce3e5f.
//
// Solidity: function getAvailableTransferTypes(uint8 walletType) view returns(uint8[])
func (_State *StateSession) GetAvailableTransferTypes(walletType uint8) ([]uint8, error) {
	return _State.Contract.GetAvailableTransferTypes(&_State.CallOpts, walletType)
}

// GetAvailableTransferTypes is a free data retrieval call binding the contract method 0x90ce3e5f.
//
// Solidity: function getAvailableTransferTypes(uint8 walletType) view returns(uint8[])
func (_State *StateCallerSession) GetAvailableTransferTypes(walletType uint8) ([]uint8, error) {
	return _State.Contract.GetAvailableTransferTypes(&_State.CallOpts, walletType)
}

// GetSender is a free data retrieval call binding the contract method 0x03fba0ce.
//
// Solidity: function getSender(address sender) view returns(string)
func (_State *StateCaller) GetSender(opts *bind.CallOpts, sender common.Address) (string, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getSender", sender)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSender is a free data retrieval call binding the contract method 0x03fba0ce.
//
// Solidity: function getSender(address sender) view returns(string)
func (_State *StateSession) GetSender(sender common.Address) (string, error) {
	return _State.Contract.GetSender(&_State.CallOpts, sender)
}

// GetSender is a free data retrieval call binding the contract method 0x03fba0ce.
//
// Solidity: function getSender(address sender) view returns(string)
func (_State *StateCallerSession) GetSender(sender common.Address) (string, error) {
	return _State.Contract.GetSender(&_State.CallOpts, sender)
}

// GetWalletStatus is a free data retrieval call binding the contract method 0x05930848.
//
// Solidity: function getWalletStatus(string walletId) view returns(uint8)
func (_State *StateCaller) GetWalletStatus(opts *bind.CallOpts, walletId string) (uint8, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getWalletStatus", walletId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetWalletStatus is a free data retrieval call binding the contract method 0x05930848.
//
// Solidity: function getWalletStatus(string walletId) view returns(uint8)
func (_State *StateSession) GetWalletStatus(walletId string) (uint8, error) {
	return _State.Contract.GetWalletStatus(&_State.CallOpts, walletId)
}

// GetWalletStatus is a free data retrieval call binding the contract method 0x05930848.
//
// Solidity: function getWalletStatus(string walletId) view returns(uint8)
func (_State *StateCallerSession) GetWalletStatus(walletId string) (uint8, error) {
	return _State.Contract.GetWalletStatus(&_State.CallOpts, walletId)
}

// GetWalletType is a free data retrieval call binding the contract method 0x4b8ac5ea.
//
// Solidity: function getWalletType(string walletId) view returns(uint8)
func (_State *StateCaller) GetWalletType(opts *bind.CallOpts, walletId string) (uint8, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getWalletType", walletId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetWalletType is a free data retrieval call binding the contract method 0x4b8ac5ea.
//
// Solidity: function getWalletType(string walletId) view returns(uint8)
func (_State *StateSession) GetWalletType(walletId string) (uint8, error) {
	return _State.Contract.GetWalletType(&_State.CallOpts, walletId)
}

// GetWalletType is a free data retrieval call binding the contract method 0x4b8ac5ea.
//
// Solidity: function getWalletType(string walletId) view returns(uint8)
func (_State *StateCallerSession) GetWalletType(walletId string) (uint8, error) {
	return _State.Contract.GetWalletType(&_State.CallOpts, walletId)
}

// MonthAmount is a free data retrieval call binding the contract method 0xce29615d.
//
// Solidity: function monthAmount(string walletId) view returns(uint256)
func (_State *StateCaller) MonthAmount(opts *bind.CallOpts, walletId string) (*big.Int, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "monthAmount", walletId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MonthAmount is a free data retrieval call binding the contract method 0xce29615d.
//
// Solidity: function monthAmount(string walletId) view returns(uint256)
func (_State *StateSession) MonthAmount(walletId string) (*big.Int, error) {
	return _State.Contract.MonthAmount(&_State.CallOpts, walletId)
}

// MonthAmount is a free data retrieval call binding the contract method 0xce29615d.
//
// Solidity: function monthAmount(string walletId) view returns(uint256)
func (_State *StateCallerSession) MonthAmount(walletId string) (*big.Int, error) {
	return _State.Contract.MonthAmount(&_State.CallOpts, walletId)
}

// Add is a paid mutator transaction binding the contract method 0x36555b85.
//
// Solidity: function add(string walletId, uint256 amount) returns()
func (_State *StateTransactor) Add(opts *bind.TransactOpts, walletId string, amount *big.Int) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "add", walletId, amount)
}

// Add is a paid mutator transaction binding the contract method 0x36555b85.
//
// Solidity: function add(string walletId, uint256 amount) returns()
func (_State *StateSession) Add(walletId string, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.Add(&_State.TransactOpts, walletId, amount)
}

// Add is a paid mutator transaction binding the contract method 0x36555b85.
//
// Solidity: function add(string walletId, uint256 amount) returns()
func (_State *StateTransactorSession) Add(walletId string, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.Add(&_State.TransactOpts, walletId, amount)
}

// AddMonth is a paid mutator transaction binding the contract method 0xcc4ef183.
//
// Solidity: function addMonth(string walletId, uint256 amount) returns()
func (_State *StateTransactor) AddMonth(opts *bind.TransactOpts, walletId string, amount *big.Int) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "addMonth", walletId, amount)
}

// AddMonth is a paid mutator transaction binding the contract method 0xcc4ef183.
//
// Solidity: function addMonth(string walletId, uint256 amount) returns()
func (_State *StateSession) AddMonth(walletId string, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.AddMonth(&_State.TransactOpts, walletId, amount)
}

// AddMonth is a paid mutator transaction binding the contract method 0xcc4ef183.
//
// Solidity: function addMonth(string walletId, uint256 amount) returns()
func (_State *StateTransactorSession) AddMonth(walletId string, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.AddMonth(&_State.TransactOpts, walletId, amount)
}

// SetSender is a paid mutator transaction binding the contract method 0xee31c4cc.
//
// Solidity: function setSender(address sender, string walletId) returns()
func (_State *StateTransactor) SetSender(opts *bind.TransactOpts, sender common.Address, walletId string) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "setSender", sender, walletId)
}

// SetSender is a paid mutator transaction binding the contract method 0xee31c4cc.
//
// Solidity: function setSender(address sender, string walletId) returns()
func (_State *StateSession) SetSender(sender common.Address, walletId string) (*types.Transaction, error) {
	return _State.Contract.SetSender(&_State.TransactOpts, sender, walletId)
}

// SetSender is a paid mutator transaction binding the contract method 0xee31c4cc.
//
// Solidity: function setSender(address sender, string walletId) returns()
func (_State *StateTransactorSession) SetSender(sender common.Address, walletId string) (*types.Transaction, error) {
	return _State.Contract.SetSender(&_State.TransactOpts, sender, walletId)
}

// Sub is a paid mutator transaction binding the contract method 0x8f178992.
//
// Solidity: function sub(string walletId, uint256 amount) returns()
func (_State *StateTransactor) Sub(opts *bind.TransactOpts, walletId string, amount *big.Int) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "sub", walletId, amount)
}

// Sub is a paid mutator transaction binding the contract method 0x8f178992.
//
// Solidity: function sub(string walletId, uint256 amount) returns()
func (_State *StateSession) Sub(walletId string, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.Sub(&_State.TransactOpts, walletId, amount)
}

// Sub is a paid mutator transaction binding the contract method 0x8f178992.
//
// Solidity: function sub(string walletId, uint256 amount) returns()
func (_State *StateTransactorSession) Sub(walletId string, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.Sub(&_State.TransactOpts, walletId, amount)
}
