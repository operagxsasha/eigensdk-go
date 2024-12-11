package eigenpod

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/ethereum/go-ethereum/common"
)

func BuildEigenPodClients(
	address common.Address,
	client eth.HttpBackend,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ChainReader, *ChainWriter, *ContractBindings, error) {
	eigenPodBindings, err := NewContractBindings(address, client)
	if err != nil {
		return nil, nil, nil, utils.WrapError("failed to create EigenPod contract bindings", err)
	}

	eigenPodChainReader := newChainReader(
		&eigenPodBindings.ContractIEigenPodCaller,
		client,
		logger,
	)

	eigenPodChainWriter := newChainWriter(
		eigenPodBindings.ContractIEigenPod,
		client,
		logger,
		txMgr,
	)

	return eigenPodChainReader, eigenPodChainWriter, eigenPodBindings, nil
}

func BuildEigenPodManagerClients(
	address common.Address,
	client eth.HttpBackend,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ManagerChainReader, *ManagerChainWriter, *ManagerContractBindings, error) {
	eigenPodManagerBindings, err := NewManagerContractBindings(address, client)
	if err != nil {
		return nil, nil, nil, utils.WrapError("failed to create EigenPod manager contract bindings", err)
	}

	eigenPodManagerChainReader := newManagerChainReader(
		&eigenPodManagerBindings.ContractIEigenPodManagerCaller,
		client,
		logger,
	)

	eigenPodManagerChainWriter := newManagerChainWriter(
		eigenPodManagerBindings.ContractIEigenPodManager,
		client,
		logger,
		txMgr,
	)

	return eigenPodManagerChainReader, eigenPodManagerChainWriter, eigenPodManagerBindings, nil
}
