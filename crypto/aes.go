package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/xoreo/flash-encrypt/fs"
	"io"
	"io/ioutil"
)

// EncryptFile encrypts a file using the AES encryption standard.
// passphrase is in plaintext.
func EncryptFile(path, passphrase string) error {
	// Read the file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Create the key
	key, err := Argon2String(passphrase)
	if err != nil {
		return err
	}

	// Create an AES block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Get the GCM of the block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Get the nonce size of the block
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return err
	}

	encryptedFile := gcm.Seal(nonce, nonce, data, nil) // Encrypt the data
	err = ioutil.WriteFile(path, encryptedFile, 0644)  // Write to file
	if err != nil {
		return err
	}

	return nil
}

// DecryptFile decrypts a file using the AES encryption standard.
// passphrase is in plaintext.
func DecryptFile(path, passphrase string) error {
	// Read the file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Create the key
	key, err := Argon2String(passphrase)
	if err != nil {
		return err
	}

	// Create a block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Get the block GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Calculate the nonce
	nonceSize := gcm.NonceSize()
	nonce, encryptedFile := data[:nonceSize], data[nonceSize:]

	// Decrypt the file data
	decryptedFile, err := gcm.Open(nil, nonce, encryptedFile, nil)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, decryptedFile, 0644) // Write to file
	if err != nil {
		return err
	}

	return nil
}

// EncryptDir encrypts an entire directory.
func EncryptDir(path, passphrase string) error {
	// Get file paths
	paths, err := fs.ListDir(path)
	if err != nil {
		return err
	}

	// Encrypt each file
	for i, path := range paths {
		err = EncryptFile(path, passphrase)
		if err != nil {
			return err
		}
		fmt.Printf("[%d] '%s' encrypted\n", i, path)
	}
	return nil
}

// DecryptDir decrypts an entire directory.
func DecryptDir(path, passphrase string) error {
	// Get file paths
	paths, err := fs.ListDir(path)
	if err != nil {
		return err
	}

	// Encrypt each file
	for i, path := range paths {
		err = DecryptFile(path, passphrase)
		if err != nil {
			return err
		}
		fmt.Printf("[%d] '%s' decrypted\n", i, path)
	}
	return nil
}