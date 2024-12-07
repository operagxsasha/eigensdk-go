package bls

type SignerConfig struct {
	// Local keystore params
	// Path to the key file
	Path string
	// Password to decrypt the key file
	Password string

	// Params for cerberus signer
	// CerberusUrl of the cerberus signer
	CerberusUrl string
	// PublicKeyHex is the public key of the cerberus signer
	PublicKeyHex string
	// CerberusPassword is the password to encrypt the key if cerberus using local keystore
	CerberusPassword string
}
