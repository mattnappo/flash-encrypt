package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/xoreo/flash-encrypt/crypto"
	"github.com/xoreo/flash-encrypt/fs"
	"os"
	"strconv"
	"strings"
)

func printHeader() {
	fmt.Println("Welcome to")
	fmt.Println("   ______   ___   ______ __    _____  ____________  _____  ______")
	fmt.Println("  / __/ /  / _ | / __/ // /___/ __/ |/ / ___/ _ \\ \\/ / _ \\/_  __/")
	fmt.Println(" / _// /__/ __ |_\\ \\/ _  /___/ _//    / /__/ , _/\\  / ___/ / /   ")
	fmt.Println("/_/ /____/_/ |_/___/_//_/   /___/_/|_/\\___/_/|_| /_/_/    /_/    ")
	fmt.Println("v2.0!")
}

func getDrives() ([]string, error) {
	drives, err := fs.GetDrivesDarwin()
	if err != nil {
		panic(err)
	}

	fmt.Println("Drives:")
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
	targetDrive = strings.TrimSuffix(targetDrive, "\n")

	for i, drive := range drives {
		if strconv.Itoa(i) == targetDrive {
			fmt.Print("Are you sure you want to encrypt " + drive + " (yes/no)? ")
			confirmation, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			confirmation = strings.TrimSuffix(confirmation, "\n")

			if confirmation == "yes" {
				fmt.Print("Passphrase: ")
				passphrase, err := reader.ReadString('\n')
				if err != nil {
					return err
				}
				confirmation = strings.TrimSuffix(confirmation, "\n")

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
	return nil
}

func decrypt() error {
	drives, err := getDrives()
	if err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Which drive do you want to decrypt? ")
	targetDrive, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	targetDrive = strings.TrimSuffix(targetDrive, "\n")

	for i, drive := range drives {
		if strconv.Itoa(i) == targetDrive {
			fmt.Print("Are you sure you want to decrypt " + drive + " (yes/no)? ")
			confirmation, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			confirmation = strings.TrimSuffix(confirmation, "\n")

			if confirmation == "yes" {
				fmt.Print("Passphrase: ")
				passphrase, err := reader.ReadString('\n')
				if err != nil {
					return err
				}
				confirmation = strings.TrimSuffix(confirmation, "\n")

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
	return nil
}

func main() {
	printHeader()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("[0] Encrypt\n[1] Decrypt")
	option, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	option = strings.TrimSuffix(option, "\n")

	switch option {
	case "0":
		err = encrypt()
		if err != nil {
			panic(err)
		}
		break

	case "1":
		err = decrypt()
		if err != nil {
			panic(err)
		}

		break
	default:
		panic(errors.New(fmt.Sprintf("'%s' is not an option", option)))
	}

}
