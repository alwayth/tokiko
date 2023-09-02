package main

import (
	"fmt"
	"os"
)

func main() {
	cli := NewCLI(os.Stdin, os.Stdout)
	input, output, err := cli.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	generator := NewGenerator(input, output)
	if err := generator.Generate(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
