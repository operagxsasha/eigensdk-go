package cerberus

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	sdkBls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/signer/bls"

	v1 "github.com/Layr-Labs/cerberus-api/pkg/api/v1"
)

type Config struct {
	URL          string
	PublicKeyHex string

	// Optional: in case if your signer uses local keystore
	Password string

	EnableTLS       bool
	TLSCertFilePath string

	SigningTimeout time.Duration
}

type Signer struct {
	client    v1.SignerClient
	pubKeyHex string
	password  string
}

func New(cfg Config) (Signer, error) {
	opts := make([]grpc.DialOption, 0)
	if cfg.EnableTLS {
		creds, err := credentials.NewClientTLSFromFile(cfg.TLSCertFilePath, "")
		if err != nil {
			log.Fatalf("could not load tls cert: %s", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.NewClient(cfg.URL, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := v1.NewSignerClient(conn)
	return Signer{
		client:    client,
		pubKeyHex: cfg.PublicKeyHex,
		password:  cfg.Password,
	}, nil
}

func (s Signer) Sign(ctx context.Context, msg []byte) ([]byte, error) {
	if len(msg) != 32 {
		return nil, bls.ErrInvalidMessageLength
	}

	var data [32]byte
	copy(data[:], msg)

	resp, err := s.client.SignGeneric(ctx, &v1.SignGenericRequest{
		Data:      data[:],
		PublicKey: s.pubKeyHex,
		Password:  s.password,
	})
	if err != nil {
		return nil, err
	}

	return resp.Signature, nil
}

func (s Signer) GetOperatorId() (string, error) {
	pkBytes, err := hex.DecodeString(s.pubKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode BLS public key: %w", err)
	}
	pubkey := new(sdkBls.G1Point)
	publicKey := pubkey.Deserialize(pkBytes)
	return publicKey.GetOperatorID(), nil
}
