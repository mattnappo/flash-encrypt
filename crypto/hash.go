package crypto

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

// argon2params stores the parameters for generating an
// Argon2 hash.
type argon2params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// Argon2String hashes an input string with Argon2.
func Argon2String(input string) (hash []byte, err error) {
	// Set hashing parameters
	p := &argon2params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	// Generate a salt
	salt, err := generateSalt(p.saltLength)
	if err != nil {
		return nil, err
	}

	// Return the hash
	hash = argon2.IDKey([]byte(input), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return hash, nil
}

/* -- BEGIN PRIVATE METHODS -- */

func generateSalt(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

/* -- END PRIVATE METHODS -- */