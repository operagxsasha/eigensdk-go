// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractRegistryCoordinator

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSApkRegistryPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// IRegistryCoordinatorOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorInfo struct {
	OperatorId [32]byte
	Status     uint8
}

// IRegistryCoordinatorOperatorKickParam is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorKickParam struct {
	QuorumNumber uint8
	Operator     common.Address
}

// IRegistryCoordinatorOperatorSetParam is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorSetParam struct {
	MaxOperatorCount        uint32
	KickBIPsOfOperatorStake uint16
	KickBIPsOfTotalStake    uint16
}

// IRegistryCoordinatorQuorumBitmapUpdate is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorQuorumBitmapUpdate struct {
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
	QuorumBitmap          *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// IStakeRegistryStrategyParams is an auto generated low-level Go binding around an user-defined struct.
type IStakeRegistryStrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ContractRegistryCoordinatorMetaData contains all meta data concerning the ContractRegistryCoordinator contract.
var ContractRegistryCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_serviceManager\",\"type\":\"address\",\"internalType\":\"contractIServiceManager\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_blsApkRegistry\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"},{\"name\":\"_indexRegistry\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"},{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_CHURN_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUBKEY_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorChurnApprovalDigestHash\",\"inputs\":[{\"name\":\"registeringOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registeringOperatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"churnApprover\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createSlashableStakeQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistry.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"lookAheadPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createTotalDelegatedStakeQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistry.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectionCooldown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ejector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"enableOperatorSets\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentQuorumBitmap\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorInfo\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromId\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorStatus\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapAtBlockNumberByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapHistoryLength\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapUpdateByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.QuorumBitmapUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBitmap\",\"type\":\"uint192\",\"internalType\":\"uint192\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"indexRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_operatorSetParams\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam[]\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"_minimumStakes\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"_strategyParams\",\"type\":\"tuple[][]\",\"internalType\":\"structIStakeRegistry.StrategyParams[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"_stakeTypes\",\"type\":\"uint8[]\",\"internalType\":\"enumStakeType[]\"},{\"name\":\"_lookAheadPeriods\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isChurnApproverSaltUsed\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isM2Quorum\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSetAVS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isUsingOperatorSets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastEjectionTimestamp\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numRegistries\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumUpdateBlockNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorWithChurn\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"churnApproverSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registries\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serviceManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIServiceManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setChurnApprover\",\"inputs\":[{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjectionCooldown\",\"inputs\":[{\"name\":\"_ejectionCooldown\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjector\",\"inputs\":[{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorsForQuorum\",\"inputs\":[{\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChurnApproverUpdated\",\"inputs\":[{\"name\":\"prevChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EjectorUpdated\",\"inputs\":[{\"name\":\"prevEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetParamsUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSocketUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumBlockNumberUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"blocknumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRegisteredForQuorums\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotChurnSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotKickOperatorAboveThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotReregisterYet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChurnApproverSaltUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientStakeForChurn\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxQuorumsReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegisteredForQuorum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAllocationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEjector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetsEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetsNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumOperatorCountMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RegistryCoordinatorSignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltAlreadyUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]}]",
	Bin: "0x610200604052348015610010575f5ffd5b506040516163a33803806163a383398101604081905261002f9161027e565b604080518082018252601681527f4156535265676973747279436f6f7264696e61746f720000000000000000000060208083019182528351808501909452600684526576302e302e3160d01b908401528151902060e08190527f6bda7e3f385e48841048390444cced5cc795af87758af67622e5f4f0882c4a996101008190524660a0528993899389938993899389939290917f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6101318184846040805160208101859052908101839052606081018290524660808201523060a08201525f9060c0016040516020818303038152906040528051906020012090509392505050565b6080523060c052610120525050506001600160a01b0382169050610168576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811661014052948516610160529284166101a0529083166101805282166101c052166101e05261019f6101aa565b505050505050610301565b5f54610100900460ff16156102155760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610265575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461027b575f5ffd5b50565b5f5f5f5f5f5f60c08789031215610293575f5ffd5b865161029e81610267565b60208801519096506102af81610267565b60408801519095506102c081610267565b60608801519094506102d181610267565b60808801519093506102e281610267565b60a08801519092506102f381610267565b809150509295509295509295565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516101c0516101e051615f4661045d5f395f6106ad01525f81816107af01528181610f440152818161159901528181612b49015281816130580152818161365d0152613b4b01525f8181610686015281816115240152818161230d01528181612a2d01528181612ac801528181612fdd015281816135bb015281816137700152613aa901525f818161064c01528181610cf40152818161156201528181612bbe01528181612f64015281816131030152818161317a015281816135420152613a3001525f818161057d01528181611a5e01528181611d6901528181611fc60152818161299901528181612ec0015261349001525f818161073b01528181610b8e0152818161121c01526120e601525f613da301525f613df201525f613dcd01525f613d2601525f613d5001525f613d7a0152615f465ff3fe608060405234801561000f575f5ffd5b5060043610610371575f3560e01c80636b3aa72e116101d4578063a4d7871f11610109578063ca4f2d97116100a9578063ee31882111610079578063ee31882114610965578063f2fde38b1461096d578063fabc1cbc14610980578063fd39105a14610993575f5ffd5b8063ca4f2d971461089b578063cabbb17f146108ae578063d72d8dd6146108bb578063e65797ad146108c3575f5ffd5b8063adcf73f7116100e4578063adcf73f714610836578063bd33ee2414610849578063c391425e14610854578063ca0de88214610874575f5ffd5b8063a4d7871f146107f8578063a50857bf1461081a578063a96f783e1461082d575f5ffd5b8063886f1195116101745780639b5d177b1161014f5780639b5d177b146107845780639d8e0c23146107975780639e9923c2146107aa5780639feab859146107d1575f5ffd5b8063886f1195146107365780638da5cb5b1461075d5780639aa1653d14610765575f5ffd5b80637fc3f886116101af5780637fc3f886146106ea5780638281ab75146106fd57806384ca521314610710578063871ef04914610723575f5ffd5b80636b3aa72e146106a85780636e3b17db146106cf578063715018a6146106e2575f5ffd5b806329d1e0c3116102aa578063595c6a671161024a5780635c975abb116102255780635c975abb1461063f5780635df45946146106475780636347c9001461066e5780636830483514610681575f5ffd5b8063595c6a67146106055780635ac86ab71461060d5780635b0b829f1461062c575f5ffd5b80633c2a7f4c116102855780633c2a7f4c1461059f5780633eef3a51146105bf5780635140a548146105d25780635865c60c146105e5575f5ffd5b806329d1e0c3146105525780632cdd1e86146105655780633998fdd314610578575f5ffd5b806313542a4e116103155780631eb812da116102f05780631eb812da146104c4578063249a0c421461050d57806328f61b311461052c578063296bb0641461053f575f5ffd5b806313542a4e14610457578063136439dd1461047f5780631478851f14610492575f5ffd5b8063054310e611610350578063054310e6146103e75780630cf4b767146104125780630d3f213414610425578063125e058414610438575f5ffd5b8062cf2ab51461037557806303fd34921461038a57806304ec6351146103bc575b5f5ffd5b610388610383366004614878565b6109ce565b005b6103a96103983660046148a9565b5f9081526098602052604090205490565b6040519081526020015b60405180910390f35b6103cf6103ca3660046148d1565b610ac3565b6040516001600160c01b0390911681526020016103b3565b609d546103fa906001600160a01b031681565b6040516001600160a01b0390911681526020016103b3565b61038861042036600461497b565b610adb565b6103886104333660046148a9565b610b6c565b6103a96104463660046149ac565b609f6020525f908152604090205481565b6103a96104653660046149ac565b6001600160a01b03165f9081526099602052604090205490565b61038861048d3660046148a9565b610b79565b6104b46104a03660046148a9565b609a6020525f908152604090205460ff1681565b60405190151581526020016103b3565b6104d76104d23660046149c7565b610c4e565b60408051825163ffffffff908116825260208085015190911690820152918101516001600160c01b0316908201526060016103b3565b6103a961051b3660046149f7565b609b6020525f908152604090205481565b609e546103fa906001600160a01b031681565b6103fa61054d3660046148a9565b610cdc565b6103886105603660046149ac565b610d65565b6103886105733660046149ac565b610d79565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b6105b26105ad3660046149ac565b610d8a565b6040516103b39190614a10565b6103886105cd366004614b31565b610e08565b6103886105e0366004614bdd565b610e47565b6105f86105f33660046149ac565b611195565b6040516103b39190614cee565b610388611207565b6104b461061b3660046149f7565b6001805460ff9092161b9081161490565b61038861063a366004614d09565b6112b6565b6001546103a9565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b6103fa61067c3660046148a9565b6112d2565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b6103886106dd366004614d3b565b6112fa565b6103886113a7565b6103886106f8366004614f99565b6113b8565b61038861070b3660046150b1565b6116b1565b6103a961071e3660046151a1565b6116c6565b6103cf6107313660046148a9565b61170f565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b6103fa611719565b6096546107729060ff1681565b60405160ff90911681526020016103b3565b61038861079236600461535e565b611731565b6103886107a5366004615456565b6119aa565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b6103a97f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de681565b6104b46108063660046149f7565b60a26020525f908152604090205460ff1681565b610388610828366004615498565b611bb7565b6103a960a05481565b610388610844366004615537565b611cb5565b60a15460ff166104b4565b6108676108623660046155a1565b611ef6565b6040516103b39190615646565b6103a97f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a81565b6103886108a936600461497b565b611f04565b60a1546104b49060ff1681565b609c546103a9565b6109316108d13660046149f7565b60408051606080820183525f808352602080840182905292840181905260ff9490941684526097825292829020825193840183525463ffffffff8116845261ffff600160201b8204811692850192909252600160301b9004169082015290565b60408051825163ffffffff16815260208084015161ffff9081169183019190915292820151909216908201526060016103b3565b610388611fa9565b61038861097b3660046149ac565b61206e565b61038861098e3660046148a9565b6120e4565b6109c16109a13660046149ac565b6001600160a01b03165f9081526099602052604090206001015460ff1690565b6040516103b39190615683565b6001546002906004908116036109f75760405163840a48d560e01b815260040160405180910390fd5b5f5b8251811015610abe575f838281518110610a1557610a15615691565b6020908102919091018101516001600160a01b0381165f9081526099835260408082208151808301909252805482526001810154939550919390929083019060ff166002811115610a6857610a68614cba565b6002811115610a7957610a79614cba565b90525080519091505f610a8b826121fb565b90505f610aa0826001600160c01b0316612207565b9050610aad8585836122d0565b5050600190930192506109f9915050565b505050565b5f610ad160988585856123b9565b90505b9392505050565b6001335f9081526099602052604090206001015460ff166002811115610b0357610b03614cba565b14610b215760405163aba4733960e01b815260040160405180910390fd5b335f90815260996020526040908190205490517fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa90610b619084906156d3565b60405180910390a250565b610b746125b0565b60a055565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610bdb573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bff91906156e5565b610c1c57604051631d77d47760e21b815260040160405180910390fd5b6001548181168114610c415760405163c61dca5d60e01b815260040160405180910390fd5b610c4a8261260f565b5050565b604080516060810182525f80825260208201819052918101919091525f838152609860205260409020805483908110610c8957610c89615691565b5f91825260209182902060408051606081018252919092015463ffffffff8082168352600160201b820416938201939093526001600160c01b03600160401b909304929092169082015290505b92915050565b6040516308f6629d60e31b8152600481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906347b314e890602401602060405180830381865afa158015610d41573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cd69190615704565b610d6d6125b0565b610d7681612646565b50565b610d816125b0565b610d76816126af565b604080518082019091525f8082526020820152610cd6610e037f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de684604051602001610de89291909182526001600160a01b0316602082015260400190565b60405160208183030381529060405280519060200120612718565b612764565b610e106125b0565b60a15460ff16610e3357604051635b77901960e01b815260040160405180910390fd5b610e418484846001856127ee565b50505050565b600154600290600490811603610e705760405163840a48d560e01b815260040160405180910390fd5b5f610eb684848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505060965460ff169150612c1c9050565b85519091508314610eda5760405163aaad13f760e01b815260040160405180910390fd5b5f5b8381101561118d575f858583818110610ef757610ef7615691565b895192013560f81c92505f9189915084908110610f1657610f16615691565b60209081029190910101516040516379a0849160e11b815260ff841660048201529091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f341092290602401602060405180830381865afa158015610f89573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fad919061571f565b63ffffffff16815114610fd357604051638e5aeee760e01b815260040160405180910390fd5b5f805b8251811015611134575f838281518110610ff257610ff2615691565b6020908102919091018101516001600160a01b0381165f9081526099835260408082208151808301909252805482526001810154939550919390929083019060ff16600281111561104557611045614cba565b600281111561105657611056614cba565b90525080519091505f611068826121fb565b905060016001600160c01b03821660ff8a161c81161461109b5760405163d053aa2160e01b815260040160405180910390fd5b856001600160a01b0316846001600160a01b0316116110cd5760405163ba50f91160e01b815260040160405180910390fd5b5061112783838e8b8f6110e182600161574e565b926110ee93929190615761565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506122d092505050565b5090925050600101610fd6565b5060ff83165f818152609b6020908152604091829020439081905591519182527f46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4910160405180910390a2505050806001019050610edc565b505050505050565b604080518082019091525f80825260208201526001600160a01b0382165f908152609960209081526040918290208251808401909352805483526001810154909183019060ff1660028111156111ed576111ed614cba565b60028111156111fe576111fe614cba565b90525092915050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611269573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061128d91906156e5565b6112aa57604051631d77d47760e21b815260040160405180910390fd5b6112b45f1961260f565b565b6112be6125b0565b816112c881612ca5565b610abe8383612cce565b609c81815481106112e1575f80fd5b5f918252602090912001546001600160a01b0316905081565b611302612d73565b6001600160a01b0382165f908152609f602090815260408083204290556099909152812080546096549192909161133d90859060ff16612c1c565b90505f611349836121fb565b905060018085015460ff16600281111561136557611365614cba565b14801561137a57506001600160c01b03821615155b801561139857506113986001600160c01b0383811690831681161490565b1561118d5761118d8686612d9e565b6113af6125b0565b6112b45f613091565b5f54610100900460ff16158080156113d657505f54600160ff909116105b806113ef5750303b1580156113ef57505f5460ff166001145b6114575760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015611478575f805461ff0019166101001790555b8451865114801561148a575083518551145b8015611497575082518451145b80156114a4575081518351145b6114c15760405163aaad13f760e01b815260040160405180910390fd5b6114ca8a613091565b6114d389612646565b6114dc8761260f565b6114e5886126af565b609c8054600181810183555f8381527faf85b9071dfafeac1409d3f1d19bafc9bc7c37974cde8df0ee6168f0086e539c92830180546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166001600160a01b03199283161790925585548085018755850180547f0000000000000000000000000000000000000000000000000000000000000000841690831617905585549384019095559190920180547f000000000000000000000000000000000000000000000000000000000000000090921691909316179091555b865181101561165f576116578782815181106115e2576115e2615691565b60200260200101518783815181106115fc576115fc615691565b602002602001015187848151811061161657611616615691565b602002602001015187858151811061163057611630615691565b602002602001015187868151811061164a5761164a615691565b60200260200101516127ee565b6001016115c4565b5080156116a5575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050505050565b6116b96125b0565b610abe8383835f5f6127ee565b5f6117057f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a8787878787604051602001610de896959493929190615788565b9695505050505050565b5f610cd6826121fb565b5f61172c6064546001600160a01b031690565b905090565b600180545f91908116036117585760405163840a48d560e01b815260040160405180910390fd5b60a15460ff161561177c57604051630b88306f60e01b815260040160405180910390fd5b8351871461179d5760405163aaad13f760e01b815260040160405180910390fd5b5f6117a833876130e2565b90506117b633828787613210565b5f6117fb33838c8c8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508e92508a91506132bb9050565b90505f5b8981101561199d575f60975f8d8d8581811061181d5761181d615691565b919091013560f81c82525060208082019290925260409081015f208151606081018352905463ffffffff811680835261ffff600160201b8304811695840195909552600160301b9091049093169181019190915284518051919350908490811061188957611889615691565b602002602001015163ffffffff1611156119945761191c8c8c848181106118b2576118b2615691565b9050013560f81c60f81b60f81c846040015184815181106118d5576118d5615691565b602002602001015133866020015186815181106118f4576118f4615691565b60200260200101518c878151811061190e5761190e615691565b6020026020010151866136e3565b61199488838151811061193157611931615691565b6020026020010151602001518d8d859086600161194e919061574e565b9261195b93929190615761565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612d9e92505050565b506001016117ff565b5050505050505050505050565b600180545f91908116036119d15760405163840a48d560e01b815260040160405180910390fd5b60a15460ff166119f457604051635b77901960e01b815260040160405180910390fd5b5f5b8251811015611a5b5760a25f848381518110611a1457611a14615691565b60209081029190910181015160ff9081168352908201929092526040015f20541615611a535760405163fd2c1f4d60e01b815260040160405180910390fd5b6001016119f6565b507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ca8aa7c76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ab8573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611adc9190615704565b6001600160a01b0316336001600160a01b031614611b0d576040516323d871a560e01b815260040160405180910390fd5b5f82516001600160401b03811115611b2757611b27614731565b6040519080825280601f01601f191660200182016040528015611b51576020820181803683370190505b5090505f5b8351811015611bac57838181518110611b7157611b71615691565b602002602001015160f81b828281518110611b8e57611b8e615691565b60200101906001600160f81b03191690815f1a905350600101611b56565b50610e418482612d9e565b600180545f9190811603611bde5760405163840a48d560e01b815260040160405180910390fd5b60a15460ff1615611c0257604051630b88306f60e01b815260040160405180910390fd5b5f611c0d33856130e2565b90505f611c1d33838989886132bb565b5190505f5b8751811015611cab575f888281518110611c3e57611c3e615691565b0160209081015160f81c5f8181526097909252604090912054845191925063ffffffff1690849084908110611c7557611c75615691565b602002602001015163ffffffff161115611ca257604051633cb89c9760e01b815260040160405180910390fd5b50600101611c22565b5050505050505050565b600180545f9190811603611cdc5760405163840a48d560e01b815260040160405180910390fd5b60a15460ff16611cff57604051635b77901960e01b815260040160405180910390fd5b5f5b8351811015611d665760a25f858381518110611d1f57611d1f615691565b60209081029190910181015160ff9081168352908201929092526040015f20541615611d5e5760405163fd2c1f4d60e01b815260040160405180910390fd5b600101611d01565b507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ca8aa7c76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611dc3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611de79190615704565b6001600160a01b0316336001600160a01b031614611e18576040516323d871a560e01b815260040160405180910390fd5b5f5f83806020019051810190611e2e9190615880565b915091505f611e3d87836130e2565b90505f86516001600160401b03811115611e5957611e59614731565b6040519080825280601f01601f191660200182016040528015611e83576020820181803683370190505b5090505f5b8751811015611ede57878181518110611ea357611ea3615691565b602002602001015160f81b828281518110611ec057611ec0615691565b60200101906001600160f81b03191690815f1a905350600101611e88565b50611eeb88838387613859565b505050505050505050565b6060610ad460988484613bd0565b60018054600290811603611f2b5760405163840a48d560e01b815260040160405180910390fd5b5f5b8251811015611f9e575f838281518110611f4957611f49615691565b016020015160a15460f89190911c915060ff161580611f78575060ff8082165f90815260a26020526040902054165b611f9557604051630b88306f60e01b815260040160405180910390fd5b50600101611f2d565b50610c4a3383612d9e565b611fb16125b0565b604051630f25f16160e41b81523060048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063f25f1610906024015f604051808303815f87803b15801561200f575f5ffd5b505af1158015612021573d5f5f3e3d5ffd5b505f925050505b60965460ff908116908216101561205e5760ff81165f90815260a260205260409020805460ff1916600190811790915501612028565b5060a1805460ff19166001179055565b6120766125b0565b6001600160a01b0381166120db5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161144e565b610d7681613091565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612140573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121649190615704565b6001600160a01b0316336001600160a01b0316146121955760405163794821ff60e01b815260040160405180910390fd5b600154801982198116146121bc5760405163c61dca5d60e01b815260040160405180910390fd5b600182905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020015b60405180910390a25050565b5f610cd6609883613c7f565b60605f5f61221484613cf0565b61ffff166001600160401b0381111561222f5761222f614731565b6040519080825280601f01601f191660200182016040528015612259576020820181803683370190505b5090505f805b825182108015612270575061010081105b156122c6576001811b9350858416156122b6578060f81b83838151811061229957612299615691565b60200101906001600160f81b03191690815f1a9053508160010191505b6122bf8161597b565b905061225f565b5090949350505050565b6001826020015160028111156122e8576122e8614cba565b146122f257505050565b81516040516333567f7f60e11b81525f906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906366acfefe9061234690889086908890600401615993565b6020604051808303815f875af1158015612362573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061238691906159c2565b90506001600160c01b038116156123b2576123b2856123ad836001600160c01b0316612207565b612d9e565b5050505050565b5f8381526020859052604081208054829190849081106123db576123db615691565b5f91825260209182902060408051606081018252929091015463ffffffff808216808552600160201b8304821695850195909552600160401b9091046001600160c01b031691830191909152909250851610156124d45760405162461bcd60e51b815260206004820152606560248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d61704174426c6f636b4e756d6265724279496e6465783a2071756f72756d4260648201527f69746d61705570646174652069732066726f6d20616674657220626c6f636b4e6084820152643ab6b132b960d91b60a482015260c40161144e565b602081015163ffffffff1615806124fa5750806020015163ffffffff168463ffffffff16105b6125a15760405162461bcd60e51b815260206004820152606660248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d61704174426c6f636b4e756d6265724279496e6465783a2071756f72756d4260648201527f69746d61705570646174652069732066726f6d206265666f726520626c6f636b608482015265273ab6b132b960d11b60a482015260c40161144e565b6040015190505b949350505050565b336125b9611719565b6001600160a01b0316146112b45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161144e565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d90602001610b61565b609d54604080516001600160a01b03928316815291831660208301527f315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c910160405180910390a1609d80546001600160a01b0319166001600160a01b0392909216919091179055565b609e54604080516001600160a01b03928316815291831660208301527f8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9910160405180910390a1609e80546001600160a01b0319166001600160a01b0392909216919091179055565b5f610cd6612724613d1a565b8360405161190160f01b602082015260228101839052604281018290525f9060620160405160208183030381529060405280519060200120905092915050565b604080518082019091525f80825260208201525f80806127915f516020615ef15f395f51905f52866159fc565b90505b61279d81613e40565b90935091505f516020615ef15f395f51905f5282830983036127d5576040805180820190915290815260208101919091529392505050565b5f516020615ef15f395f51905f52600182089050612794565b60965460ff1660c0811061281557604051633cb89c9760e01b815260040160405180910390fd5b612820816001615a0f565b6096805460ff191660ff929092169190911790558061283f8188612cce565b60a15460ff168015612862575060ff8082165f90815260a2602052604090205416155b156129fe576040805160018082528183019092525f91816020015b604080518082019091525f81526060602082015281526020019060019003908161287d5790505090505f86516001600160401b038111156128c0576128c0614731565b6040519080825280602002602001820160405280156128e9578160200160208202803683370190505b5090505f5b87518110156129465787818151811061290957612909615691565b60200260200101515f015182828151811061292657612926615691565b6001600160a01b03909216602092830291909101909101526001016128ee565b5060405180604001604052808460ff1663ffffffff16815260200182815250825f8151811061297757612977615691565b602090810291909101015260405163847d634f60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063847d634f906129ce908590600401615a28565b5f604051808303815f87803b1580156129e5575f5ffd5b505af11580156129f7573d5f5f3e3d5ffd5b5050505050505b5f846001811115612a1157612a11614cba565b03612a9857604051633aea0b9d60e11b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906375d4173a90612a669084908a908a90600401615b33565b5f604051808303815f87803b158015612a7d575f5ffd5b505af1158015612a8f573d5f5f3e3d5ffd5b50505050612b31565b6001846001811115612aac57612aac614cba565b03612b3157604051630662d3e160e51b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063cc5a7c2090612b039084908a9088908b90600401615b5d565b5f604051808303815f87803b158015612b1a575f5ffd5b505af1158015612b2c573d5f5f3e3d5ffd5b505050505b60405163136ca0f960e11b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906326d941f2906024015f604051808303815f87803b158015612b92575f5ffd5b505af1158015612ba4573d5f5f3e3d5ffd5b505060405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f291506024015b5f604051808303815f87803b158015612c0a575f5ffd5b505af115801561199d573d5f5f3e3d5ffd5b5f5f612c2784613ebc565b9050808360ff166001901b11610ad45760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c756500606482015260840161144e565b60965460ff90811690821610610d7657604051637310cff560e11b815260040160405180910390fd5b60ff82165f81815260976020908152604091829020845181548684018051888701805163ffffffff90951665ffffffffffff199094168417600160201b61ffff938416021767ffff0000000000001916600160301b95831695909502949094179094558551918252518316938101939093525116918101919091527f3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac906060016121ef565b609e546001600160a01b031633146112b4576040516376d8ab1760e11b815260040160405180910390fd5b6001600160a01b0382165f908152609960205260409020805460018083015460ff166002811115612dd157612dd1614cba565b14612def5760405163aba4733960e01b815260040160405180910390fd5b6096545f90612e0290859060ff16612c1c565b90505f612e0e836121fb565b90506001600160c01b038216612e37576040516368b6a87560e11b815260040160405180910390fd5b612e4e6001600160c01b0383811690831681161490565b612e6b5760405163d053aa2160e01b815260040160405180910390fd5b6001600160c01b0382811619821616612e84848261403f565b6001600160c01b038116612f4d5760018501805460ff191660021790556040516351b27a6d60e11b81526001600160a01b0388811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015f604051808303815f87803b158015612f01575f5ffd5b505af1158015612f13573d5f5f3e3d5ffd5b50506040518692506001600160a01b038a1691507f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4905f90a35b60405163f4e24fe560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f4e24fe590612f9b908a908a90600401615b93565b5f604051808303815f87803b158015612fb2575f5ffd5b505af1158015612fc4573d5f5f3e3d5ffd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd91506130169087908a90600401615bb6565b5f604051808303815f87803b15801561302d575f5ffd5b505af115801561303f573d5f5f3e3d5ffd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd9150612bf39087908a90600401615bb6565b606480546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6040516309aa152760e11b81526001600160a01b0383811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906313542a4e90602401602060405180830381865afa15801561314a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061316e9190615bce565b90505f819003610cd6577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bf79ce5884846131b287610d8a565b6040518463ffffffff1660e01b81526004016131d093929190615c07565b6020604051808303815f875af11580156131ec573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ad49190615bce565b6020808201515f908152609a909152604090205460ff161561324557604051636fbefec360e11b815260040160405180910390fd5b428160400151101561326a57604051630819bdcd60e01b815260040160405180910390fd5b602080820180515f908152609a909252604091829020805460ff19166001179055609d54905191830151610e41926001600160a01b03909216916132b491889188918891906116c6565b835161404b565b6132df60405180606001604052806060815260200160608152602001606081525090565b6096545f906132f290869060ff16612c1c565b90505f6132fe876121fb565b90506001600160c01b038216613327576040516313ca465760e01b815260040160405180910390fd5b8082166001600160c01b03161561335157604051630c6816cd60e01b815260040160405180910390fd5b60a0546001600160a01b0389165f908152609f60205260409020546001600160c01b0383811690851617914291613388919061574e565b106133a657604051631968677d60e11b815260040160405180910390fd5b6133b0888261403f565b877fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa876040516133e091906156d3565b60405180910390a260016001600160a01b038a165f9081526099602052604090206001015460ff16600281111561341957613419614cba565b1461352b57604080518082018252898152600160208083018281526001600160a01b038e165f908152609990925293902082518155925183820180549394939192909160ff19169083600281111561347357613473614cba565b021790555050604051639926ee7d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169150639926ee7d906134c8908c908990600401615c80565b5f604051808303815f87803b1580156134df575f5ffd5b505af11580156134f1573d5f5f3e3d5ffd5b50506040518a92506001600160a01b038c1691507fe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe905f90a35b604051631fd93ca960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633fb2795290613579908c908b90600401615b93565b5f604051808303815f87803b158015613590575f5ffd5b505af11580156135a2573d5f5f3e3d5ffd5b5050604051632550477760e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169250632550477791506135f6908c908c908c90600401615993565b5f604051808303815f875af1158015613611573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526136389190810190615d2e565b60408087019190915260208601919091525162bff04d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062bff04d90613693908b908b90600401615bb6565b5f604051808303815f875af11580156136ae573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526136d59190810190615d87565b845250505095945050505050565b6020808301516001600160a01b038082165f8181526099909452604090932054919290871603613726576040516356168b4160e11b815260040160405180910390fd5b8760ff16845f015160ff161461374f57604051638e5aeee760e01b815260040160405180910390fd5b604051635401ed2760e01b81526004810182905260ff891660248201525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635401ed2790604401602060405180830381865afa1580156137bd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137e19190615e16565b90506137ed8185614073565b6001600160601b0316866001600160601b03161161381e57604051634c44995d60e01b815260040160405180910390fd5b6138288885614096565b6001600160601b0316816001600160601b031610611eeb5760405163b187e86960e01b815260040160405180910390fd5b61387d60405180606001604052806060815260200160608152602001606081525090565b6096545f9061389090859060ff16612c1c565b90505f61389c866121fb565b90506001600160c01b0382166138c5576040516313ca465760e01b815260040160405180910390fd5b8082166001600160c01b0316156138ef57604051630c6816cd60e01b815260040160405180910390fd5b60a0546001600160a01b0388165f908152609f60205260409020546001600160c01b0383811690851617914291613926919061574e565b1061394457604051631968677d60e11b815260040160405180910390fd5b61394e878261403f565b867fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa8660405161397e91906156d3565b60405180910390a260016001600160a01b0389165f9081526099602052604090206001015460ff1660028111156139b7576139b7614cba565b14613a1957604080518082018252888152600160208083018281526001600160a01b038d165f908152609990925293902082518155925183820180549394939192909160ff191690836002811115613a1157613a11614cba565b021790555050505b604051631fd93ca960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633fb2795290613a67908b908a90600401615b93565b5f604051808303815f87803b158015613a7e575f5ffd5b505af1158015613a90573d5f5f3e3d5ffd5b5050604051632550477760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063255047779150613ae4908b908b908b90600401615993565b5f604051808303815f875af1158015613aff573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613b269190810190615d2e565b60408087019190915260208601919091525162bff04d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062bff04d90613b81908a908a90600401615bb6565b5f604051808303815f875af1158015613b9c573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613bc39190810190615d87565b8452505050949350505050565b60605f82516001600160401b03811115613bec57613bec614731565b604051908082528060200260200182016040528015613c15578160200160208202803683370190505b5090505f5b8351811015613c7657613c478686868481518110613c3a57613c3a615691565b60200260200101516140af565b828281518110613c5957613c59615691565b63ffffffff90921660209283029190910190910152600101613c1a565b50949350505050565b5f81815260208390526040812054808203613c9d575f915050610cd6565b5f838152602085905260409020613cb5600183615e31565b81548110613cc557613cc5615691565b5f91825260209091200154600160401b90046001600160c01b03169150610cd69050565b5092915050565b5f805b8215610cd657613d04600184615e31565b9092169180613d1281615e44565b915050613cf3565b5f306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015613d7257507f000000000000000000000000000000000000000000000000000000000000000046145b15613d9c57507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b5f80805f516020615ef15f395f51905f5260035f516020615ef15f395f51905f52865f516020615ef15f395f51905f52888909090890505f613eb0827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f525f516020615ef15f395f51905f526141c7565b91959194509092505050565b5f61010082511115613f445760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a40161144e565b81515f03613f5357505f919050565b5f5f835f81518110613f6757613f67615691565b0160200151600160f89190911c81901b92505b845181101561403657848181518110613f9557613f95615691565b0160200151600160f89190911c1b915082821161402a5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a40161144e565b91811791600101613f7a565b50909392505050565b610c4a6098838361426f565b614056838383614428565b610abe57604051638baa579f60e01b815260040160405180910390fd5b60208101515f906127109061408c9061ffff1685615e64565b610ad49190615e86565b60408101515f906127109061408c9061ffff1685615e64565b5f81815260208490526040812054815b818110156141325760016140d38284615e31565b6140dd9190615e31565b92508463ffffffff16865f8681526020019081526020015f208463ffffffff168154811061410d5761410d615691565b5f9182526020909120015463ffffffff161161412a575050610ad4565b6001016140bf565b5060405162461bcd60e51b815260206004820152605c60248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d6170496e6465784174426c6f636b4e756d6265723a206e6f206269746d617060648201527f2075706461746520666f756e6420666f72206f70657261746f72496400000000608482015260a40161144e565b5f5f6141d16146f5565b6141d9614713565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828061421657fe5b50826142645760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c757265000000000000604482015260640161144e565b505195945050505050565b5f8281526020849052604081205490819003614313575f83815260208581526040808320815160608101835263ffffffff43811682528185018681526001600160c01b03808a16958401958652845460018101865594885295909620915191909201805495519351909416600160401b026001600160401b03938316600160201b0267ffffffffffffffff1990961691909216179390931716919091179055610e41565b5f83815260208590526040812061432b600184615e31565b8154811061433b5761433b615691565b5f918252602090912001805490915063ffffffff43811691160361437c5780546001600160401b0316600160401b6001600160c01b038516021781556123b2565b805463ffffffff438116600160201b81810267ffffffff00000000199094169390931784555f8781526020898152604080832081516060810183529485528483018481526001600160c01b03808c1693870193845282546001810184559286529390942094519401805493519151909216600160401b026001600160401b0391861690960267ffffffffffffffff19909316939094169290921717919091169190911790555050505050565b5f5f5f614435858561456d565b90925090505f81600481111561444d5761444d614cba565b14801561446b5750856001600160a01b0316826001600160a01b0316145b1561447b57600192505050610ad4565b5f5f876001600160a01b0316631626ba7e60e01b88886040516024016144a2929190615bb6565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516144e09190615eb3565b5f60405180830381855afa9150503d805f8114614518576040519150601f19603f3d011682016040523d82523d5f602084013e61451d565b606091505b5091509150818015614530575080516020145b801561456157508051630b135d3f60e11b906145559083016020908101908401615ec9565b6001600160e01b031916145b98975050505050505050565b5f5f82516041036145a1576020830151604084015160608501515f1a614595878285856145d8565b945094505050506145d1565b82516040036145ca57602083015160408401516145bf8683836146bd565b9350935050506145d1565b505f905060025b9250929050565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561460d57505f905060036146b4565b8460ff16601b1415801561462557508460ff16601c14155b1561463557505f905060046146b4565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614686573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166146ae575f600192509250506146b4565b91505f90505b94509492505050565b5f806001600160ff1b038316816146d960ff86901c601b61574e565b90506146e7878288856145d8565b935093505050935093915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b038111828210171561476757614767614731565b60405290565b604080519081016001600160401b038111828210171561476757614767614731565b604051601f8201601f191681016001600160401b03811182821017156147b7576147b7614731565b604052919050565b5f6001600160401b038211156147d7576147d7614731565b5060051b60200190565b6001600160a01b0381168114610d76575f5ffd5b8035614800816147e1565b919050565b5f82601f830112614814575f5ffd5b8135614827614822826147bf565b61478f565b8082825260208201915060208360051b860101925085831115614848575f5ffd5b602085015b8381101561486e578035614860816147e1565b83526020928301920161484d565b5095945050505050565b5f60208284031215614888575f5ffd5b81356001600160401b0381111561489d575f5ffd5b6125a884828501614805565b5f602082840312156148b9575f5ffd5b5035919050565b63ffffffff81168114610d76575f5ffd5b5f5f5f606084860312156148e3575f5ffd5b8335925060208401356148f5816148c0565b929592945050506040919091013590565b5f6001600160401b0382111561491e5761491e614731565b50601f01601f191660200190565b5f82601f83011261493b575f5ffd5b8135602083015f61494e61482284614906565b9050828152858383011115614961575f5ffd5b828260208301375f92810160200192909252509392505050565b5f6020828403121561498b575f5ffd5b81356001600160401b038111156149a0575f5ffd5b6125a88482850161492c565b5f602082840312156149bc575f5ffd5b8135610ad4816147e1565b5f5f604083850312156149d8575f5ffd5b50508035926020909101359150565b803560ff81168114614800575f5ffd5b5f60208284031215614a07575f5ffd5b610ad4826149e7565b815181526020808301519082015260408101610cd6565b803561ffff81168114614800575f5ffd5b5f60608284031215614a48575f5ffd5b614a50614745565b90508135614a5d816148c0565b8152614a6b60208301614a27565b6020820152614a7c60408301614a27565b604082015292915050565b6001600160601b0381168114610d76575f5ffd5b5f82601f830112614aaa575f5ffd5b8135614ab8614822826147bf565b8082825260208201915060208360061b860101925085831115614ad9575f5ffd5b602085015b8381101561486e5760408188031215614af5575f5ffd5b614afd61476d565b8135614b08816147e1565b81526020820135614b1881614a87565b6020828101919091529084529290920191604001614ade565b5f5f5f5f60c08587031215614b44575f5ffd5b614b4e8686614a38565b93506060850135614b5e81614a87565b925060808501356001600160401b03811115614b78575f5ffd5b614b8487828801614a9b565b92505060a0850135614b95816148c0565b939692955090935050565b5f5f83601f840112614bb0575f5ffd5b5081356001600160401b03811115614bc6575f5ffd5b6020830191508360208285010111156145d1575f5ffd5b5f5f5f60408486031215614bef575f5ffd5b83356001600160401b03811115614c04575f5ffd5b8401601f81018613614c14575f5ffd5b8035614c22614822826147bf565b8082825260208201915060208360051b850101925088831115614c43575f5ffd5b602084015b83811015614c835780356001600160401b03811115614c65575f5ffd5b614c748b602083890101614805565b84525060209283019201614c48565b50955050505060208401356001600160401b03811115614ca1575f5ffd5b614cad86828701614ba0565b9497909650939450505050565b634e487b7160e01b5f52602160045260245ffd5b60038110614cea57634e487b7160e01b5f52602160045260245ffd5b9052565b815181526020808301516040830191613ce990840182614cce565b5f5f60808385031215614d1a575f5ffd5b614d23836149e7565b9150614d328460208501614a38565b90509250929050565b5f5f60408385031215614d4c575f5ffd5b8235614d57816147e1565b915060208301356001600160401b03811115614d71575f5ffd5b614d7d8582860161492c565b9150509250929050565b5f82601f830112614d96575f5ffd5b8135614da4614822826147bf565b80828252602082019150602060608402860101925085831115614dc5575f5ffd5b602085015b8381101561486e57614ddc8782614a38565b8352602090920191606001614dca565b5f82601f830112614dfb575f5ffd5b8135614e09614822826147bf565b8082825260208201915060208360051b860101925085831115614e2a575f5ffd5b602085015b8381101561486e578035614e4281614a87565b835260209283019201614e2f565b5f82601f830112614e5f575f5ffd5b8135614e6d614822826147bf565b8082825260208201915060208360051b860101925085831115614e8e575f5ffd5b602085015b8381101561486e5780356001600160401b03811115614eb0575f5ffd5b614ebf886020838a0101614a9b565b84525060209283019201614e93565b5f82601f830112614edd575f5ffd5b8135614eeb614822826147bf565b8082825260208201915060208360051b860101925085831115614f0c575f5ffd5b602085015b8381101561486e57803560028110614f27575f5ffd5b835260209283019201614f11565b5f82601f830112614f44575f5ffd5b8135614f52614822826147bf565b8082825260208201915060208360051b860101925085831115614f73575f5ffd5b602085015b8381101561486e578035614f8b816148c0565b835260209283019201614f78565b5f5f5f5f5f5f5f5f5f6101208a8c031215614fb2575f5ffd5b614fbb8a6147f5565b9850614fc960208b016147f5565b9750614fd760408b016147f5565b965060608a0135955060808a01356001600160401b03811115614ff8575f5ffd5b6150048c828d01614d87565b95505060a08a01356001600160401b0381111561501f575f5ffd5b61502b8c828d01614dec565b94505060c08a01356001600160401b03811115615046575f5ffd5b6150528c828d01614e50565b93505060e08a01356001600160401b0381111561506d575f5ffd5b6150798c828d01614ece565b9250506101008a01356001600160401b03811115615095575f5ffd5b6150a18c828d01614f35565b9150509295985092959850929598565b5f5f5f60a084860312156150c3575f5ffd5b6150cd8585614a38565b925060608401356150dd81614a87565b915060808401356001600160401b038111156150f7575f5ffd5b61510386828701614a9b565b9150509250925092565b5f82601f83011261511c575f5ffd5b813561512a614822826147bf565b8082825260208201915060208360061b86010192508583111561514b575f5ffd5b602085015b8381101561486e5760408188031215615167575f5ffd5b61516f61476d565b615178826149e7565b81526020820135615188816147e1565b6020828101919091529084529290920191604001615150565b5f5f5f5f5f60a086880312156151b5575f5ffd5b85356151c0816147e1565b94506020860135935060408601356001600160401b038111156151e1575f5ffd5b6151ed8882890161510d565b9598949750949560608101359550608001359392505050565b5f60408284031215615216575f5ffd5b61521e61476d565b823581526020928301359281019290925250919050565b5f82601f830112615244575f5ffd5b61524e604061478f565b80604084018581111561525f575f5ffd5b845b81811015615279578035845260209384019301615261565b509095945050505050565b5f818303610100811215615296575f5ffd5b61529e614745565b91506152aa8484615206565b82526152b98460408501615206565b60208301526080607f19820112156152cf575f5ffd5b506152d861476d565b6152e58460808501615235565b81526152f48460c08501615235565b6020820152604082015292915050565b5f60608284031215615314575f5ffd5b61531c614745565b905081356001600160401b03811115615333575f5ffd5b61533f8482850161492c565b8252506020828101359082015260409182013591810191909152919050565b5f5f5f5f5f5f5f6101a0888a031215615375575f5ffd5b87356001600160401b0381111561538a575f5ffd5b6153968a828b01614ba0565b90985096505060208801356001600160401b038111156153b4575f5ffd5b6153c08a828b0161492c565b9550506153d08960408a01615284565b93506101408801356001600160401b038111156153eb575f5ffd5b6153f78a828b0161510d565b9350506101608801356001600160401b03811115615413575f5ffd5b61541f8a828b01615304565b9250506101808801356001600160401b0381111561543b575f5ffd5b6154478a828b01615304565b91505092959891949750929550565b5f5f60408385031215615467575f5ffd5b8235615472816147e1565b915060208301356001600160401b0381111561548c575f5ffd5b614d7d85828601614f35565b5f5f5f5f61016085870312156154ac575f5ffd5b84356001600160401b038111156154c1575f5ffd5b6154cd8782880161492c565b94505060208501356001600160401b038111156154e8575f5ffd5b6154f48782880161492c565b9350506155048660408701615284565b91506101408501356001600160401b0381111561551f575f5ffd5b61552b87828801615304565b91505092959194509250565b5f5f5f60608486031215615549575f5ffd5b8335615554816147e1565b925060208401356001600160401b0381111561556e575f5ffd5b61557a86828701614f35565b92505060408401356001600160401b03811115615595575f5ffd5b6151038682870161492c565b5f5f604083850312156155b2575f5ffd5b82356155bd816148c0565b915060208301356001600160401b038111156155d7575f5ffd5b8301601f810185136155e7575f5ffd5b80356155f5614822826147bf565b8082825260208201915060208360051b850101925087831115615616575f5ffd5b6020840193505b8284101561563857833582526020938401939091019061561d565b809450505050509250929050565b602080825282518282018190525f918401906040840190835b8181101561527957835163ffffffff1683526020938401939092019160010161565f565b60208101610cd68284614cce565b634e487b7160e01b5f52603260045260245ffd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610ad460208301846156a5565b5f602082840312156156f5575f5ffd5b81518015158114610ad4575f5ffd5b5f60208284031215615714575f5ffd5b8151610ad4816147e1565b5f6020828403121561572f575f5ffd5b8151610ad4816148c0565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610cd657610cd661573a565b5f5f8585111561576f575f5ffd5b8386111561577b575f5ffd5b5050820193919092039150565b5f60c0820188835260018060a01b038816602084015286604084015260c0606084015280865180835260e0850191506020880192505f5b818110156157f5578351805160ff1684526020908101516001600160a01b031681850152909301926040909201916001016157bf565b50506080840195909552505060a00152949350505050565b5f6040828403121561581d575f5ffd5b61582561476d565b825181526020928301519281019290925250919050565b5f82601f83011261584b575f5ffd5b615855604061478f565b806040840185811115615866575f5ffd5b845b81811015615279578051845260209384019301615868565b5f5f828403610120811215615893575f5ffd5b83516001600160401b038111156158a8575f5ffd5b8401601f810186136158b8575f5ffd5b80516158c661482282614906565b8181528760208385010111156158da575f5ffd5b8160208401602083015e5f91810160200191909152935050610100601f1982011215615904575f5ffd5b61590c614745565b615919866020870161580d565b8152615928866060870161580d565b60208201526080609f198301121561593e575f5ffd5b61594661476d565b91506159558660a0870161583c565b82526159648660e0870161583c565b602083015281604082015280925050509250929050565b5f6001820161598c5761598c61573a565b5060010190565b60018060a01b0384168152826020820152606060408201525f6159b960608301846156a5565b95945050505050565b5f602082840312156159d2575f5ffd5b81516001600160c01b0381168114610ad4575f5ffd5b634e487b7160e01b5f52601260045260245ffd5b5f82615a0a57615a0a6159e8565b500690565b60ff8181168382160190811115610cd657610cd661573a565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015615acf57868503603f190184528151805163ffffffff168652602090810151604082880181905281519088018190529101905f9060608801905b80831015615ab75783516001600160a01b031682526020938401936001939093019290910190615a8c565b50965050506020938401939190910190600101615a4e565b50929695505050505050565b5f8151808452602084019350602083015f5b82811015615b2957815180516001600160a01b031687526020908101516001600160601b03168188015260409096019590910190600101615aed565b5093949350505050565b60ff841681526001600160601b0383166020820152606060408201525f6159b96060830184615adb565b60ff851681526001600160601b038416602082015263ffffffff83166040820152608060608201525f6117056080830184615adb565b6001600160a01b03831681526040602082018190525f90610ad1908301846156a5565b828152604060208201525f610ad160408301846156a5565b5f60208284031215615bde575f5ffd5b5051919050565b805f5b6002811015610e41578151845260209384019390910190600101615be8565b6001600160a01b03841681528251805160208084019190915201516040820152610160810160208481015180516060850152908101516080840152506040840151615c5660a084018251615be5565b60200151615c6760e0840182615be5565b50825161012083015260208301516101408301526125a8565b60018060a01b0383168152604060208201525f825160606040840152615ca960a08401826156a5565b90506020840151606084015260408401516080840152809150509392505050565b5f82601f830112615cd9575f5ffd5b8151615ce7614822826147bf565b8082825260208201915060208360051b860101925085831115615d08575f5ffd5b602085015b8381101561486e578051615d2081614a87565b835260209283019201615d0d565b5f5f60408385031215615d3f575f5ffd5b82516001600160401b03811115615d54575f5ffd5b615d6085828601615cca565b92505060208301516001600160401b03811115615d7b575f5ffd5b614d7d85828601615cca565b5f60208284031215615d97575f5ffd5b81516001600160401b03811115615dac575f5ffd5b8201601f81018413615dbc575f5ffd5b8051615dca614822826147bf565b8082825260208201915060208360051b850101925086831115615deb575f5ffd5b6020840193505b82841015611705578351615e05816148c0565b825260209384019390910190615df2565b5f60208284031215615e26575f5ffd5b8151610ad481614a87565b81810381811115610cd657610cd661573a565b5f61ffff821661ffff8103615e5b57615e5b61573a565b60010192915050565b6001600160601b038181168382160290811690818114613ce957613ce961573a565b5f6001600160601b03831680615e9e57615e9e6159e8565b806001600160601b0384160491505092915050565b5f82518060208501845e5f920191825250919050565b5f60208284031215615ed9575f5ffd5b81516001600160e01b031981168114610ad4575f5ffdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220cd57101616c87a68734963c038945c4cd8640aeefc8e21e3cb7441307f42124f64736f6c634300081b0033",
}

// ContractRegistryCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractRegistryCoordinatorMetaData.ABI instead.
var ContractRegistryCoordinatorABI = ContractRegistryCoordinatorMetaData.ABI

// ContractRegistryCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractRegistryCoordinatorMetaData.Bin instead.
var ContractRegistryCoordinatorBin = ContractRegistryCoordinatorMetaData.Bin

// DeployContractRegistryCoordinator deploys a new Ethereum contract, binding an instance of ContractRegistryCoordinator to it.
func DeployContractRegistryCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _serviceManager common.Address, _stakeRegistry common.Address, _blsApkRegistry common.Address, _indexRegistry common.Address, _avsDirectory common.Address, _pauserRegistry common.Address) (common.Address, *types.Transaction, *ContractRegistryCoordinator, error) {
	parsed, err := ContractRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractRegistryCoordinatorBin), backend, _serviceManager, _stakeRegistry, _blsApkRegistry, _indexRegistry, _avsDirectory, _pauserRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractRegistryCoordinator{ContractRegistryCoordinatorCaller: ContractRegistryCoordinatorCaller{contract: contract}, ContractRegistryCoordinatorTransactor: ContractRegistryCoordinatorTransactor{contract: contract}, ContractRegistryCoordinatorFilterer: ContractRegistryCoordinatorFilterer{contract: contract}}, nil
}

// ContractRegistryCoordinatorMethods is an auto generated interface around an Ethereum contract.
type ContractRegistryCoordinatorMethods interface {
	ContractRegistryCoordinatorCalls
	ContractRegistryCoordinatorTransacts
	ContractRegistryCoordinatorFilters
}

// ContractRegistryCoordinatorCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractRegistryCoordinatorCalls interface {
	OPERATORCHURNAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	AvsDirectory(opts *bind.CallOpts) (common.Address, error)

	BlsApkRegistry(opts *bind.CallOpts) (common.Address, error)

	CalculateOperatorChurnApprovalDigestHash(opts *bind.CallOpts, registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error)

	ChurnApprover(opts *bind.CallOpts) (common.Address, error)

	EjectionCooldown(opts *bind.CallOpts) (*big.Int, error)

	Ejector(opts *bind.CallOpts) (common.Address, error)

	GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error)

	GetOperator(opts *bind.CallOpts, operator common.Address) (IRegistryCoordinatorOperatorInfo, error)

	GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error)

	GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error)

	GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error)

	GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error)

	GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error)

	GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error)

	GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error)

	IndexRegistry(opts *bind.CallOpts) (common.Address, error)

	IsChurnApproverSaltUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error)

	IsM2Quorum(opts *bind.CallOpts, arg0 uint8) (bool, error)

	IsOperatorSetAVS(opts *bind.CallOpts) (bool, error)

	IsUsingOperatorSets(opts *bind.CallOpts) (bool, error)

	LastEjectionTimestamp(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	NumRegistries(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error)

	QuorumCount(opts *bind.CallOpts) (uint8, error)

	QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error)

	Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error)

	ServiceManager(opts *bind.CallOpts) (common.Address, error)

	StakeRegistry(opts *bind.CallOpts) (common.Address, error)
}

// ContractRegistryCoordinatorTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractRegistryCoordinatorTransacts interface {
	CreateSlashableStakeQuorum(opts *bind.TransactOpts, operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error)

	CreateTotalDelegatedStakeQuorum(opts *bind.TransactOpts, operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error)

	DeregisterOperator(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error)

	DeregisterOperator0(opts *bind.TransactOpts, quorumNumbers []byte) (*types.Transaction, error)

	EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)

	EnableOperatorSets(opts *bind.TransactOpts) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IRegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IStakeRegistryStrategyParams, _stakeTypes []uint8, _lookAheadPeriods []uint32) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RegisterOperator(opts *bind.TransactOpts, quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RegisterOperator0(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error)

	RegisterOperatorWithChurn(opts *bind.TransactOpts, quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetChurnApprover(opts *bind.TransactOpts, _churnApprover common.Address) (*types.Transaction, error)

	SetEjectionCooldown(opts *bind.TransactOpts, _ejectionCooldown *big.Int) (*types.Transaction, error)

	SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error)

	SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error)

	UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error)

	UpdateSocket(opts *bind.TransactOpts, socket string) (*types.Transaction, error)
}

// ContractRegistryCoordinatorFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractRegistryCoordinatorFilters interface {
	FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractRegistryCoordinatorChurnApproverUpdatedIterator, error)
	WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorChurnApproverUpdated) (event.Subscription, error)
	ParseChurnApproverUpdated(log types.Log) (*ContractRegistryCoordinatorChurnApproverUpdated, error)

	FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractRegistryCoordinatorEjectorUpdatedIterator, error)
	WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorEjectorUpdated) (event.Subscription, error)
	ParseEjectorUpdated(log types.Log) (*ContractRegistryCoordinatorEjectorUpdated, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractRegistryCoordinatorInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractRegistryCoordinatorInitialized, error)

	FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractRegistryCoordinatorOperatorDeregisteredIterator, error)
	WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorDeregistered(log types.Log) (*ContractRegistryCoordinatorOperatorDeregistered, error)

	FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractRegistryCoordinatorOperatorRegisteredIterator, error)
	WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorRegistered(log types.Log) (*ContractRegistryCoordinatorOperatorRegistered, error)

	FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator, error)
	WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error)
	ParseOperatorSetParamsUpdated(log types.Log) (*ContractRegistryCoordinatorOperatorSetParamsUpdated, error)

	FilterOperatorSocketUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractRegistryCoordinatorOperatorSocketUpdateIterator, error)
	WatchOperatorSocketUpdate(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorSocketUpdate, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorSocketUpdate(log types.Log) (*ContractRegistryCoordinatorOperatorSocketUpdate, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractRegistryCoordinatorOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractRegistryCoordinatorOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractRegistryCoordinatorPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractRegistryCoordinatorPaused, error)

	FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator, error)
	WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error)
	ParseQuorumBlockNumberUpdated(log types.Log) (*ContractRegistryCoordinatorQuorumBlockNumberUpdated, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractRegistryCoordinatorUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractRegistryCoordinatorUnpaused, error)
}

// ContractRegistryCoordinator is an auto generated Go binding around an Ethereum contract.
type ContractRegistryCoordinator struct {
	ContractRegistryCoordinatorCaller     // Read-only binding to the contract
	ContractRegistryCoordinatorTransactor // Write-only binding to the contract
	ContractRegistryCoordinatorFilterer   // Log filterer for contract events
}

// ContractRegistryCoordinator implements the ContractRegistryCoordinatorMethods interface.
var _ ContractRegistryCoordinatorMethods = (*ContractRegistryCoordinator)(nil)

// ContractRegistryCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractRegistryCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryCoordinatorCaller implements the ContractRegistryCoordinatorCalls interface.
var _ ContractRegistryCoordinatorCalls = (*ContractRegistryCoordinatorCaller)(nil)

// ContractRegistryCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractRegistryCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryCoordinatorTransactor implements the ContractRegistryCoordinatorTransacts interface.
var _ ContractRegistryCoordinatorTransacts = (*ContractRegistryCoordinatorTransactor)(nil)

// ContractRegistryCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractRegistryCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryCoordinatorFilterer implements the ContractRegistryCoordinatorFilters interface.
var _ ContractRegistryCoordinatorFilters = (*ContractRegistryCoordinatorFilterer)(nil)

// ContractRegistryCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractRegistryCoordinatorSession struct {
	Contract     *ContractRegistryCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractRegistryCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractRegistryCoordinatorCallerSession struct {
	Contract *ContractRegistryCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ContractRegistryCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractRegistryCoordinatorTransactorSession struct {
	Contract     *ContractRegistryCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractRegistryCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRegistryCoordinatorRaw struct {
	Contract *ContractRegistryCoordinator // Generic contract binding to access the raw methods on
}

// ContractRegistryCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractRegistryCoordinatorCallerRaw struct {
	Contract *ContractRegistryCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// ContractRegistryCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractRegistryCoordinatorTransactorRaw struct {
	Contract *ContractRegistryCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractRegistryCoordinator creates a new instance of ContractRegistryCoordinator, bound to a specific deployed contract.
func NewContractRegistryCoordinator(address common.Address, backend bind.ContractBackend) (*ContractRegistryCoordinator, error) {
	contract, err := bindContractRegistryCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinator{ContractRegistryCoordinatorCaller: ContractRegistryCoordinatorCaller{contract: contract}, ContractRegistryCoordinatorTransactor: ContractRegistryCoordinatorTransactor{contract: contract}, ContractRegistryCoordinatorFilterer: ContractRegistryCoordinatorFilterer{contract: contract}}, nil
}

// NewContractRegistryCoordinatorCaller creates a new read-only instance of ContractRegistryCoordinator, bound to a specific deployed contract.
func NewContractRegistryCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*ContractRegistryCoordinatorCaller, error) {
	contract, err := bindContractRegistryCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorCaller{contract: contract}, nil
}

// NewContractRegistryCoordinatorTransactor creates a new write-only instance of ContractRegistryCoordinator, bound to a specific deployed contract.
func NewContractRegistryCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractRegistryCoordinatorTransactor, error) {
	contract, err := bindContractRegistryCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorTransactor{contract: contract}, nil
}

// NewContractRegistryCoordinatorFilterer creates a new log filterer instance of ContractRegistryCoordinator, bound to a specific deployed contract.
func NewContractRegistryCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractRegistryCoordinatorFilterer, error) {
	contract, err := bindContractRegistryCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorFilterer{contract: contract}, nil
}

// bindContractRegistryCoordinator binds a generic wrapper to an already deployed contract.
func bindContractRegistryCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractRegistryCoordinator.Contract.ContractRegistryCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.ContractRegistryCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.ContractRegistryCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractRegistryCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.contract.Transact(opts, method, params...)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) OPERATORCHURNAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "OPERATOR_CHURN_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractRegistryCoordinator.CallOpts)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractRegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "PUBKEY_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractRegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractRegistryCoordinator.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) AvsDirectory() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.AvsDirectory(&_ContractRegistryCoordinator.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.AvsDirectory(&_ContractRegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) BlsApkRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.BlsApkRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.BlsApkRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) CalculateOperatorChurnApprovalDigestHash(opts *bind.CallOpts, registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "calculateOperatorChurnApprovalDigestHash", registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractRegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractRegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) ChurnApprover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "churnApprover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) ChurnApprover() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.ChurnApprover(&_ContractRegistryCoordinator.CallOpts)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) ChurnApprover() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.ChurnApprover(&_ContractRegistryCoordinator.CallOpts)
}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) EjectionCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "ejectionCooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) EjectionCooldown() (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.EjectionCooldown(&_ContractRegistryCoordinator.CallOpts)
}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) EjectionCooldown() (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.EjectionCooldown(&_ContractRegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) Ejector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "ejector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Ejector() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.Ejector(&_ContractRegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) Ejector() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.Ejector(&_ContractRegistryCoordinator.CallOpts)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getCurrentQuorumBitmap", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractRegistryCoordinator.CallOpts, operatorId)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractRegistryCoordinator.CallOpts, operatorId)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetOperator(opts *bind.CallOpts, operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getOperator", operator)

	if err != nil {
		return *new(IRegistryCoordinatorOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorOperatorInfo)).(*IRegistryCoordinatorOperatorInfo)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetOperator(operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	return _ContractRegistryCoordinator.Contract.GetOperator(&_ContractRegistryCoordinator.CallOpts, operator)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetOperator(operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	return _ContractRegistryCoordinator.Contract.GetOperator(&_ContractRegistryCoordinator.CallOpts, operator)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getOperatorFromId", operatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorFromId(&_ContractRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorFromId(&_ContractRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorId(&_ContractRegistryCoordinator.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorId(&_ContractRegistryCoordinator.CallOpts, operator)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getOperatorSetParams", quorumNumber)

	if err != nil {
		return *new(IRegistryCoordinatorOperatorSetParam), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorOperatorSetParam)).(*IRegistryCoordinatorOperatorSetParam)

	return out0, err

}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetOperatorSetParams(quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetOperatorSetParams(quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getOperatorStatus", operator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorStatus(&_ContractRegistryCoordinator.CallOpts, operator)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorStatus(&_ContractRegistryCoordinator.CallOpts, operator)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapAtBlockNumberByIndex", operatorId, blockNumber, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapHistoryLength", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapIndicesAtBlockNumber", blockNumber, operatorIds)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapUpdateByIndex", operatorId, index)

	if err != nil {
		return *new(IRegistryCoordinatorQuorumBitmapUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorQuorumBitmapUpdate)).(*IRegistryCoordinatorQuorumBitmapUpdate)

	return out0, err

}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractRegistryCoordinator.CallOpts, operatorId, index)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractRegistryCoordinator.CallOpts, operatorId, index)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) IndexRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "indexRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) IndexRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.IndexRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) IndexRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.IndexRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) IsChurnApproverSaltUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "isChurnApproverSaltUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// IsM2Quorum is a free data retrieval call binding the contract method 0xa4d7871f.
//
// Solidity: function isM2Quorum(uint8 ) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) IsM2Quorum(opts *bind.CallOpts, arg0 uint8) (bool, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "isM2Quorum", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsM2Quorum is a free data retrieval call binding the contract method 0xa4d7871f.
//
// Solidity: function isM2Quorum(uint8 ) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) IsM2Quorum(arg0 uint8) (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsM2Quorum(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// IsM2Quorum is a free data retrieval call binding the contract method 0xa4d7871f.
//
// Solidity: function isM2Quorum(uint8 ) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) IsM2Quorum(arg0 uint8) (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsM2Quorum(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0xcabbb17f.
//
// Solidity: function isOperatorSetAVS() view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) IsOperatorSetAVS(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "isOperatorSetAVS")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0xcabbb17f.
//
// Solidity: function isOperatorSetAVS() view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) IsOperatorSetAVS() (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsOperatorSetAVS(&_ContractRegistryCoordinator.CallOpts)
}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0xcabbb17f.
//
// Solidity: function isOperatorSetAVS() view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) IsOperatorSetAVS() (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsOperatorSetAVS(&_ContractRegistryCoordinator.CallOpts)
}

// IsUsingOperatorSets is a free data retrieval call binding the contract method 0xbd33ee24.
//
// Solidity: function isUsingOperatorSets() view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) IsUsingOperatorSets(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "isUsingOperatorSets")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUsingOperatorSets is a free data retrieval call binding the contract method 0xbd33ee24.
//
// Solidity: function isUsingOperatorSets() view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) IsUsingOperatorSets() (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsUsingOperatorSets(&_ContractRegistryCoordinator.CallOpts)
}

// IsUsingOperatorSets is a free data retrieval call binding the contract method 0xbd33ee24.
//
// Solidity: function isUsingOperatorSets() view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) IsUsingOperatorSets() (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsUsingOperatorSets(&_ContractRegistryCoordinator.CallOpts)
}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) LastEjectionTimestamp(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "lastEjectionTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) LastEjectionTimestamp(arg0 common.Address) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.LastEjectionTimestamp(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) LastEjectionTimestamp(arg0 common.Address) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.LastEjectionTimestamp(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) NumRegistries(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "numRegistries")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) NumRegistries() (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.NumRegistries(&_ContractRegistryCoordinator.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) NumRegistries() (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.NumRegistries(&_ContractRegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Owner() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.Owner(&_ContractRegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) Owner() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.Owner(&_ContractRegistryCoordinator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Paused(index uint8) (bool, error) {
	return _ContractRegistryCoordinator.Contract.Paused(&_ContractRegistryCoordinator.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) Paused(index uint8) (bool, error) {
	return _ContractRegistryCoordinator.Contract.Paused(&_ContractRegistryCoordinator.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Paused0() (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.Paused0(&_ContractRegistryCoordinator.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) Paused0() (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.Paused0(&_ContractRegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) PauserRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.PauserRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.PauserRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "pubkeyRegistrationMessageHash", operator)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractRegistryCoordinator.CallOpts, operator)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractRegistryCoordinator.CallOpts, operator)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) QuorumCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "quorumCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) QuorumCount() (uint8, error) {
	return _ContractRegistryCoordinator.Contract.QuorumCount(&_ContractRegistryCoordinator.CallOpts)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) QuorumCount() (uint8, error) {
	return _ContractRegistryCoordinator.Contract.QuorumCount(&_ContractRegistryCoordinator.CallOpts)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "quorumUpdateBlockNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "registries", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.Registries(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.Registries(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "serviceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) ServiceManager() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.ServiceManager(&_ContractRegistryCoordinator.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) ServiceManager() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.ServiceManager(&_ContractRegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) StakeRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.StakeRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.StakeRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) CreateSlashableStakeQuorum(opts *bind.TransactOpts, operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "createSlashableStakeQuorum", operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) CreateSlashableStakeQuorum(operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.CreateSlashableStakeQuorum(&_ContractRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) CreateSlashableStakeQuorum(operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.CreateSlashableStakeQuorum(&_ContractRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) CreateTotalDelegatedStakeQuorum(opts *bind.TransactOpts, operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "createTotalDelegatedStakeQuorum", operatorSetParams, minimumStake, strategyParams)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) CreateTotalDelegatedStakeQuorum(operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.CreateTotalDelegatedStakeQuorum(&_ContractRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) CreateTotalDelegatedStakeQuorum(operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.CreateTotalDelegatedStakeQuorum(&_ContractRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x9d8e0c23.
//
// Solidity: function deregisterOperator(address operator, uint32[] operatorSetIds) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "deregisterOperator", operator, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x9d8e0c23.
//
// Solidity: function deregisterOperator(address operator, uint32[] operatorSetIds) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) DeregisterOperator(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.DeregisterOperator(&_ContractRegistryCoordinator.TransactOpts, operator, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x9d8e0c23.
//
// Solidity: function deregisterOperator(address operator, uint32[] operatorSetIds) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) DeregisterOperator(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.DeregisterOperator(&_ContractRegistryCoordinator.TransactOpts, operator, operatorSetIds)
}

// DeregisterOperator0 is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) DeregisterOperator0(opts *bind.TransactOpts, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "deregisterOperator0", quorumNumbers)
}

// DeregisterOperator0 is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) DeregisterOperator0(quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.DeregisterOperator0(&_ContractRegistryCoordinator.TransactOpts, quorumNumbers)
}

// DeregisterOperator0 is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) DeregisterOperator0(quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.DeregisterOperator0(&_ContractRegistryCoordinator.TransactOpts, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "ejectOperator", operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.EjectOperator(&_ContractRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.EjectOperator(&_ContractRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// EnableOperatorSets is a paid mutator transaction binding the contract method 0xee318821.
//
// Solidity: function enableOperatorSets() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) EnableOperatorSets(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "enableOperatorSets")
}

// EnableOperatorSets is a paid mutator transaction binding the contract method 0xee318821.
//
// Solidity: function enableOperatorSets() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) EnableOperatorSets() (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.EnableOperatorSets(&_ContractRegistryCoordinator.TransactOpts)
}

// EnableOperatorSets is a paid mutator transaction binding the contract method 0xee318821.
//
// Solidity: function enableOperatorSets() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) EnableOperatorSets() (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.EnableOperatorSets(&_ContractRegistryCoordinator.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fc3f886.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, (uint32,uint16,uint16)[] _operatorSetParams, uint96[] _minimumStakes, (address,uint96)[][] _strategyParams, uint8[] _stakeTypes, uint32[] _lookAheadPeriods) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IRegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IStakeRegistryStrategyParams, _stakeTypes []uint8, _lookAheadPeriods []uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "initialize", _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams, _stakeTypes, _lookAheadPeriods)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fc3f886.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, (uint32,uint16,uint16)[] _operatorSetParams, uint96[] _minimumStakes, (address,uint96)[][] _strategyParams, uint8[] _stakeTypes, uint32[] _lookAheadPeriods) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IRegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IStakeRegistryStrategyParams, _stakeTypes []uint8, _lookAheadPeriods []uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.Initialize(&_ContractRegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams, _stakeTypes, _lookAheadPeriods)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fc3f886.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, (uint32,uint16,uint16)[] _operatorSetParams, uint96[] _minimumStakes, (address,uint96)[][] _strategyParams, uint8[] _stakeTypes, uint32[] _lookAheadPeriods) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IRegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IStakeRegistryStrategyParams, _stakeTypes []uint8, _lookAheadPeriods []uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.Initialize(&_ContractRegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams, _stakeTypes, _lookAheadPeriods)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.Pause(&_ContractRegistryCoordinator.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.Pause(&_ContractRegistryCoordinator.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) PauseAll() (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.PauseAll(&_ContractRegistryCoordinator.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.PauseAll(&_ContractRegistryCoordinator.TransactOpts)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xa50857bf.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) RegisterOperator(opts *bind.TransactOpts, quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "registerOperator", quorumNumbers, socket, params, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xa50857bf.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) RegisterOperator(quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RegisterOperator(&_ContractRegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xa50857bf.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) RegisterOperator(quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RegisterOperator(&_ContractRegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorSignature)
}

// RegisterOperator0 is a paid mutator transaction binding the contract method 0xadcf73f7.
//
// Solidity: function registerOperator(address operator, uint32[] operatorSetIds, bytes data) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) RegisterOperator0(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "registerOperator0", operator, operatorSetIds, data)
}

// RegisterOperator0 is a paid mutator transaction binding the contract method 0xadcf73f7.
//
// Solidity: function registerOperator(address operator, uint32[] operatorSetIds, bytes data) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) RegisterOperator0(operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RegisterOperator0(&_ContractRegistryCoordinator.TransactOpts, operator, operatorSetIds, data)
}

// RegisterOperator0 is a paid mutator transaction binding the contract method 0xadcf73f7.
//
// Solidity: function registerOperator(address operator, uint32[] operatorSetIds, bytes data) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) RegisterOperator0(operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RegisterOperator0(&_ContractRegistryCoordinator.TransactOpts, operator, operatorSetIds, data)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x9b5d177b.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) RegisterOperatorWithChurn(opts *bind.TransactOpts, quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "registerOperatorWithChurn", quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x9b5d177b.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) RegisterOperatorWithChurn(quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RegisterOperatorWithChurn(&_ContractRegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x9b5d177b.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) RegisterOperatorWithChurn(quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RegisterOperatorWithChurn(&_ContractRegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RenounceOwnership(&_ContractRegistryCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RenounceOwnership(&_ContractRegistryCoordinator.TransactOpts)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) SetChurnApprover(opts *bind.TransactOpts, _churnApprover common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "setChurnApprover", _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetChurnApprover(&_ContractRegistryCoordinator.TransactOpts, _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetChurnApprover(&_ContractRegistryCoordinator.TransactOpts, _churnApprover)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) SetEjectionCooldown(opts *bind.TransactOpts, _ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "setEjectionCooldown", _ejectionCooldown)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) SetEjectionCooldown(_ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetEjectionCooldown(&_ContractRegistryCoordinator.TransactOpts, _ejectionCooldown)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) SetEjectionCooldown(_ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetEjectionCooldown(&_ContractRegistryCoordinator.TransactOpts, _ejectionCooldown)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "setEjector", _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetEjector(&_ContractRegistryCoordinator.TransactOpts, _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetEjector(&_ContractRegistryCoordinator.TransactOpts, _ejector)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "setOperatorSetParams", quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetOperatorSetParams(&_ContractRegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetOperatorSetParams(&_ContractRegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.TransferOwnership(&_ContractRegistryCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.TransferOwnership(&_ContractRegistryCoordinator.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.Unpause(&_ContractRegistryCoordinator.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.Unpause(&_ContractRegistryCoordinator.TransactOpts, newPausedStatus)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) UpdateOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "updateOperators", operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.UpdateOperators(&_ContractRegistryCoordinator.TransactOpts, operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.UpdateOperators(&_ContractRegistryCoordinator.TransactOpts, operators)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "updateOperatorsForQuorum", operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractRegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractRegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) UpdateSocket(opts *bind.TransactOpts, socket string) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "updateSocket", socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.UpdateSocket(&_ContractRegistryCoordinator.TransactOpts, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.UpdateSocket(&_ContractRegistryCoordinator.TransactOpts, socket)
}

// ContractRegistryCoordinatorChurnApproverUpdatedIterator is returned from FilterChurnApproverUpdated and is used to iterate over the raw logs and unpacked data for ChurnApproverUpdated events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorChurnApproverUpdatedIterator struct {
	Event *ContractRegistryCoordinatorChurnApproverUpdated // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorChurnApproverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorChurnApproverUpdated)
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
		it.Event = new(ContractRegistryCoordinatorChurnApproverUpdated)
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
func (it *ContractRegistryCoordinatorChurnApproverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorChurnApproverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorChurnApproverUpdated represents a ChurnApproverUpdated event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorChurnApproverUpdated struct {
	PrevChurnApprover common.Address
	NewChurnApprover  common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChurnApproverUpdated is a free log retrieval operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractRegistryCoordinatorChurnApproverUpdatedIterator, error) {

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorChurnApproverUpdatedIterator{contract: _ContractRegistryCoordinator.contract, event: "ChurnApproverUpdated", logs: logs, sub: sub}, nil
}

// WatchChurnApproverUpdated is a free log subscription operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorChurnApproverUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorChurnApproverUpdated)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
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

// ParseChurnApproverUpdated is a log parse operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseChurnApproverUpdated(log types.Log) (*ContractRegistryCoordinatorChurnApproverUpdated, error) {
	event := new(ContractRegistryCoordinatorChurnApproverUpdated)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorEjectorUpdatedIterator is returned from FilterEjectorUpdated and is used to iterate over the raw logs and unpacked data for EjectorUpdated events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorEjectorUpdatedIterator struct {
	Event *ContractRegistryCoordinatorEjectorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorEjectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorEjectorUpdated)
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
		it.Event = new(ContractRegistryCoordinatorEjectorUpdated)
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
func (it *ContractRegistryCoordinatorEjectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorEjectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorEjectorUpdated represents a EjectorUpdated event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorEjectorUpdated struct {
	PrevEjector common.Address
	NewEjector  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEjectorUpdated is a free log retrieval operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractRegistryCoordinatorEjectorUpdatedIterator, error) {

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorEjectorUpdatedIterator{contract: _ContractRegistryCoordinator.contract, event: "EjectorUpdated", logs: logs, sub: sub}, nil
}

// WatchEjectorUpdated is a free log subscription operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorEjectorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorEjectorUpdated)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
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

// ParseEjectorUpdated is a log parse operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseEjectorUpdated(log types.Log) (*ContractRegistryCoordinatorEjectorUpdated, error) {
	event := new(ContractRegistryCoordinatorEjectorUpdated)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorInitializedIterator struct {
	Event *ContractRegistryCoordinatorInitialized // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorInitialized)
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
		it.Event = new(ContractRegistryCoordinatorInitialized)
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
func (it *ContractRegistryCoordinatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorInitialized represents a Initialized event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractRegistryCoordinatorInitializedIterator, error) {

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorInitializedIterator{contract: _ContractRegistryCoordinator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorInitialized)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseInitialized(log types.Log) (*ContractRegistryCoordinatorInitialized, error) {
	event := new(ContractRegistryCoordinatorInitialized)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorDeregisteredIterator struct {
	Event *ContractRegistryCoordinatorOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorOperatorDeregistered)
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
		it.Event = new(ContractRegistryCoordinatorOperatorDeregistered)
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
func (it *ContractRegistryCoordinatorOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorOperatorDeregistered represents a OperatorDeregistered event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorDeregistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractRegistryCoordinatorOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorOperatorDeregisteredIterator{contract: _ContractRegistryCoordinator.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorOperatorDeregistered)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseOperatorDeregistered(log types.Log) (*ContractRegistryCoordinatorOperatorDeregistered, error) {
	event := new(ContractRegistryCoordinatorOperatorDeregistered)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorRegisteredIterator struct {
	Event *ContractRegistryCoordinatorOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorOperatorRegistered)
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
		it.Event = new(ContractRegistryCoordinatorOperatorRegistered)
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
func (it *ContractRegistryCoordinatorOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorOperatorRegistered represents a OperatorRegistered event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorRegistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractRegistryCoordinatorOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorOperatorRegisteredIterator{contract: _ContractRegistryCoordinator.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorOperatorRegistered)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseOperatorRegistered(log types.Log) (*ContractRegistryCoordinatorOperatorRegistered, error) {
	event := new(ContractRegistryCoordinatorOperatorRegistered)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator is returned from FilterOperatorSetParamsUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetParamsUpdated events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator struct {
	Event *ContractRegistryCoordinatorOperatorSetParamsUpdated // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorOperatorSetParamsUpdated)
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
		it.Event = new(ContractRegistryCoordinatorOperatorSetParamsUpdated)
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
func (it *ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorOperatorSetParamsUpdated represents a OperatorSetParamsUpdated event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorSetParamsUpdated struct {
	QuorumNumber      uint8
	OperatorSetParams IRegistryCoordinatorOperatorSetParam
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetParamsUpdated is a free log retrieval operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator{contract: _ContractRegistryCoordinator.contract, event: "OperatorSetParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetParamsUpdated is a free log subscription operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorOperatorSetParamsUpdated)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
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

// ParseOperatorSetParamsUpdated is a log parse operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseOperatorSetParamsUpdated(log types.Log) (*ContractRegistryCoordinatorOperatorSetParamsUpdated, error) {
	event := new(ContractRegistryCoordinatorOperatorSetParamsUpdated)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorOperatorSocketUpdateIterator is returned from FilterOperatorSocketUpdate and is used to iterate over the raw logs and unpacked data for OperatorSocketUpdate events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorSocketUpdateIterator struct {
	Event *ContractRegistryCoordinatorOperatorSocketUpdate // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorOperatorSocketUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorOperatorSocketUpdate)
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
		it.Event = new(ContractRegistryCoordinatorOperatorSocketUpdate)
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
func (it *ContractRegistryCoordinatorOperatorSocketUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorOperatorSocketUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorOperatorSocketUpdate represents a OperatorSocketUpdate event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorSocketUpdate struct {
	OperatorId [32]byte
	Socket     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketUpdate is a free log retrieval operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterOperatorSocketUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractRegistryCoordinatorOperatorSocketUpdateIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorOperatorSocketUpdateIterator{contract: _ContractRegistryCoordinator.contract, event: "OperatorSocketUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketUpdate is a free log subscription operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchOperatorSocketUpdate(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorSocketUpdate, operatorId [][32]byte) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorOperatorSocketUpdate)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
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

// ParseOperatorSocketUpdate is a log parse operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseOperatorSocketUpdate(log types.Log) (*ContractRegistryCoordinatorOperatorSocketUpdate, error) {
	event := new(ContractRegistryCoordinatorOperatorSocketUpdate)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOwnershipTransferredIterator struct {
	Event *ContractRegistryCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorOwnershipTransferred)
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
		it.Event = new(ContractRegistryCoordinatorOwnershipTransferred)
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
func (it *ContractRegistryCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractRegistryCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorOwnershipTransferredIterator{contract: _ContractRegistryCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorOwnershipTransferred)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*ContractRegistryCoordinatorOwnershipTransferred, error) {
	event := new(ContractRegistryCoordinatorOwnershipTransferred)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorPausedIterator struct {
	Event *ContractRegistryCoordinatorPaused // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorPaused)
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
		it.Event = new(ContractRegistryCoordinatorPaused)
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
func (it *ContractRegistryCoordinatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorPaused represents a Paused event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractRegistryCoordinatorPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorPausedIterator{contract: _ContractRegistryCoordinator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorPaused)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParsePaused(log types.Log) (*ContractRegistryCoordinatorPaused, error) {
	event := new(ContractRegistryCoordinatorPaused)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator is returned from FilterQuorumBlockNumberUpdated and is used to iterate over the raw logs and unpacked data for QuorumBlockNumberUpdated events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator struct {
	Event *ContractRegistryCoordinatorQuorumBlockNumberUpdated // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorQuorumBlockNumberUpdated)
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
		it.Event = new(ContractRegistryCoordinatorQuorumBlockNumberUpdated)
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
func (it *ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorQuorumBlockNumberUpdated represents a QuorumBlockNumberUpdated event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorQuorumBlockNumberUpdated struct {
	QuorumNumber uint8
	Blocknumber  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumBlockNumberUpdated is a free log retrieval operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator{contract: _ContractRegistryCoordinator.contract, event: "QuorumBlockNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumBlockNumberUpdated is a free log subscription operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorQuorumBlockNumberUpdated)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
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

// ParseQuorumBlockNumberUpdated is a log parse operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseQuorumBlockNumberUpdated(log types.Log) (*ContractRegistryCoordinatorQuorumBlockNumberUpdated, error) {
	event := new(ContractRegistryCoordinatorQuorumBlockNumberUpdated)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorUnpausedIterator struct {
	Event *ContractRegistryCoordinatorUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorUnpaused)
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
		it.Event = new(ContractRegistryCoordinatorUnpaused)
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
func (it *ContractRegistryCoordinatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorUnpaused represents a Unpaused event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractRegistryCoordinatorUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorUnpausedIterator{contract: _ContractRegistryCoordinator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorUnpaused)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseUnpaused(log types.Log) (*ContractRegistryCoordinatorUnpaused, error) {
	event := new(ContractRegistryCoordinatorUnpaused)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
