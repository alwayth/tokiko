package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"tokiko/cmd"

	"github.com/stretchr/testify/assert"
)

func TestCLI_Run(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := cmd.NewCmd(
			new(bytes.Buffer),
			new(bytes.Buffer),
		)

		args := strings.Split("tokiko -i sample.gif -o sample.lgtm", " ")
		err := c.Execute(args)
		assert.NoError(t, err)
	})

	t.Run("invalid input option", func(t *testing.T) {
		c := cmd.NewCmd(
			new(bytes.Buffer),
			new(bytes.Buffer),
		)

		args := strings.Split("tokiko -x sample.gif -o sample.lgtm", " ")
		err := c.Execute(args)
		assert.Error(t, err)
	})

	t.Run("invalid output option", func(t *testing.T) {
		c := cmd.NewCmd(
			new(bytes.Buffer),
			new(bytes.Buffer),
		)

		args := strings.Split("tokiko -i sample.gif -x sample.lgtm", " ")
		err := c.Execute(args)
		assert.Error(t, err)
	})
}
