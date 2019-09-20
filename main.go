package main

import (
	"fmt"
	"github.com/xoreo/flash-encrypt/cli"
)

func main() {
	// Start the CLI
	err := cli.NewCLI()
	if err != nil { // Check err
		fmt.Println()
	}
}
