package main

import (
	"flag"
	"io"
)

type CLI struct {
	stdIn, stdOut io.Writer
}

func NewCLI(
	stdIn, stdOut io.Writer,
) CLI {
	return CLI{
		stdIn:  stdIn,
		stdOut: stdOut,
	}
}

func (c *CLI) Run(
	args []string,
) (
	inputFilePathFlag string,
	outputFilePathFlag string,
	err error,
) {
	fs := flag.NewFlagSet("tokiko", flag.ExitOnError)
	fs.SetOutput(c.stdOut)

	fs.StringVar(&inputFilePathFlag, "i", "input.gif", "input git file path.")
	fs.StringVar(&outputFilePathFlag, "o", "output.gif", "output gif file path.")

	if err = fs.Parse(args[1:]); err != nil {
		return "", "", err
	}

	return inputFilePathFlag, outputFilePathFlag, nil
}
