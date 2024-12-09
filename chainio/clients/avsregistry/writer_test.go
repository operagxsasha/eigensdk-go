package avsregistry_test

import (
	"context"
	"testing"

	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriterMethods(t *testing.T) {
	clients, _ := testclients.BuildTestClients(t)
	chainWriter := clients.AvsRegistryChainWriter

	keypair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)

	addr := gethcommon.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	ecdsaPrivKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ecdsaPrivateKey, err := crypto.HexToECDSA(ecdsaPrivKeyHex)
	require.NoError(t, err)

	quorumNumbers := types.QuorumNums{0}

	t.Run("register operator", func(t *testing.T) {
		receipt, err := chainWriter.RegisterOperator(
			context.Background(),
			ecdsaPrivateKey,
			keypair,
			quorumNumbers,
			"",
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("update stake of operator subset", func(t *testing.T) {
		receipt, err := chainWriter.UpdateStakesOfOperatorSubsetForAllQuorums(
			context.Background(),
			[]gethcommon.Address{addr},
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("update stake of entire operator set", func(t *testing.T) {
		receipt, err := chainWriter.UpdateStakesOfEntireOperatorSetForQuorums(
			context.Background(),
			[][]gethcommon.Address{{addr}},
			quorumNumbers,
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("deregister operator", func(t *testing.T) {
		receipt, err := chainWriter.DeregisterOperator(
			context.Background(),
			quorumNumbers,
			chainioutils.ConvertToBN254G1Point(keypair.PubKey),
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})
}

func TestBlsSignature(t *testing.T) {
	// read input from JSON if available, otherwise use default values
	var defaultInput = struct {
		Message    string `json:"message"`
		BlsPrivKey string `json:"bls_priv_key"`
	}{
		Message:    "Hello, world!Hello, world!123456",
		BlsPrivKey: "12248929636257230549931416853095037629726205319386239410403476017439825112537",
	}

	testData := testutils.NewTestData(defaultInput)
	// The message to sign
	messageArray := []byte(testData.Input.Message)

	var messageArray32 [32]byte
	copy(messageArray32[:], messageArray)

	// The private key as a string
	privKey, _ := bls.NewPrivateKey(testData.Input.BlsPrivKey)
	keyPair := bls.NewKeyPair(privKey)

	sig := keyPair.SignMessage(messageArray32)

	x := sig.G1Affine.X.String()
	y := sig.G1Affine.Y.String()

	assert.Equal(t, x, "15790168376429033610067099039091292283117017641532256477437243974517959682102")
	assert.Equal(t, y, "4960450323239587206117776989095741074887370703941588742100855592356200866613")
}
