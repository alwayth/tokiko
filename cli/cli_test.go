package cli_test

import (
	"bytes"
	"strings"
	"testing"

	"tokiko/cli"

	"github.com/stretchr/testify/assert"
)

func TestCLI_Run(t *testing.T) {
	t.Run("", func(t *testing.T) {
		c := cli.NewCLI(
			new(bytes.Buffer),
			new(bytes.Buffer),
		)

		args := strings.Split("tokiko -i sample.gif -o sample.lgtm", " ")
		_, _, err := c.Run(args)
		assert.Error(t, err)
	})
}
