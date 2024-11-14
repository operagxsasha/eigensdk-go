package bn254

import "context"

type Signer interface {
	SignMessage(ctx context.Context, message [32]byte) ([]byte, error)

	GetOperatorId() (string, error)
}
