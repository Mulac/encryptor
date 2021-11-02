package encryptor

import (
	"log"
	"strings"
)

func caesarCipherEncryptor() (Encryptor, error) {
	return &caesarCipher{EncryptorType: Caesar}, nil
}

type caesarCipher struct {
	EncryptorType
	encryptorOptions
}

// Value reciever because we don't want our key to persist between calls
func (e caesarCipher) Encrypt(message string, opts ...EncryptorOption) (string, error) {
	e.apply(opts) // Grab our key
	if e.key == 0 {
		log.Printf("WARN|caesarEncryptor.Encrypt(%s, %d)|key is 0|no enciphering has taken place", message, e.key)
	}

	mapping := func(r rune) rune {
		return caesar(r, e.key)
	}

	// TODO(calum): is there a quicker way of doing this?  Use benchmark in _test.go
	return strings.Map(mapping, message), nil
}

// Value reciever because we don't want our key to persist between calls
func (e caesarCipher) Decrypt(message string, opts ...EncryptorOption) (string, error) {
	e.apply(opts) // Grab our key
	if e.key == 0 {
		log.Printf("WARN|caesarEncryptor.Encrypt(%s, %d)|key is 0|no enciphering has taken place", message, e.key)
	}

	mapping := func(r rune) rune {
		return caesar(r, -e.key)
	}

	// TODO(calum): is there a quicker way of doing this?  Use benchmark in _test.go
	return strings.Map(mapping, message), nil
}

// caesar cipher with letters a-z and A-Z
func caesar(r rune, key int) rune {
	if r >= 'a' && r <= 'z' {
		return shift(r, key, 'z', 'a')
	}
	if r >= 'A' && r <= 'Z' {
		return shift(r, key, 'Z', 'A')
	}
	return r
}

// shifts the rune by the key wrapping between boundries upper and lower
func shift(r rune, key, upper, lower int) rune {
	s := int(r) + key
	if s > upper {
		return rune(s - (upper - lower + 1))
	} else if s < lower {
		return rune(s + (upper - lower + 1))
	}
	return rune(s)
}
