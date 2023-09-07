package main

import (
	"fmt"
	"os"
	"tokiko/cli"
	"tokiko/generate"
)

func main() {
	cli := cli.NewCLI(os.Stdin, os.Stdout)
	input, output, err := cli.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	generator := generate.NewGenerator()
	if err := generator.Generate(input, output); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
