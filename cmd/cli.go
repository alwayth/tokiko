package cmd

import (
	"flag"
	"io"
	"os"

	"tokiko/generate"
)

var (
	inputFile  string
	outputFile string
)

type Cmd struct {
	stdIn, stdOut io.Writer
}

func NewCmd(
	stdIn, stdOut io.Writer,
) Cmd {
	return Cmd{
		stdIn:  stdIn,
		stdOut: stdOut,
	}
}

func (c *Cmd) Execute(args []string) error {
	fs := flag.NewFlagSet("tokiko", flag.ContinueOnError)
	fs.SetOutput(c.stdOut)

	fs.StringVar(&inputFile, "i", "input.gif", "input git file")
	fs.StringVar(&outputFile, "o", "output.gif", "output gif file")

	if err := fs.Parse(args[1:]); err != nil {
		return err
	}

	if err := generate.Generate(inputFile, outputFile); err != nil {
		os.Exit(0)
	}

	return nil
}
