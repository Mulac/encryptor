package encryptor

// Encryptor is anything that can encrypt and decrypt text
type Encryptor interface {
	Encrypt(string, ...EncryptorOption) (string, error)
	Decrypt(string, ...EncryptorOption) (string, error)
}

type EncryptorType string

const (
	Caesar EncryptorType = "caesar"
)

// NewEncryptor is a very simple factory for generating Encryptors of different types.
//
// It could be expanded to use the singleton pattern or encapsulated in
// it's own factory type for more complex encryptors that require more setup.
func NewEncryptor(eType EncryptorType) (Encryptor, error) {
	switch eType {
	case Caesar:
		return caesarCipherEncryptor()
	default:
		return unknownEncryptor(eType)
	}
}

// EncrptorOption allows us to add much more complex behaviour to our Encryptors in the future.
//
// Useful for more advanced encryption methods which need more keys or parameters than our simple
// caesar cipher.
type EncryptorOption struct {
	f func(*encryptorOptions)
}

type encryptorOptions struct {
	key int
}

func (eo *encryptorOptions) apply(opts []EncryptorOption) {
	for _, opt := range opts {
		opt.f(eo)
	}
}

// Key is our basic integer key used for encryptors like caesar cipher
func Key(key int) EncryptorOption {
	return EncryptorOption{func(eo *encryptorOptions) {
		eo.key = key
	}}
}
