package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

// EncryptFile encrypts a file using the AES encryption standard.
// passphrase is in plaintext.
func EncryptFile(path string, passphrase string) ([]byte, error) {
	// Read the file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := aes.NewCipher(Argon2String(passphrase))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}
