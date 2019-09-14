package main

import (
	"github.com/xoreo/flash-encrypt/cli"
)

func main() {
	err := cli.NewCLI()
	if err != nil {
		panic(err)
	}
}
