package keystore

import (
	"context"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
)

type Signer struct {
	keyPair *bls.KeyPair
}

func New(path, password string) (*Signer, error) {
	keyPair, err := bls.ReadPrivateKeyFromFile(path, password)
	if err != nil {
		return nil, err
	}
	return &Signer{
		keyPair: keyPair,
	}, nil
}

func (s *Signer) SignMessage(ctx context.Context, message [32]byte) ([]byte, error) {
	signatureBytes := s.keyPair.SignMessage(message).Bytes()
	var signature []byte
	copy(signature[:], signatureBytes[:])
	return signature, nil
}

func (s *Signer) GetOperatorId() (string, error) {
	return types.OperatorIdFromG1Pubkey(s.keyPair.PubKey).LogValue().String(), nil
}
