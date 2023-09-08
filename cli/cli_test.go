package cli_test

import (
	"bytes"
	"strings"
	"testing"

	"tokiko/cli"

	"github.com/stretchr/testify/assert"
)

func TestCLI_Run(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := cli.NewCLI(
			new(bytes.Buffer),
			new(bytes.Buffer),
		)

		args := strings.Split("tokiko -i sample.gif -o sample.lgtm", " ")
		_, _, err := c.Run(args)
		assert.NoError(t, err)
	})

	t.Run("invalid input option", func(t *testing.T) {
		c := cli.NewCLI(
			new(bytes.Buffer),
			new(bytes.Buffer),
		)

		args := strings.Split("tokiko -x sample.gif -o sample.lgtm", " ")
		_, _, err := c.Run(args)
		assert.Error(t, err)
	})

	t.Run("invalid output option", func(t *testing.T) {
		c := cli.NewCLI(
			new(bytes.Buffer),
			new(bytes.Buffer),
		)

		args := strings.Split("tokiko -i sample.gif -x sample.lgtm", " ")
		_, _, err := c.Run(args)
		assert.Error(t, err)
	})
}
