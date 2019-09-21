package main

import (
	"fmt"
	"github.com/xoreo/flash-encrypt/cli"
)

// //go:generate go run scripts/includetxt.go

func main() {
	// Start the CLI
	err := cli.NewCLI()
	if err != nil { // Check err
		fmt.Println("Bye!")
	}
}
