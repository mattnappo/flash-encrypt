package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/xoreo/flash-encrypt/fs"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

// GenerateKey generates a new AES key.
func GenerateKey(passphrase string) (*Key, error) {
	// Create the key
	key, err := Argon2String(passphrase)
	if err != nil {
		return nil, err
	}

	// Create an AES block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Get the GCM of the block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Get the nonce size of the block
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	// Return a new key
	return &Key{
		GCM:       gcm,
		Nonce:     nonce,
		NonceSize: gcm.NonceSize(),
	}, nil

}

// EncryptFile encrypts a file using the AES encryption standard.
// passphrase is in plaintext.
func EncryptFile(path string, key *Key) error {
	if strings.Contains(path, "flash-encrypt") {
		return nil
	}

	// Read the file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	encryptedFile := (*key).GCM.Seal((*key).Nonce, (*key).Nonce, data, nil) // Encrypt the data
	err = ioutil.WriteFile(path, encryptedFile, 0644)                       // Write to file
	if err != nil {
		return err
	}

	return nil
}

// DecryptFile decrypts a file using the AES encryption standard.
// passphrase is in plaintext.
func DecryptFile(path string, key *Key) error {
	if strings.Contains(path, "flash-encrypt") {
		return nil
	}

	// Read the file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Calculate the nonce
	nonceSize := (*key).NonceSize
	nonce, encryptedFile := data[:nonceSize], data[nonceSize:]

	// Decrypt the file data
	decryptedFile, err := (*key).GCM.Open(nil, nonce, encryptedFile, nil)
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
func EncryptDir(rootPath, passphrase string) error {
	// Get file paths
	paths, err := fs.ListDir(rootPath)
	if err != nil {
		return err
	}

	// Generate a new key
	key, err := GenerateKey(passphrase)
	if err != nil {
		return err
	}

	start := time.Now()

	// Encrypt all the files synchronously
	var wg sync.WaitGroup // Initialize the waitgroup
	for i, path := range paths {
		wg.Add(1)

		errChan := make(chan error)
		go func() {
			defer wg.Done()

			errChan <- EncryptFile(path, key) // Encrypt the file

			fmt.Printf("[%d] '%s' encrypted\n", i, path)
		}()

		if <-errChan != nil {
			return err
		}
	}

	wg.Wait() // Wait for the waitgroup to finish

	elapsed := time.Since(start)
	log.Printf("Encryption took %s", elapsed)

	return nil
}

// DecryptDir decrypts an entire directory.
func DecryptDir(path, passphrase string) error {
	// Get file paths
	paths, err := fs.ListDir(path)
	if err != nil {
		return err
	}

	// Generate a new key
	key, err := GenerateKey(passphrase)
	if err != nil {
		return err
	}

	start := time.Now()

	// Decrypt all the files synchronously
	var wg sync.WaitGroup // Initialize the waitgroup
	for i, path := range paths {
		wg.Add(1)

		errChan := make(chan error)
		go func() {
			defer wg.Done()

			errChan <- DecryptFile(path, key) // Encrypt the file

			fmt.Printf("[%d] '%s' decrypted\n", i, path)
		}()

		if <-errChan != nil {
			return err
		}
	}

	wg.Wait() // Wait for the waitgroup to finish

	elapsed := time.Since(start)
	log.Printf("Decryption took %s", elapsed)

	return nil
}
