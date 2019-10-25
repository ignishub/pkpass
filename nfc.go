package pkpass

type NFCDict struct {
	Message             string `json:"message,omitempty"`
	EncryptionPublicKey string `json:"encryptionPublicKey,omitempty"`
}
