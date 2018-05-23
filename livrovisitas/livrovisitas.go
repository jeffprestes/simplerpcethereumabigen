// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package livrovisitas

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// LivroVisitasABI is the input ABI used to generate the binding from.
const LivroVisitasABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"visitor\",\"type\":\"address\"}],\"name\":\"getMessageOfVisit\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"whoIsTheOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"visitor\",\"type\":\"address\"},{\"name\":\"message\",\"type\":\"string\"}],\"name\":\"recordVisit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"visitor\",\"type\":\"address\"}],\"name\":\"NewVisitor\",\"type\":\"event\"}]"

// LivroVisitasBin is the compiled bytecode used for deploying new contracts.
const LivroVisitasBin = `0x608060405260008054600160a060020a033316600160a060020a03199091161790556104ca806100306000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630fb66da38114610071578063220f52c51461010757806341c0e1b514610138578063780d23281461014f578063a6f9dae1146101ca575b600080fd5b34801561007d57600080fd5b50610092600160a060020a03600435166101eb565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100cc5781810151838201526020016100b4565b50505050905090810190601f1680156100f95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561011357600080fd5b5061011c6102e3565b60408051600160a060020a039092168252519081900360200190f35b34801561014457600080fd5b5061014d6102f3565b005b34801561015b57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101b6958335600160a060020a031695369560449491939091019190819084018382808284375094975061031a9650505050505050565b604080519115158252519081900360200190f35b3480156101d657600080fd5b506101b6600160a060020a03600435166103a0565b600160a060020a03811660009081526001602081905260409091205460609160026000196101008484161502019092169190910411156102cd57600160a060020a03821660009081526001602081815260409283902080548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156102c15780601f10610296576101008083540402835291602001916102c1565b820191906000526020600020905b8154815290600101906020018083116102a457829003601f168201915b505050505090506102de565b506040805160208101909152600081525b919050565b600054600160a060020a03165b90565b60005433600160a060020a039081169116141561031857600054600160a060020a0316ff5b565b6000600160a060020a038316151561033157600080fd5b600160a060020a0383166000908152600160209081526040909120835161035a92850190610406565b5060408051600160a060020a038516815290517f9a8d0466990bbfbc04d4c6bdba8806fc67473744d61b52e75c629eea10ad90ea9181900360200190a150600192915050565b6000600160a060020a03821615156103b757600080fd5b60005433600160a060020a03908116911614156103fe57506000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03831617905560016102de565b506000919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061044757805160ff1916838001178555610474565b82800160010185558215610474579182015b82811115610474578251825591602001919060010190610459565b50610480929150610484565b5090565b6102f091905b80821115610480576000815560010161048a5600a165627a7a723058202afb018ebfff2567632cc33fec6b0d0db71301b17ae098c599ec8233060d0e630029`

// DeployLivroVisitas deploys a new Ethereum contract, binding an instance of LivroVisitas to it.
func DeployLivroVisitas(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LivroVisitas, error) {
	parsed, err := abi.JSON(strings.NewReader(LivroVisitasABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LivroVisitasBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LivroVisitas{LivroVisitasCaller: LivroVisitasCaller{contract: contract}, LivroVisitasTransactor: LivroVisitasTransactor{contract: contract}, LivroVisitasFilterer: LivroVisitasFilterer{contract: contract}}, nil
}

// LivroVisitas is an auto generated Go binding around an Ethereum contract.
type LivroVisitas struct {
	LivroVisitasCaller     // Read-only binding to the contract
	LivroVisitasTransactor // Write-only binding to the contract
	LivroVisitasFilterer   // Log filterer for contract events
}

// LivroVisitasCaller is an auto generated read-only Go binding around an Ethereum contract.
type LivroVisitasCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LivroVisitasTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LivroVisitasTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LivroVisitasFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LivroVisitasFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LivroVisitasSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LivroVisitasSession struct {
	Contract     *LivroVisitas     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LivroVisitasCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LivroVisitasCallerSession struct {
	Contract *LivroVisitasCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LivroVisitasTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LivroVisitasTransactorSession struct {
	Contract     *LivroVisitasTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LivroVisitasRaw is an auto generated low-level Go binding around an Ethereum contract.
type LivroVisitasRaw struct {
	Contract *LivroVisitas // Generic contract binding to access the raw methods on
}

// LivroVisitasCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LivroVisitasCallerRaw struct {
	Contract *LivroVisitasCaller // Generic read-only contract binding to access the raw methods on
}

// LivroVisitasTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LivroVisitasTransactorRaw struct {
	Contract *LivroVisitasTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLivroVisitas creates a new instance of LivroVisitas, bound to a specific deployed contract.
func NewLivroVisitas(address common.Address, backend bind.ContractBackend) (*LivroVisitas, error) {
	contract, err := bindLivroVisitas(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LivroVisitas{LivroVisitasCaller: LivroVisitasCaller{contract: contract}, LivroVisitasTransactor: LivroVisitasTransactor{contract: contract}, LivroVisitasFilterer: LivroVisitasFilterer{contract: contract}}, nil
}

// NewLivroVisitasCaller creates a new read-only instance of LivroVisitas, bound to a specific deployed contract.
func NewLivroVisitasCaller(address common.Address, caller bind.ContractCaller) (*LivroVisitasCaller, error) {
	contract, err := bindLivroVisitas(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LivroVisitasCaller{contract: contract}, nil
}

// NewLivroVisitasTransactor creates a new write-only instance of LivroVisitas, bound to a specific deployed contract.
func NewLivroVisitasTransactor(address common.Address, transactor bind.ContractTransactor) (*LivroVisitasTransactor, error) {
	contract, err := bindLivroVisitas(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LivroVisitasTransactor{contract: contract}, nil
}

// NewLivroVisitasFilterer creates a new log filterer instance of LivroVisitas, bound to a specific deployed contract.
func NewLivroVisitasFilterer(address common.Address, filterer bind.ContractFilterer) (*LivroVisitasFilterer, error) {
	contract, err := bindLivroVisitas(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LivroVisitasFilterer{contract: contract}, nil
}

// bindLivroVisitas binds a generic wrapper to an already deployed contract.
func bindLivroVisitas(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LivroVisitasABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LivroVisitas *LivroVisitasRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LivroVisitas.Contract.LivroVisitasCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LivroVisitas *LivroVisitasRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LivroVisitas.Contract.LivroVisitasTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LivroVisitas *LivroVisitasRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LivroVisitas.Contract.LivroVisitasTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LivroVisitas *LivroVisitasCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LivroVisitas.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LivroVisitas *LivroVisitasTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LivroVisitas.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LivroVisitas *LivroVisitasTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LivroVisitas.Contract.contract.Transact(opts, method, params...)
}

// GetMessageOfVisit is a free data retrieval call binding the contract method 0x0fb66da3.
//
// Solidity: function getMessageOfVisit(visitor address) constant returns(string)
func (_LivroVisitas *LivroVisitasCaller) GetMessageOfVisit(opts *bind.CallOpts, visitor common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LivroVisitas.contract.Call(opts, out, "getMessageOfVisit", visitor)
	return *ret0, err
}

// GetMessageOfVisit is a free data retrieval call binding the contract method 0x0fb66da3.
//
// Solidity: function getMessageOfVisit(visitor address) constant returns(string)
func (_LivroVisitas *LivroVisitasSession) GetMessageOfVisit(visitor common.Address) (string, error) {
	return _LivroVisitas.Contract.GetMessageOfVisit(&_LivroVisitas.CallOpts, visitor)
}

// GetMessageOfVisit is a free data retrieval call binding the contract method 0x0fb66da3.
//
// Solidity: function getMessageOfVisit(visitor address) constant returns(string)
func (_LivroVisitas *LivroVisitasCallerSession) GetMessageOfVisit(visitor common.Address) (string, error) {
	return _LivroVisitas.Contract.GetMessageOfVisit(&_LivroVisitas.CallOpts, visitor)
}

// WhoIsTheOwner is a free data retrieval call binding the contract method 0x220f52c5.
//
// Solidity: function whoIsTheOwner() constant returns(address)
func (_LivroVisitas *LivroVisitasCaller) WhoIsTheOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LivroVisitas.contract.Call(opts, out, "whoIsTheOwner")
	return *ret0, err
}

// WhoIsTheOwner is a free data retrieval call binding the contract method 0x220f52c5.
//
// Solidity: function whoIsTheOwner() constant returns(address)
func (_LivroVisitas *LivroVisitasSession) WhoIsTheOwner() (common.Address, error) {
	return _LivroVisitas.Contract.WhoIsTheOwner(&_LivroVisitas.CallOpts)
}

// WhoIsTheOwner is a free data retrieval call binding the contract method 0x220f52c5.
//
// Solidity: function whoIsTheOwner() constant returns(address)
func (_LivroVisitas *LivroVisitasCallerSession) WhoIsTheOwner() (common.Address, error) {
	return _LivroVisitas.Contract.WhoIsTheOwner(&_LivroVisitas.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns(bool)
func (_LivroVisitas *LivroVisitasTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LivroVisitas.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns(bool)
func (_LivroVisitas *LivroVisitasSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _LivroVisitas.Contract.ChangeOwner(&_LivroVisitas.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns(bool)
func (_LivroVisitas *LivroVisitasTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _LivroVisitas.Contract.ChangeOwner(&_LivroVisitas.TransactOpts, newOwner)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_LivroVisitas *LivroVisitasTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LivroVisitas.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_LivroVisitas *LivroVisitasSession) Kill() (*types.Transaction, error) {
	return _LivroVisitas.Contract.Kill(&_LivroVisitas.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_LivroVisitas *LivroVisitasTransactorSession) Kill() (*types.Transaction, error) {
	return _LivroVisitas.Contract.Kill(&_LivroVisitas.TransactOpts)
}

// RecordVisit is a paid mutator transaction binding the contract method 0x780d2328.
//
// Solidity: function recordVisit(visitor address, message string) returns(bool)
func (_LivroVisitas *LivroVisitasTransactor) RecordVisit(opts *bind.TransactOpts, visitor common.Address, message string) (*types.Transaction, error) {
	return _LivroVisitas.contract.Transact(opts, "recordVisit", visitor, message)
}

// RecordVisit is a paid mutator transaction binding the contract method 0x780d2328.
//
// Solidity: function recordVisit(visitor address, message string) returns(bool)
func (_LivroVisitas *LivroVisitasSession) RecordVisit(visitor common.Address, message string) (*types.Transaction, error) {
	return _LivroVisitas.Contract.RecordVisit(&_LivroVisitas.TransactOpts, visitor, message)
}

// RecordVisit is a paid mutator transaction binding the contract method 0x780d2328.
//
// Solidity: function recordVisit(visitor address, message string) returns(bool)
func (_LivroVisitas *LivroVisitasTransactorSession) RecordVisit(visitor common.Address, message string) (*types.Transaction, error) {
	return _LivroVisitas.Contract.RecordVisit(&_LivroVisitas.TransactOpts, visitor, message)
}

// LivroVisitasNewVisitorIterator is returned from FilterNewVisitor and is used to iterate over the raw logs and unpacked data for NewVisitor events raised by the LivroVisitas contract.
type LivroVisitasNewVisitorIterator struct {
	Event *LivroVisitasNewVisitor // Event containing the contract specifics and raw log

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
func (it *LivroVisitasNewVisitorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LivroVisitasNewVisitor)
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
		it.Event = new(LivroVisitasNewVisitor)
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
func (it *LivroVisitasNewVisitorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LivroVisitasNewVisitorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LivroVisitasNewVisitor represents a NewVisitor event raised by the LivroVisitas contract.
type LivroVisitasNewVisitor struct {
	Visitor common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewVisitor is a free log retrieval operation binding the contract event 0x9a8d0466990bbfbc04d4c6bdba8806fc67473744d61b52e75c629eea10ad90ea.
//
// Solidity: e NewVisitor(visitor address)
func (_LivroVisitas *LivroVisitasFilterer) FilterNewVisitor(opts *bind.FilterOpts) (*LivroVisitasNewVisitorIterator, error) {

	logs, sub, err := _LivroVisitas.contract.FilterLogs(opts, "NewVisitor")
	if err != nil {
		return nil, err
	}
	return &LivroVisitasNewVisitorIterator{contract: _LivroVisitas.contract, event: "NewVisitor", logs: logs, sub: sub}, nil
}

// WatchNewVisitor is a free log subscription operation binding the contract event 0x9a8d0466990bbfbc04d4c6bdba8806fc67473744d61b52e75c629eea10ad90ea.
//
// Solidity: e NewVisitor(visitor address)
func (_LivroVisitas *LivroVisitasFilterer) WatchNewVisitor(opts *bind.WatchOpts, sink chan<- *LivroVisitasNewVisitor) (event.Subscription, error) {

	logs, sub, err := _LivroVisitas.contract.WatchLogs(opts, "NewVisitor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				dataEvent := new(common.Address)
				if err := _LivroVisitas.contract.UnpackLog(dataEvent, "NewVisitor", log); err != nil {
					return err
				}
				event := new(LivroVisitasNewVisitor)
				event.Raw = log
				event.Visitor = *dataEvent
				//fmt.Printf("log watcher: %+v\n", event)
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

// MortalABI is the input ABI used to generate the binding from.
const MortalABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"whoIsTheOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MortalBin is the compiled bytecode used for deploying new contracts.
const MortalBin = `0x608060405260008054600160a060020a033316600160a060020a03199091161790556101a1806100306000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663220f52c5811461005b57806341c0e1b51461008c578063a6f9dae1146100a3575b600080fd5b34801561006757600080fd5b506100706100d8565b60408051600160a060020a039092168252519081900360200190f35b34801561009857600080fd5b506100a16100e7565b005b3480156100af57600080fd5b506100c4600160a060020a036004351661010e565b604080519115158252519081900360200190f35b600054600160a060020a031690565b60005433600160a060020a039081169116141561010c57600054600160a060020a0316ff5b565b6000600160a060020a038216151561012557600080fd5b60005433600160a060020a039081169116141561016c57506000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790556001610170565b5060005b9190505600a165627a7a72305820d466bcf3da1862a746570ff0e258d401526e953ee145f8dad88ee1ecc82323110029`

// DeployMortal deploys a new Ethereum contract, binding an instance of Mortal to it.
func DeployMortal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mortal, error) {
	parsed, err := abi.JSON(strings.NewReader(MortalABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MortalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mortal{MortalCaller: MortalCaller{contract: contract}, MortalTransactor: MortalTransactor{contract: contract}, MortalFilterer: MortalFilterer{contract: contract}}, nil
}

// Mortal is an auto generated Go binding around an Ethereum contract.
type Mortal struct {
	MortalCaller     // Read-only binding to the contract
	MortalTransactor // Write-only binding to the contract
	MortalFilterer   // Log filterer for contract events
}

// MortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type MortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MortalSession struct {
	Contract     *Mortal           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MortalCallerSession struct {
	Contract *MortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MortalTransactorSession struct {
	Contract     *MortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type MortalRaw struct {
	Contract *Mortal // Generic contract binding to access the raw methods on
}

// MortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MortalCallerRaw struct {
	Contract *MortalCaller // Generic read-only contract binding to access the raw methods on
}

// MortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MortalTransactorRaw struct {
	Contract *MortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMortal creates a new instance of Mortal, bound to a specific deployed contract.
func NewMortal(address common.Address, backend bind.ContractBackend) (*Mortal, error) {
	contract, err := bindMortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mortal{MortalCaller: MortalCaller{contract: contract}, MortalTransactor: MortalTransactor{contract: contract}, MortalFilterer: MortalFilterer{contract: contract}}, nil
}

// NewMortalCaller creates a new read-only instance of Mortal, bound to a specific deployed contract.
func NewMortalCaller(address common.Address, caller bind.ContractCaller) (*MortalCaller, error) {
	contract, err := bindMortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MortalCaller{contract: contract}, nil
}

// NewMortalTransactor creates a new write-only instance of Mortal, bound to a specific deployed contract.
func NewMortalTransactor(address common.Address, transactor bind.ContractTransactor) (*MortalTransactor, error) {
	contract, err := bindMortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MortalTransactor{contract: contract}, nil
}

// NewMortalFilterer creates a new log filterer instance of Mortal, bound to a specific deployed contract.
func NewMortalFilterer(address common.Address, filterer bind.ContractFilterer) (*MortalFilterer, error) {
	contract, err := bindMortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MortalFilterer{contract: contract}, nil
}

// bindMortal binds a generic wrapper to an already deployed contract.
func bindMortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MortalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mortal *MortalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mortal.Contract.MortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mortal *MortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortal.Contract.MortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mortal *MortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mortal.Contract.MortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mortal *MortalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mortal *MortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mortal *MortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mortal.Contract.contract.Transact(opts, method, params...)
}

// WhoIsTheOwner is a free data retrieval call binding the contract method 0x220f52c5.
//
// Solidity: function whoIsTheOwner() constant returns(address)
func (_Mortal *MortalCaller) WhoIsTheOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mortal.contract.Call(opts, out, "whoIsTheOwner")
	return *ret0, err
}

// WhoIsTheOwner is a free data retrieval call binding the contract method 0x220f52c5.
//
// Solidity: function whoIsTheOwner() constant returns(address)
func (_Mortal *MortalSession) WhoIsTheOwner() (common.Address, error) {
	return _Mortal.Contract.WhoIsTheOwner(&_Mortal.CallOpts)
}

// WhoIsTheOwner is a free data retrieval call binding the contract method 0x220f52c5.
//
// Solidity: function whoIsTheOwner() constant returns(address)
func (_Mortal *MortalCallerSession) WhoIsTheOwner() (common.Address, error) {
	return _Mortal.Contract.WhoIsTheOwner(&_Mortal.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns(bool)
func (_Mortal *MortalTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mortal.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns(bool)
func (_Mortal *MortalSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Mortal.Contract.ChangeOwner(&_Mortal.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns(bool)
func (_Mortal *MortalTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Mortal.Contract.ChangeOwner(&_Mortal.TransactOpts, newOwner)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Mortal *MortalTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortal.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Mortal *MortalSession) Kill() (*types.Transaction, error) {
	return _Mortal.Contract.Kill(&_Mortal.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Mortal *MortalTransactorSession) Kill() (*types.Transaction, error) {
	return _Mortal.Contract.Kill(&_Mortal.TransactOpts)
}

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"whoIsTheOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a033316600160a060020a03199091161790556101a68061003d6000396000f30060806040526004361061004b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663220f52c58114610050578063a6f9dae11461008e575b600080fd5b34801561005c57600080fd5b506100656100d0565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561009a57600080fd5b506100bc73ffffffffffffffffffffffffffffffffffffffff600435166100ec565b604080519115158252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1690565b600073ffffffffffffffffffffffffffffffffffffffff8216151561011057600080fd5b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116141561017157506000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790556001610175565b5060005b9190505600a165627a7a72305820e2593ab12f07977a8b3a047b37e52c87b56699eefc4225ff8bdd8e705f85c1130029`

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
	OwnedFilterer   // Log filterer for contract events
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// NewOwnedFilterer creates a new log filterer instance of Owned, bound to a specific deployed contract.
func NewOwnedFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnedFilterer, error) {
	contract, err := bindOwned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnedFilterer{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// WhoIsTheOwner is a free data retrieval call binding the contract method 0x220f52c5.
//
// Solidity: function whoIsTheOwner() constant returns(address)
func (_Owned *OwnedCaller) WhoIsTheOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.Call(opts, out, "whoIsTheOwner")
	return *ret0, err
}

// WhoIsTheOwner is a free data retrieval call binding the contract method 0x220f52c5.
//
// Solidity: function whoIsTheOwner() constant returns(address)
func (_Owned *OwnedSession) WhoIsTheOwner() (common.Address, error) {
	return _Owned.Contract.WhoIsTheOwner(&_Owned.CallOpts)
}

// WhoIsTheOwner is a free data retrieval call binding the contract method 0x220f52c5.
//
// Solidity: function whoIsTheOwner() constant returns(address)
func (_Owned *OwnedCallerSession) WhoIsTheOwner() (common.Address, error) {
	return _Owned.Contract.WhoIsTheOwner(&_Owned.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns(bool)
func (_Owned *OwnedTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns(bool)
func (_Owned *OwnedSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns(bool)
func (_Owned *OwnedTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, newOwner)
}
