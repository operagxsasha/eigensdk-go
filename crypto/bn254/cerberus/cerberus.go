package cerberus

type Signer struct {
}

func New(signerUrl, pubKeyHex, password string) *Signer {
	return &Signer{}
}

func (s *Signer) SignMessage(message [32]byte) ([]byte, error) {
	return nil, nil
}

func (s *Signer) GetOperatorId() (string, error) {
	return "", nil
}
