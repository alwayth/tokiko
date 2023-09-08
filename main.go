package main

import (
	"fmt"
	"os"

	"tokiko/cli"
	"tokiko/generate"
)

func main() {
	c := cli.NewCLI(os.Stdin, os.Stdout)
	input, output, err := c.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if err := generate.Generate(input, output); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
