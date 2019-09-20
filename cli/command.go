package cli

import (
	"errors"
	"strings"
)

// Command represents a command.
type Command struct {
	Receiver string   `json:"receiver"` // The receiver of the command
	Method   string   `json:"method"`   // The method of the command
	Params   []string `json:"params"`   // The parameters of the command
}

// NewCommand constructs a new command.
func NewCommand(receiver, method string, params []string) (Command, error) {
	// Check arguments
	if method == "" {
		return Command{}, errors.New("invalid params to construct Command")
	}

	// Return a new command
	return Command{
		Receiver: strings.ToLower(receiver),
		Method:   strings.ToLower(method),
		Params:   params,
	}, nil
}
