package encryptor

import (
	"errors"
	"fmt"
)

// ErrEKey should be wrapped for invalid Key EncryptorOption
var ErrKey = errors.New("invalid encryptor key")

// unknownEncryptor is used when the Encryptor type given is not recognised
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
