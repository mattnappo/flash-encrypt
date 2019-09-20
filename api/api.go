package api

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/xoreo/flash-encrypt/crypto"
	"github.com/xoreo/flash-encrypt/fs"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
	"strings"
	"syscall"
)

// Encrypt is an abstracted method to encrypt a directory.
func Encrypt(targetDriveID string) error {
	reader := bufio.NewReader(os.Stdin)

	// Get the connected drives
	drives, err := fs.GetDrivesDarwin()
	if err != nil {
		return err
	}

	found := false
	// Find the drive name
	for i, drive := range drives {
		if strconv.Itoa(i) == targetDriveID {
			found = true
			// Confirm
			fmt.Print("Are you sure you want to encrypt " + drive + " (yes/no)? ")
			confirmation, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			confirmation = strings.TrimSuffix(confirmation, "\n")

			if confirmation == "yes" {
				// Ask for passphrase
				fmt.Print("Passphrase: ")
				passphraseBytes, err := terminal.ReadPassword(int(syscall.Stdin))
				if err != nil {
					return err
				}
				passphrase := strings.TrimSpace(string(passphraseBytes))

				// Encrypt the entire flash drive
				err = crypto.EncryptDir(fs.GetDrivePath(drive), passphrase)
				if err != nil {
					return err
				}

				fmt.Printf("Encrypted %s\n", drive)

			} else {
				return nil
			}

		}
	}

	// Validate input
	if found == false {
		return errors.New(fmt.Sprintf("drive with id '%s' could not be found", targetDriveID))
	}

	return nil
}

// Decrypt is an abstracted method to decrypt a directory.
func Decrypt(targetDriveID string) error {
	reader := bufio.NewReader(os.Stdin)

	// Get the connected drives
	drives, err := fs.GetDrivesDarwin()
	if err != nil {
		return err
	}

	found := false
	// Get the drive name
	for i, drive := range drives {
		if strconv.Itoa(i) == targetDriveID {
			found = true
			// Ask for confirmation
			fmt.Print("Are you sure you want to decrypt " + drive + " (yes/no)? ")
			confirmation, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			confirmation = strings.TrimSuffix(confirmation, "\n")

			if confirmation == "yes" {
				// Ask for passphrase
				fmt.Print("Passphrase: ")
				passphraseBytes, err := terminal.ReadPassword(int(syscall.Stdin))
				if err != nil {
					return err
				}
				passphrase := strings.TrimSpace(string(passphraseBytes))

				// Decrypt the entire directory
				err = crypto.DecryptDir(fs.GetDrivePath(drive), passphrase)
				if err != nil {
					return err
				}

				fmt.Printf("Decrypted %s\n", drive)

			} else {
				return nil
			}

		}
	}

	// Validate input
	if found == false {
		return errors.New(fmt.Sprintf("drive with id '%s' could not be found", targetDriveID))
	}

	return nil
}

func ListDrives() error {
	// Get connected drives
	drives, err := fs.GetDrivesDarwin()
	if err != nil {
		panic(err)
	}

	if len(drives) > 0 {
		// Print these drives
		fmt.Println("Connected drives:")
		for i, drive := range drives {
			fmt.Printf("[%d] %s\n", i, drive)
		}
	} else {
		fmt.Printf("There are no connected drives.")
	}


	return nil
}

func Status() error {
	return nil
}
