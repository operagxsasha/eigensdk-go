package local

import (
	"context"

	sdkBls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/signer/bls"
)

type Config struct {
	Path     string
	Password string
}

type Signer struct {
	key *sdkBls.KeyPair
}

func New(cfg Config) *Signer {
	keyPair, err := sdkBls.ReadPrivateKeyFromFile(cfg.Path, cfg.Password)
	if err != nil {
		return nil
	}
	return &Signer{
		key: keyPair,
	}
}

func (s *Signer) Sign(ctx context.Context, msg []byte) ([]byte, error) {
	if len(msg) != 32 {
		return nil, bls.ErrInvalidMessageLength
	}

	var data [32]byte
	copy(data[:], msg)

	return s.key.SignMessage(data).Serialize(), nil
}

func (s *Signer) GetOperatorId() (string, error) {
	return s.key.PubKey.GetOperatorID(), nil
}
