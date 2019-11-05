package common

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

var (
	// ProhibitedFiles represents the filenames
	// that will not be encrypted or decrypted.
	ProhibitedFiles = []string{
		"System Volume Information",
		".Trashes",
		".Spotlight-V100",
		".fseventsd",
	}

	// OSSlash represents the filepath slash delimiter for the current OS.
	OSSlash = "/"
)

const (
	// EncryptionDir the directory name that the standalone CLI will encrypt.
	EncryptionDir = "./secure"
)

// ByteSliceEqual returns true if two []byte a and b are equal,
// false if they are not.
func ByteSliceEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// CreateDirIfDoesNotExist creates a directory if it does not already exist.
func CreateDirIfDoesNotExist(dir string) error {
	dir = filepath.FromSlash(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetPassphrase will read a passphrase from stdin.
func GetPassphrase(confirm bool) (string, error) {
	// Ask for passphrase
	fmt.Print("Passphrase: ")
	passphraseBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	fmt.Println()

	// If the user wants password confirmation
	if confirm {
		// Ask again
		fmt.Print("Confirm passphrase: ")
		confirmationBytes, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return "", err
		}

		fmt.Println()

		// Compare passwords
		if !ByteSliceEqual(passphraseBytes, confirmationBytes) {
			return "", errors.New("passphrases do not match")
		}
	}

	passphrase := strings.TrimSpace(string(passphraseBytes))

	return passphrase, nil
}
