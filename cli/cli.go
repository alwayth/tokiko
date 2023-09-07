package cli

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
	inputFile string,
	outputFile string,
	err error,
) {
	fs := flag.NewFlagSet("tokiko", flag.ExitOnError)
	fs.SetOutput(c.stdOut)

	fs.StringVar(&inputFile, "i", "input.gif", "input git file")
	fs.StringVar(&outputFile, "o", "output.gif", "output gif file")

	if err = fs.Parse(args[1:]); err != nil {
		return "", "", err
	}

	return inputFile, outputFile, nil
}
