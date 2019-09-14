package common

import (
	"errors"
	"strings"
)

var (
	// ErrNilInput is the error thrown when the input is nil
	ErrNilInput = errors.New("input can not be nil")
)

// ParseCommand parses a command from the CLI input.
func ParseCommand(input string) (string, string, []string, error) {
	if input == "" { // Check for errors
		return "", "", []string{}, ErrNilInput // Return found error
	} else if !strings.Contains(input, "(") || !strings.Contains(input, ")") {
		input = input + "()" // Fetch receiver methods
	}

	var method string
	if strings.Contains(input, ".") { // Check for nil receiver
		method = strings.Split(strings.Split(input, "(")[0], ".")[1] // Fetch method
	}

	receiver := GetReceiver(input)

	var params []string // Init buffer

	if strings.Contains(input, ",") || !strings.Contains(input, "()") { // Check for nil params
		params, _ = ParseParams(input) // Fetch params
	}

	return receiver, method, params, nil // No error occurred, return parsed method+params
}

// GetReceiver returns the receiver of a method call.
func GetReceiver(input string) string {
	return strings.Split(strings.Split(input, "(")[0], ".")[0] // Return split string
}

// ParseParams returns the parameters of a method call.
func ParseParams(input string) ([]string, error) {
	if input == "" { // Check for errors
		return []string{}, ErrNilInput // Return found error
	}

	parenthesesStripped := StripReceiver(input) // Strip parentheses
	params := strings.Split(parenthesesStripped, ", ") // Split by ', '

	return params, nil // No error occurred, return split params
}

// StripReceiver strips the receiver from a method call.
func StripReceiver(input string) string {
	openParenthIndex := strings.Index(input, "(")      // Get open parent index
	closeParenthIndex := strings.LastIndex(input, ")") // Get close parent index

	return input[openParenthIndex+1 : closeParenthIndex] // Strip receiver
}