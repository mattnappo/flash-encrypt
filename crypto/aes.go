package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)

// EncryptFile encrypts a file using the AES encryption standard.
// passphrase is in plaintext.
func EncryptFile(path string, passphrase string) error {
	// Read the file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	key := Argon2String(passphrase) // Create a new key

	// Create an AES block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize()) // Get the nonce size of the block
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return err
	}

	encryptedFile := gcm.Seal(nonce, nonce, data, nil) // Encrypt the data
	err = ioutil.WriteFile(path, encryptedFile, 0644) // Write to file
	if err != nil {
		return err
	}

	return nil
}
