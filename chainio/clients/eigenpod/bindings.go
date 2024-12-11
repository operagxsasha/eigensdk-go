package eigenpod

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	ieigenpod "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IEigenPod"
	ieigenpodmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IEigenPodManager"
	"github.com/ethereum/go-ethereum/common"
)

type ContractBindings struct {
	Address common.Address
	*ieigenpod.ContractIEigenPod
}

type ContractCallerBindings struct {
	Address common.Address
	*ieigenpod.ContractIEigenPodCaller
}

type ManagerContractBindings struct {
	Address common.Address
	*ieigenpodmanager.ContractIEigenPodManager
}

type ManagerContractCallerBindings struct {
	Address common.Address
	*ieigenpodmanager.ContractIEigenPodManagerCaller
}

func NewContractBindings(
	address common.Address,
	ethClient eth.HttpBackend,
) (*ContractBindings, error) {
	pod, err := ieigenpod.NewContractIEigenPod(address, ethClient)
	if err != nil {
		return nil, err
	}
	return &ContractBindings{
		Address:           address,
		ContractIEigenPod: pod,
	}, nil
}

func NewContractCallerBindings(
	address common.Address,
	ethClient eth.HttpBackend,
) (*ContractCallerBindings, error) {
	pod, err := ieigenpod.NewContractIEigenPodCaller(address, ethClient)
	if err != nil {
		return nil, err
	}
	return &ContractCallerBindings{
		Address:                 address,
		ContractIEigenPodCaller: pod,
	}, nil
}

func NewManagerContractBindings(
	address common.Address,
	ethClient eth.HttpBackend,
) (*ManagerContractBindings, error) {
	manager, err := ieigenpodmanager.NewContractIEigenPodManager(address, ethClient)
	if err != nil {
		return nil, err
	}
	return &ManagerContractBindings{
		Address:                  address,
		ContractIEigenPodManager: manager,
	}, nil
}

func NewManagerContractCallerBindings(
	address common.Address,
	ethClient eth.HttpBackend,
) (*ManagerContractCallerBindings, error) {
	manager, err := ieigenpodmanager.NewContractIEigenPodManagerCaller(address, ethClient)
	if err != nil {
		return nil, err
	}
	return &ManagerContractCallerBindings{
		Address:                        address,
		ContractIEigenPodManagerCaller: manager,
	}, nil
}
