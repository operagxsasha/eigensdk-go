package txmgr_test

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestSendWithRetry(t *testing.T) {
	ecdsaPrivKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ecdsaPrivateKey, err := crypto.HexToECDSA(ecdsaPrivKeyHex)
	require.NoError(t, err)

	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer("")
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})

	ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
	require.NoError(t, err)

	rpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainid, err := ethHttpClient.ChainID(rpcCtx)
	require.NoError(t, err)

	signerV2, addr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainid)
	require.NoError(t, err)

	pkWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, addr, logger)
	require.NoError(t, err)

	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethHttpClient, logger, addr)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainid,
		Nonce:     0,
		GasTipCap: big.NewInt(1),
		GasFeeCap: big.NewInt(1_000_000_000),
		Gas:       21000,
		To:        testutils.ZeroAddress(),
		Value:     big.NewInt(1),
	})
	retryTimeout := 3 * time.Second // TODO: this is enough to fetch receipt on first attempt. If lower, fails with nonce too low. Fix it
	maxRetries := uint32(10)

	receipt, err := txMgr.SendWithRetry(context.Background(), tx, retryTimeout, maxRetries)
	require.NoError(t, err)

	t.Log(receipt)
}
