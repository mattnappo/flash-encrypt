package crypto

import "crypto/cipher"

// Key represents necessary AES key information.
type Key struct {
	GCM cipher.AEAD
	Nonce []byte
}