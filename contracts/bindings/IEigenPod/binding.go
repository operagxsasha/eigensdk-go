// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIEigenPod

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

// BeaconChainProofsBalanceContainerProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsBalanceContainerProof struct {
	BalanceContainerRoot [32]byte
	Proof                []byte
}

// BeaconChainProofsBalanceProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsBalanceProof struct {
	PubkeyHash  [32]byte
	BalanceRoot [32]byte
	Proof       []byte
}

// BeaconChainProofsStateRootProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsStateRootProof struct {
	BeaconStateRoot [32]byte
	Proof           []byte
}

// BeaconChainProofsValidatorProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsValidatorProof struct {
	ValidatorFields [][32]byte
	Proof           []byte
}

// IEigenPodCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodCheckpoint struct {
	BeaconBlockRoot   [32]byte
	ProofsRemaining   *big.Int
	PodBalanceGwei    uint64
	BalanceDeltasGwei *big.Int
}

// IEigenPodValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodValidatorInfo struct {
	ValidatorIndex      uint64
	RestakedBalanceGwei uint64
	LastCheckpointedAt  uint64
	Status              uint8
}

// ContractIEigenPodMetaData contains all meta data concerning the ContractIEigenPod contract.
var ContractIEigenPodMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"activeValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointBalanceExitedGwei\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentCheckpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPod.Checkpoint\",\"components\":[{\"name\":\"beaconBlockRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proofsRemaining\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"podBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"balanceDeltasGwei\",\"type\":\"int128\",\"internalType\":\"int128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentCheckpointTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParentBlockRoot\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastCheckpointTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"podOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proofSubmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverTokens\",\"inputs\":[{\"name\":\"tokenList\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"amountsToWithdraw\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProofSubmitter\",\"inputs\":[{\"name\":\"newProofSubmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"startCheckpoint\",\"inputs\":[{\"name\":\"revertIfNoBalance\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorPubkeyHashToInfo\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPod.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastCheckpointedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorPubkeyToInfo\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPod.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastCheckpointedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCheckpointProofs\",\"inputs\":[{\"name\":\"balanceContainerProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.BalanceContainerProof\",\"components\":[{\"name\":\"balanceContainerRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"proofs\",\"type\":\"tuple[]\",\"internalType\":\"structBeaconChainProofs.BalanceProof[]\",\"components\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"balanceRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyStaleBalance\",\"inputs\":[{\"name\":\"beaconTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.ValidatorProof\",\"components\":[{\"name\":\"validatorFields\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyWithdrawalCredentials\",\"inputs\":[{\"name\":\"beaconTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"validatorIndices\",\"type\":\"uint40[]\",\"internalType\":\"uint40[]\"},{\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawRestakedBeaconChainETH\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawableRestakedExecutionLayerGwei\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CheckpointCreated\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"beaconBlockRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validatorCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointFinalized\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"totalShareDeltaWei\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EigenPodStaked\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonBeaconChainETHReceived\",\"inputs\":[{\"name\":\"amountReceived\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProofSubmitterUpdated\",\"inputs\":[{\"name\":\"prevProofSubmitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newProofSubmitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RestakedBeaconChainETHWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorBalanceUpdated\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"balanceTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newValidatorBalanceGwei\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorCheckpointed\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":true,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRestaked\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorWithdrawn\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":true,\"internalType\":\"uint40\"}],\"anonymous\":false}]",
}

// ContractIEigenPodABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIEigenPodMetaData.ABI instead.
var ContractIEigenPodABI = ContractIEigenPodMetaData.ABI

// ContractIEigenPodMethods is an auto generated interface around an Ethereum contract.
type ContractIEigenPodMethods interface {
	ContractIEigenPodCalls
	ContractIEigenPodTransacts
	ContractIEigenPodFilters
}

// ContractIEigenPodCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractIEigenPodCalls interface {
	ActiveValidatorCount(opts *bind.CallOpts) (*big.Int, error)

	CheckpointBalanceExitedGwei(opts *bind.CallOpts, arg0 uint64) (uint64, error)

	CurrentCheckpoint(opts *bind.CallOpts) (IEigenPodCheckpoint, error)

	CurrentCheckpointTimestamp(opts *bind.CallOpts) (uint64, error)

	EigenPodManager(opts *bind.CallOpts) (common.Address, error)

	GetParentBlockRoot(opts *bind.CallOpts, timestamp uint64) ([32]byte, error)

	LastCheckpointTimestamp(opts *bind.CallOpts) (uint64, error)

	PodOwner(opts *bind.CallOpts) (common.Address, error)

	ProofSubmitter(opts *bind.CallOpts) (common.Address, error)

	ValidatorPubkeyHashToInfo(opts *bind.CallOpts, validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error)

	ValidatorPubkeyToInfo(opts *bind.CallOpts, validatorPubkey []byte) (IEigenPodValidatorInfo, error)

	ValidatorStatus(opts *bind.CallOpts, validatorPubkey []byte) (uint8, error)

	ValidatorStatus0(opts *bind.CallOpts, pubkeyHash [32]byte) (uint8, error)

	WithdrawableRestakedExecutionLayerGwei(opts *bind.CallOpts) (uint64, error)
}

// ContractIEigenPodTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIEigenPodTransacts interface {
	Initialize(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error)

	RecoverTokens(opts *bind.TransactOpts, tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error)

	SetProofSubmitter(opts *bind.TransactOpts, newProofSubmitter common.Address) (*types.Transaction, error)

	Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error)

	StartCheckpoint(opts *bind.TransactOpts, revertIfNoBalance bool) (*types.Transaction, error)

	VerifyCheckpointProofs(opts *bind.TransactOpts, balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error)

	VerifyStaleBalance(opts *bind.TransactOpts, beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, proof BeaconChainProofsValidatorProof) (*types.Transaction, error)

	VerifyWithdrawalCredentials(opts *bind.TransactOpts, beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error)

	WithdrawRestakedBeaconChainETH(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)
}

// ContractIEigenPodFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIEigenPodFilters interface {
	FilterCheckpointCreated(opts *bind.FilterOpts, checkpointTimestamp []uint64, beaconBlockRoot [][32]byte) (*ContractIEigenPodCheckpointCreatedIterator, error)
	WatchCheckpointCreated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodCheckpointCreated, checkpointTimestamp []uint64, beaconBlockRoot [][32]byte) (event.Subscription, error)
	ParseCheckpointCreated(log types.Log) (*ContractIEigenPodCheckpointCreated, error)

	FilterCheckpointFinalized(opts *bind.FilterOpts, checkpointTimestamp []uint64) (*ContractIEigenPodCheckpointFinalizedIterator, error)
	WatchCheckpointFinalized(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodCheckpointFinalized, checkpointTimestamp []uint64) (event.Subscription, error)
	ParseCheckpointFinalized(log types.Log) (*ContractIEigenPodCheckpointFinalized, error)

	FilterEigenPodStaked(opts *bind.FilterOpts) (*ContractIEigenPodEigenPodStakedIterator, error)
	WatchEigenPodStaked(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodEigenPodStaked) (event.Subscription, error)
	ParseEigenPodStaked(log types.Log) (*ContractIEigenPodEigenPodStaked, error)

	FilterNonBeaconChainETHReceived(opts *bind.FilterOpts) (*ContractIEigenPodNonBeaconChainETHReceivedIterator, error)
	WatchNonBeaconChainETHReceived(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodNonBeaconChainETHReceived) (event.Subscription, error)
	ParseNonBeaconChainETHReceived(log types.Log) (*ContractIEigenPodNonBeaconChainETHReceived, error)

	FilterProofSubmitterUpdated(opts *bind.FilterOpts) (*ContractIEigenPodProofSubmitterUpdatedIterator, error)
	WatchProofSubmitterUpdated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodProofSubmitterUpdated) (event.Subscription, error)
	ParseProofSubmitterUpdated(log types.Log) (*ContractIEigenPodProofSubmitterUpdated, error)

	FilterRestakedBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator, error)
	WatchRestakedBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodRestakedBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error)
	ParseRestakedBeaconChainETHWithdrawn(log types.Log) (*ContractIEigenPodRestakedBeaconChainETHWithdrawn, error)

	FilterValidatorBalanceUpdated(opts *bind.FilterOpts) (*ContractIEigenPodValidatorBalanceUpdatedIterator, error)
	WatchValidatorBalanceUpdated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodValidatorBalanceUpdated) (event.Subscription, error)
	ParseValidatorBalanceUpdated(log types.Log) (*ContractIEigenPodValidatorBalanceUpdated, error)

	FilterValidatorCheckpointed(opts *bind.FilterOpts, checkpointTimestamp []uint64, validatorIndex []*big.Int) (*ContractIEigenPodValidatorCheckpointedIterator, error)
	WatchValidatorCheckpointed(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodValidatorCheckpointed, checkpointTimestamp []uint64, validatorIndex []*big.Int) (event.Subscription, error)
	ParseValidatorCheckpointed(log types.Log) (*ContractIEigenPodValidatorCheckpointed, error)

	FilterValidatorRestaked(opts *bind.FilterOpts) (*ContractIEigenPodValidatorRestakedIterator, error)
	WatchValidatorRestaked(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodValidatorRestaked) (event.Subscription, error)
	ParseValidatorRestaked(log types.Log) (*ContractIEigenPodValidatorRestaked, error)

	FilterValidatorWithdrawn(opts *bind.FilterOpts, checkpointTimestamp []uint64, validatorIndex []*big.Int) (*ContractIEigenPodValidatorWithdrawnIterator, error)
	WatchValidatorWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodValidatorWithdrawn, checkpointTimestamp []uint64, validatorIndex []*big.Int) (event.Subscription, error)
	ParseValidatorWithdrawn(log types.Log) (*ContractIEigenPodValidatorWithdrawn, error)
}

// ContractIEigenPod is an auto generated Go binding around an Ethereum contract.
type ContractIEigenPod struct {
	ContractIEigenPodCaller     // Read-only binding to the contract
	ContractIEigenPodTransactor // Write-only binding to the contract
	ContractIEigenPodFilterer   // Log filterer for contract events
}

// ContractIEigenPod implements the ContractIEigenPodMethods interface.
var _ ContractIEigenPodMethods = (*ContractIEigenPod)(nil)

// ContractIEigenPodCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIEigenPodCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIEigenPodCaller implements the ContractIEigenPodCalls interface.
var _ ContractIEigenPodCalls = (*ContractIEigenPodCaller)(nil)

// ContractIEigenPodTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIEigenPodTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIEigenPodTransactor implements the ContractIEigenPodTransacts interface.
var _ ContractIEigenPodTransacts = (*ContractIEigenPodTransactor)(nil)

// ContractIEigenPodFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIEigenPodFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIEigenPodFilterer implements the ContractIEigenPodFilters interface.
var _ ContractIEigenPodFilters = (*ContractIEigenPodFilterer)(nil)

// ContractIEigenPodSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIEigenPodSession struct {
	Contract     *ContractIEigenPod // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContractIEigenPodCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIEigenPodCallerSession struct {
	Contract *ContractIEigenPodCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ContractIEigenPodTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIEigenPodTransactorSession struct {
	Contract     *ContractIEigenPodTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractIEigenPodRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIEigenPodRaw struct {
	Contract *ContractIEigenPod // Generic contract binding to access the raw methods on
}

// ContractIEigenPodCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIEigenPodCallerRaw struct {
	Contract *ContractIEigenPodCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIEigenPodTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIEigenPodTransactorRaw struct {
	Contract *ContractIEigenPodTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIEigenPod creates a new instance of ContractIEigenPod, bound to a specific deployed contract.
func NewContractIEigenPod(address common.Address, backend bind.ContractBackend) (*ContractIEigenPod, error) {
	contract, err := bindContractIEigenPod(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPod{ContractIEigenPodCaller: ContractIEigenPodCaller{contract: contract}, ContractIEigenPodTransactor: ContractIEigenPodTransactor{contract: contract}, ContractIEigenPodFilterer: ContractIEigenPodFilterer{contract: contract}}, nil
}

// NewContractIEigenPodCaller creates a new read-only instance of ContractIEigenPod, bound to a specific deployed contract.
func NewContractIEigenPodCaller(address common.Address, caller bind.ContractCaller) (*ContractIEigenPodCaller, error) {
	contract, err := bindContractIEigenPod(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodCaller{contract: contract}, nil
}

// NewContractIEigenPodTransactor creates a new write-only instance of ContractIEigenPod, bound to a specific deployed contract.
func NewContractIEigenPodTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIEigenPodTransactor, error) {
	contract, err := bindContractIEigenPod(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodTransactor{contract: contract}, nil
}

// NewContractIEigenPodFilterer creates a new log filterer instance of ContractIEigenPod, bound to a specific deployed contract.
func NewContractIEigenPodFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIEigenPodFilterer, error) {
	contract, err := bindContractIEigenPod(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodFilterer{contract: contract}, nil
}

// bindContractIEigenPod binds a generic wrapper to an already deployed contract.
func bindContractIEigenPod(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIEigenPodMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIEigenPod *ContractIEigenPodRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIEigenPod.Contract.ContractIEigenPodCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIEigenPod *ContractIEigenPodRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.ContractIEigenPodTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIEigenPod *ContractIEigenPodRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.ContractIEigenPodTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIEigenPod *ContractIEigenPodCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIEigenPod.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIEigenPod *ContractIEigenPodTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIEigenPod *ContractIEigenPodTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.contract.Transact(opts, method, params...)
}

// ActiveValidatorCount is a free data retrieval call binding the contract method 0x2340e8d3.
//
// Solidity: function activeValidatorCount() view returns(uint256)
func (_ContractIEigenPod *ContractIEigenPodCaller) ActiveValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "activeValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveValidatorCount is a free data retrieval call binding the contract method 0x2340e8d3.
//
// Solidity: function activeValidatorCount() view returns(uint256)
func (_ContractIEigenPod *ContractIEigenPodSession) ActiveValidatorCount() (*big.Int, error) {
	return _ContractIEigenPod.Contract.ActiveValidatorCount(&_ContractIEigenPod.CallOpts)
}

// ActiveValidatorCount is a free data retrieval call binding the contract method 0x2340e8d3.
//
// Solidity: function activeValidatorCount() view returns(uint256)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) ActiveValidatorCount() (*big.Int, error) {
	return _ContractIEigenPod.Contract.ActiveValidatorCount(&_ContractIEigenPod.CallOpts)
}

// CheckpointBalanceExitedGwei is a free data retrieval call binding the contract method 0x52396a59.
//
// Solidity: function checkpointBalanceExitedGwei(uint64 ) view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCaller) CheckpointBalanceExitedGwei(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "checkpointBalanceExitedGwei", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CheckpointBalanceExitedGwei is a free data retrieval call binding the contract method 0x52396a59.
//
// Solidity: function checkpointBalanceExitedGwei(uint64 ) view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodSession) CheckpointBalanceExitedGwei(arg0 uint64) (uint64, error) {
	return _ContractIEigenPod.Contract.CheckpointBalanceExitedGwei(&_ContractIEigenPod.CallOpts, arg0)
}

// CheckpointBalanceExitedGwei is a free data retrieval call binding the contract method 0x52396a59.
//
// Solidity: function checkpointBalanceExitedGwei(uint64 ) view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) CheckpointBalanceExitedGwei(arg0 uint64) (uint64, error) {
	return _ContractIEigenPod.Contract.CheckpointBalanceExitedGwei(&_ContractIEigenPod.CallOpts, arg0)
}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int128))
func (_ContractIEigenPod *ContractIEigenPodCaller) CurrentCheckpoint(opts *bind.CallOpts) (IEigenPodCheckpoint, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "currentCheckpoint")

	if err != nil {
		return *new(IEigenPodCheckpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodCheckpoint)).(*IEigenPodCheckpoint)

	return out0, err

}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int128))
func (_ContractIEigenPod *ContractIEigenPodSession) CurrentCheckpoint() (IEigenPodCheckpoint, error) {
	return _ContractIEigenPod.Contract.CurrentCheckpoint(&_ContractIEigenPod.CallOpts)
}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int128))
func (_ContractIEigenPod *ContractIEigenPodCallerSession) CurrentCheckpoint() (IEigenPodCheckpoint, error) {
	return _ContractIEigenPod.Contract.CurrentCheckpoint(&_ContractIEigenPod.CallOpts)
}

// CurrentCheckpointTimestamp is a free data retrieval call binding the contract method 0x42ecff2a.
//
// Solidity: function currentCheckpointTimestamp() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCaller) CurrentCheckpointTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "currentCheckpointTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrentCheckpointTimestamp is a free data retrieval call binding the contract method 0x42ecff2a.
//
// Solidity: function currentCheckpointTimestamp() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodSession) CurrentCheckpointTimestamp() (uint64, error) {
	return _ContractIEigenPod.Contract.CurrentCheckpointTimestamp(&_ContractIEigenPod.CallOpts)
}

// CurrentCheckpointTimestamp is a free data retrieval call binding the contract method 0x42ecff2a.
//
// Solidity: function currentCheckpointTimestamp() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) CurrentCheckpointTimestamp() (uint64, error) {
	return _ContractIEigenPod.Contract.CurrentCheckpointTimestamp(&_ContractIEigenPod.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodSession) EigenPodManager() (common.Address, error) {
	return _ContractIEigenPod.Contract.EigenPodManager(&_ContractIEigenPod.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) EigenPodManager() (common.Address, error) {
	return _ContractIEigenPod.Contract.EigenPodManager(&_ContractIEigenPod.CallOpts)
}

// GetParentBlockRoot is a free data retrieval call binding the contract method 0x6c0d2d5a.
//
// Solidity: function getParentBlockRoot(uint64 timestamp) view returns(bytes32)
func (_ContractIEigenPod *ContractIEigenPodCaller) GetParentBlockRoot(opts *bind.CallOpts, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "getParentBlockRoot", timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetParentBlockRoot is a free data retrieval call binding the contract method 0x6c0d2d5a.
//
// Solidity: function getParentBlockRoot(uint64 timestamp) view returns(bytes32)
func (_ContractIEigenPod *ContractIEigenPodSession) GetParentBlockRoot(timestamp uint64) ([32]byte, error) {
	return _ContractIEigenPod.Contract.GetParentBlockRoot(&_ContractIEigenPod.CallOpts, timestamp)
}

// GetParentBlockRoot is a free data retrieval call binding the contract method 0x6c0d2d5a.
//
// Solidity: function getParentBlockRoot(uint64 timestamp) view returns(bytes32)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) GetParentBlockRoot(timestamp uint64) ([32]byte, error) {
	return _ContractIEigenPod.Contract.GetParentBlockRoot(&_ContractIEigenPod.CallOpts, timestamp)
}

// LastCheckpointTimestamp is a free data retrieval call binding the contract method 0xee94d67c.
//
// Solidity: function lastCheckpointTimestamp() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCaller) LastCheckpointTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "lastCheckpointTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastCheckpointTimestamp is a free data retrieval call binding the contract method 0xee94d67c.
//
// Solidity: function lastCheckpointTimestamp() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodSession) LastCheckpointTimestamp() (uint64, error) {
	return _ContractIEigenPod.Contract.LastCheckpointTimestamp(&_ContractIEigenPod.CallOpts)
}

// LastCheckpointTimestamp is a free data retrieval call binding the contract method 0xee94d67c.
//
// Solidity: function lastCheckpointTimestamp() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) LastCheckpointTimestamp() (uint64, error) {
	return _ContractIEigenPod.Contract.LastCheckpointTimestamp(&_ContractIEigenPod.CallOpts)
}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodCaller) PodOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "podOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodSession) PodOwner() (common.Address, error) {
	return _ContractIEigenPod.Contract.PodOwner(&_ContractIEigenPod.CallOpts)
}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) PodOwner() (common.Address, error) {
	return _ContractIEigenPod.Contract.PodOwner(&_ContractIEigenPod.CallOpts)
}

// ProofSubmitter is a free data retrieval call binding the contract method 0x58753357.
//
// Solidity: function proofSubmitter() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodCaller) ProofSubmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "proofSubmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProofSubmitter is a free data retrieval call binding the contract method 0x58753357.
//
// Solidity: function proofSubmitter() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodSession) ProofSubmitter() (common.Address, error) {
	return _ContractIEigenPod.Contract.ProofSubmitter(&_ContractIEigenPod.CallOpts)
}

// ProofSubmitter is a free data retrieval call binding the contract method 0x58753357.
//
// Solidity: function proofSubmitter() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) ProofSubmitter() (common.Address, error) {
	return _ContractIEigenPod.Contract.ProofSubmitter(&_ContractIEigenPod.CallOpts)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_ContractIEigenPod *ContractIEigenPodCaller) ValidatorPubkeyHashToInfo(opts *bind.CallOpts, validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "validatorPubkeyHashToInfo", validatorPubkeyHash)

	if err != nil {
		return *new(IEigenPodValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodValidatorInfo)).(*IEigenPodValidatorInfo)

	return out0, err

}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_ContractIEigenPod *ContractIEigenPodSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	return _ContractIEigenPod.Contract.ValidatorPubkeyHashToInfo(&_ContractIEigenPod.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_ContractIEigenPod *ContractIEigenPodCallerSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	return _ContractIEigenPod.Contract.ValidatorPubkeyHashToInfo(&_ContractIEigenPod.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_ContractIEigenPod *ContractIEigenPodCaller) ValidatorPubkeyToInfo(opts *bind.CallOpts, validatorPubkey []byte) (IEigenPodValidatorInfo, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "validatorPubkeyToInfo", validatorPubkey)

	if err != nil {
		return *new(IEigenPodValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodValidatorInfo)).(*IEigenPodValidatorInfo)

	return out0, err

}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_ContractIEigenPod *ContractIEigenPodSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodValidatorInfo, error) {
	return _ContractIEigenPod.Contract.ValidatorPubkeyToInfo(&_ContractIEigenPod.CallOpts, validatorPubkey)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_ContractIEigenPod *ContractIEigenPodCallerSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodValidatorInfo, error) {
	return _ContractIEigenPod.Contract.ValidatorPubkeyToInfo(&_ContractIEigenPod.CallOpts, validatorPubkey)
}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_ContractIEigenPod *ContractIEigenPodCaller) ValidatorStatus(opts *bind.CallOpts, validatorPubkey []byte) (uint8, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "validatorStatus", validatorPubkey)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_ContractIEigenPod *ContractIEigenPodSession) ValidatorStatus(validatorPubkey []byte) (uint8, error) {
	return _ContractIEigenPod.Contract.ValidatorStatus(&_ContractIEigenPod.CallOpts, validatorPubkey)
}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) ValidatorStatus(validatorPubkey []byte) (uint8, error) {
	return _ContractIEigenPod.Contract.ValidatorStatus(&_ContractIEigenPod.CallOpts, validatorPubkey)
}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_ContractIEigenPod *ContractIEigenPodCaller) ValidatorStatus0(opts *bind.CallOpts, pubkeyHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "validatorStatus0", pubkeyHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_ContractIEigenPod *ContractIEigenPodSession) ValidatorStatus0(pubkeyHash [32]byte) (uint8, error) {
	return _ContractIEigenPod.Contract.ValidatorStatus0(&_ContractIEigenPod.CallOpts, pubkeyHash)
}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) ValidatorStatus0(pubkeyHash [32]byte) (uint8, error) {
	return _ContractIEigenPod.Contract.ValidatorStatus0(&_ContractIEigenPod.CallOpts, pubkeyHash)
}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCaller) WithdrawableRestakedExecutionLayerGwei(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "withdrawableRestakedExecutionLayerGwei")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodSession) WithdrawableRestakedExecutionLayerGwei() (uint64, error) {
	return _ContractIEigenPod.Contract.WithdrawableRestakedExecutionLayerGwei(&_ContractIEigenPod.CallOpts)
}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) WithdrawableRestakedExecutionLayerGwei() (uint64, error) {
	return _ContractIEigenPod.Contract.WithdrawableRestakedExecutionLayerGwei(&_ContractIEigenPod.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) Initialize(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "initialize", owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.Initialize(&_ContractIEigenPod.TransactOpts, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.Initialize(&_ContractIEigenPod.TransactOpts, owner)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) RecoverTokens(opts *bind.TransactOpts, tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "recoverTokens", tokenList, amountsToWithdraw, recipient)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.RecoverTokens(&_ContractIEigenPod.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.RecoverTokens(&_ContractIEigenPod.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address newProofSubmitter) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) SetProofSubmitter(opts *bind.TransactOpts, newProofSubmitter common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "setProofSubmitter", newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address newProofSubmitter) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) SetProofSubmitter(newProofSubmitter common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.SetProofSubmitter(&_ContractIEigenPod.TransactOpts, newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address newProofSubmitter) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) SetProofSubmitter(newProofSubmitter common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.SetProofSubmitter(&_ContractIEigenPod.TransactOpts, newProofSubmitter)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "stake", pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_ContractIEigenPod *ContractIEigenPodSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.Stake(&_ContractIEigenPod.TransactOpts, pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.Stake(&_ContractIEigenPod.TransactOpts, pubkey, signature, depositDataRoot)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x88676cad.
//
// Solidity: function startCheckpoint(bool revertIfNoBalance) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) StartCheckpoint(opts *bind.TransactOpts, revertIfNoBalance bool) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "startCheckpoint", revertIfNoBalance)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x88676cad.
//
// Solidity: function startCheckpoint(bool revertIfNoBalance) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) StartCheckpoint(revertIfNoBalance bool) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.StartCheckpoint(&_ContractIEigenPod.TransactOpts, revertIfNoBalance)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x88676cad.
//
// Solidity: function startCheckpoint(bool revertIfNoBalance) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) StartCheckpoint(revertIfNoBalance bool) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.StartCheckpoint(&_ContractIEigenPod.TransactOpts, revertIfNoBalance)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) VerifyCheckpointProofs(opts *bind.TransactOpts, balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "verifyCheckpointProofs", balanceContainerProof, proofs)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) VerifyCheckpointProofs(balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.VerifyCheckpointProofs(&_ContractIEigenPod.TransactOpts, balanceContainerProof, proofs)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) VerifyCheckpointProofs(balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.VerifyCheckpointProofs(&_ContractIEigenPod.TransactOpts, balanceContainerProof, proofs)
}

// VerifyStaleBalance is a paid mutator transaction binding the contract method 0x039157d2.
//
// Solidity: function verifyStaleBalance(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, (bytes32[],bytes) proof) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) VerifyStaleBalance(opts *bind.TransactOpts, beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, proof BeaconChainProofsValidatorProof) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "verifyStaleBalance", beaconTimestamp, stateRootProof, proof)
}

// VerifyStaleBalance is a paid mutator transaction binding the contract method 0x039157d2.
//
// Solidity: function verifyStaleBalance(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, (bytes32[],bytes) proof) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) VerifyStaleBalance(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, proof BeaconChainProofsValidatorProof) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.VerifyStaleBalance(&_ContractIEigenPod.TransactOpts, beaconTimestamp, stateRootProof, proof)
}

// VerifyStaleBalance is a paid mutator transaction binding the contract method 0x039157d2.
//
// Solidity: function verifyStaleBalance(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, (bytes32[],bytes) proof) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) VerifyStaleBalance(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, proof BeaconChainProofsValidatorProof) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.VerifyStaleBalance(&_ContractIEigenPod.TransactOpts, beaconTimestamp, stateRootProof, proof)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) VerifyWithdrawalCredentials(opts *bind.TransactOpts, beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "verifyWithdrawalCredentials", beaconTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) VerifyWithdrawalCredentials(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.VerifyWithdrawalCredentials(&_ContractIEigenPod.TransactOpts, beaconTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) VerifyWithdrawalCredentials(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.VerifyWithdrawalCredentials(&_ContractIEigenPod.TransactOpts, beaconTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) WithdrawRestakedBeaconChainETH(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "withdrawRestakedBeaconChainETH", recipient, amount)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) WithdrawRestakedBeaconChainETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.WithdrawRestakedBeaconChainETH(&_ContractIEigenPod.TransactOpts, recipient, amount)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) WithdrawRestakedBeaconChainETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.WithdrawRestakedBeaconChainETH(&_ContractIEigenPod.TransactOpts, recipient, amount)
}

// ContractIEigenPodCheckpointCreatedIterator is returned from FilterCheckpointCreated and is used to iterate over the raw logs and unpacked data for CheckpointCreated events raised by the ContractIEigenPod contract.
type ContractIEigenPodCheckpointCreatedIterator struct {
	Event *ContractIEigenPodCheckpointCreated // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodCheckpointCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodCheckpointCreated)
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
		it.Event = new(ContractIEigenPodCheckpointCreated)
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
func (it *ContractIEigenPodCheckpointCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodCheckpointCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodCheckpointCreated represents a CheckpointCreated event raised by the ContractIEigenPod contract.
type ContractIEigenPodCheckpointCreated struct {
	CheckpointTimestamp uint64
	BeaconBlockRoot     [32]byte
	ValidatorCount      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCheckpointCreated is a free log retrieval operation binding the contract event 0x575796133bbed337e5b39aa49a30dc2556a91e0c6c2af4b7b886ae77ebef1076.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot, uint256 validatorCount)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterCheckpointCreated(opts *bind.FilterOpts, checkpointTimestamp []uint64, beaconBlockRoot [][32]byte) (*ContractIEigenPodCheckpointCreatedIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var beaconBlockRootRule []interface{}
	for _, beaconBlockRootItem := range beaconBlockRoot {
		beaconBlockRootRule = append(beaconBlockRootRule, beaconBlockRootItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "CheckpointCreated", checkpointTimestampRule, beaconBlockRootRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodCheckpointCreatedIterator{contract: _ContractIEigenPod.contract, event: "CheckpointCreated", logs: logs, sub: sub}, nil
}

// WatchCheckpointCreated is a free log subscription operation binding the contract event 0x575796133bbed337e5b39aa49a30dc2556a91e0c6c2af4b7b886ae77ebef1076.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot, uint256 validatorCount)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchCheckpointCreated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodCheckpointCreated, checkpointTimestamp []uint64, beaconBlockRoot [][32]byte) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var beaconBlockRootRule []interface{}
	for _, beaconBlockRootItem := range beaconBlockRoot {
		beaconBlockRootRule = append(beaconBlockRootRule, beaconBlockRootItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "CheckpointCreated", checkpointTimestampRule, beaconBlockRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodCheckpointCreated)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "CheckpointCreated", log); err != nil {
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

// ParseCheckpointCreated is a log parse operation binding the contract event 0x575796133bbed337e5b39aa49a30dc2556a91e0c6c2af4b7b886ae77ebef1076.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot, uint256 validatorCount)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseCheckpointCreated(log types.Log) (*ContractIEigenPodCheckpointCreated, error) {
	event := new(ContractIEigenPodCheckpointCreated)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "CheckpointCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodCheckpointFinalizedIterator is returned from FilterCheckpointFinalized and is used to iterate over the raw logs and unpacked data for CheckpointFinalized events raised by the ContractIEigenPod contract.
type ContractIEigenPodCheckpointFinalizedIterator struct {
	Event *ContractIEigenPodCheckpointFinalized // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodCheckpointFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodCheckpointFinalized)
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
		it.Event = new(ContractIEigenPodCheckpointFinalized)
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
func (it *ContractIEigenPodCheckpointFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodCheckpointFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodCheckpointFinalized represents a CheckpointFinalized event raised by the ContractIEigenPod contract.
type ContractIEigenPodCheckpointFinalized struct {
	CheckpointTimestamp uint64
	TotalShareDeltaWei  *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCheckpointFinalized is a free log retrieval operation binding the contract event 0x525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e44.
//
// Solidity: event CheckpointFinalized(uint64 indexed checkpointTimestamp, int256 totalShareDeltaWei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterCheckpointFinalized(opts *bind.FilterOpts, checkpointTimestamp []uint64) (*ContractIEigenPodCheckpointFinalizedIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "CheckpointFinalized", checkpointTimestampRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodCheckpointFinalizedIterator{contract: _ContractIEigenPod.contract, event: "CheckpointFinalized", logs: logs, sub: sub}, nil
}

// WatchCheckpointFinalized is a free log subscription operation binding the contract event 0x525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e44.
//
// Solidity: event CheckpointFinalized(uint64 indexed checkpointTimestamp, int256 totalShareDeltaWei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchCheckpointFinalized(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodCheckpointFinalized, checkpointTimestamp []uint64) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "CheckpointFinalized", checkpointTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodCheckpointFinalized)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "CheckpointFinalized", log); err != nil {
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

// ParseCheckpointFinalized is a log parse operation binding the contract event 0x525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e44.
//
// Solidity: event CheckpointFinalized(uint64 indexed checkpointTimestamp, int256 totalShareDeltaWei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseCheckpointFinalized(log types.Log) (*ContractIEigenPodCheckpointFinalized, error) {
	event := new(ContractIEigenPodCheckpointFinalized)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "CheckpointFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodEigenPodStakedIterator is returned from FilterEigenPodStaked and is used to iterate over the raw logs and unpacked data for EigenPodStaked events raised by the ContractIEigenPod contract.
type ContractIEigenPodEigenPodStakedIterator struct {
	Event *ContractIEigenPodEigenPodStaked // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodEigenPodStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodEigenPodStaked)
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
		it.Event = new(ContractIEigenPodEigenPodStaked)
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
func (it *ContractIEigenPodEigenPodStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodEigenPodStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodEigenPodStaked represents a EigenPodStaked event raised by the ContractIEigenPod contract.
type ContractIEigenPodEigenPodStaked struct {
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEigenPodStaked is a free log retrieval operation binding the contract event 0x606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e23.
//
// Solidity: event EigenPodStaked(bytes pubkey)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterEigenPodStaked(opts *bind.FilterOpts) (*ContractIEigenPodEigenPodStakedIterator, error) {

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "EigenPodStaked")
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodEigenPodStakedIterator{contract: _ContractIEigenPod.contract, event: "EigenPodStaked", logs: logs, sub: sub}, nil
}

// WatchEigenPodStaked is a free log subscription operation binding the contract event 0x606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e23.
//
// Solidity: event EigenPodStaked(bytes pubkey)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchEigenPodStaked(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodEigenPodStaked) (event.Subscription, error) {

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "EigenPodStaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodEigenPodStaked)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "EigenPodStaked", log); err != nil {
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

// ParseEigenPodStaked is a log parse operation binding the contract event 0x606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e23.
//
// Solidity: event EigenPodStaked(bytes pubkey)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseEigenPodStaked(log types.Log) (*ContractIEigenPodEigenPodStaked, error) {
	event := new(ContractIEigenPodEigenPodStaked)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "EigenPodStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodNonBeaconChainETHReceivedIterator is returned from FilterNonBeaconChainETHReceived and is used to iterate over the raw logs and unpacked data for NonBeaconChainETHReceived events raised by the ContractIEigenPod contract.
type ContractIEigenPodNonBeaconChainETHReceivedIterator struct {
	Event *ContractIEigenPodNonBeaconChainETHReceived // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodNonBeaconChainETHReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodNonBeaconChainETHReceived)
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
		it.Event = new(ContractIEigenPodNonBeaconChainETHReceived)
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
func (it *ContractIEigenPodNonBeaconChainETHReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodNonBeaconChainETHReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodNonBeaconChainETHReceived represents a NonBeaconChainETHReceived event raised by the ContractIEigenPod contract.
type ContractIEigenPodNonBeaconChainETHReceived struct {
	AmountReceived *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNonBeaconChainETHReceived is a free log retrieval operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterNonBeaconChainETHReceived(opts *bind.FilterOpts) (*ContractIEigenPodNonBeaconChainETHReceivedIterator, error) {

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "NonBeaconChainETHReceived")
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodNonBeaconChainETHReceivedIterator{contract: _ContractIEigenPod.contract, event: "NonBeaconChainETHReceived", logs: logs, sub: sub}, nil
}

// WatchNonBeaconChainETHReceived is a free log subscription operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchNonBeaconChainETHReceived(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodNonBeaconChainETHReceived) (event.Subscription, error) {

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "NonBeaconChainETHReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodNonBeaconChainETHReceived)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "NonBeaconChainETHReceived", log); err != nil {
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

// ParseNonBeaconChainETHReceived is a log parse operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseNonBeaconChainETHReceived(log types.Log) (*ContractIEigenPodNonBeaconChainETHReceived, error) {
	event := new(ContractIEigenPodNonBeaconChainETHReceived)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "NonBeaconChainETHReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodProofSubmitterUpdatedIterator is returned from FilterProofSubmitterUpdated and is used to iterate over the raw logs and unpacked data for ProofSubmitterUpdated events raised by the ContractIEigenPod contract.
type ContractIEigenPodProofSubmitterUpdatedIterator struct {
	Event *ContractIEigenPodProofSubmitterUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodProofSubmitterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodProofSubmitterUpdated)
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
		it.Event = new(ContractIEigenPodProofSubmitterUpdated)
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
func (it *ContractIEigenPodProofSubmitterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodProofSubmitterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodProofSubmitterUpdated represents a ProofSubmitterUpdated event raised by the ContractIEigenPod contract.
type ContractIEigenPodProofSubmitterUpdated struct {
	PrevProofSubmitter common.Address
	NewProofSubmitter  common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProofSubmitterUpdated is a free log retrieval operation binding the contract event 0xfb8129080a19d34dceac04ba253fc50304dc86c729bd63cdca4a969ad19a5eac.
//
// Solidity: event ProofSubmitterUpdated(address prevProofSubmitter, address newProofSubmitter)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterProofSubmitterUpdated(opts *bind.FilterOpts) (*ContractIEigenPodProofSubmitterUpdatedIterator, error) {

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "ProofSubmitterUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodProofSubmitterUpdatedIterator{contract: _ContractIEigenPod.contract, event: "ProofSubmitterUpdated", logs: logs, sub: sub}, nil
}

// WatchProofSubmitterUpdated is a free log subscription operation binding the contract event 0xfb8129080a19d34dceac04ba253fc50304dc86c729bd63cdca4a969ad19a5eac.
//
// Solidity: event ProofSubmitterUpdated(address prevProofSubmitter, address newProofSubmitter)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchProofSubmitterUpdated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodProofSubmitterUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "ProofSubmitterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodProofSubmitterUpdated)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "ProofSubmitterUpdated", log); err != nil {
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

// ParseProofSubmitterUpdated is a log parse operation binding the contract event 0xfb8129080a19d34dceac04ba253fc50304dc86c729bd63cdca4a969ad19a5eac.
//
// Solidity: event ProofSubmitterUpdated(address prevProofSubmitter, address newProofSubmitter)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseProofSubmitterUpdated(log types.Log) (*ContractIEigenPodProofSubmitterUpdated, error) {
	event := new(ContractIEigenPodProofSubmitterUpdated)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "ProofSubmitterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator is returned from FilterRestakedBeaconChainETHWithdrawn and is used to iterate over the raw logs and unpacked data for RestakedBeaconChainETHWithdrawn events raised by the ContractIEigenPod contract.
type ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator struct {
	Event *ContractIEigenPodRestakedBeaconChainETHWithdrawn // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodRestakedBeaconChainETHWithdrawn)
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
		it.Event = new(ContractIEigenPodRestakedBeaconChainETHWithdrawn)
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
func (it *ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodRestakedBeaconChainETHWithdrawn represents a RestakedBeaconChainETHWithdrawn event raised by the ContractIEigenPod contract.
type ContractIEigenPodRestakedBeaconChainETHWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRestakedBeaconChainETHWithdrawn is a free log retrieval operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterRestakedBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "RestakedBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator{contract: _ContractIEigenPod.contract, event: "RestakedBeaconChainETHWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRestakedBeaconChainETHWithdrawn is a free log subscription operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchRestakedBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodRestakedBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "RestakedBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodRestakedBeaconChainETHWithdrawn)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "RestakedBeaconChainETHWithdrawn", log); err != nil {
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

// ParseRestakedBeaconChainETHWithdrawn is a log parse operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseRestakedBeaconChainETHWithdrawn(log types.Log) (*ContractIEigenPodRestakedBeaconChainETHWithdrawn, error) {
	event := new(ContractIEigenPodRestakedBeaconChainETHWithdrawn)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "RestakedBeaconChainETHWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodValidatorBalanceUpdatedIterator is returned from FilterValidatorBalanceUpdated and is used to iterate over the raw logs and unpacked data for ValidatorBalanceUpdated events raised by the ContractIEigenPod contract.
type ContractIEigenPodValidatorBalanceUpdatedIterator struct {
	Event *ContractIEigenPodValidatorBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodValidatorBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodValidatorBalanceUpdated)
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
		it.Event = new(ContractIEigenPodValidatorBalanceUpdated)
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
func (it *ContractIEigenPodValidatorBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodValidatorBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodValidatorBalanceUpdated represents a ValidatorBalanceUpdated event raised by the ContractIEigenPod contract.
type ContractIEigenPodValidatorBalanceUpdated struct {
	ValidatorIndex          *big.Int
	BalanceTimestamp        uint64
	NewValidatorBalanceGwei uint64
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterValidatorBalanceUpdated is a free log retrieval operation binding the contract event 0x0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df.
//
// Solidity: event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterValidatorBalanceUpdated(opts *bind.FilterOpts) (*ContractIEigenPodValidatorBalanceUpdatedIterator, error) {

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "ValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodValidatorBalanceUpdatedIterator{contract: _ContractIEigenPod.contract, event: "ValidatorBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorBalanceUpdated is a free log subscription operation binding the contract event 0x0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df.
//
// Solidity: event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchValidatorBalanceUpdated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodValidatorBalanceUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "ValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodValidatorBalanceUpdated)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "ValidatorBalanceUpdated", log); err != nil {
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

// ParseValidatorBalanceUpdated is a log parse operation binding the contract event 0x0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df.
//
// Solidity: event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseValidatorBalanceUpdated(log types.Log) (*ContractIEigenPodValidatorBalanceUpdated, error) {
	event := new(ContractIEigenPodValidatorBalanceUpdated)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "ValidatorBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodValidatorCheckpointedIterator is returned from FilterValidatorCheckpointed and is used to iterate over the raw logs and unpacked data for ValidatorCheckpointed events raised by the ContractIEigenPod contract.
type ContractIEigenPodValidatorCheckpointedIterator struct {
	Event *ContractIEigenPodValidatorCheckpointed // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodValidatorCheckpointedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodValidatorCheckpointed)
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
		it.Event = new(ContractIEigenPodValidatorCheckpointed)
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
func (it *ContractIEigenPodValidatorCheckpointedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodValidatorCheckpointedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodValidatorCheckpointed represents a ValidatorCheckpointed event raised by the ContractIEigenPod contract.
type ContractIEigenPodValidatorCheckpointed struct {
	CheckpointTimestamp uint64
	ValidatorIndex      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorCheckpointed is a free log retrieval operation binding the contract event 0xa91c59033c3423e18b54d0acecebb4972f9ea95aedf5f4cae3b677b02eaf3a3f.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterValidatorCheckpointed(opts *bind.FilterOpts, checkpointTimestamp []uint64, validatorIndex []*big.Int) (*ContractIEigenPodValidatorCheckpointedIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "ValidatorCheckpointed", checkpointTimestampRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodValidatorCheckpointedIterator{contract: _ContractIEigenPod.contract, event: "ValidatorCheckpointed", logs: logs, sub: sub}, nil
}

// WatchValidatorCheckpointed is a free log subscription operation binding the contract event 0xa91c59033c3423e18b54d0acecebb4972f9ea95aedf5f4cae3b677b02eaf3a3f.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchValidatorCheckpointed(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodValidatorCheckpointed, checkpointTimestamp []uint64, validatorIndex []*big.Int) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "ValidatorCheckpointed", checkpointTimestampRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodValidatorCheckpointed)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "ValidatorCheckpointed", log); err != nil {
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

// ParseValidatorCheckpointed is a log parse operation binding the contract event 0xa91c59033c3423e18b54d0acecebb4972f9ea95aedf5f4cae3b677b02eaf3a3f.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseValidatorCheckpointed(log types.Log) (*ContractIEigenPodValidatorCheckpointed, error) {
	event := new(ContractIEigenPodValidatorCheckpointed)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "ValidatorCheckpointed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodValidatorRestakedIterator is returned from FilterValidatorRestaked and is used to iterate over the raw logs and unpacked data for ValidatorRestaked events raised by the ContractIEigenPod contract.
type ContractIEigenPodValidatorRestakedIterator struct {
	Event *ContractIEigenPodValidatorRestaked // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodValidatorRestakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodValidatorRestaked)
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
		it.Event = new(ContractIEigenPodValidatorRestaked)
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
func (it *ContractIEigenPodValidatorRestakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodValidatorRestakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodValidatorRestaked represents a ValidatorRestaked event raised by the ContractIEigenPod contract.
type ContractIEigenPodValidatorRestaked struct {
	ValidatorIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorRestaked is a free log retrieval operation binding the contract event 0x2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449.
//
// Solidity: event ValidatorRestaked(uint40 validatorIndex)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterValidatorRestaked(opts *bind.FilterOpts) (*ContractIEigenPodValidatorRestakedIterator, error) {

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "ValidatorRestaked")
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodValidatorRestakedIterator{contract: _ContractIEigenPod.contract, event: "ValidatorRestaked", logs: logs, sub: sub}, nil
}

// WatchValidatorRestaked is a free log subscription operation binding the contract event 0x2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449.
//
// Solidity: event ValidatorRestaked(uint40 validatorIndex)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchValidatorRestaked(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodValidatorRestaked) (event.Subscription, error) {

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "ValidatorRestaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodValidatorRestaked)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "ValidatorRestaked", log); err != nil {
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

// ParseValidatorRestaked is a log parse operation binding the contract event 0x2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449.
//
// Solidity: event ValidatorRestaked(uint40 validatorIndex)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseValidatorRestaked(log types.Log) (*ContractIEigenPodValidatorRestaked, error) {
	event := new(ContractIEigenPodValidatorRestaked)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "ValidatorRestaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodValidatorWithdrawnIterator is returned from FilterValidatorWithdrawn and is used to iterate over the raw logs and unpacked data for ValidatorWithdrawn events raised by the ContractIEigenPod contract.
type ContractIEigenPodValidatorWithdrawnIterator struct {
	Event *ContractIEigenPodValidatorWithdrawn // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodValidatorWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodValidatorWithdrawn)
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
		it.Event = new(ContractIEigenPodValidatorWithdrawn)
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
func (it *ContractIEigenPodValidatorWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodValidatorWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodValidatorWithdrawn represents a ValidatorWithdrawn event raised by the ContractIEigenPod contract.
type ContractIEigenPodValidatorWithdrawn struct {
	CheckpointTimestamp uint64
	ValidatorIndex      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorWithdrawn is a free log retrieval operation binding the contract event 0x2a02361ffa66cf2c2da4682c2355a6adcaa9f6c227b6e6563e68480f9587626a.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterValidatorWithdrawn(opts *bind.FilterOpts, checkpointTimestamp []uint64, validatorIndex []*big.Int) (*ContractIEigenPodValidatorWithdrawnIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "ValidatorWithdrawn", checkpointTimestampRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodValidatorWithdrawnIterator{contract: _ContractIEigenPod.contract, event: "ValidatorWithdrawn", logs: logs, sub: sub}, nil
}

// WatchValidatorWithdrawn is a free log subscription operation binding the contract event 0x2a02361ffa66cf2c2da4682c2355a6adcaa9f6c227b6e6563e68480f9587626a.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchValidatorWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodValidatorWithdrawn, checkpointTimestamp []uint64, validatorIndex []*big.Int) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "ValidatorWithdrawn", checkpointTimestampRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodValidatorWithdrawn)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
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

// ParseValidatorWithdrawn is a log parse operation binding the contract event 0x2a02361ffa66cf2c2da4682c2355a6adcaa9f6c227b6e6563e68480f9587626a.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseValidatorWithdrawn(log types.Log) (*ContractIEigenPodValidatorWithdrawn, error) {
	event := new(ContractIEigenPodValidatorWithdrawn)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
