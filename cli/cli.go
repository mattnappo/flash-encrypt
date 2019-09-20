package cli

import (
	"bufio"
	"errors"
	"fmt"

	"io/ioutil"
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
func NewCLI() error {
	// Print the header information
	printHeader()

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
		err = handleCommand(command)
		if err.Error() == "exit" {
			return err
		}

		if err != nil {
			fmt.Println(err.Error())
		}
	}

}

// handleCommand determines which receiver to use to execute the command.
func handleCommand(command Command) error {
	// The receiver handler
	switch command.Receiver {
	case "":
		err := handleNoReceiver(command)
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
func handleNoReceiver(command Command) error {
	switch command.Method {
	case "encrypt":
		// Check params
		if len(command.Params) != 1 {
			return ErrInvalidParamCount
		}

		// Run the code to encrypt
		err := api.Encrypt(command.Params[0])
		if err != nil {
			return err
		}
		break

	case "decrypt":
		// Check params
		if len(command.Params) != 1 {
			return ErrInvalidParamCount
		}

		// Run the code to decrypt
		err := api.Decrypt(command.Params[0])
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
		// Run the help command
		err := printHelp()
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

// printHelp prints the help screen of the CLI.
func printHelp() error {
	helpFile := fmt.Sprintf("cli%shelp.txt", common.OSSlash)

	// Read from the help file
	help, err := ioutil.ReadFile(helpFile)
	if err != nil {
		return err
	}

	// Print the help file
	helpString := string(help)
	fmt.Println(helpString)

	return nil
}
