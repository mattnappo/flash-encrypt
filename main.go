package main

import (
	"bufio"
	"fmt"
	"github.com/xoreo/flash-encrypt/crypto"
	"github.com/xoreo/flash-encrypt/fs"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func printHeader() {
	exec.Command("clear")
	fmt.Println("Welcome to")
	fmt.Println("   ______   ___   ______ __    _____  ____________  _____  ______")
	fmt.Println("  / __/ /  / _ | / __/ // /___/ __/ |/ / ___/ _ \\ \\/ / _ \\/_  __/")
	fmt.Println(" / _// /__/ __ |_\\ \\/ _  /___/ _//    / /__/ , _/\\  / ___/ / /   ")
	fmt.Println("/_/ /____/_/ |_/___/_//_/   /___/_/|_/\\___/_/|_| /_/_/    /_/    ")
	fmt.Println("v2.0!")
	fmt.Println("Run 'help' for help!")
}

func printHelp() {
	fmt.Println("This is the help menu.")
}

func printDrives() ([]string, error) {
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
	drives, err := printDrives()
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
	drives, err := printDrives()
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
	_, err := printDrives()
	if err != nil {
		panic(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\n> ")

		command, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		command = strings.TrimSuffix(command, "\n")

		switch strings.ToLower(command) {
		case "encrypt":
			err = encrypt()
			if err != nil {
				panic(err)
			}
			break

		case "decrypt":
			err = decrypt()
			if err != nil {
				panic(err)
			}
			break

		case "exit":
			break

		case "help":
			printHelp()
			break

		default:
			fmt.Printf("'%s' is not a valid command. Run 'help' for help.", command)
		}
	}

}
