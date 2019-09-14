package cli

import (
	"errors"
	"strings"
)

var (
	// ErrNilInput is the error thrown when the input is nil
	ErrNilInput = errors.New("input can not be nil")
)

// ParseCommand parses a command from the CLI input.
func ParseCommand(input string) (Command, error) {
	if input == "" { // Check for errors
		return Command{}, ErrNilInput // Return found error
	} else if !strings.Contains(input, "(") || !strings.Contains(input, ")") {
		input = input + "()" // Fetch receiver methods
	}

	var method, receiver string
	if !strings.Contains(input, ".") { // Check for nil receiver
		method = strings.Split(input, "(")[0] // Fetch method
		receiver = ""
	} else {
		method = strings.Split(strings.Split(input, "(")[0], ".")[1] // Fetch method
		receiver = getReceiver(input)
	}

	var params []string // Init buffer

	if strings.Contains(input, ",") || !strings.Contains(input, "()") { // Check for nil params
		params, _ = parseParams(input) // Fetch params
	}

	command, err := NewCommand(receiver, method, params)
	if err != nil {
		return command, err
	}

	return command, nil // No error occurred, return parsed method+params
}

/* -- BEGIN INTERNAL METHODS -- */

// getReceiver returns the receiver of a method call.
func getReceiver(input string) string {
	return strings.Split(strings.Split(input, "(")[0], ".")[0] // Return split string
}

// parseParams returns the parameters of a method call.
func parseParams(input string) ([]string, error) {
	if input == "" { // Check for errors
		return []string{}, ErrNilInput // Return found error
	}

	parenthesesStripped := stripReceiver(input)        // Strip parentheses
	params := strings.Split(parenthesesStripped, ", ") // Split by ', '

	return params, nil // No error occurred, return split params
}

// stripReceiver strips the receiver from a method call.
func stripReceiver(input string) string {
	openParenthIndex := strings.Index(input, "(")      // Get open parent index
	closeParenthIndex := strings.LastIndex(input, ")") // Get close parent index

	return input[openParenthIndex+1 : closeParenthIndex] // Strip receiver
}

/* -- END INTERNAL METHODS -- */
