// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractServiceManagerBase

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

// IAllocationManagerTypesCreateSetParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesCreateSetParams struct {
	OperatorSetId uint32
	Strategies    []common.Address
}

// IAllocationManagerTypesSlashingParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesSlashingParams struct {
	Operator      common.Address
	OperatorSetId uint32
	Strategies    []common.Address
	WadsToSlash   []*big.Int
	Description   string
}

// IRewardsCoordinatorTypesRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorTypesStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IRewardsCoordinatorTypesStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractServiceManagerBaseMetaData contains all meta data concerning the ContractServiceManagerBase contract.
var ContractServiceManagerBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"SLASHER_PROPOSAL_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptProposedSlasher\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategyToOperatorSet\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParams[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrationFinalized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeNewSlasher\",\"inputs\":[{\"name\":\"newSlasher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proposedSlasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsInitiator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAVSRegistrar\",\"inputs\":[{\"name\":\"registrar\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsInitiator\",\"inputs\":[{\"name\":\"newRewardsInitiator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadsToSlash\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slasherProposalTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsInitiatorUpdated\",\"inputs\":[{\"name\":\"prevRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlasherProposed\",\"inputs\":[{\"name\":\"newSlasher\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"slasherProposalTimestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlasherUpdated\",\"inputs\":[{\"name\":\"prevSlasher\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newSlasher\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// ContractServiceManagerBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractServiceManagerBaseMetaData.ABI instead.
var ContractServiceManagerBaseABI = ContractServiceManagerBaseMetaData.ABI

// ContractServiceManagerBaseMethods is an auto generated interface around an Ethereum contract.
type ContractServiceManagerBaseMethods interface {
	ContractServiceManagerBaseCalls
	ContractServiceManagerBaseTransacts
	ContractServiceManagerBaseFilters
}

// ContractServiceManagerBaseCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractServiceManagerBaseCalls interface {
	SLASHERPROPOSALDELAY(opts *bind.CallOpts) (*big.Int, error)

	AllocationManager(opts *bind.CallOpts) (common.Address, error)

	AvsDirectory(opts *bind.CallOpts) (common.Address, error)

	GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error)

	GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error)

	MigrationFinalized(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	ProposedSlasher(opts *bind.CallOpts) (common.Address, error)

	RewardsInitiator(opts *bind.CallOpts) (common.Address, error)

	Slasher(opts *bind.CallOpts) (common.Address, error)

	SlasherProposalTimestamp(opts *bind.CallOpts) (*big.Int, error)
}

// ContractServiceManagerBaseTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractServiceManagerBaseTransacts interface {
	AcceptProposedSlasher(opts *bind.TransactOpts) (*types.Transaction, error)

	AddStrategyToOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error)

	CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error)

	CreateOperatorSets(opts *bind.TransactOpts, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error)

	DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error)

	DeregisterOperatorFromOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error)

	ProposeNewSlasher(opts *bind.TransactOpts, newSlasher common.Address) (*types.Transaction, error)

	RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RemoveStrategiesFromOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAVSRegistrar(opts *bind.TransactOpts, registrar common.Address) (*types.Transaction, error)

	SetRewardsInitiator(opts *bind.TransactOpts, newRewardsInitiator common.Address) (*types.Transaction, error)

	SlashOperator(opts *bind.TransactOpts, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error)
}

// ContractServiceManagerBaseFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractServiceManagerBaseFilters interface {
	FilterInitialized(opts *bind.FilterOpts) (*ContractServiceManagerBaseInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractServiceManagerBaseInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractServiceManagerBaseInitialized, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractServiceManagerBaseOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractServiceManagerBaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractServiceManagerBaseOwnershipTransferred, error)

	FilterRewardsInitiatorUpdated(opts *bind.FilterOpts) (*ContractServiceManagerBaseRewardsInitiatorUpdatedIterator, error)
	WatchRewardsInitiatorUpdated(opts *bind.WatchOpts, sink chan<- *ContractServiceManagerBaseRewardsInitiatorUpdated) (event.Subscription, error)
	ParseRewardsInitiatorUpdated(log types.Log) (*ContractServiceManagerBaseRewardsInitiatorUpdated, error)

	FilterSlasherProposed(opts *bind.FilterOpts) (*ContractServiceManagerBaseSlasherProposedIterator, error)
	WatchSlasherProposed(opts *bind.WatchOpts, sink chan<- *ContractServiceManagerBaseSlasherProposed) (event.Subscription, error)
	ParseSlasherProposed(log types.Log) (*ContractServiceManagerBaseSlasherProposed, error)

	FilterSlasherUpdated(opts *bind.FilterOpts) (*ContractServiceManagerBaseSlasherUpdatedIterator, error)
	WatchSlasherUpdated(opts *bind.WatchOpts, sink chan<- *ContractServiceManagerBaseSlasherUpdated) (event.Subscription, error)
	ParseSlasherUpdated(log types.Log) (*ContractServiceManagerBaseSlasherUpdated, error)
}

// ContractServiceManagerBase is an auto generated Go binding around an Ethereum contract.
type ContractServiceManagerBase struct {
	ContractServiceManagerBaseCaller     // Read-only binding to the contract
	ContractServiceManagerBaseTransactor // Write-only binding to the contract
	ContractServiceManagerBaseFilterer   // Log filterer for contract events
}

// ContractServiceManagerBase implements the ContractServiceManagerBaseMethods interface.
var _ ContractServiceManagerBaseMethods = (*ContractServiceManagerBase)(nil)

// ContractServiceManagerBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractServiceManagerBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractServiceManagerBaseCaller implements the ContractServiceManagerBaseCalls interface.
var _ ContractServiceManagerBaseCalls = (*ContractServiceManagerBaseCaller)(nil)

// ContractServiceManagerBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractServiceManagerBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractServiceManagerBaseTransactor implements the ContractServiceManagerBaseTransacts interface.
var _ ContractServiceManagerBaseTransacts = (*ContractServiceManagerBaseTransactor)(nil)

// ContractServiceManagerBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractServiceManagerBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractServiceManagerBaseFilterer implements the ContractServiceManagerBaseFilters interface.
var _ ContractServiceManagerBaseFilters = (*ContractServiceManagerBaseFilterer)(nil)

// ContractServiceManagerBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractServiceManagerBaseSession struct {
	Contract     *ContractServiceManagerBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractServiceManagerBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractServiceManagerBaseCallerSession struct {
	Contract *ContractServiceManagerBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractServiceManagerBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractServiceManagerBaseTransactorSession struct {
	Contract     *ContractServiceManagerBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractServiceManagerBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractServiceManagerBaseRaw struct {
	Contract *ContractServiceManagerBase // Generic contract binding to access the raw methods on
}

// ContractServiceManagerBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractServiceManagerBaseCallerRaw struct {
	Contract *ContractServiceManagerBaseCaller // Generic read-only contract binding to access the raw methods on
}

// ContractServiceManagerBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractServiceManagerBaseTransactorRaw struct {
	Contract *ContractServiceManagerBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractServiceManagerBase creates a new instance of ContractServiceManagerBase, bound to a specific deployed contract.
func NewContractServiceManagerBase(address common.Address, backend bind.ContractBackend) (*ContractServiceManagerBase, error) {
	contract, err := bindContractServiceManagerBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerBase{ContractServiceManagerBaseCaller: ContractServiceManagerBaseCaller{contract: contract}, ContractServiceManagerBaseTransactor: ContractServiceManagerBaseTransactor{contract: contract}, ContractServiceManagerBaseFilterer: ContractServiceManagerBaseFilterer{contract: contract}}, nil
}

// NewContractServiceManagerBaseCaller creates a new read-only instance of ContractServiceManagerBase, bound to a specific deployed contract.
func NewContractServiceManagerBaseCaller(address common.Address, caller bind.ContractCaller) (*ContractServiceManagerBaseCaller, error) {
	contract, err := bindContractServiceManagerBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerBaseCaller{contract: contract}, nil
}

// NewContractServiceManagerBaseTransactor creates a new write-only instance of ContractServiceManagerBase, bound to a specific deployed contract.
func NewContractServiceManagerBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractServiceManagerBaseTransactor, error) {
	contract, err := bindContractServiceManagerBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerBaseTransactor{contract: contract}, nil
}

// NewContractServiceManagerBaseFilterer creates a new log filterer instance of ContractServiceManagerBase, bound to a specific deployed contract.
func NewContractServiceManagerBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractServiceManagerBaseFilterer, error) {
	contract, err := bindContractServiceManagerBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerBaseFilterer{contract: contract}, nil
}

// bindContractServiceManagerBase binds a generic wrapper to an already deployed contract.
func bindContractServiceManagerBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractServiceManagerBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractServiceManagerBase *ContractServiceManagerBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractServiceManagerBase.Contract.ContractServiceManagerBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractServiceManagerBase *ContractServiceManagerBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.ContractServiceManagerBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractServiceManagerBase *ContractServiceManagerBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.ContractServiceManagerBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractServiceManagerBase *ContractServiceManagerBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractServiceManagerBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.contract.Transact(opts, method, params...)
}

// SLASHERPROPOSALDELAY is a free data retrieval call binding the contract method 0x67940c89.
//
// Solidity: function SLASHER_PROPOSAL_DELAY() view returns(uint256)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCaller) SLASHERPROPOSALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractServiceManagerBase.contract.Call(opts, &out, "SLASHER_PROPOSAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLASHERPROPOSALDELAY is a free data retrieval call binding the contract method 0x67940c89.
//
// Solidity: function SLASHER_PROPOSAL_DELAY() view returns(uint256)
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) SLASHERPROPOSALDELAY() (*big.Int, error) {
	return _ContractServiceManagerBase.Contract.SLASHERPROPOSALDELAY(&_ContractServiceManagerBase.CallOpts)
}

// SLASHERPROPOSALDELAY is a free data retrieval call binding the contract method 0x67940c89.
//
// Solidity: function SLASHER_PROPOSAL_DELAY() view returns(uint256)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCallerSession) SLASHERPROPOSALDELAY() (*big.Int, error) {
	return _ContractServiceManagerBase.Contract.SLASHERPROPOSALDELAY(&_ContractServiceManagerBase.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractServiceManagerBase.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) AllocationManager() (common.Address, error) {
	return _ContractServiceManagerBase.Contract.AllocationManager(&_ContractServiceManagerBase.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCallerSession) AllocationManager() (common.Address, error) {
	return _ContractServiceManagerBase.Contract.AllocationManager(&_ContractServiceManagerBase.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractServiceManagerBase.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) AvsDirectory() (common.Address, error) {
	return _ContractServiceManagerBase.Contract.AvsDirectory(&_ContractServiceManagerBase.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractServiceManagerBase.Contract.AvsDirectory(&_ContractServiceManagerBase.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractServiceManagerBase *ContractServiceManagerBaseCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractServiceManagerBase.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractServiceManagerBase.Contract.GetOperatorRestakedStrategies(&_ContractServiceManagerBase.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractServiceManagerBase *ContractServiceManagerBaseCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractServiceManagerBase.Contract.GetOperatorRestakedStrategies(&_ContractServiceManagerBase.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractServiceManagerBase *ContractServiceManagerBaseCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractServiceManagerBase.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractServiceManagerBase.Contract.GetRestakeableStrategies(&_ContractServiceManagerBase.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractServiceManagerBase *ContractServiceManagerBaseCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractServiceManagerBase.Contract.GetRestakeableStrategies(&_ContractServiceManagerBase.CallOpts)
}

// MigrationFinalized is a free data retrieval call binding the contract method 0x8d68349a.
//
// Solidity: function migrationFinalized() view returns(bool)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCaller) MigrationFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractServiceManagerBase.contract.Call(opts, &out, "migrationFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MigrationFinalized is a free data retrieval call binding the contract method 0x8d68349a.
//
// Solidity: function migrationFinalized() view returns(bool)
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) MigrationFinalized() (bool, error) {
	return _ContractServiceManagerBase.Contract.MigrationFinalized(&_ContractServiceManagerBase.CallOpts)
}

// MigrationFinalized is a free data retrieval call binding the contract method 0x8d68349a.
//
// Solidity: function migrationFinalized() view returns(bool)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCallerSession) MigrationFinalized() (bool, error) {
	return _ContractServiceManagerBase.Contract.MigrationFinalized(&_ContractServiceManagerBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractServiceManagerBase.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) Owner() (common.Address, error) {
	return _ContractServiceManagerBase.Contract.Owner(&_ContractServiceManagerBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCallerSession) Owner() (common.Address, error) {
	return _ContractServiceManagerBase.Contract.Owner(&_ContractServiceManagerBase.CallOpts)
}

// ProposedSlasher is a free data retrieval call binding the contract method 0xe46f1816.
//
// Solidity: function proposedSlasher() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCaller) ProposedSlasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractServiceManagerBase.contract.Call(opts, &out, "proposedSlasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposedSlasher is a free data retrieval call binding the contract method 0xe46f1816.
//
// Solidity: function proposedSlasher() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) ProposedSlasher() (common.Address, error) {
	return _ContractServiceManagerBase.Contract.ProposedSlasher(&_ContractServiceManagerBase.CallOpts)
}

// ProposedSlasher is a free data retrieval call binding the contract method 0xe46f1816.
//
// Solidity: function proposedSlasher() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCallerSession) ProposedSlasher() (common.Address, error) {
	return _ContractServiceManagerBase.Contract.ProposedSlasher(&_ContractServiceManagerBase.CallOpts)
}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCaller) RewardsInitiator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractServiceManagerBase.contract.Call(opts, &out, "rewardsInitiator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) RewardsInitiator() (common.Address, error) {
	return _ContractServiceManagerBase.Contract.RewardsInitiator(&_ContractServiceManagerBase.CallOpts)
}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCallerSession) RewardsInitiator() (common.Address, error) {
	return _ContractServiceManagerBase.Contract.RewardsInitiator(&_ContractServiceManagerBase.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractServiceManagerBase.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) Slasher() (common.Address, error) {
	return _ContractServiceManagerBase.Contract.Slasher(&_ContractServiceManagerBase.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCallerSession) Slasher() (common.Address, error) {
	return _ContractServiceManagerBase.Contract.Slasher(&_ContractServiceManagerBase.CallOpts)
}

// SlasherProposalTimestamp is a free data retrieval call binding the contract method 0xfcd1c375.
//
// Solidity: function slasherProposalTimestamp() view returns(uint256)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCaller) SlasherProposalTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractServiceManagerBase.contract.Call(opts, &out, "slasherProposalTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlasherProposalTimestamp is a free data retrieval call binding the contract method 0xfcd1c375.
//
// Solidity: function slasherProposalTimestamp() view returns(uint256)
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) SlasherProposalTimestamp() (*big.Int, error) {
	return _ContractServiceManagerBase.Contract.SlasherProposalTimestamp(&_ContractServiceManagerBase.CallOpts)
}

// SlasherProposalTimestamp is a free data retrieval call binding the contract method 0xfcd1c375.
//
// Solidity: function slasherProposalTimestamp() view returns(uint256)
func (_ContractServiceManagerBase *ContractServiceManagerBaseCallerSession) SlasherProposalTimestamp() (*big.Int, error) {
	return _ContractServiceManagerBase.Contract.SlasherProposalTimestamp(&_ContractServiceManagerBase.CallOpts)
}

// AcceptProposedSlasher is a paid mutator transaction binding the contract method 0x26f017e2.
//
// Solidity: function acceptProposedSlasher() returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) AcceptProposedSlasher(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "acceptProposedSlasher")
}

// AcceptProposedSlasher is a paid mutator transaction binding the contract method 0x26f017e2.
//
// Solidity: function acceptProposedSlasher() returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) AcceptProposedSlasher() (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.AcceptProposedSlasher(&_ContractServiceManagerBase.TransactOpts)
}

// AcceptProposedSlasher is a paid mutator transaction binding the contract method 0x26f017e2.
//
// Solidity: function acceptProposedSlasher() returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) AcceptProposedSlasher() (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.AcceptProposedSlasher(&_ContractServiceManagerBase.TransactOpts)
}

// AddStrategyToOperatorSet is a paid mutator transaction binding the contract method 0x6ecbccfe.
//
// Solidity: function addStrategyToOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) AddStrategyToOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "addStrategyToOperatorSet", operatorSetId, strategies)
}

// AddStrategyToOperatorSet is a paid mutator transaction binding the contract method 0x6ecbccfe.
//
// Solidity: function addStrategyToOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) AddStrategyToOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.AddStrategyToOperatorSet(&_ContractServiceManagerBase.TransactOpts, operatorSetId, strategies)
}

// AddStrategyToOperatorSet is a paid mutator transaction binding the contract method 0x6ecbccfe.
//
// Solidity: function addStrategyToOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) AddStrategyToOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.AddStrategyToOperatorSet(&_ContractServiceManagerBase.TransactOpts, operatorSetId, strategies)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "createAVSRewardsSubmission", rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.CreateAVSRewardsSubmission(&_ContractServiceManagerBase.TransactOpts, rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.CreateAVSRewardsSubmission(&_ContractServiceManagerBase.TransactOpts, rewardsSubmissions)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x847d634f.
//
// Solidity: function createOperatorSets((uint32,address[])[] params) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) CreateOperatorSets(opts *bind.TransactOpts, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "createOperatorSets", params)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x847d634f.
//
// Solidity: function createOperatorSets((uint32,address[])[] params) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) CreateOperatorSets(params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.CreateOperatorSets(&_ContractServiceManagerBase.TransactOpts, params)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x847d634f.
//
// Solidity: function createOperatorSets((uint32,address[])[] params) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) CreateOperatorSets(params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.CreateOperatorSets(&_ContractServiceManagerBase.TransactOpts, params)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.DeregisterOperatorFromAVS(&_ContractServiceManagerBase.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.DeregisterOperatorFromAVS(&_ContractServiceManagerBase.TransactOpts, operator)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) DeregisterOperatorFromOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "deregisterOperatorFromOperatorSets", operator, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) DeregisterOperatorFromOperatorSets(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.DeregisterOperatorFromOperatorSets(&_ContractServiceManagerBase.TransactOpts, operator, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) DeregisterOperatorFromOperatorSets(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.DeregisterOperatorFromOperatorSets(&_ContractServiceManagerBase.TransactOpts, operator, operatorSetIds)
}

// ProposeNewSlasher is a paid mutator transaction binding the contract method 0x8999817f.
//
// Solidity: function proposeNewSlasher(address newSlasher) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) ProposeNewSlasher(opts *bind.TransactOpts, newSlasher common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "proposeNewSlasher", newSlasher)
}

// ProposeNewSlasher is a paid mutator transaction binding the contract method 0x8999817f.
//
// Solidity: function proposeNewSlasher(address newSlasher) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) ProposeNewSlasher(newSlasher common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.ProposeNewSlasher(&_ContractServiceManagerBase.TransactOpts, newSlasher)
}

// ProposeNewSlasher is a paid mutator transaction binding the contract method 0x8999817f.
//
// Solidity: function proposeNewSlasher(address newSlasher) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) ProposeNewSlasher(newSlasher common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.ProposeNewSlasher(&_ContractServiceManagerBase.TransactOpts, newSlasher)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.RegisterOperatorToAVS(&_ContractServiceManagerBase.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.RegisterOperatorToAVS(&_ContractServiceManagerBase.TransactOpts, operator, operatorSignature)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xce7b5e4b.
//
// Solidity: function removeStrategiesFromOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) RemoveStrategiesFromOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "removeStrategiesFromOperatorSet", operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xce7b5e4b.
//
// Solidity: function removeStrategiesFromOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) RemoveStrategiesFromOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.RemoveStrategiesFromOperatorSet(&_ContractServiceManagerBase.TransactOpts, operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xce7b5e4b.
//
// Solidity: function removeStrategiesFromOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) RemoveStrategiesFromOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.RemoveStrategiesFromOperatorSet(&_ContractServiceManagerBase.TransactOpts, operatorSetId, strategies)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.RenounceOwnership(&_ContractServiceManagerBase.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.RenounceOwnership(&_ContractServiceManagerBase.TransactOpts)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xf25f1610.
//
// Solidity: function setAVSRegistrar(address registrar) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) SetAVSRegistrar(opts *bind.TransactOpts, registrar common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "setAVSRegistrar", registrar)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xf25f1610.
//
// Solidity: function setAVSRegistrar(address registrar) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) SetAVSRegistrar(registrar common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.SetAVSRegistrar(&_ContractServiceManagerBase.TransactOpts, registrar)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xf25f1610.
//
// Solidity: function setAVSRegistrar(address registrar) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) SetAVSRegistrar(registrar common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.SetAVSRegistrar(&_ContractServiceManagerBase.TransactOpts, registrar)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) SetRewardsInitiator(opts *bind.TransactOpts, newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "setRewardsInitiator", newRewardsInitiator)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) SetRewardsInitiator(newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.SetRewardsInitiator(&_ContractServiceManagerBase.TransactOpts, newRewardsInitiator)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) SetRewardsInitiator(newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.SetRewardsInitiator(&_ContractServiceManagerBase.TransactOpts, newRewardsInitiator)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x3d071422.
//
// Solidity: function slashOperator((address,uint32,address[],uint256[],string) params) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) SlashOperator(opts *bind.TransactOpts, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "slashOperator", params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x3d071422.
//
// Solidity: function slashOperator((address,uint32,address[],uint256[],string) params) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) SlashOperator(params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.SlashOperator(&_ContractServiceManagerBase.TransactOpts, params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x3d071422.
//
// Solidity: function slashOperator((address,uint32,address[],uint256[],string) params) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) SlashOperator(params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.SlashOperator(&_ContractServiceManagerBase.TransactOpts, params)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.TransferOwnership(&_ContractServiceManagerBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.TransferOwnership(&_ContractServiceManagerBase.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractServiceManagerBase.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.UpdateAVSMetadataURI(&_ContractServiceManagerBase.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractServiceManagerBase *ContractServiceManagerBaseTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractServiceManagerBase.Contract.UpdateAVSMetadataURI(&_ContractServiceManagerBase.TransactOpts, _metadataURI)
}

// ContractServiceManagerBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractServiceManagerBase contract.
type ContractServiceManagerBaseInitializedIterator struct {
	Event *ContractServiceManagerBaseInitialized // Event containing the contract specifics and raw log

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
func (it *ContractServiceManagerBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractServiceManagerBaseInitialized)
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
		it.Event = new(ContractServiceManagerBaseInitialized)
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
func (it *ContractServiceManagerBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractServiceManagerBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractServiceManagerBaseInitialized represents a Initialized event raised by the ContractServiceManagerBase contract.
type ContractServiceManagerBaseInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractServiceManagerBaseInitializedIterator, error) {

	logs, sub, err := _ContractServiceManagerBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerBaseInitializedIterator{contract: _ContractServiceManagerBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractServiceManagerBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractServiceManagerBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractServiceManagerBaseInitialized)
				if err := _ContractServiceManagerBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) ParseInitialized(log types.Log) (*ContractServiceManagerBaseInitialized, error) {
	event := new(ContractServiceManagerBaseInitialized)
	if err := _ContractServiceManagerBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractServiceManagerBaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractServiceManagerBase contract.
type ContractServiceManagerBaseOwnershipTransferredIterator struct {
	Event *ContractServiceManagerBaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractServiceManagerBaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractServiceManagerBaseOwnershipTransferred)
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
		it.Event = new(ContractServiceManagerBaseOwnershipTransferred)
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
func (it *ContractServiceManagerBaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractServiceManagerBaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractServiceManagerBaseOwnershipTransferred represents a OwnershipTransferred event raised by the ContractServiceManagerBase contract.
type ContractServiceManagerBaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractServiceManagerBaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractServiceManagerBase.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerBaseOwnershipTransferredIterator{contract: _ContractServiceManagerBase.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractServiceManagerBaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractServiceManagerBase.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractServiceManagerBaseOwnershipTransferred)
				if err := _ContractServiceManagerBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) ParseOwnershipTransferred(log types.Log) (*ContractServiceManagerBaseOwnershipTransferred, error) {
	event := new(ContractServiceManagerBaseOwnershipTransferred)
	if err := _ContractServiceManagerBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractServiceManagerBaseRewardsInitiatorUpdatedIterator is returned from FilterRewardsInitiatorUpdated and is used to iterate over the raw logs and unpacked data for RewardsInitiatorUpdated events raised by the ContractServiceManagerBase contract.
type ContractServiceManagerBaseRewardsInitiatorUpdatedIterator struct {
	Event *ContractServiceManagerBaseRewardsInitiatorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractServiceManagerBaseRewardsInitiatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractServiceManagerBaseRewardsInitiatorUpdated)
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
		it.Event = new(ContractServiceManagerBaseRewardsInitiatorUpdated)
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
func (it *ContractServiceManagerBaseRewardsInitiatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractServiceManagerBaseRewardsInitiatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractServiceManagerBaseRewardsInitiatorUpdated represents a RewardsInitiatorUpdated event raised by the ContractServiceManagerBase contract.
type ContractServiceManagerBaseRewardsInitiatorUpdated struct {
	PrevRewardsInitiator common.Address
	NewRewardsInitiator  common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRewardsInitiatorUpdated is a free log retrieval operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) FilterRewardsInitiatorUpdated(opts *bind.FilterOpts) (*ContractServiceManagerBaseRewardsInitiatorUpdatedIterator, error) {

	logs, sub, err := _ContractServiceManagerBase.contract.FilterLogs(opts, "RewardsInitiatorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerBaseRewardsInitiatorUpdatedIterator{contract: _ContractServiceManagerBase.contract, event: "RewardsInitiatorUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsInitiatorUpdated is a free log subscription operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) WatchRewardsInitiatorUpdated(opts *bind.WatchOpts, sink chan<- *ContractServiceManagerBaseRewardsInitiatorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractServiceManagerBase.contract.WatchLogs(opts, "RewardsInitiatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractServiceManagerBaseRewardsInitiatorUpdated)
				if err := _ContractServiceManagerBase.contract.UnpackLog(event, "RewardsInitiatorUpdated", log); err != nil {
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

// ParseRewardsInitiatorUpdated is a log parse operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) ParseRewardsInitiatorUpdated(log types.Log) (*ContractServiceManagerBaseRewardsInitiatorUpdated, error) {
	event := new(ContractServiceManagerBaseRewardsInitiatorUpdated)
	if err := _ContractServiceManagerBase.contract.UnpackLog(event, "RewardsInitiatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractServiceManagerBaseSlasherProposedIterator is returned from FilterSlasherProposed and is used to iterate over the raw logs and unpacked data for SlasherProposed events raised by the ContractServiceManagerBase contract.
type ContractServiceManagerBaseSlasherProposedIterator struct {
	Event *ContractServiceManagerBaseSlasherProposed // Event containing the contract specifics and raw log

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
func (it *ContractServiceManagerBaseSlasherProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractServiceManagerBaseSlasherProposed)
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
		it.Event = new(ContractServiceManagerBaseSlasherProposed)
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
func (it *ContractServiceManagerBaseSlasherProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractServiceManagerBaseSlasherProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractServiceManagerBaseSlasherProposed represents a SlasherProposed event raised by the ContractServiceManagerBase contract.
type ContractServiceManagerBaseSlasherProposed struct {
	NewSlasher               common.Address
	SlasherProposalTimestamp *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSlasherProposed is a free log retrieval operation binding the contract event 0x2f8afc8a78fd958f3301c0233aa326b9c4b9a2884a7483227d6b0555aaa03adb.
//
// Solidity: event SlasherProposed(address newSlasher, uint256 slasherProposalTimestamp)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) FilterSlasherProposed(opts *bind.FilterOpts) (*ContractServiceManagerBaseSlasherProposedIterator, error) {

	logs, sub, err := _ContractServiceManagerBase.contract.FilterLogs(opts, "SlasherProposed")
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerBaseSlasherProposedIterator{contract: _ContractServiceManagerBase.contract, event: "SlasherProposed", logs: logs, sub: sub}, nil
}

// WatchSlasherProposed is a free log subscription operation binding the contract event 0x2f8afc8a78fd958f3301c0233aa326b9c4b9a2884a7483227d6b0555aaa03adb.
//
// Solidity: event SlasherProposed(address newSlasher, uint256 slasherProposalTimestamp)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) WatchSlasherProposed(opts *bind.WatchOpts, sink chan<- *ContractServiceManagerBaseSlasherProposed) (event.Subscription, error) {

	logs, sub, err := _ContractServiceManagerBase.contract.WatchLogs(opts, "SlasherProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractServiceManagerBaseSlasherProposed)
				if err := _ContractServiceManagerBase.contract.UnpackLog(event, "SlasherProposed", log); err != nil {
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

// ParseSlasherProposed is a log parse operation binding the contract event 0x2f8afc8a78fd958f3301c0233aa326b9c4b9a2884a7483227d6b0555aaa03adb.
//
// Solidity: event SlasherProposed(address newSlasher, uint256 slasherProposalTimestamp)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) ParseSlasherProposed(log types.Log) (*ContractServiceManagerBaseSlasherProposed, error) {
	event := new(ContractServiceManagerBaseSlasherProposed)
	if err := _ContractServiceManagerBase.contract.UnpackLog(event, "SlasherProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractServiceManagerBaseSlasherUpdatedIterator is returned from FilterSlasherUpdated and is used to iterate over the raw logs and unpacked data for SlasherUpdated events raised by the ContractServiceManagerBase contract.
type ContractServiceManagerBaseSlasherUpdatedIterator struct {
	Event *ContractServiceManagerBaseSlasherUpdated // Event containing the contract specifics and raw log

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
func (it *ContractServiceManagerBaseSlasherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractServiceManagerBaseSlasherUpdated)
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
		it.Event = new(ContractServiceManagerBaseSlasherUpdated)
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
func (it *ContractServiceManagerBaseSlasherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractServiceManagerBaseSlasherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractServiceManagerBaseSlasherUpdated represents a SlasherUpdated event raised by the ContractServiceManagerBase contract.
type ContractServiceManagerBaseSlasherUpdated struct {
	PrevSlasher common.Address
	NewSlasher  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSlasherUpdated is a free log retrieval operation binding the contract event 0xe0d49a54274423183dadecbdf239eaac6e06ba88320b26fe8cc5ec9d050a6395.
//
// Solidity: event SlasherUpdated(address prevSlasher, address newSlasher)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) FilterSlasherUpdated(opts *bind.FilterOpts) (*ContractServiceManagerBaseSlasherUpdatedIterator, error) {

	logs, sub, err := _ContractServiceManagerBase.contract.FilterLogs(opts, "SlasherUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerBaseSlasherUpdatedIterator{contract: _ContractServiceManagerBase.contract, event: "SlasherUpdated", logs: logs, sub: sub}, nil
}

// WatchSlasherUpdated is a free log subscription operation binding the contract event 0xe0d49a54274423183dadecbdf239eaac6e06ba88320b26fe8cc5ec9d050a6395.
//
// Solidity: event SlasherUpdated(address prevSlasher, address newSlasher)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) WatchSlasherUpdated(opts *bind.WatchOpts, sink chan<- *ContractServiceManagerBaseSlasherUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractServiceManagerBase.contract.WatchLogs(opts, "SlasherUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractServiceManagerBaseSlasherUpdated)
				if err := _ContractServiceManagerBase.contract.UnpackLog(event, "SlasherUpdated", log); err != nil {
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

// ParseSlasherUpdated is a log parse operation binding the contract event 0xe0d49a54274423183dadecbdf239eaac6e06ba88320b26fe8cc5ec9d050a6395.
//
// Solidity: event SlasherUpdated(address prevSlasher, address newSlasher)
func (_ContractServiceManagerBase *ContractServiceManagerBaseFilterer) ParseSlasherUpdated(log types.Log) (*ContractServiceManagerBaseSlasherUpdated, error) {
	event := new(ContractServiceManagerBaseSlasherUpdated)
	if err := _ContractServiceManagerBase.contract.UnpackLog(event, "SlasherUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
