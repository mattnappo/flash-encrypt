package main

import (
	"flag"
	"fmt"

	"github.com/xoreo/flash-encrypt/cli"
)

// //go:generate go run scripts/includetxt.go

// StandaloneFlag determines which mode to run the program in.
var StandaloneFlag = flag.Bool("standalone", false, "Run flash-encrypt independently")

func main() {
	flag.Parse()

	if *StandaloneFlag {
		cli.StandaloneCLI()
	} else {
		// Start the CLI
		err := cli.NewCLI()
		if err != nil { // Check err
			fmt.Println("Bye!")
		}
	}

}
