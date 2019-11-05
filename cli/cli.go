package cli

import (
	"bufio"
	"errors"
	"fmt"

	"os"
	"os/exec"
	"strings"

	"github.com/xoreo/flash-encrypt/api"
	"github.com/xoreo/flash-encrypt/common"
)

var (
	// ErrInvalidParamCount is thrown when an invalid number of parameters is used in a command.
	ErrInvalidParamCount = errors.New("invalid number of parameters to call method")
)

// NewCLI creates a new CLI.
func NewCLI(isStandalone bool) error {
	// Print the header information
	printHeader(isStandalone)

	// Set stdin input buffer
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		// Read input from user
		commandString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		commandString = strings.TrimSuffix(commandString, "\n")

		// Parse the command
		command, err := ParseCommand(commandString)
		if err != nil {
			fmt.Println(err.Error())
		}

		// Handle the command
		err = handleCommand(command, isStandalone)
		if err != nil {
			// If exit
			if err.Error() == "exit" {
				return err
			}

			fmt.Println(err.Error()) // Else, print the error
		}
	}
}

// handleCommand determines which receiver to use to execute the command.
func handleCommand(command Command, isStandalone bool) error {
	// The receiver handler
	switch command.Receiver {
	case "":
		err := handleNoReceiver(command, isStandalone)
		if err != nil {
			return err
		}

		break

	default:
		fmt.Printf("'%s' is an unknown receiver.\n", command.Receiver)
	}

	return nil
}

// handleNoReceiver handles commands with no receiver.
func handleNoReceiver(command Command, isStandalone bool) error {
	switch command.Method {
	case "encrypt":
		if isStandalone { // If this is a standalone CLI
			// Make sure that the directory exists, then encrypt it.
			common.CreateDirIfDoesNotExist(common.EncryptionDir)
			err := api.EncryptDir(common.EncryptionDir)
			if err != nil {
				return err
			}

			return nil
		}

		// Check params
		if len(command.Params) != 1 {
			return ErrInvalidParamCount
		}

		// Run the code to encrypt
		err := api.EncryptDrive(command.Params[0])
		if err != nil {
			return err
		}
		break

	case "decrypt":
		if isStandalone { // If this is a standalone CLI
			// Make sure that the directory exists, then encrypt it.
			common.CreateDirIfDoesNotExist(common.EncryptionDir)
			err := api.DecryptDir(common.EncryptionDir)
			if err != nil {
				return err
			}

			break
		}

		// Check params
		if len(command.Params) != 1 {
			return ErrInvalidParamCount
		}

		// Run the code to decrypt
		err := api.DecryptDrive(command.Params[0])
		if err != nil {
			return err
		}
		break

	case "drives":
		// Run the code to list the current drives
		err := api.ListDrives()
		if err != nil {
			return err
		}
		break

	case "help":
		if isStandalone {
			// Run the help command
			err := printHelp(isStandalone)
			if err != nil {
				return err
			}
			break
		}

		// Run the help command
		err := printHelp(isStandalone)
		if err != nil {
			return err
		}
		break

	case "exit":
		return errors.New("exit")

	default:
		fmt.Printf("'%s' is not a valid command. Run 'help' for help.\n", command.Method)
	}

	return nil
}

// printHeader prints the header of the CLI.
func printHeader(isStandalone bool) {
	exec.Command("clear")
	fmt.Println("Welcome to")
	fmt.Println("   ______   ___   ______ __    _____  ____________  _____  ______")
	fmt.Println("  / __/ /  / _ | / __/ // /___/ __/ |/ / ___/ _ \\ \\/ / _ \\/_  __/")
	fmt.Println(" / _// /__/ __ |_\\ \\/ _  /___/ _//    / /__/ , _/\\  / ___/ / /   ")
	fmt.Println("/_/ /____/_/ |_/___/_//_/   /___/_/|_/\\___/_/|_| /_/_/    /_/    ")
	if isStandalone {
		fmt.Println("-- Standalone Mode --")
	}
	fmt.Println("v2.0!")
	fmt.Println("Run 'help' for help!")
}

// printHelp prints the help screen of the CLI.
func printHelp(isStandalone bool) error {
	if isStandalone {
		// Print the standalone help file
		helpMenu, err := Asset("assets/help_standalone.txt")
		if err != nil {
			return err
		}
		fmt.Println(string(helpMenu))
		return nil
	}

	// Print the help file
	helpMenu, err := Asset("assets/help.txt")
	if err != nil {
		return err
	}
	fmt.Println(string(helpMenu))

	return nil
}
