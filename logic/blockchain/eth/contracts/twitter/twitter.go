// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package twitter

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

// TwitterMetaData contains all meta data concerning the Twitter contract.
var TwitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callbackAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"callbackFunctionSignature\",\"type\":\"bytes4\"}],\"name\":\"TwitterFollowerNumRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"followerCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listedCount\",\"type\":\"uint256\"}],\"name\":\"handleCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"followerCountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listedCountThreshold\",\"type\":\"uint256\"}],\"name\":\"requestAsyncCall\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdToRequest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"userName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"followerCountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listedCountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRequestCompleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"results\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TwitterABI is the input ABI used to generate the binding from.
// Deprecated: Use TwitterMetaData.ABI instead.
var TwitterABI = TwitterMetaData.ABI

// Twitter is an auto generated Go binding around an Ethereum contract.
type Twitter struct {
	TwitterCaller     // Read-only binding to the contract
	TwitterTransactor // Write-only binding to the contract
	TwitterFilterer   // Log filterer for contract events
}

// TwitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TwitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TwitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TwitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TwitterSession struct {
	Contract     *Twitter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TwitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TwitterCallerSession struct {
	Contract *TwitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TwitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TwitterTransactorSession struct {
	Contract     *TwitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TwitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TwitterRaw struct {
	Contract *Twitter // Generic contract binding to access the raw methods on
}

// TwitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TwitterCallerRaw struct {
	Contract *TwitterCaller // Generic read-only contract binding to access the raw methods on
}

// TwitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TwitterTransactorRaw struct {
	Contract *TwitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTwitter creates a new instance of Twitter, bound to a specific deployed contract.
func NewTwitter(address common.Address, backend bind.ContractBackend) (*Twitter, error) {
	contract, err := bindTwitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Twitter{TwitterCaller: TwitterCaller{contract: contract}, TwitterTransactor: TwitterTransactor{contract: contract}, TwitterFilterer: TwitterFilterer{contract: contract}}, nil
}

// NewTwitterCaller creates a new read-only instance of Twitter, bound to a specific deployed contract.
func NewTwitterCaller(address common.Address, caller bind.ContractCaller) (*TwitterCaller, error) {
	contract, err := bindTwitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TwitterCaller{contract: contract}, nil
}

// NewTwitterTransactor creates a new write-only instance of Twitter, bound to a specific deployed contract.
func NewTwitterTransactor(address common.Address, transactor bind.ContractTransactor) (*TwitterTransactor, error) {
	contract, err := bindTwitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TwitterTransactor{contract: contract}, nil
}

// NewTwitterFilterer creates a new log filterer instance of Twitter, bound to a specific deployed contract.
func NewTwitterFilterer(address common.Address, filterer bind.ContractFilterer) (*TwitterFilterer, error) {
	contract, err := bindTwitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TwitterFilterer{contract: contract}, nil
}

// bindTwitter binds a generic wrapper to an already deployed contract.
func bindTwitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TwitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Twitter *TwitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Twitter.Contract.TwitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Twitter *TwitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Twitter.Contract.TwitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Twitter *TwitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Twitter.Contract.TwitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Twitter *TwitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Twitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Twitter *TwitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Twitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Twitter *TwitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Twitter.Contract.contract.Transact(opts, method, params...)
}

// RequestIdToRequest is a free data retrieval call binding the contract method 0x01e5bbb9.
//
// Solidity: function requestIdToRequest(bytes32 ) view returns(address requester, string userName, uint256 followerCountThreshold, uint256 listedCountThreshold, bool isRequestCompleted, bool exists)
func (_Twitter *TwitterCaller) RequestIdToRequest(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Requester              common.Address
	UserName               string
	FollowerCountThreshold *big.Int
	ListedCountThreshold   *big.Int
	IsRequestCompleted     bool
	Exists                 bool
}, error) {
	var out []interface{}
	err := _Twitter.contract.Call(opts, &out, "requestIdToRequest", arg0)

	outstruct := new(struct {
		Requester              common.Address
		UserName               string
		FollowerCountThreshold *big.Int
		ListedCountThreshold   *big.Int
		IsRequestCompleted     bool
		Exists                 bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requester = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.UserName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.FollowerCountThreshold = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ListedCountThreshold = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsRequestCompleted = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Exists = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// RequestIdToRequest is a free data retrieval call binding the contract method 0x01e5bbb9.
//
// Solidity: function requestIdToRequest(bytes32 ) view returns(address requester, string userName, uint256 followerCountThreshold, uint256 listedCountThreshold, bool isRequestCompleted, bool exists)
func (_Twitter *TwitterSession) RequestIdToRequest(arg0 [32]byte) (struct {
	Requester              common.Address
	UserName               string
	FollowerCountThreshold *big.Int
	ListedCountThreshold   *big.Int
	IsRequestCompleted     bool
	Exists                 bool
}, error) {
	return _Twitter.Contract.RequestIdToRequest(&_Twitter.CallOpts, arg0)
}

// RequestIdToRequest is a free data retrieval call binding the contract method 0x01e5bbb9.
//
// Solidity: function requestIdToRequest(bytes32 ) view returns(address requester, string userName, uint256 followerCountThreshold, uint256 listedCountThreshold, bool isRequestCompleted, bool exists)
func (_Twitter *TwitterCallerSession) RequestIdToRequest(arg0 [32]byte) (struct {
	Requester              common.Address
	UserName               string
	FollowerCountThreshold *big.Int
	ListedCountThreshold   *big.Int
	IsRequestCompleted     bool
	Exists                 bool
}, error) {
	return _Twitter.Contract.RequestIdToRequest(&_Twitter.CallOpts, arg0)
}

// Results is a free data retrieval call binding the contract method 0x4c6b25b1.
//
// Solidity: function results(bytes32 ) view returns(string)
func (_Twitter *TwitterCaller) Results(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _Twitter.contract.Call(opts, &out, "results", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Results is a free data retrieval call binding the contract method 0x4c6b25b1.
//
// Solidity: function results(bytes32 ) view returns(string)
func (_Twitter *TwitterSession) Results(arg0 [32]byte) (string, error) {
	return _Twitter.Contract.Results(&_Twitter.CallOpts, arg0)
}

// Results is a free data retrieval call binding the contract method 0x4c6b25b1.
//
// Solidity: function results(bytes32 ) view returns(string)
func (_Twitter *TwitterCallerSession) Results(arg0 [32]byte) (string, error) {
	return _Twitter.Contract.Results(&_Twitter.CallOpts, arg0)
}

// HandleCallback is a paid mutator transaction binding the contract method 0xb707bacd.
//
// Solidity: function handleCallback(bytes32 requestId, uint256 followerCount, uint256 listedCount) returns()
func (_Twitter *TwitterTransactor) HandleCallback(opts *bind.TransactOpts, requestId [32]byte, followerCount *big.Int, listedCount *big.Int) (*types.Transaction, error) {
	return _Twitter.contract.Transact(opts, "handleCallback", requestId, followerCount, listedCount)
}

// HandleCallback is a paid mutator transaction binding the contract method 0xb707bacd.
//
// Solidity: function handleCallback(bytes32 requestId, uint256 followerCount, uint256 listedCount) returns()
func (_Twitter *TwitterSession) HandleCallback(requestId [32]byte, followerCount *big.Int, listedCount *big.Int) (*types.Transaction, error) {
	return _Twitter.Contract.HandleCallback(&_Twitter.TransactOpts, requestId, followerCount, listedCount)
}

// HandleCallback is a paid mutator transaction binding the contract method 0xb707bacd.
//
// Solidity: function handleCallback(bytes32 requestId, uint256 followerCount, uint256 listedCount) returns()
func (_Twitter *TwitterTransactorSession) HandleCallback(requestId [32]byte, followerCount *big.Int, listedCount *big.Int) (*types.Transaction, error) {
	return _Twitter.Contract.HandleCallback(&_Twitter.TransactOpts, requestId, followerCount, listedCount)
}

// RequestAsyncCall is a paid mutator transaction binding the contract method 0x19e1600b.
//
// Solidity: function requestAsyncCall(string userName, uint256 followerCountThreshold, uint256 listedCountThreshold) returns(bytes32 requestId)
func (_Twitter *TwitterTransactor) RequestAsyncCall(opts *bind.TransactOpts, userName string, followerCountThreshold *big.Int, listedCountThreshold *big.Int) (*types.Transaction, error) {
	return _Twitter.contract.Transact(opts, "requestAsyncCall", userName, followerCountThreshold, listedCountThreshold)
}

// RequestAsyncCall is a paid mutator transaction binding the contract method 0x19e1600b.
//
// Solidity: function requestAsyncCall(string userName, uint256 followerCountThreshold, uint256 listedCountThreshold) returns(bytes32 requestId)
func (_Twitter *TwitterSession) RequestAsyncCall(userName string, followerCountThreshold *big.Int, listedCountThreshold *big.Int) (*types.Transaction, error) {
	return _Twitter.Contract.RequestAsyncCall(&_Twitter.TransactOpts, userName, followerCountThreshold, listedCountThreshold)
}

// RequestAsyncCall is a paid mutator transaction binding the contract method 0x19e1600b.
//
// Solidity: function requestAsyncCall(string userName, uint256 followerCountThreshold, uint256 listedCountThreshold) returns(bytes32 requestId)
func (_Twitter *TwitterTransactorSession) RequestAsyncCall(userName string, followerCountThreshold *big.Int, listedCountThreshold *big.Int) (*types.Transaction, error) {
	return _Twitter.Contract.RequestAsyncCall(&_Twitter.TransactOpts, userName, followerCountThreshold, listedCountThreshold)
}

// TwitterTwitterFollowerNumRequestIterator is returned from FilterTwitterFollowerNumRequest and is used to iterate over the raw logs and unpacked data for TwitterFollowerNumRequest events raised by the Twitter contract.
type TwitterTwitterFollowerNumRequestIterator struct {
	Event *TwitterTwitterFollowerNumRequest // Event containing the contract specifics and raw log

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
func (it *TwitterTwitterFollowerNumRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwitterTwitterFollowerNumRequest)
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
		it.Event = new(TwitterTwitterFollowerNumRequest)
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
func (it *TwitterTwitterFollowerNumRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwitterTwitterFollowerNumRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwitterTwitterFollowerNumRequest represents a TwitterFollowerNumRequest event raised by the Twitter contract.
type TwitterTwitterFollowerNumRequest struct {
	RequestId                 [32]byte
	Requester                 common.Address
	Username                  string
	CallbackAddr              common.Address
	CallbackFunctionSignature [4]byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterTwitterFollowerNumRequest is a free log retrieval operation binding the contract event 0x2d7ca72a8e2c202dd595ddc87c781c7c8665ed662cdae9bf8472083937ff0873.
//
// Solidity: event TwitterFollowerNumRequest(bytes32 indexed requestId, address indexed requester, string username, address callbackAddr, bytes4 callbackFunctionSignature)
func (_Twitter *TwitterFilterer) FilterTwitterFollowerNumRequest(opts *bind.FilterOpts, requestId [][32]byte, requester []common.Address) (*TwitterTwitterFollowerNumRequestIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Twitter.contract.FilterLogs(opts, "TwitterFollowerNumRequest", requestIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &TwitterTwitterFollowerNumRequestIterator{contract: _Twitter.contract, event: "TwitterFollowerNumRequest", logs: logs, sub: sub}, nil
}

// WatchTwitterFollowerNumRequest is a free log subscription operation binding the contract event 0x2d7ca72a8e2c202dd595ddc87c781c7c8665ed662cdae9bf8472083937ff0873.
//
// Solidity: event TwitterFollowerNumRequest(bytes32 indexed requestId, address indexed requester, string username, address callbackAddr, bytes4 callbackFunctionSignature)
func (_Twitter *TwitterFilterer) WatchTwitterFollowerNumRequest(opts *bind.WatchOpts, sink chan<- *TwitterTwitterFollowerNumRequest, requestId [][32]byte, requester []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Twitter.contract.WatchLogs(opts, "TwitterFollowerNumRequest", requestIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwitterTwitterFollowerNumRequest)
				if err := _Twitter.contract.UnpackLog(event, "TwitterFollowerNumRequest", log); err != nil {
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

// ParseTwitterFollowerNumRequest is a log parse operation binding the contract event 0x2d7ca72a8e2c202dd595ddc87c781c7c8665ed662cdae9bf8472083937ff0873.
//
// Solidity: event TwitterFollowerNumRequest(bytes32 indexed requestId, address indexed requester, string username, address callbackAddr, bytes4 callbackFunctionSignature)
func (_Twitter *TwitterFilterer) ParseTwitterFollowerNumRequest(log types.Log) (*TwitterTwitterFollowerNumRequest, error) {
	event := new(TwitterTwitterFollowerNumRequest)
	if err := _Twitter.contract.UnpackLog(event, "TwitterFollowerNumRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
