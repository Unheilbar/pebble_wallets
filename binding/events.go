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

// EventsMetaData contains all meta data concerning the Events contract.
var EventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fromWallet\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"toWallet\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"TransferTokens\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fromWalletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"toWalletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"emitTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608080604052346100155760c9908161001a8239f35b5f80fdfe60808060405260043610156011575f80fd5b5f90813560e01c637ee3830f146025575f80fd5b34608f576060366003190112608f576044356fffffffffffffffffffffffffffffffff8116809103608b577fd58fcf656b4c751f996aadecdfbd77bf81cebfdc8950af8de3a571bef7102af991606091600435825260243560208301526040820152a180f35b8280fd5b5080fdfea2646970667358221220bf2f898b7f65218f650f5bc90d100906d912be3713db2e00981023e6c2cd169464736f6c63430008150033",
}

// EventsABI is the input ABI used to generate the binding from.
// Deprecated: Use EventsMetaData.ABI instead.
var EventsABI = EventsMetaData.ABI

// EventsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EventsMetaData.Bin instead.
var EventsBin = EventsMetaData.Bin

// DeployEvents deploys a new Ethereum contract, binding an instance of Events to it.
func DeployEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Events, error) {
	parsed, err := EventsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Events{EventsCaller: EventsCaller{contract: contract}, EventsTransactor: EventsTransactor{contract: contract}, EventsFilterer: EventsFilterer{contract: contract}}, nil
}

// Events is an auto generated Go binding around an Ethereum contract.
type Events struct {
	EventsCaller     // Read-only binding to the contract
	EventsTransactor // Write-only binding to the contract
	EventsFilterer   // Log filterer for contract events
}

// EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventsSession struct {
	Contract     *Events           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventsCallerSession struct {
	Contract *EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventsTransactorSession struct {
	Contract     *EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventsRaw struct {
	Contract *Events // Generic contract binding to access the raw methods on
}

// EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventsCallerRaw struct {
	Contract *EventsCaller // Generic read-only contract binding to access the raw methods on
}

// EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventsTransactorRaw struct {
	Contract *EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvents creates a new instance of Events, bound to a specific deployed contract.
func NewEvents(address common.Address, backend bind.ContractBackend) (*Events, error) {
	contract, err := bindEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Events{EventsCaller: EventsCaller{contract: contract}, EventsTransactor: EventsTransactor{contract: contract}, EventsFilterer: EventsFilterer{contract: contract}}, nil
}

// NewEventsCaller creates a new read-only instance of Events, bound to a specific deployed contract.
func NewEventsCaller(address common.Address, caller bind.ContractCaller) (*EventsCaller, error) {
	contract, err := bindEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventsCaller{contract: contract}, nil
}

// NewEventsTransactor creates a new write-only instance of Events, bound to a specific deployed contract.
func NewEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*EventsTransactor, error) {
	contract, err := bindEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventsTransactor{contract: contract}, nil
}

// NewEventsFilterer creates a new log filterer instance of Events, bound to a specific deployed contract.
func NewEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*EventsFilterer, error) {
	contract, err := bindEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventsFilterer{contract: contract}, nil
}

// bindEvents binds a generic wrapper to an already deployed contract.
func bindEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Events.Contract.EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.contract.Transact(opts, method, params...)
}

// EmitTransfer is a paid mutator transaction binding the contract method 0x7ee3830f.
//
// Solidity: function emitTransfer(bytes32 fromWalletId, bytes32 toWalletId, uint128 amount) returns()
func (_Events *EventsTransactor) EmitTransfer(opts *bind.TransactOpts, fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Events.contract.Transact(opts, "emitTransfer", fromWalletId, toWalletId, amount)
}

// EmitTransfer is a paid mutator transaction binding the contract method 0x7ee3830f.
//
// Solidity: function emitTransfer(bytes32 fromWalletId, bytes32 toWalletId, uint128 amount) returns()
func (_Events *EventsSession) EmitTransfer(fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Events.Contract.EmitTransfer(&_Events.TransactOpts, fromWalletId, toWalletId, amount)
}

// EmitTransfer is a paid mutator transaction binding the contract method 0x7ee3830f.
//
// Solidity: function emitTransfer(bytes32 fromWalletId, bytes32 toWalletId, uint128 amount) returns()
func (_Events *EventsTransactorSession) EmitTransfer(fromWalletId [32]byte, toWalletId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Events.Contract.EmitTransfer(&_Events.TransactOpts, fromWalletId, toWalletId, amount)
}

// EventsTransferTokensIterator is returned from FilterTransferTokens and is used to iterate over the raw logs and unpacked data for TransferTokens events raised by the Events contract.
type EventsTransferTokensIterator struct {
	Event *EventsTransferTokens // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsTransferTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsTransferTokens)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsTransferTokens)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsTransferTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsTransferTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsTransferTokens represents a TransferTokens event raised by the Events contract.
type EventsTransferTokens struct {
	FromWallet [32]byte
	ToWallet   [32]byte
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferTokens is a free log retrieval operation binding the contract event 0xd58fcf656b4c751f996aadecdfbd77bf81cebfdc8950af8de3a571bef7102af9.
//
// Solidity: event TransferTokens(bytes32 fromWallet, bytes32 toWallet, uint128 amount)
func (_Events *EventsFilterer) FilterTransferTokens(opts *bind.FilterOpts) (*EventsTransferTokensIterator, error) {

	logs, sub, err := _Events.contract.FilterLogs(opts, "TransferTokens")
	if err != nil {
		return nil, err
	}
	return &EventsTransferTokensIterator{contract: _Events.contract, event: "TransferTokens", logs: logs, sub: sub}, nil
}

// WatchTransferTokens is a free log subscription operation binding the contract event 0xd58fcf656b4c751f996aadecdfbd77bf81cebfdc8950af8de3a571bef7102af9.
//
// Solidity: event TransferTokens(bytes32 fromWallet, bytes32 toWallet, uint128 amount)
func (_Events *EventsFilterer) WatchTransferTokens(opts *bind.WatchOpts, sink chan<- *EventsTransferTokens) (event.Subscription, error) {

	logs, sub, err := _Events.contract.WatchLogs(opts, "TransferTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsTransferTokens)
				if err := _Events.contract.UnpackLog(event, "TransferTokens", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferTokens is a log parse operation binding the contract event 0xd58fcf656b4c751f996aadecdfbd77bf81cebfdc8950af8de3a571bef7102af9.
//
// Solidity: event TransferTokens(bytes32 fromWallet, bytes32 toWallet, uint128 amount)
func (_Events *EventsFilterer) ParseTransferTokens(log types.Log) (*EventsTransferTokens, error) {
	event := new(EventsTransferTokens)
	if err := _Events.contract.UnpackLog(event, "TransferTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
