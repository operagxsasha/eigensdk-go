package bls

import "context"

type Signer interface {
	// Sign signs the message using the BLS signature scheme
	Sign(ctx context.Context, msg []byte) ([]byte, error)

	// GetOperatorId returns the operator ID of the signer.
	// This is hash of the G1 public key of the signer
	GetOperatorId() (string, error)
}
