package crypto

import "crypto/cipher"

// Key represents necessary AES key information.
type Key struct {
	GCM       cipher.AEAD // The GCM of the cipher block
	Nonce     []byte      // The nonce
	NonceSize int         // The size of the nonce
}
