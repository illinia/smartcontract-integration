// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
)

// ActivatableMetaData contains all meta data concerning the Activatable contract.
var ActivatableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Activate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Deactivate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0f15f4c0": "activate()",
		"02fb0c5e": "active()",
		"51b42b00": "deactivate()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x60806040526000805460ff60a01b1916905534801561001d57600080fd5b506100273361002c565b61007c565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6103c88061008b6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806302fb0c5e146100675780630f15f4c01461009057806351b42b001461009a578063715018a6146100a25780638da5cb5b146100aa578063f2fde38b146100c5575b600080fd5b60005461007b90600160a01b900460ff1681565b60405190151581526020015b60405180910390f35b6100986100d8565b005b610098610178565b61009861020c565b6000546040516001600160a01b039091168152602001610087565b6100986100d336600461032d565b610242565b600054600160a01b900460ff16156101115760405162461bcd60e51b815260206004820152600060248201526044015b60405180910390fd5b6000546001600160a01b0316331461013b5760405162461bcd60e51b81526004016101089061035d565b6000805460ff60a01b1916600160a01b17815560405133917ff7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba1591a2565b600054600160a01b900460ff166101ab5760405162461bcd60e51b81526020600482015260006024820152604401610108565b6000546001600160a01b031633146101d55760405162461bcd60e51b81526004016101089061035d565b6000805460ff60a01b1916815560405133917f238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c91a2565b6000546001600160a01b031633146102365760405162461bcd60e51b81526004016101089061035d565b61024060006102dd565b565b6000546001600160a01b0316331461026c5760405162461bcd60e51b81526004016101089061035d565b6001600160a01b0381166102d15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610108565b6102da816102dd565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561033f57600080fd5b81356001600160a01b038116811461035657600080fd5b9392505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260408201526060019056fea2646970667358221220a51cdae461b66c14b6493bc4cbc1ef225361e693bb8b104348eacd653a117d8f64736f6c63430008090033",
}

// ActivatableABI is the input ABI used to generate the binding from.
// Deprecated: Use ActivatableMetaData.ABI instead.
var ActivatableABI = ActivatableMetaData.ABI

// Deprecated: Use ActivatableMetaData.Sigs instead.
// ActivatableFuncSigs maps the 4-byte function signature to its string representation.
var ActivatableFuncSigs = ActivatableMetaData.Sigs

// ActivatableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ActivatableMetaData.Bin instead.
var ActivatableBin = ActivatableMetaData.Bin

// DeployActivatable deploys a new Ethereum contract, binding an instance of Activatable to it.
func DeployActivatable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Activatable, error) {
	parsed, err := ActivatableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ActivatableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Activatable{ActivatableCaller: ActivatableCaller{contract: contract}, ActivatableTransactor: ActivatableTransactor{contract: contract}, ActivatableFilterer: ActivatableFilterer{contract: contract}}, nil
}

// Activatable is an auto generated Go binding around an Ethereum contract.
type Activatable struct {
	ActivatableCaller     // Read-only binding to the contract
	ActivatableTransactor // Write-only binding to the contract
	ActivatableFilterer   // Log filterer for contract events
}

// ActivatableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivatableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivatableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivatableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivatableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivatableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivatableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivatableSession struct {
	Contract     *Activatable      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActivatableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivatableCallerSession struct {
	Contract *ActivatableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ActivatableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivatableTransactorSession struct {
	Contract     *ActivatableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ActivatableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivatableRaw struct {
	Contract *Activatable // Generic contract binding to access the raw methods on
}

// ActivatableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivatableCallerRaw struct {
	Contract *ActivatableCaller // Generic read-only contract binding to access the raw methods on
}

// ActivatableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivatableTransactorRaw struct {
	Contract *ActivatableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivatable creates a new instance of Activatable, bound to a specific deployed contract.
func NewActivatable(address common.Address, backend bind.ContractBackend) (*Activatable, error) {
	contract, err := bindActivatable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Activatable{ActivatableCaller: ActivatableCaller{contract: contract}, ActivatableTransactor: ActivatableTransactor{contract: contract}, ActivatableFilterer: ActivatableFilterer{contract: contract}}, nil
}

// NewActivatableCaller creates a new read-only instance of Activatable, bound to a specific deployed contract.
func NewActivatableCaller(address common.Address, caller bind.ContractCaller) (*ActivatableCaller, error) {
	contract, err := bindActivatable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivatableCaller{contract: contract}, nil
}

// NewActivatableTransactor creates a new write-only instance of Activatable, bound to a specific deployed contract.
func NewActivatableTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivatableTransactor, error) {
	contract, err := bindActivatable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivatableTransactor{contract: contract}, nil
}

// NewActivatableFilterer creates a new log filterer instance of Activatable, bound to a specific deployed contract.
func NewActivatableFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivatableFilterer, error) {
	contract, err := bindActivatable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivatableFilterer{contract: contract}, nil
}

// bindActivatable binds a generic wrapper to an already deployed contract.
func bindActivatable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivatableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Activatable *ActivatableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Activatable.Contract.ActivatableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Activatable *ActivatableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.Contract.ActivatableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Activatable *ActivatableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Activatable.Contract.ActivatableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Activatable *ActivatableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Activatable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Activatable *ActivatableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Activatable *ActivatableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Activatable.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Activatable *ActivatableCaller) Active(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Activatable.contract.Call(opts, &out, "active")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Activatable *ActivatableSession) Active() (bool, error) {
	return _Activatable.Contract.Active(&_Activatable.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Activatable *ActivatableCallerSession) Active() (bool, error) {
	return _Activatable.Contract.Active(&_Activatable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Activatable *ActivatableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Activatable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Activatable *ActivatableSession) Owner() (common.Address, error) {
	return _Activatable.Contract.Owner(&_Activatable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Activatable *ActivatableCallerSession) Owner() (common.Address, error) {
	return _Activatable.Contract.Owner(&_Activatable.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Activatable *ActivatableTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Activatable *ActivatableSession) Activate() (*types.Transaction, error) {
	return _Activatable.Contract.Activate(&_Activatable.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Activatable *ActivatableTransactorSession) Activate() (*types.Transaction, error) {
	return _Activatable.Contract.Activate(&_Activatable.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Activatable *ActivatableTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Activatable *ActivatableSession) Deactivate() (*types.Transaction, error) {
	return _Activatable.Contract.Deactivate(&_Activatable.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Activatable *ActivatableTransactorSession) Deactivate() (*types.Transaction, error) {
	return _Activatable.Contract.Deactivate(&_Activatable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Activatable *ActivatableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Activatable *ActivatableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Activatable.Contract.RenounceOwnership(&_Activatable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Activatable *ActivatableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Activatable.Contract.RenounceOwnership(&_Activatable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Activatable *ActivatableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Activatable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Activatable *ActivatableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Activatable.Contract.TransferOwnership(&_Activatable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Activatable *ActivatableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Activatable.Contract.TransferOwnership(&_Activatable.TransactOpts, newOwner)
}

// ActivatableActivateIterator is returned from FilterActivate and is used to iterate over the raw logs and unpacked data for Activate events raised by the Activatable contract.
type ActivatableActivateIterator struct {
	Event *ActivatableActivate // Event containing the contract specifics and raw log

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
func (it *ActivatableActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivatableActivate)
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
		it.Event = new(ActivatableActivate)
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
func (it *ActivatableActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivatableActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivatableActivate represents a Activate event raised by the Activatable contract.
type ActivatableActivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterActivate is a free log retrieval operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Activatable *ActivatableFilterer) FilterActivate(opts *bind.FilterOpts, _sender []common.Address) (*ActivatableActivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Activatable.contract.FilterLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ActivatableActivateIterator{contract: _Activatable.contract, event: "Activate", logs: logs, sub: sub}, nil
}

// WatchActivate is a free log subscription operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Activatable *ActivatableFilterer) WatchActivate(opts *bind.WatchOpts, sink chan<- *ActivatableActivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Activatable.contract.WatchLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivatableActivate)
				if err := _Activatable.contract.UnpackLog(event, "Activate", log); err != nil {
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

// ParseActivate is a log parse operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Activatable *ActivatableFilterer) ParseActivate(log types.Log) (*ActivatableActivate, error) {
	event := new(ActivatableActivate)
	if err := _Activatable.contract.UnpackLog(event, "Activate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActivatableDeactivateIterator is returned from FilterDeactivate and is used to iterate over the raw logs and unpacked data for Deactivate events raised by the Activatable contract.
type ActivatableDeactivateIterator struct {
	Event *ActivatableDeactivate // Event containing the contract specifics and raw log

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
func (it *ActivatableDeactivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivatableDeactivate)
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
		it.Event = new(ActivatableDeactivate)
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
func (it *ActivatableDeactivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivatableDeactivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivatableDeactivate represents a Deactivate event raised by the Activatable contract.
type ActivatableDeactivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeactivate is a free log retrieval operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Activatable *ActivatableFilterer) FilterDeactivate(opts *bind.FilterOpts, _sender []common.Address) (*ActivatableDeactivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Activatable.contract.FilterLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ActivatableDeactivateIterator{contract: _Activatable.contract, event: "Deactivate", logs: logs, sub: sub}, nil
}

// WatchDeactivate is a free log subscription operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Activatable *ActivatableFilterer) WatchDeactivate(opts *bind.WatchOpts, sink chan<- *ActivatableDeactivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Activatable.contract.WatchLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivatableDeactivate)
				if err := _Activatable.contract.UnpackLog(event, "Deactivate", log); err != nil {
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

// ParseDeactivate is a log parse operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Activatable *ActivatableFilterer) ParseDeactivate(log types.Log) (*ActivatableDeactivate, error) {
	event := new(ActivatableDeactivate)
	if err := _Activatable.contract.UnpackLog(event, "Deactivate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActivatableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Activatable contract.
type ActivatableOwnershipTransferredIterator struct {
	Event *ActivatableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ActivatableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivatableOwnershipTransferred)
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
		it.Event = new(ActivatableOwnershipTransferred)
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
func (it *ActivatableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivatableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivatableOwnershipTransferred represents a OwnershipTransferred event raised by the Activatable contract.
type ActivatableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Activatable *ActivatableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ActivatableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Activatable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ActivatableOwnershipTransferredIterator{contract: _Activatable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Activatable *ActivatableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ActivatableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Activatable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivatableOwnershipTransferred)
				if err := _Activatable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Activatable *ActivatableFilterer) ParseOwnershipTransferred(log types.Log) (*ActivatableOwnershipTransferred, error) {
	event := new(ActivatableOwnershipTransferred)
	if err := _Activatable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Deprecated: Use OwnableMetaData.Sigs instead.
// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = OwnableMetaData.Sigs

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableMetaData contains all meta data concerning the Pausable contract.
var PausableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5c975abb": "paused()",
	},
}

// PausableABI is the input ABI used to generate the binding from.
// Deprecated: Use PausableMetaData.ABI instead.
var PausableABI = PausableMetaData.ABI

// Deprecated: Use PausableMetaData.Sigs instead.
// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = PausableMetaData.Sigs

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

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
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
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
		it.Event = new(PausablePaused)
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
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) ParsePaused(log types.Log) (*PausablePaused, error) {
	event := new(PausablePaused)
	if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
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
		it.Event = new(PausableUnpaused)
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
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) ParseUnpaused(log types.Log) (*PausableUnpaused, error) {
	event := new(PausableUnpaused)
	if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomMetaData contains all meta data concerning the Room contract.
var RoomMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Activate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Deactivate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositedValue\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_refundedBalance\",\"type\":\"uint256\"}],\"name\":\"RefundedToOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"RewardSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundToOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardSent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"sendReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0f15f4c0": "activate()",
		"02fb0c5e": "active()",
		"51b42b00": "deactivate()",
		"d0e30db0": "deposit()",
		"8da5cb5b": "owner()",
		"5c975abb": "paused()",
		"3e4e10d6": "refundToOwner()",
		"715018a6": "renounceOwnership()",
		"8f134025": "rewardSent(uint256)",
		"2673bd74": "sendReward(uint256,address,uint256)",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060408190526000805460ff60a81b198116909155600180546001600160a01b0319166001600160a01b039092169190911790556109b038819003908190833981016040819052610050916100db565b6100593361008b565b6000805460ff60a01b19169055600180546001600160a01b0319166001600160a01b039290921691909117905561010b565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100ed57600080fd5b81516001600160a01b038116811461010457600080fd5b9392505050565b6108968061011a6000396000f3fe60806040526004361061009c5760003560e01c80635c975abb116100645780635c975abb14610138578063715018a6146101575780638da5cb5b1461016c5780638f13402514610194578063d0e30db0146101c4578063f2fde38b146101cc57600080fd5b806302fb0c5e146100a15780630f15f4c0146100d75780632673bd74146100ee5780633e4e10d61461010e57806351b42b0014610123575b600080fd5b3480156100ad57600080fd5b506000546100c290600160a81b900460ff1681565b60405190151581526020015b60405180910390f35b3480156100e357600080fd5b506100ec6101ec565b005b3480156100fa57600080fd5b506100ec6101093660046107bb565b61028c565b34801561011a57600080fd5b506100ec610441565b34801561012f57600080fd5b506100ec61053c565b34801561014457600080fd5b50600054600160a01b900460ff166100c2565b34801561016357600080fd5b506100ec6105d0565b34801561017857600080fd5b506000546040516001600160a01b0390911681526020016100ce565b3480156101a057600080fd5b506100c26101af3660046107f0565b60026020526000908152604090205460ff1681565b6100ec610606565b3480156101d857600080fd5b506100ec6101e7366004610809565b6106b4565b600054600160a81b900460ff16156102255760405162461bcd60e51b815260206004820152600060248201526044015b60405180910390fd5b6000546001600160a01b0316331461024f5760405162461bcd60e51b815260040161021c9061082b565b6000805460ff60a81b1916600160a81b17815560405133917ff7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba1591a2565b6000546001600160a01b031633146102b65760405162461bcd60e51b815260040161021c9061082b565b60008181526002602052604090205460ff16156102ef5760405162461bcd60e51b8152602060048201526000602482015260440161021c565b600083116103195760405162461bcd60e51b8152602060048201526000602482015260440161021c565b824710156103435760405162461bcd60e51b8152602060048201526000602482015260440161021c565b6001600160a01b0382166103735760405162461bcd60e51b8152602060048201526000602482015260440161021c565b6001546001600160a01b03838116911614156103ab5760405162461bcd60e51b8152602060048201526000602482015260440161021c565b600081815260026020526040808220805460ff19166001179055516001600160a01b0384169185156108fc02918691818181858888f193505050501580156103f7573d6000803e3d6000fd5b5060408051848152602081018390526001600160a01b038416917f6379339f0ae63e95e65fad18ca2a7ec4e7e3f05f3cc5f7079f4d8da8cec34faa910160405180910390a2505050565b600054600160a81b900460ff16156104755760405162461bcd60e51b8152602060048201526000602482015260440161021c565b6000546001600160a01b0316331461049f5760405162461bcd60e51b815260040161021c9061082b565b600047116104c95760405162461bcd60e51b8152602060048201526000602482015260440161021c565b60015460405147916001600160a01b03169082156108fc029083906000818181858888f19350505050158015610503573d6000803e3d6000fd5b5060405181815233907fa9304f0ac97f820b11d5a66f29245fb1e76b1f1f93f46a84a71682e69d095e809060200160405180910390a250565b600054600160a81b900460ff1661056f5760405162461bcd60e51b8152602060048201526000602482015260440161021c565b6000546001600160a01b031633146105995760405162461bcd60e51b815260040161021c9061082b565b6000805460ff60a81b1916815560405133917f238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c91a2565b6000546001600160a01b031633146105fa5760405162461bcd60e51b815260040161021c9061082b565b610604600061074f565b565b600054600160a01b900460ff16156106535760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161021c565b6000341161067d5760405162461bcd60e51b8152602060048201526000602482015260440161021c565b60405134815233907f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c49060200160405180910390a2565b6000546001600160a01b031633146106de5760405162461bcd60e51b815260040161021c9061082b565b6001600160a01b0381166107435760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161021c565b61074c8161074f565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80356001600160a01b03811681146107b657600080fd5b919050565b6000806000606084860312156107d057600080fd5b833592506107e06020850161079f565b9150604084013590509250925092565b60006020828403121561080257600080fd5b5035919050565b60006020828403121561081b57600080fd5b6108248261079f565b9392505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260408201526060019056fea2646970667358221220280167c5526444f7843187b65bd6989c15e4050d4b9fd3817bc59e82cef7f4c964736f6c63430008090033",
}

// RoomABI is the input ABI used to generate the binding from.
// Deprecated: Use RoomMetaData.ABI instead.
var RoomABI = RoomMetaData.ABI

// Deprecated: Use RoomMetaData.Sigs instead.
// RoomFuncSigs maps the 4-byte function signature to its string representation.
var RoomFuncSigs = RoomMetaData.Sigs

// RoomBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoomMetaData.Bin instead.
var RoomBin = RoomMetaData.Bin

// DeployRoom deploys a new Ethereum contract, binding an instance of Room to it.
func DeployRoom(auth *bind.TransactOpts, backend bind.ContractBackend, _creator common.Address) (common.Address, *types.Transaction, *Room, error) {
	parsed, err := RoomMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoomBin), backend, _creator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Room{RoomCaller: RoomCaller{contract: contract}, RoomTransactor: RoomTransactor{contract: contract}, RoomFilterer: RoomFilterer{contract: contract}}, nil
}

// Room is an auto generated Go binding around an Ethereum contract.
type Room struct {
	RoomCaller     // Read-only binding to the contract
	RoomTransactor // Write-only binding to the contract
	RoomFilterer   // Log filterer for contract events
}

// RoomCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoomCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoomTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoomFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoomSession struct {
	Contract     *Room             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoomCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoomCallerSession struct {
	Contract *RoomCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RoomTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoomTransactorSession struct {
	Contract     *RoomTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoomRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoomRaw struct {
	Contract *Room // Generic contract binding to access the raw methods on
}

// RoomCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoomCallerRaw struct {
	Contract *RoomCaller // Generic read-only contract binding to access the raw methods on
}

// RoomTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoomTransactorRaw struct {
	Contract *RoomTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoom creates a new instance of Room, bound to a specific deployed contract.
func NewRoom(address common.Address, backend bind.ContractBackend) (*Room, error) {
	contract, err := bindRoom(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Room{RoomCaller: RoomCaller{contract: contract}, RoomTransactor: RoomTransactor{contract: contract}, RoomFilterer: RoomFilterer{contract: contract}}, nil
}

// NewRoomCaller creates a new read-only instance of Room, bound to a specific deployed contract.
func NewRoomCaller(address common.Address, caller bind.ContractCaller) (*RoomCaller, error) {
	contract, err := bindRoom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoomCaller{contract: contract}, nil
}

// NewRoomTransactor creates a new write-only instance of Room, bound to a specific deployed contract.
func NewRoomTransactor(address common.Address, transactor bind.ContractTransactor) (*RoomTransactor, error) {
	contract, err := bindRoom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoomTransactor{contract: contract}, nil
}

// NewRoomFilterer creates a new log filterer instance of Room, bound to a specific deployed contract.
func NewRoomFilterer(address common.Address, filterer bind.ContractFilterer) (*RoomFilterer, error) {
	contract, err := bindRoom(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoomFilterer{contract: contract}, nil
}

// bindRoom binds a generic wrapper to an already deployed contract.
func bindRoom(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoomABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Room *RoomRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Room.Contract.RoomCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Room *RoomRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.Contract.RoomTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Room *RoomRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Room.Contract.RoomTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Room *RoomCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Room.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Room *RoomTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Room *RoomTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Room.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Room *RoomCaller) Active(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Room.contract.Call(opts, &out, "active")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Room *RoomSession) Active() (bool, error) {
	return _Room.Contract.Active(&_Room.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Room *RoomCallerSession) Active() (bool, error) {
	return _Room.Contract.Active(&_Room.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Room *RoomCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Room.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Room *RoomSession) Owner() (common.Address, error) {
	return _Room.Contract.Owner(&_Room.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Room *RoomCallerSession) Owner() (common.Address, error) {
	return _Room.Contract.Owner(&_Room.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Room *RoomCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Room.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Room *RoomSession) Paused() (bool, error) {
	return _Room.Contract.Paused(&_Room.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Room *RoomCallerSession) Paused() (bool, error) {
	return _Room.Contract.Paused(&_Room.CallOpts)
}

// RewardSent is a free data retrieval call binding the contract method 0x8f134025.
//
// Solidity: function rewardSent(uint256 ) view returns(bool)
func (_Room *RoomCaller) RewardSent(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Room.contract.Call(opts, &out, "rewardSent", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardSent is a free data retrieval call binding the contract method 0x8f134025.
//
// Solidity: function rewardSent(uint256 ) view returns(bool)
func (_Room *RoomSession) RewardSent(arg0 *big.Int) (bool, error) {
	return _Room.Contract.RewardSent(&_Room.CallOpts, arg0)
}

// RewardSent is a free data retrieval call binding the contract method 0x8f134025.
//
// Solidity: function rewardSent(uint256 ) view returns(bool)
func (_Room *RoomCallerSession) RewardSent(arg0 *big.Int) (bool, error) {
	return _Room.Contract.RewardSent(&_Room.CallOpts, arg0)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Room *RoomTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Room *RoomSession) Activate() (*types.Transaction, error) {
	return _Room.Contract.Activate(&_Room.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Room *RoomTransactorSession) Activate() (*types.Transaction, error) {
	return _Room.Contract.Activate(&_Room.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Room *RoomTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Room *RoomSession) Deactivate() (*types.Transaction, error) {
	return _Room.Contract.Deactivate(&_Room.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Room *RoomTransactorSession) Deactivate() (*types.Transaction, error) {
	return _Room.Contract.Deactivate(&_Room.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Room *RoomTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Room *RoomSession) Deposit() (*types.Transaction, error) {
	return _Room.Contract.Deposit(&_Room.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Room *RoomTransactorSession) Deposit() (*types.Transaction, error) {
	return _Room.Contract.Deposit(&_Room.TransactOpts)
}

// RefundToOwner is a paid mutator transaction binding the contract method 0x3e4e10d6.
//
// Solidity: function refundToOwner() returns()
func (_Room *RoomTransactor) RefundToOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "refundToOwner")
}

// RefundToOwner is a paid mutator transaction binding the contract method 0x3e4e10d6.
//
// Solidity: function refundToOwner() returns()
func (_Room *RoomSession) RefundToOwner() (*types.Transaction, error) {
	return _Room.Contract.RefundToOwner(&_Room.TransactOpts)
}

// RefundToOwner is a paid mutator transaction binding the contract method 0x3e4e10d6.
//
// Solidity: function refundToOwner() returns()
func (_Room *RoomTransactorSession) RefundToOwner() (*types.Transaction, error) {
	return _Room.Contract.RefundToOwner(&_Room.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Room *RoomTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Room *RoomSession) RenounceOwnership() (*types.Transaction, error) {
	return _Room.Contract.RenounceOwnership(&_Room.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Room *RoomTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Room.Contract.RenounceOwnership(&_Room.TransactOpts)
}

// SendReward is a paid mutator transaction binding the contract method 0x2673bd74.
//
// Solidity: function sendReward(uint256 _reward, address _dest, uint256 _id) returns()
func (_Room *RoomTransactor) SendReward(opts *bind.TransactOpts, _reward *big.Int, _dest common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "sendReward", _reward, _dest, _id)
}

// SendReward is a paid mutator transaction binding the contract method 0x2673bd74.
//
// Solidity: function sendReward(uint256 _reward, address _dest, uint256 _id) returns()
func (_Room *RoomSession) SendReward(_reward *big.Int, _dest common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Room.Contract.SendReward(&_Room.TransactOpts, _reward, _dest, _id)
}

// SendReward is a paid mutator transaction binding the contract method 0x2673bd74.
//
// Solidity: function sendReward(uint256 _reward, address _dest, uint256 _id) returns()
func (_Room *RoomTransactorSession) SendReward(_reward *big.Int, _dest common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Room.Contract.SendReward(&_Room.TransactOpts, _reward, _dest, _id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Room *RoomTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Room *RoomSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Room.Contract.TransferOwnership(&_Room.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Room *RoomTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Room.Contract.TransferOwnership(&_Room.TransactOpts, newOwner)
}

// RoomActivateIterator is returned from FilterActivate and is used to iterate over the raw logs and unpacked data for Activate events raised by the Room contract.
type RoomActivateIterator struct {
	Event *RoomActivate // Event containing the contract specifics and raw log

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
func (it *RoomActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomActivate)
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
		it.Event = new(RoomActivate)
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
func (it *RoomActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomActivate represents a Activate event raised by the Room contract.
type RoomActivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterActivate is a free log retrieval operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Room *RoomFilterer) FilterActivate(opts *bind.FilterOpts, _sender []common.Address) (*RoomActivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &RoomActivateIterator{contract: _Room.contract, event: "Activate", logs: logs, sub: sub}, nil
}

// WatchActivate is a free log subscription operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Room *RoomFilterer) WatchActivate(opts *bind.WatchOpts, sink chan<- *RoomActivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomActivate)
				if err := _Room.contract.UnpackLog(event, "Activate", log); err != nil {
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

// ParseActivate is a log parse operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Room *RoomFilterer) ParseActivate(log types.Log) (*RoomActivate, error) {
	event := new(RoomActivate)
	if err := _Room.contract.UnpackLog(event, "Activate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomDeactivateIterator is returned from FilterDeactivate and is used to iterate over the raw logs and unpacked data for Deactivate events raised by the Room contract.
type RoomDeactivateIterator struct {
	Event *RoomDeactivate // Event containing the contract specifics and raw log

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
func (it *RoomDeactivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomDeactivate)
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
		it.Event = new(RoomDeactivate)
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
func (it *RoomDeactivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomDeactivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomDeactivate represents a Deactivate event raised by the Room contract.
type RoomDeactivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeactivate is a free log retrieval operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Room *RoomFilterer) FilterDeactivate(opts *bind.FilterOpts, _sender []common.Address) (*RoomDeactivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &RoomDeactivateIterator{contract: _Room.contract, event: "Deactivate", logs: logs, sub: sub}, nil
}

// WatchDeactivate is a free log subscription operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Room *RoomFilterer) WatchDeactivate(opts *bind.WatchOpts, sink chan<- *RoomDeactivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomDeactivate)
				if err := _Room.contract.UnpackLog(event, "Deactivate", log); err != nil {
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

// ParseDeactivate is a log parse operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Room *RoomFilterer) ParseDeactivate(log types.Log) (*RoomDeactivate, error) {
	event := new(RoomDeactivate)
	if err := _Room.contract.UnpackLog(event, "Deactivate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Room contract.
type RoomDepositedIterator struct {
	Event *RoomDeposited // Event containing the contract specifics and raw log

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
func (it *RoomDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomDeposited)
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
		it.Event = new(RoomDeposited)
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
func (it *RoomDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomDeposited represents a Deposited event raised by the Room contract.
type RoomDeposited struct {
	Depositor      common.Address
	DepositedValue *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed _depositor, uint256 _depositedValue)
func (_Room *RoomFilterer) FilterDeposited(opts *bind.FilterOpts, _depositor []common.Address) (*RoomDepositedIterator, error) {

	var _depositorRule []interface{}
	for _, _depositorItem := range _depositor {
		_depositorRule = append(_depositorRule, _depositorItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "Deposited", _depositorRule)
	if err != nil {
		return nil, err
	}
	return &RoomDepositedIterator{contract: _Room.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed _depositor, uint256 _depositedValue)
func (_Room *RoomFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *RoomDeposited, _depositor []common.Address) (event.Subscription, error) {

	var _depositorRule []interface{}
	for _, _depositorItem := range _depositor {
		_depositorRule = append(_depositorRule, _depositorItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "Deposited", _depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomDeposited)
				if err := _Room.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed _depositor, uint256 _depositedValue)
func (_Room *RoomFilterer) ParseDeposited(log types.Log) (*RoomDeposited, error) {
	event := new(RoomDeposited)
	if err := _Room.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Room contract.
type RoomOwnershipTransferredIterator struct {
	Event *RoomOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RoomOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomOwnershipTransferred)
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
		it.Event = new(RoomOwnershipTransferred)
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
func (it *RoomOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomOwnershipTransferred represents a OwnershipTransferred event raised by the Room contract.
type RoomOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Room *RoomFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RoomOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RoomOwnershipTransferredIterator{contract: _Room.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Room *RoomFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RoomOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomOwnershipTransferred)
				if err := _Room.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Room *RoomFilterer) ParseOwnershipTransferred(log types.Log) (*RoomOwnershipTransferred, error) {
	event := new(RoomOwnershipTransferred)
	if err := _Room.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Room contract.
type RoomPausedIterator struct {
	Event *RoomPaused // Event containing the contract specifics and raw log

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
func (it *RoomPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomPaused)
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
		it.Event = new(RoomPaused)
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
func (it *RoomPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomPaused represents a Paused event raised by the Room contract.
type RoomPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Room *RoomFilterer) FilterPaused(opts *bind.FilterOpts) (*RoomPausedIterator, error) {

	logs, sub, err := _Room.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RoomPausedIterator{contract: _Room.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Room *RoomFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RoomPaused) (event.Subscription, error) {

	logs, sub, err := _Room.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomPaused)
				if err := _Room.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Room *RoomFilterer) ParsePaused(log types.Log) (*RoomPaused, error) {
	event := new(RoomPaused)
	if err := _Room.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomRefundedToOwnerIterator is returned from FilterRefundedToOwner and is used to iterate over the raw logs and unpacked data for RefundedToOwner events raised by the Room contract.
type RoomRefundedToOwnerIterator struct {
	Event *RoomRefundedToOwner // Event containing the contract specifics and raw log

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
func (it *RoomRefundedToOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomRefundedToOwner)
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
		it.Event = new(RoomRefundedToOwner)
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
func (it *RoomRefundedToOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomRefundedToOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomRefundedToOwner represents a RefundedToOwner event raised by the Room contract.
type RoomRefundedToOwner struct {
	Dest            common.Address
	RefundedBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRefundedToOwner is a free log retrieval operation binding the contract event 0xa9304f0ac97f820b11d5a66f29245fb1e76b1f1f93f46a84a71682e69d095e80.
//
// Solidity: event RefundedToOwner(address indexed _dest, uint256 _refundedBalance)
func (_Room *RoomFilterer) FilterRefundedToOwner(opts *bind.FilterOpts, _dest []common.Address) (*RoomRefundedToOwnerIterator, error) {

	var _destRule []interface{}
	for _, _destItem := range _dest {
		_destRule = append(_destRule, _destItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "RefundedToOwner", _destRule)
	if err != nil {
		return nil, err
	}
	return &RoomRefundedToOwnerIterator{contract: _Room.contract, event: "RefundedToOwner", logs: logs, sub: sub}, nil
}

// WatchRefundedToOwner is a free log subscription operation binding the contract event 0xa9304f0ac97f820b11d5a66f29245fb1e76b1f1f93f46a84a71682e69d095e80.
//
// Solidity: event RefundedToOwner(address indexed _dest, uint256 _refundedBalance)
func (_Room *RoomFilterer) WatchRefundedToOwner(opts *bind.WatchOpts, sink chan<- *RoomRefundedToOwner, _dest []common.Address) (event.Subscription, error) {

	var _destRule []interface{}
	for _, _destItem := range _dest {
		_destRule = append(_destRule, _destItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "RefundedToOwner", _destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomRefundedToOwner)
				if err := _Room.contract.UnpackLog(event, "RefundedToOwner", log); err != nil {
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

// ParseRefundedToOwner is a log parse operation binding the contract event 0xa9304f0ac97f820b11d5a66f29245fb1e76b1f1f93f46a84a71682e69d095e80.
//
// Solidity: event RefundedToOwner(address indexed _dest, uint256 _refundedBalance)
func (_Room *RoomFilterer) ParseRefundedToOwner(log types.Log) (*RoomRefundedToOwner, error) {
	event := new(RoomRefundedToOwner)
	if err := _Room.contract.UnpackLog(event, "RefundedToOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomRewardSentIterator is returned from FilterRewardSent and is used to iterate over the raw logs and unpacked data for RewardSent events raised by the Room contract.
type RoomRewardSentIterator struct {
	Event *RoomRewardSent // Event containing the contract specifics and raw log

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
func (it *RoomRewardSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomRewardSent)
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
		it.Event = new(RoomRewardSent)
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
func (it *RoomRewardSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomRewardSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomRewardSent represents a RewardSent event raised by the Room contract.
type RoomRewardSent struct {
	Dest   common.Address
	Reward *big.Int
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardSent is a free log retrieval operation binding the contract event 0x6379339f0ae63e95e65fad18ca2a7ec4e7e3f05f3cc5f7079f4d8da8cec34faa.
//
// Solidity: event RewardSent(address indexed _dest, uint256 _reward, uint256 _id)
func (_Room *RoomFilterer) FilterRewardSent(opts *bind.FilterOpts, _dest []common.Address) (*RoomRewardSentIterator, error) {

	var _destRule []interface{}
	for _, _destItem := range _dest {
		_destRule = append(_destRule, _destItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "RewardSent", _destRule)
	if err != nil {
		return nil, err
	}
	return &RoomRewardSentIterator{contract: _Room.contract, event: "RewardSent", logs: logs, sub: sub}, nil
}

// WatchRewardSent is a free log subscription operation binding the contract event 0x6379339f0ae63e95e65fad18ca2a7ec4e7e3f05f3cc5f7079f4d8da8cec34faa.
//
// Solidity: event RewardSent(address indexed _dest, uint256 _reward, uint256 _id)
func (_Room *RoomFilterer) WatchRewardSent(opts *bind.WatchOpts, sink chan<- *RoomRewardSent, _dest []common.Address) (event.Subscription, error) {

	var _destRule []interface{}
	for _, _destItem := range _dest {
		_destRule = append(_destRule, _destItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "RewardSent", _destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomRewardSent)
				if err := _Room.contract.UnpackLog(event, "RewardSent", log); err != nil {
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

// ParseRewardSent is a log parse operation binding the contract event 0x6379339f0ae63e95e65fad18ca2a7ec4e7e3f05f3cc5f7079f4d8da8cec34faa.
//
// Solidity: event RewardSent(address indexed _dest, uint256 _reward, uint256 _id)
func (_Room *RoomFilterer) ParseRewardSent(log types.Log) (*RoomRewardSent, error) {
	event := new(RoomRewardSent)
	if err := _Room.contract.UnpackLog(event, "RewardSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Room contract.
type RoomUnpausedIterator struct {
	Event *RoomUnpaused // Event containing the contract specifics and raw log

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
func (it *RoomUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomUnpaused)
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
		it.Event = new(RoomUnpaused)
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
func (it *RoomUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomUnpaused represents a Unpaused event raised by the Room contract.
type RoomUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Room *RoomFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RoomUnpausedIterator, error) {

	logs, sub, err := _Room.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RoomUnpausedIterator{contract: _Room.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Room *RoomFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RoomUnpaused) (event.Subscription, error) {

	logs, sub, err := _Room.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomUnpaused)
				if err := _Room.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Room *RoomFilterer) ParseUnpaused(log types.Log) (*RoomUnpaused, error) {
	event := new(RoomUnpaused)
	if err := _Room.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomFactoryMetaData contains all meta data concerning the RoomFactory contract.
var RoomFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_room\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositedValue\",\"type\":\"uint256\"}],\"name\":\"RoomCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"createRoom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3be272aa": "createRoom()",
		"8da5cb5b": "owner()",
		"5c975abb": "paused()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061001a3361002c565b6000805460ff60a01b1916905561007c565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610d628061008b6000396000f3fe60806040526004361061004a5760003560e01c80633be272aa1461004f5780635c975abb14610059578063715018a6146100885780638da5cb5b1461009d578063f2fde38b146100c5575b600080fd5b6100576100e5565b005b34801561006557600080fd5b50600054600160a01b900460ff1660405190151581526020015b60405180910390f35b34801561009457600080fd5b506100576101be565b3480156100a957600080fd5b506000546040516001600160a01b03909116815260200161007f565b3480156100d157600080fd5b506100576100e036600461034c565b610224565b600054600160a01b900460ff16156101375760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064015b60405180910390fd5b600034336040516101479061033f565b6001600160a01b0390911681526020016040518091039082f0905080158015610174573d6000803e3d6000fd5b50604080516001600160a01b038316815234602082015291925033917f6849f7a409ad97d39c5ffa074bf765330bf1a574da99d4c4831196ecd77ea8da910160405180910390a250565b6000546001600160a01b031633146102185760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161012e565b61022260006102ef565b565b6000546001600160a01b0316331461027e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161012e565b6001600160a01b0381166102e35760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161012e565b6102ec816102ef565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6109b08061037d83390190565b60006020828403121561035e57600080fd5b81356001600160a01b038116811461037557600080fd5b939250505056fe608060408190526000805460ff60a81b198116909155600180546001600160a01b0319166001600160a01b039092169190911790556109b038819003908190833981016040819052610050916100db565b6100593361008b565b6000805460ff60a01b19169055600180546001600160a01b0319166001600160a01b039290921691909117905561010b565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100ed57600080fd5b81516001600160a01b038116811461010457600080fd5b9392505050565b6108968061011a6000396000f3fe60806040526004361061009c5760003560e01c80635c975abb116100645780635c975abb14610138578063715018a6146101575780638da5cb5b1461016c5780638f13402514610194578063d0e30db0146101c4578063f2fde38b146101cc57600080fd5b806302fb0c5e146100a15780630f15f4c0146100d75780632673bd74146100ee5780633e4e10d61461010e57806351b42b0014610123575b600080fd5b3480156100ad57600080fd5b506000546100c290600160a81b900460ff1681565b60405190151581526020015b60405180910390f35b3480156100e357600080fd5b506100ec6101ec565b005b3480156100fa57600080fd5b506100ec6101093660046107bb565b61028c565b34801561011a57600080fd5b506100ec610441565b34801561012f57600080fd5b506100ec61053c565b34801561014457600080fd5b50600054600160a01b900460ff166100c2565b34801561016357600080fd5b506100ec6105d0565b34801561017857600080fd5b506000546040516001600160a01b0390911681526020016100ce565b3480156101a057600080fd5b506100c26101af3660046107f0565b60026020526000908152604090205460ff1681565b6100ec610606565b3480156101d857600080fd5b506100ec6101e7366004610809565b6106b4565b600054600160a81b900460ff16156102255760405162461bcd60e51b815260206004820152600060248201526044015b60405180910390fd5b6000546001600160a01b0316331461024f5760405162461bcd60e51b815260040161021c9061082b565b6000805460ff60a81b1916600160a81b17815560405133917ff7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba1591a2565b6000546001600160a01b031633146102b65760405162461bcd60e51b815260040161021c9061082b565b60008181526002602052604090205460ff16156102ef5760405162461bcd60e51b8152602060048201526000602482015260440161021c565b600083116103195760405162461bcd60e51b8152602060048201526000602482015260440161021c565b824710156103435760405162461bcd60e51b8152602060048201526000602482015260440161021c565b6001600160a01b0382166103735760405162461bcd60e51b8152602060048201526000602482015260440161021c565b6001546001600160a01b03838116911614156103ab5760405162461bcd60e51b8152602060048201526000602482015260440161021c565b600081815260026020526040808220805460ff19166001179055516001600160a01b0384169185156108fc02918691818181858888f193505050501580156103f7573d6000803e3d6000fd5b5060408051848152602081018390526001600160a01b038416917f6379339f0ae63e95e65fad18ca2a7ec4e7e3f05f3cc5f7079f4d8da8cec34faa910160405180910390a2505050565b600054600160a81b900460ff16156104755760405162461bcd60e51b8152602060048201526000602482015260440161021c565b6000546001600160a01b0316331461049f5760405162461bcd60e51b815260040161021c9061082b565b600047116104c95760405162461bcd60e51b8152602060048201526000602482015260440161021c565b60015460405147916001600160a01b03169082156108fc029083906000818181858888f19350505050158015610503573d6000803e3d6000fd5b5060405181815233907fa9304f0ac97f820b11d5a66f29245fb1e76b1f1f93f46a84a71682e69d095e809060200160405180910390a250565b600054600160a81b900460ff1661056f5760405162461bcd60e51b8152602060048201526000602482015260440161021c565b6000546001600160a01b031633146105995760405162461bcd60e51b815260040161021c9061082b565b6000805460ff60a81b1916815560405133917f238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c91a2565b6000546001600160a01b031633146105fa5760405162461bcd60e51b815260040161021c9061082b565b610604600061074f565b565b600054600160a01b900460ff16156106535760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161021c565b6000341161067d5760405162461bcd60e51b8152602060048201526000602482015260440161021c565b60405134815233907f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c49060200160405180910390a2565b6000546001600160a01b031633146106de5760405162461bcd60e51b815260040161021c9061082b565b6001600160a01b0381166107435760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161021c565b61074c8161074f565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80356001600160a01b03811681146107b657600080fd5b919050565b6000806000606084860312156107d057600080fd5b833592506107e06020850161079f565b9150604084013590509250925092565b60006020828403121561080257600080fd5b5035919050565b60006020828403121561081b57600080fd5b6108248261079f565b9392505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260408201526060019056fea2646970667358221220280167c5526444f7843187b65bd6989c15e4050d4b9fd3817bc59e82cef7f4c964736f6c63430008090033a26469706673582212209821515ba0fe37fe390df1c0d7d39ab91037bc88aa25fb1b0e64207b571bbbb764736f6c63430008090033",
}

// RoomFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use RoomFactoryMetaData.ABI instead.
var RoomFactoryABI = RoomFactoryMetaData.ABI

// Deprecated: Use RoomFactoryMetaData.Sigs instead.
// RoomFactoryFuncSigs maps the 4-byte function signature to its string representation.
var RoomFactoryFuncSigs = RoomFactoryMetaData.Sigs

// RoomFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoomFactoryMetaData.Bin instead.
var RoomFactoryBin = RoomFactoryMetaData.Bin

// DeployRoomFactory deploys a new Ethereum contract, binding an instance of RoomFactory to it.
func DeployRoomFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RoomFactory, error) {
	parsed, err := RoomFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoomFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoomFactory{RoomFactoryCaller: RoomFactoryCaller{contract: contract}, RoomFactoryTransactor: RoomFactoryTransactor{contract: contract}, RoomFactoryFilterer: RoomFactoryFilterer{contract: contract}}, nil
}

// RoomFactory is an auto generated Go binding around an Ethereum contract.
type RoomFactory struct {
	RoomFactoryCaller     // Read-only binding to the contract
	RoomFactoryTransactor // Write-only binding to the contract
	RoomFactoryFilterer   // Log filterer for contract events
}

// RoomFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoomFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoomFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoomFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoomFactorySession struct {
	Contract     *RoomFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoomFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoomFactoryCallerSession struct {
	Contract *RoomFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RoomFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoomFactoryTransactorSession struct {
	Contract     *RoomFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RoomFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoomFactoryRaw struct {
	Contract *RoomFactory // Generic contract binding to access the raw methods on
}

// RoomFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoomFactoryCallerRaw struct {
	Contract *RoomFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// RoomFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoomFactoryTransactorRaw struct {
	Contract *RoomFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoomFactory creates a new instance of RoomFactory, bound to a specific deployed contract.
func NewRoomFactory(address common.Address, backend bind.ContractBackend) (*RoomFactory, error) {
	contract, err := bindRoomFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoomFactory{RoomFactoryCaller: RoomFactoryCaller{contract: contract}, RoomFactoryTransactor: RoomFactoryTransactor{contract: contract}, RoomFactoryFilterer: RoomFactoryFilterer{contract: contract}}, nil
}

// NewRoomFactoryCaller creates a new read-only instance of RoomFactory, bound to a specific deployed contract.
func NewRoomFactoryCaller(address common.Address, caller bind.ContractCaller) (*RoomFactoryCaller, error) {
	contract, err := bindRoomFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryCaller{contract: contract}, nil
}

// NewRoomFactoryTransactor creates a new write-only instance of RoomFactory, bound to a specific deployed contract.
func NewRoomFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*RoomFactoryTransactor, error) {
	contract, err := bindRoomFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryTransactor{contract: contract}, nil
}

// NewRoomFactoryFilterer creates a new log filterer instance of RoomFactory, bound to a specific deployed contract.
func NewRoomFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*RoomFactoryFilterer, error) {
	contract, err := bindRoomFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryFilterer{contract: contract}, nil
}

// bindRoomFactory binds a generic wrapper to an already deployed contract.
func bindRoomFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoomFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoomFactory *RoomFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoomFactory.Contract.RoomFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoomFactory *RoomFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.Contract.RoomFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoomFactory *RoomFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoomFactory.Contract.RoomFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoomFactory *RoomFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoomFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoomFactory *RoomFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoomFactory *RoomFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoomFactory.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoomFactory *RoomFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoomFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoomFactory *RoomFactorySession) Owner() (common.Address, error) {
	return _RoomFactory.Contract.Owner(&_RoomFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoomFactory *RoomFactoryCallerSession) Owner() (common.Address, error) {
	return _RoomFactory.Contract.Owner(&_RoomFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RoomFactory *RoomFactoryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RoomFactory.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RoomFactory *RoomFactorySession) Paused() (bool, error) {
	return _RoomFactory.Contract.Paused(&_RoomFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RoomFactory *RoomFactoryCallerSession) Paused() (bool, error) {
	return _RoomFactory.Contract.Paused(&_RoomFactory.CallOpts)
}

// CreateRoom is a paid mutator transaction binding the contract method 0x3be272aa.
//
// Solidity: function createRoom() payable returns()
func (_RoomFactory *RoomFactoryTransactor) CreateRoom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "createRoom")
}

// CreateRoom is a paid mutator transaction binding the contract method 0x3be272aa.
//
// Solidity: function createRoom() payable returns()
func (_RoomFactory *RoomFactorySession) CreateRoom() (*types.Transaction, error) {
	return _RoomFactory.Contract.CreateRoom(&_RoomFactory.TransactOpts)
}

// CreateRoom is a paid mutator transaction binding the contract method 0x3be272aa.
//
// Solidity: function createRoom() payable returns()
func (_RoomFactory *RoomFactoryTransactorSession) CreateRoom() (*types.Transaction, error) {
	return _RoomFactory.Contract.CreateRoom(&_RoomFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoomFactory *RoomFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoomFactory *RoomFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _RoomFactory.Contract.RenounceOwnership(&_RoomFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoomFactory *RoomFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RoomFactory.Contract.RenounceOwnership(&_RoomFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RoomFactory *RoomFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RoomFactory *RoomFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RoomFactory.Contract.TransferOwnership(&_RoomFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RoomFactory *RoomFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RoomFactory.Contract.TransferOwnership(&_RoomFactory.TransactOpts, newOwner)
}

// RoomFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RoomFactory contract.
type RoomFactoryOwnershipTransferredIterator struct {
	Event *RoomFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RoomFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryOwnershipTransferred)
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
		it.Event = new(RoomFactoryOwnershipTransferred)
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
func (it *RoomFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the RoomFactory contract.
type RoomFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoomFactory *RoomFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RoomFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryOwnershipTransferredIterator{contract: _RoomFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoomFactory *RoomFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RoomFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryOwnershipTransferred)
				if err := _RoomFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoomFactory *RoomFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*RoomFactoryOwnershipTransferred, error) {
	event := new(RoomFactoryOwnershipTransferred)
	if err := _RoomFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomFactoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the RoomFactory contract.
type RoomFactoryPausedIterator struct {
	Event *RoomFactoryPaused // Event containing the contract specifics and raw log

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
func (it *RoomFactoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryPaused)
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
		it.Event = new(RoomFactoryPaused)
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
func (it *RoomFactoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryPaused represents a Paused event raised by the RoomFactory contract.
type RoomFactoryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RoomFactory *RoomFactoryFilterer) FilterPaused(opts *bind.FilterOpts) (*RoomFactoryPausedIterator, error) {

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RoomFactoryPausedIterator{contract: _RoomFactory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RoomFactory *RoomFactoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RoomFactoryPaused) (event.Subscription, error) {

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryPaused)
				if err := _RoomFactory.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RoomFactory *RoomFactoryFilterer) ParsePaused(log types.Log) (*RoomFactoryPaused, error) {
	event := new(RoomFactoryPaused)
	if err := _RoomFactory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomFactoryRoomCreatedIterator is returned from FilterRoomCreated and is used to iterate over the raw logs and unpacked data for RoomCreated events raised by the RoomFactory contract.
type RoomFactoryRoomCreatedIterator struct {
	Event *RoomFactoryRoomCreated // Event containing the contract specifics and raw log

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
func (it *RoomFactoryRoomCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryRoomCreated)
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
		it.Event = new(RoomFactoryRoomCreated)
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
func (it *RoomFactoryRoomCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryRoomCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryRoomCreated represents a RoomCreated event raised by the RoomFactory contract.
type RoomFactoryRoomCreated struct {
	Creator        common.Address
	Room           common.Address
	DepositedValue *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRoomCreated is a free log retrieval operation binding the contract event 0x6849f7a409ad97d39c5ffa074bf765330bf1a574da99d4c4831196ecd77ea8da.
//
// Solidity: event RoomCreated(address indexed _creator, address _room, uint256 _depositedValue)
func (_RoomFactory *RoomFactoryFilterer) FilterRoomCreated(opts *bind.FilterOpts, _creator []common.Address) (*RoomFactoryRoomCreatedIterator, error) {

	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "RoomCreated", _creatorRule)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryRoomCreatedIterator{contract: _RoomFactory.contract, event: "RoomCreated", logs: logs, sub: sub}, nil
}

// WatchRoomCreated is a free log subscription operation binding the contract event 0x6849f7a409ad97d39c5ffa074bf765330bf1a574da99d4c4831196ecd77ea8da.
//
// Solidity: event RoomCreated(address indexed _creator, address _room, uint256 _depositedValue)
func (_RoomFactory *RoomFactoryFilterer) WatchRoomCreated(opts *bind.WatchOpts, sink chan<- *RoomFactoryRoomCreated, _creator []common.Address) (event.Subscription, error) {

	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "RoomCreated", _creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryRoomCreated)
				if err := _RoomFactory.contract.UnpackLog(event, "RoomCreated", log); err != nil {
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

// ParseRoomCreated is a log parse operation binding the contract event 0x6849f7a409ad97d39c5ffa074bf765330bf1a574da99d4c4831196ecd77ea8da.
//
// Solidity: event RoomCreated(address indexed _creator, address _room, uint256 _depositedValue)
func (_RoomFactory *RoomFactoryFilterer) ParseRoomCreated(log types.Log) (*RoomFactoryRoomCreated, error) {
	event := new(RoomFactoryRoomCreated)
	if err := _RoomFactory.contract.UnpackLog(event, "RoomCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomFactoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the RoomFactory contract.
type RoomFactoryUnpausedIterator struct {
	Event *RoomFactoryUnpaused // Event containing the contract specifics and raw log

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
func (it *RoomFactoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryUnpaused)
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
		it.Event = new(RoomFactoryUnpaused)
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
func (it *RoomFactoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryUnpaused represents a Unpaused event raised by the RoomFactory contract.
type RoomFactoryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RoomFactory *RoomFactoryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RoomFactoryUnpausedIterator, error) {

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RoomFactoryUnpausedIterator{contract: _RoomFactory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RoomFactory *RoomFactoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RoomFactoryUnpaused) (event.Subscription, error) {

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryUnpaused)
				if err := _RoomFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RoomFactory *RoomFactoryFilterer) ParseUnpaused(log types.Log) (*RoomFactoryUnpaused, error) {
	event := new(RoomFactoryUnpaused)
	if err := _RoomFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
