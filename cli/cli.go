package cli

import (
	"bufio"
	"fmt"
	"github.com/xoreo/flash-encrypt/api"
	"os"
	"os/exec"
	"strings"
)

// NewCLI creates a new CLI.
func NewCLI() error {
	// Print the header information
	printHeader()

	// Set stdin input buffer
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n> ")

		// Read input from user
		commandString, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		commandString = strings.TrimSuffix(commandString, "\n")

		// Parse the command
		command, err := ParseCommand(commandString)
		if err != nil {
			return err
		}

		// Handle the command
		err = handleCommand(command)
		if err != nil {
			return err
		}

	}
}

func handleCommand(command Command) error {
	switch command.Receiver {
	case "":
		err := handleNoReceiver(command)
		if err != nil {
			return err
		}
		break
	}

	return nil
}

func handleNoReceiver(command Command) error {
	switch command.Method {
	case "encrypt":
		// Run the code to encrypt
		err := api.Encrypt()
		if err != nil {
			return err
		}
		break

	case "decrypt":
		// Run the code to decrypt
		err := api.Decrypt()
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

	case "status":
		// Run the code to check the status of a drive
		err := api.Status()
		if err != nil {
			return err
		}
		break

	case "help":
		printHelp()
		break

	case "exit":
		return nil

	default:
		fmt.Printf("'%s' is not a valid command. Run 'help' for help.", command.Method)
	}

	return nil
}

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
	fmt.Println("this is the help")
}