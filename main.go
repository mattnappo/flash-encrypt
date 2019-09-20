package main

import (
	"fmt"
	"github.com/xoreo/flash-encrypt/cli"
)

func main() {
	err := cli.NewCLI()
	if err != nil {
		fmt.Println()
	}
}
