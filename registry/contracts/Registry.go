// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package listener

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ListenerABI is the input ABI used to generate the binding from.
const ListenerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_accountToRemove\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_accountToWhiteList\",\"type\":\"address\"}],\"name\":\"whiteList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_accountChecked\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"name\":\"whitelisted\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_whiteListedAccount\",\"type\":\"address\"}],\"name\":\"_WhiteList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_removedAccount\",\"type\":\"address\"}],\"name\":\"_Remove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"applicant\",\"type\":\"address\"}],\"name\":\"_Application\",\"type\":\"event\"}]"

// Listener is an auto generated Go binding around an Ethereum contract.
type Listener struct {
	ListenerCaller     // Read-only binding to the contract
	ListenerTransactor // Write-only binding to the contract
	ListenerFilterer   // Log filterer for contract events
}

// ListenerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ListenerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListenerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ListenerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListenerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ListenerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListenerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ListenerSession struct {
	Contract     *Listener         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ListenerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ListenerCallerSession struct {
	Contract *ListenerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ListenerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ListenerTransactorSession struct {
	Contract     *ListenerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ListenerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ListenerRaw struct {
	Contract *Listener // Generic contract binding to access the raw methods on
}

// ListenerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ListenerCallerRaw struct {
	Contract *ListenerCaller // Generic read-only contract binding to access the raw methods on
}

// ListenerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ListenerTransactorRaw struct {
	Contract *ListenerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewListener creates a new instance of Listener, bound to a specific deployed contract.
func NewListener(address common.Address, backend bind.ContractBackend) (*Listener, error) {
	contract, err := bindListener(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Listener{ListenerCaller: ListenerCaller{contract: contract}, ListenerTransactor: ListenerTransactor{contract: contract}, ListenerFilterer: ListenerFilterer{contract: contract}}, nil
}

// NewListenerCaller creates a new read-only instance of Listener, bound to a specific deployed contract.
func NewListenerCaller(address common.Address, caller bind.ContractCaller) (*ListenerCaller, error) {
	contract, err := bindListener(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ListenerCaller{contract: contract}, nil
}

// NewListenerTransactor creates a new write-only instance of Listener, bound to a specific deployed contract.
func NewListenerTransactor(address common.Address, transactor bind.ContractTransactor) (*ListenerTransactor, error) {
	contract, err := bindListener(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ListenerTransactor{contract: contract}, nil
}

// NewListenerFilterer creates a new log filterer instance of Listener, bound to a specific deployed contract.
func NewListenerFilterer(address common.Address, filterer bind.ContractFilterer) (*ListenerFilterer, error) {
	contract, err := bindListener(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ListenerFilterer{contract: contract}, nil
}

// bindListener binds a generic wrapper to an already deployed contract.
func bindListener(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ListenerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Listener *ListenerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Listener.Contract.ListenerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Listener *ListenerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Listener.Contract.ListenerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Listener *ListenerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Listener.Contract.ListenerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Listener *ListenerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Listener.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Listener *ListenerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Listener.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Listener *ListenerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Listener.Contract.contract.Transact(opts, method, params...)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _accountChecked) constant returns(bool whitelisted)
func (_Listener *ListenerCaller) IsWhitelisted(opts *bind.CallOpts, _accountChecked common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Listener.contract.Call(opts, out, "isWhitelisted", _accountChecked)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _accountChecked) constant returns(bool whitelisted)
func (_Listener *ListenerSession) IsWhitelisted(_accountChecked common.Address) (bool, error) {
	return _Listener.Contract.IsWhitelisted(&_Listener.CallOpts, _accountChecked)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _accountChecked) constant returns(bool whitelisted)
func (_Listener *ListenerCallerSession) IsWhitelisted(_accountChecked common.Address) (bool, error) {
	return _Listener.Contract.IsWhitelisted(&_Listener.CallOpts, _accountChecked)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address _accountToRemove) returns()
func (_Listener *ListenerTransactor) Remove(opts *bind.TransactOpts, _accountToRemove common.Address) (*types.Transaction, error) {
	return _Listener.contract.Transact(opts, "remove", _accountToRemove)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address _accountToRemove) returns()
func (_Listener *ListenerSession) Remove(_accountToRemove common.Address) (*types.Transaction, error) {
	return _Listener.Contract.Remove(&_Listener.TransactOpts, _accountToRemove)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address _accountToRemove) returns()
func (_Listener *ListenerTransactorSession) Remove(_accountToRemove common.Address) (*types.Transaction, error) {
	return _Listener.Contract.Remove(&_Listener.TransactOpts, _accountToRemove)
}

// WhiteList is a paid mutator transaction binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address _accountToWhiteList) returns()
func (_Listener *ListenerTransactor) WhiteList(opts *bind.TransactOpts, _accountToWhiteList common.Address) (*types.Transaction, error) {
	return _Listener.contract.Transact(opts, "whiteList", _accountToWhiteList)
}

// WhiteList is a paid mutator transaction binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address _accountToWhiteList) returns()
func (_Listener *ListenerSession) WhiteList(_accountToWhiteList common.Address) (*types.Transaction, error) {
	return _Listener.Contract.WhiteList(&_Listener.TransactOpts, _accountToWhiteList)
}

// WhiteList is a paid mutator transaction binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address _accountToWhiteList) returns()
func (_Listener *ListenerTransactorSession) WhiteList(_accountToWhiteList common.Address) (*types.Transaction, error) {
	return _Listener.Contract.WhiteList(&_Listener.TransactOpts, _accountToWhiteList)
}

// ListenerApplicationIterator is returned from FilterApplication and is used to iterate over the raw logs and unpacked data for Application events raised by the Listener contract.
type ListenerApplicationIterator struct {
	Event *ListenerApplication // Event containing the contract specifics and raw log

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
func (it *ListenerApplicationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListenerApplication)
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
		it.Event = new(ListenerApplication)
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
func (it *ListenerApplicationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListenerApplicationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListenerApplication represents a Application event raised by the Listener contract.
type ListenerApplication struct {
	ListingHash [32]byte
	Deposit     *big.Int
	Data        string
	Applicant   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterApplication is a free log retrieval operation binding the contract event 0xe1cdd401be51ef2c8159945a96274b69ba482f2a5e2131700beeda434fd6d58f.
//
// Solidity: event _Application(bytes32 indexed listingHash, uint256 deposit, string data, address indexed applicant)
func (_Listener *ListenerFilterer) FilterApplication(opts *bind.FilterOpts, listingHash [][32]byte, applicant []common.Address) (*ListenerApplicationIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	var applicantRule []interface{}
	for _, applicantItem := range applicant {
		applicantRule = append(applicantRule, applicantItem)
	}

	logs, sub, err := _Listener.contract.FilterLogs(opts, "_Application", listingHashRule, applicantRule)
	if err != nil {
		return nil, err
	}
	return &ListenerApplicationIterator{contract: _Listener.contract, event: "_Application", logs: logs, sub: sub}, nil
}

// WatchApplication is a free log subscription operation binding the contract event 0xe1cdd401be51ef2c8159945a96274b69ba482f2a5e2131700beeda434fd6d58f.
//
// Solidity: event _Application(bytes32 indexed listingHash, uint256 deposit, string data, address indexed applicant)
func (_Listener *ListenerFilterer) WatchApplication(opts *bind.WatchOpts, sink chan<- *ListenerApplication, listingHash [][32]byte, applicant []common.Address) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	var applicantRule []interface{}
	for _, applicantItem := range applicant {
		applicantRule = append(applicantRule, applicantItem)
	}

	logs, sub, err := _Listener.contract.WatchLogs(opts, "_Application", listingHashRule, applicantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListenerApplication)
				if err := _Listener.contract.UnpackLog(event, "_Application", log); err != nil {
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

// ListenerRemoveIterator is returned from FilterRemove and is used to iterate over the raw logs and unpacked data for Remove events raised by the Listener contract.
type ListenerRemoveIterator struct {
	Event *ListenerRemove // Event containing the contract specifics and raw log

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
func (it *ListenerRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListenerRemove)
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
		it.Event = new(ListenerRemove)
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
func (it *ListenerRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListenerRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListenerRemove represents a Remove event raised by the Listener contract.
type ListenerRemove struct {
	RemovedAccount common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRemove is a free log retrieval operation binding the contract event 0x4dad83fa2f662e801184bd81b75b1894fdfd16026aa67a6a65d5715476a61f53.
//
// Solidity: event _Remove(address _removedAccount)
func (_Listener *ListenerFilterer) FilterRemove(opts *bind.FilterOpts) (*ListenerRemoveIterator, error) {

	logs, sub, err := _Listener.contract.FilterLogs(opts, "_Remove")
	if err != nil {
		return nil, err
	}
	return &ListenerRemoveIterator{contract: _Listener.contract, event: "_Remove", logs: logs, sub: sub}, nil
}

// WatchRemove is a free log subscription operation binding the contract event 0x4dad83fa2f662e801184bd81b75b1894fdfd16026aa67a6a65d5715476a61f53.
//
// Solidity: event _Remove(address _removedAccount)
func (_Listener *ListenerFilterer) WatchRemove(opts *bind.WatchOpts, sink chan<- *ListenerRemove) (event.Subscription, error) {

	logs, sub, err := _Listener.contract.WatchLogs(opts, "_Remove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListenerRemove)
				if err := _Listener.contract.UnpackLog(event, "_Remove", log); err != nil {
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

// ListenerWhiteListIterator is returned from FilterWhiteList and is used to iterate over the raw logs and unpacked data for WhiteList events raised by the Listener contract.
type ListenerWhiteListIterator struct {
	Event *ListenerWhiteList // Event containing the contract specifics and raw log

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
func (it *ListenerWhiteListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListenerWhiteList)
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
		it.Event = new(ListenerWhiteList)
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
func (it *ListenerWhiteListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListenerWhiteListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListenerWhiteList represents a WhiteList event raised by the Listener contract.
type ListenerWhiteList struct {
	WhiteListedAccount common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWhiteList is a free log retrieval operation binding the contract event 0x14c99c95843c55c4e6e3cd95be615861493e90f3919ff664169bfbe2839dbfb0.
//
// Solidity: event _WhiteList(address _whiteListedAccount)
func (_Listener *ListenerFilterer) FilterWhiteList(opts *bind.FilterOpts) (*ListenerWhiteListIterator, error) {

	logs, sub, err := _Listener.contract.FilterLogs(opts, "_WhiteList")
	if err != nil {
		return nil, err
	}
	return &ListenerWhiteListIterator{contract: _Listener.contract, event: "_WhiteList", logs: logs, sub: sub}, nil
}

// WatchWhiteList is a free log subscription operation binding the contract event 0x14c99c95843c55c4e6e3cd95be615861493e90f3919ff664169bfbe2839dbfb0.
//
// Solidity: event _WhiteList(address _whiteListedAccount)
func (_Listener *ListenerFilterer) WatchWhiteList(opts *bind.WatchOpts, sink chan<- *ListenerWhiteList) (event.Subscription, error) {

	logs, sub, err := _Listener.contract.WatchLogs(opts, "_WhiteList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListenerWhiteList)
				if err := _Listener.contract.UnpackLog(event, "_WhiteList", log); err != nil {
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
