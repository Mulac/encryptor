package encryptor

import (
	"fmt"
)

func unknownEncryptor(eType EncryptorType) (Encryptor, error) {
	return &unknown{eType}, fmt.Errorf("ERROR|unknownEncryptor()|encryptor type %s not implemented", eType)
}

type unknown struct {
	EncryptorType
}

func (e unknown) Encrypt(message string, opts ...EncryptorOption) (string, error) {
	return "", fmt.Errorf("ERROR|unknown.Encrypt(%s)|encryptor type %s not implemented", message, e.EncryptorType)
}

func (e unknown) Decrypt(message string, opts ...EncryptorOption) (string, error) {
	return "", fmt.Errorf("ERROR|unknown.Decrypt(%s)|encryptor type %s not implemented", message, e.EncryptorType)
}
