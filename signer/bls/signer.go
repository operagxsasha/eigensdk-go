package bls

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/signer/bls/cerberus"
	"github.com/Layr-Labs/eigensdk-go/signer/bls/local"
	"github.com/Layr-Labs/eigensdk-go/signer/bls/types"
)

type Signer interface {
	// Sign signs the message using the BLS signature scheme
	Sign(ctx context.Context, msg []byte) ([]byte, error)

	// GetOperatorId returns the operator ID of the signer.
	// This is hash of the G1 public key of the signer
	GetOperatorId() (string, error)
}

// NewSigner creates a new signer based on the config
func NewSigner(cfg types.SignerConfig) (Signer, error) {
	switch cfg.SignerType {
	case types.Local:
		return local.New(local.Config{
			Path:     cfg.Path,
			Password: cfg.Password,
		})
	case types.Cerberus:
		return cerberus.New(cerberus.Config{
			URL:             cfg.CerberusUrl,
			PublicKeyHex:    cfg.PublicKeyHex,
			Password:        cfg.CerberusPassword,
			EnableTLS:       cfg.EnableTLS,
			TLSCertFilePath: cfg.TLSCertFilePath,
		})
	default:
		return nil, types.ErrInvalidSignerType
	}
}
