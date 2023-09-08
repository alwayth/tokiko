package main

import (
	"os"

	"tokiko/cmd"
)

func main() {
	c := cmd.NewCmd(os.Stdin, os.Stdout)
	if err := c.Execute(os.Args); err != nil {
		os.Exit(0)
	}
}
