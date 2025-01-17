// Package elcontracts bindings.go contains functions that create contract bindings for the Eigenlayer Core contracts
// These functions are meant to be used by constructors of the chainio package.
package elcontracts

import (
	permissioncontroller "github.com/Layr-Labs/eigensdk-go/contracts/bindings/PermissionController"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AVSDirectory"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

// ContractBindings contains the contract bindings for the EigenLayer Core contracts
//
// Unclear why geth bindings don't store and expose the contract address,
// so we also store them here in case the different constructors that use this struct need them
type ContractBindings struct {
	StrategyManagerAddr       gethcommon.Address
	DelegationManagerAddr     gethcommon.Address
	AvsDirectoryAddr          gethcommon.Address
	RewardsCoordinatorAddress gethcommon.Address
	AllocationManagerAddr     gethcommon.Address
	DelegationManager         *delegationmanager.ContractDelegationManager
	StrategyManager           *strategymanager.ContractStrategyManager
	AvsDirectory              *avsdirectory.ContractAVSDirectory
	RewardsCoordinator        *rewardscoordinator.ContractIRewardsCoordinator
	AllocationManager         *allocationmanager.ContractAllocationManager
	PermissionController      *permissioncontroller.ContractPermissionController
}

func NewBindingsFromConfig(
	cfg Config,
	client eth.HttpBackend,
	logger logging.Logger,
) (*ContractBindings, error) {
	var (
		err error

		contractDelegationManager *delegationmanager.ContractDelegationManager
		contractStrategyManager   *strategymanager.ContractStrategyManager
		contractAllocationManager *allocationmanager.ContractAllocationManager
		strategyManagerAddr       gethcommon.Address
		allocationManagerAddr     gethcommon.Address
		avsDirectory              *avsdirectory.ContractAVSDirectory
		rewardsCoordinator        *rewardscoordinator.ContractIRewardsCoordinator
		permissionController      *permissioncontroller.ContractPermissionController
	)

	if isZeroAddress(cfg.DelegationManagerAddress) {
		logger.Debug("DelegationManager address not provided, the calls to the contract will not work")
	} else {
		contractDelegationManager, err = delegationmanager.NewContractDelegationManager(cfg.DelegationManagerAddress, client)
		if err != nil {
			return nil, utils.WrapError("Failed to create DelegationManager contract", err)
		}

		strategyManagerAddr, err = contractDelegationManager.StrategyManager(&bind.CallOpts{})
		if err != nil {
			return nil, utils.WrapError("Failed to fetch StrategyManager address", err)
		}
		contractStrategyManager, err = strategymanager.NewContractStrategyManager(strategyManagerAddr, client)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch StrategyManager contract", err)
		}

		allocationManagerAddr, err = contractDelegationManager.AllocationManager(&bind.CallOpts{})
		if err != nil {
			return nil, utils.WrapError("Failed to fetch AllocationManager address", err)
		}
		contractAllocationManager, err = allocationmanager.NewContractAllocationManager(allocationManagerAddr, client)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch AllocationManager contract", err)
		}
	}

	if isZeroAddress(cfg.PermissionsControllerAddress) {
		logger.Debug("PermissionController address not provided, the calls to the contract will not work")
	} else {
		permissionController, err = permissioncontroller.NewContractPermissionController(
			cfg.PermissionsControllerAddress,
			client,
		)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch PermissionController contract", err)
		}
	}

	if isZeroAddress(cfg.AvsDirectoryAddress) {
		logger.Debug("AVSDirectory address not provided, the calls to the contract will not work")
	} else {
		avsDirectory, err = avsdirectory.NewContractAVSDirectory(cfg.AvsDirectoryAddress, client)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch AVSDirectory contract", err)
		}
	}

	if isZeroAddress(cfg.RewardsCoordinatorAddress) {
		logger.Debug("RewardsCoordinator address not provided, the calls to the contract will not work")
	} else {
		rewardsCoordinator, err = rewardscoordinator.NewContractIRewardsCoordinator(cfg.RewardsCoordinatorAddress, client)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch RewardsCoordinator contract", err)
		}
	}

	return &ContractBindings{
		StrategyManagerAddr:       strategyManagerAddr,
		DelegationManagerAddr:     cfg.DelegationManagerAddress,
		AvsDirectoryAddr:          cfg.AvsDirectoryAddress,
		RewardsCoordinatorAddress: cfg.RewardsCoordinatorAddress,
		StrategyManager:           contractStrategyManager,
		DelegationManager:         contractDelegationManager,
		AvsDirectory:              avsDirectory,
		RewardsCoordinator:        rewardsCoordinator,
		AllocationManager:         contractAllocationManager,
		AllocationManagerAddr:     allocationManagerAddr,
		PermissionController:      permissionController,
	}, nil
}
func isZeroAddress(address gethcommon.Address) bool {
	return address == gethcommon.Address{}
}
