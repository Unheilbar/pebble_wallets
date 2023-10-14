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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addMonth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"walletType\",\"type\":\"uint8\"}],\"name\":\"getAvailableTransferTypes\",\"outputs\":[{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSender\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"monthAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"setSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608080604052346100a8575f80805260056020527f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc8054919068010000000000000000831015610094576001830180825583101561008057906020918152208160051c019060ff60f883549260031b161b1916905561092590816100ad8239f35b634e487b7160e01b82526032600452602482fd5b634e487b7160e01b82526041600452602482fd5b5f80fdfe6080604081815260049182361015610015575f80fd5b5f9260e08435811c92836303fba0ce1461086a575082630fa5acbd1461083f5782632ee9d8c2146107f1578263442fb824146107c85782635ffa4bce1461078c57826371bae41d1461075057826389eba4211461072857826390ce3e5f1461010457508163bfb5dcf9146100d5575063e00e3c6f14610092575f80fd5b346100d1576100a0366108b8565b9190913233146100cd576100bf908385526002602052828520546108ce565b918352600260205282205580f35b8380fd5b5080fd5b905034610100576020366003190112610100578160209360ff92358152600385522054169051908152f35b8280fd5b90929150346100cd576020806003193601126107245783359460ff958681168091036100d157815260058252838120938281518095828854928381520191829886528386208b87915b83601f840110610599579086915494848410610585575b84841061056e575b848410610557575b848410610540575b84841061052a575b848410610513575b8484106104fc575b8484106104e5575b8484106104cf575b8484106104b8575b8484106104a1575b84841061048a575b848410610473575b84841061045c575b848410610445575b84841061042e575b848410610417575b848410610400575b8484106103e9575b8484106103d2575b8484106103bb575b8484106103a4575b84841061038d575b848410610376575b84841061035f575b848410610348575b848410610331575b84841061031a575b848410610305575b50508282106102f0575b8282106102db575b50106102cd575b50859003601f01601f1916850196905067ffffffffffffffff8711858810176102ba57508581528286529251828601819052928501939291905b8281106102a45785850386f35b8351871685529381019392810192600101610297565b634e487b7160e01b835260419052602482fd5b60f81c81520183905f61025d565b600191948d8560f01c16815201930184610256565b600191948d8560e81c1681520193018461024e565b85901c16855290930192600101848c5f610244565b929582600191838860d81c1681520196019261023c565b929582600191838860d01c16815201960192610234565b929582600191838860c81c1681520196019261022c565b929582600191838860c01c16815201960192610224565b929582600191838860b81c1681520196019261021c565b929582600191838860b01c16815201960192610214565b929582600191838860a81c1681520196019261020c565b929582600191838860a01c16815201960192610204565b929582600191838860981c168152019601926101fc565b929582600191838860901c168152019601926101f4565b929582600191838860881c168152019601926101ec565b929582600191838860801c168152019601926101e4565b929582600191838860781c168152019601926101dc565b929582600191838860701c168152019601926101d4565b929582600191838860681c168152019601926101cc565b929582600191838860601c168152019601926101c4565b929582600191838860581c168152019601926101bc565b929582600191838860501c168152019601926101b4565b929582600191838860481c168152019601926101ac565b92958260019183888c1c168152019601926101a4565b929582600191838860381c1681520196019261019c565b929582600191838860301c16815201960192610194565b929582600191838860281c1681520196019261018c565b9295826001918388831c16815201960192610184565b929582600191838860181c1681520196019261017c565b929582600191838860101c16815201960192610174565b929582600191838860081c1681520196019261016c565b929582600191838816815201960192610164565b95919293946104009082806001948b8a8c54948580938282168952828260081c168c8a0152828260101c16818a015282828d82828d606090838360181c168282015283836080961c1685820152838360a099828260281c168b85015260c09d8e848460301c1690860152838360381c16908501521c16610100820152838360481c16610120820152838360501c16610140820152610160848460581c169101521c166101808d0152828260681c166101a08d0152828260701c166101c08d0152828260781c166101e08d01521c166102008a0152828260881c166102208a0152828260901c166102408a0152828260981c166102608a01521c16610280870152828260a81c166102a0870152828260b01c166102c0870152828260b81c166102e08701521c16610300840152808260c81c16610320840152808260d01c16610340840152808260d81c1661036084015280828a1c16610380840152808260e81c166103a08401528160f01c166103c083015260f81c6103e082015201950191018b889594939261014d565b8480fd5b5090503461010057602036600319011261010057602092829135815280845220549051908152f35b505050346100d157806003193601126100d15761076b61089e565b323314610100576001600160a01b0316825260016020528120602435905580f35b505050346100d15761079d366108b8565b9190913233146100cd576107bb9083855284602052828520546108ce565b9183528260205282205580f35b509050346101005760203660031901126101005760209282913581526002845220549051908152f35b8482853461010057610802366108b8565b9232331461072457818552846020528285205493840393841161082c575083528260205282205580f35b634e487b7160e01b855260119052602484fd5b5090503461010057602036600319011261010057816020938260ff9335825285522054169051908152f35b85908534610100576020366003190112610100576020926001600160a01b0361089161089e565b1681526001845220548152f35b600435906001600160a01b03821682036108b457565b5f80fd5b60409060031901126108b4576004359060243590565b919082018092116108db57565b634e487b7160e01b5f52601160045260245ffdfea264697066735822122044a965ca9af2faa8ee307ea5093e7685d7dac0454a44788257cc3e38a50c4be064736f6c63430008150033",
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

// Balance is a free data retrieval call binding the contract method 0x89eba421.
//
// Solidity: function balance(bytes32 walletId) view returns(uint256)
func (_State *StateCaller) Balance(opts *bind.CallOpts, walletId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "balance", walletId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x89eba421.
//
// Solidity: function balance(bytes32 walletId) view returns(uint256)
func (_State *StateSession) Balance(walletId [32]byte) (*big.Int, error) {
	return _State.Contract.Balance(&_State.CallOpts, walletId)
}

// Balance is a free data retrieval call binding the contract method 0x89eba421.
//
// Solidity: function balance(bytes32 walletId) view returns(uint256)
func (_State *StateCallerSession) Balance(walletId [32]byte) (*big.Int, error) {
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
// Solidity: function getSender(address sender) view returns(bytes32)
func (_State *StateCaller) GetSender(opts *bind.CallOpts, sender common.Address) ([32]byte, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getSender", sender)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSender is a free data retrieval call binding the contract method 0x03fba0ce.
//
// Solidity: function getSender(address sender) view returns(bytes32)
func (_State *StateSession) GetSender(sender common.Address) ([32]byte, error) {
	return _State.Contract.GetSender(&_State.CallOpts, sender)
}

// GetSender is a free data retrieval call binding the contract method 0x03fba0ce.
//
// Solidity: function getSender(address sender) view returns(bytes32)
func (_State *StateCallerSession) GetSender(sender common.Address) ([32]byte, error) {
	return _State.Contract.GetSender(&_State.CallOpts, sender)
}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 walletId) view returns(uint8)
func (_State *StateCaller) GetWalletStatus(opts *bind.CallOpts, walletId [32]byte) (uint8, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getWalletStatus", walletId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 walletId) view returns(uint8)
func (_State *StateSession) GetWalletStatus(walletId [32]byte) (uint8, error) {
	return _State.Contract.GetWalletStatus(&_State.CallOpts, walletId)
}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 walletId) view returns(uint8)
func (_State *StateCallerSession) GetWalletStatus(walletId [32]byte) (uint8, error) {
	return _State.Contract.GetWalletStatus(&_State.CallOpts, walletId)
}

// GetWalletType is a free data retrieval call binding the contract method 0x0fa5acbd.
//
// Solidity: function getWalletType(bytes32 walletId) view returns(uint8)
func (_State *StateCaller) GetWalletType(opts *bind.CallOpts, walletId [32]byte) (uint8, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getWalletType", walletId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetWalletType is a free data retrieval call binding the contract method 0x0fa5acbd.
//
// Solidity: function getWalletType(bytes32 walletId) view returns(uint8)
func (_State *StateSession) GetWalletType(walletId [32]byte) (uint8, error) {
	return _State.Contract.GetWalletType(&_State.CallOpts, walletId)
}

// GetWalletType is a free data retrieval call binding the contract method 0x0fa5acbd.
//
// Solidity: function getWalletType(bytes32 walletId) view returns(uint8)
func (_State *StateCallerSession) GetWalletType(walletId [32]byte) (uint8, error) {
	return _State.Contract.GetWalletType(&_State.CallOpts, walletId)
}

// MonthAmount is a free data retrieval call binding the contract method 0x442fb824.
//
// Solidity: function monthAmount(bytes32 walletId) view returns(uint256)
func (_State *StateCaller) MonthAmount(opts *bind.CallOpts, walletId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "monthAmount", walletId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MonthAmount is a free data retrieval call binding the contract method 0x442fb824.
//
// Solidity: function monthAmount(bytes32 walletId) view returns(uint256)
func (_State *StateSession) MonthAmount(walletId [32]byte) (*big.Int, error) {
	return _State.Contract.MonthAmount(&_State.CallOpts, walletId)
}

// MonthAmount is a free data retrieval call binding the contract method 0x442fb824.
//
// Solidity: function monthAmount(bytes32 walletId) view returns(uint256)
func (_State *StateCallerSession) MonthAmount(walletId [32]byte) (*big.Int, error) {
	return _State.Contract.MonthAmount(&_State.CallOpts, walletId)
}

// Add is a paid mutator transaction binding the contract method 0x5ffa4bce.
//
// Solidity: function add(bytes32 walletId, uint256 amount) returns()
func (_State *StateTransactor) Add(opts *bind.TransactOpts, walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "add", walletId, amount)
}

// Add is a paid mutator transaction binding the contract method 0x5ffa4bce.
//
// Solidity: function add(bytes32 walletId, uint256 amount) returns()
func (_State *StateSession) Add(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.Add(&_State.TransactOpts, walletId, amount)
}

// Add is a paid mutator transaction binding the contract method 0x5ffa4bce.
//
// Solidity: function add(bytes32 walletId, uint256 amount) returns()
func (_State *StateTransactorSession) Add(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.Add(&_State.TransactOpts, walletId, amount)
}

// AddMonth is a paid mutator transaction binding the contract method 0xe00e3c6f.
//
// Solidity: function addMonth(bytes32 walletId, uint256 amount) returns()
func (_State *StateTransactor) AddMonth(opts *bind.TransactOpts, walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "addMonth", walletId, amount)
}

// AddMonth is a paid mutator transaction binding the contract method 0xe00e3c6f.
//
// Solidity: function addMonth(bytes32 walletId, uint256 amount) returns()
func (_State *StateSession) AddMonth(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.AddMonth(&_State.TransactOpts, walletId, amount)
}

// AddMonth is a paid mutator transaction binding the contract method 0xe00e3c6f.
//
// Solidity: function addMonth(bytes32 walletId, uint256 amount) returns()
func (_State *StateTransactorSession) AddMonth(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.AddMonth(&_State.TransactOpts, walletId, amount)
}

// SetSender is a paid mutator transaction binding the contract method 0x71bae41d.
//
// Solidity: function setSender(address sender, bytes32 walletId) returns()
func (_State *StateTransactor) SetSender(opts *bind.TransactOpts, sender common.Address, walletId [32]byte) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "setSender", sender, walletId)
}

// SetSender is a paid mutator transaction binding the contract method 0x71bae41d.
//
// Solidity: function setSender(address sender, bytes32 walletId) returns()
func (_State *StateSession) SetSender(sender common.Address, walletId [32]byte) (*types.Transaction, error) {
	return _State.Contract.SetSender(&_State.TransactOpts, sender, walletId)
}

// SetSender is a paid mutator transaction binding the contract method 0x71bae41d.
//
// Solidity: function setSender(address sender, bytes32 walletId) returns()
func (_State *StateTransactorSession) SetSender(sender common.Address, walletId [32]byte) (*types.Transaction, error) {
	return _State.Contract.SetSender(&_State.TransactOpts, sender, walletId)
}

// Sub is a paid mutator transaction binding the contract method 0x2ee9d8c2.
//
// Solidity: function sub(bytes32 walletId, uint256 amount) returns()
func (_State *StateTransactor) Sub(opts *bind.TransactOpts, walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "sub", walletId, amount)
}

// Sub is a paid mutator transaction binding the contract method 0x2ee9d8c2.
//
// Solidity: function sub(bytes32 walletId, uint256 amount) returns()
func (_State *StateSession) Sub(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.Sub(&_State.TransactOpts, walletId, amount)
}

// Sub is a paid mutator transaction binding the contract method 0x2ee9d8c2.
//
// Solidity: function sub(bytes32 walletId, uint256 amount) returns()
func (_State *StateTransactorSession) Sub(walletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _State.Contract.Sub(&_State.TransactOpts, walletId, amount)
}
