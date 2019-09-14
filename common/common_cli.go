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

	receiver := StringFetchCallReceiver(input)

	var params []string // Init buffer

	if strings.Contains(input, ",") || !strings.Contains(input, "()") { // Check for nil params
		params, _ = ParseStringParams(input) // Fetch params
	}

	return receiver, method, params, nil // No error occurred, return parsed method+params
}

// StringFetchCallReceiver - attempt to fetch receiver from string, as if it were an x.y(..., ..., ...) style method call
func StringFetchCallReceiver(input string) string {
	return strings.Split(strings.Split(input, "(")[0], ".")[0] // Return split string
}

// ParseStringParams - attempt to fetch string parameters from (..., ..., ...) style call
func ParseStringParams(input string) ([]string, error) {
	if input == "" { // Check for errors
		return []string{}, ErrNilInput // Return found error
	}

	parenthesesStripped := StringStripReceiverCall(input) // Strip parentheses

	params := strings.Split(parenthesesStripped, ", ") // Split by ', '

	return params, nil // No error occurred, return split params
}

// StringStripReceiverCall - strip receiver from string method call
func StringStripReceiverCall(input string) string {
	openParenthIndex := strings.Index(input, "(")      // Get open parent index
	closeParenthIndex := strings.LastIndex(input, ")") // Get close parent index

	return input[openParenthIndex+1 : closeParenthIndex] // Strip receiver
}
