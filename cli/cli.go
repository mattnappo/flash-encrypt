package cli

import "github.com/xoreo/flash-encrypt/api"

// NewCLI creates a new CLI.
func NewCLI() {

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
		err := api.Encrypt()
		if err != nil {
			return err
		}
		break

	case "decrypt":
		err := api.Decrypt()
		if err != nil {
			return err
		}
		break
	}
}