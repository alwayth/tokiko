package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCLI_Run(t *testing.T) {
	t.Run("", func(t *testing.T) {
		c := NewCLI(
			new(bytes.Buffer),
			new(bytes.Buffer),
		)

		args := strings.Split("tokiko -i sample.gif -o sample.lgtm", " ")
		_, _, err := c.Run(args)
		if err != nil {
			t.FailNow()
		}
	})
}
