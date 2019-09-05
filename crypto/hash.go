package crypto

import (
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
		saltLength:  0,
		keyLength:   32,
	}

	// Return the hash
	hash = argon2.IDKey([]byte(input), []byte{}, p.iterations, p.memory, p.parallelism, p.keyLength)

	return hash, nil
}