package crypto

import "crypto/cipher"

// Key represents necessary AES key information.
type Key struct {
	GCM       cipher.AEAD `json:"gcm"`        // The GCM of the cipher block
	Nonce     []byte      `json:"nonce"`      // The nonce
	NonceSize int         `json:"nonce_size"` // The size of the nonce
}
