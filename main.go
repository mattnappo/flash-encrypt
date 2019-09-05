package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/xoreo/flash-encrypt/crypto"
	"github.com/xoreo/flash-encrypt/fs"
	"os"
)

func getDrives() ([]string, error) {
	drives, err := fs.GetDrivesDarwin()
	if err != nil {
		panic(err)
	}

	for i, drive := range drives {
		fmt.Printf("[%d] %s\n", i, drive)
	}

	return drives, nil
}

func encrypt() error {
	drives, err := getDrives()
	if err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Which drive do you want to encrypt? ")
	targetDrive, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	for i, drive := range drives {
		if string(i) == targetDrive {
			fmt.Print("Are you sure you want to encrypt " + drive + " (yes/no)? ")
			confirmation, err := reader.ReadString('\n')
			if err != nil {
				return err
			}

			if confirmation == "yes" {
				fmt.Print("passphrase: ")
				passphrase, err := reader.ReadString('\n')
				if err != nil {
					return err
				}
				err = crypto.EncryptDir(fs.GetDrivePath(drive), passphrase)
				if err != nil {
					return err
				}
				fmt.Printf("Encrypted %s\n", drive)
			}

		}
	}

	fmt.Println(text)
}

func decrypt() {

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("[0] Encrypt\n[1] Decrypt")
	whereto, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	switch whereto {
	case "0":
		encrypt()
		break
	case "1":
		decrypt()
		break
	default:
		panic(errors.New(fmt.Sprintf("'%s' is not an option", whereto)))
	}

}
