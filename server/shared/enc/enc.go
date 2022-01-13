package enc

import "github.com/alexedwards/argon2id"

// Enc defines a encrypt struct.
type Enc struct {
}

// New creates a Enc object.
func New() *Enc {
	return &Enc{}
}

// Encrypt encrypts a password.
func (e Enc) Encrypt(password string) (string, error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

// Compare compares password and hash.
func (e Enc) Compare(password, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}
