package generate_test

import (
	"errors"
	"image"
	"image/gif"
	"image/png"
	"io"
	"os"
	"testing"

	"tokiko/generate"

	"github.com/stretchr/testify/assert"
)

type MockOSWrapper struct{}

func (m *MockOSWrapper) Open(name string) (*os.File, error)   { return nil, errors.New("err") }
func (m *MockOSWrapper) Create(name string) (*os.File, error) { return os.Create(name) }

type MockGIFWrapper struct{}

func (m *MockGIFWrapper) EncodeAll(w io.Writer, g *gif.GIF) error { return gif.EncodeAll(w, g) }

type MockPNGWrapper struct{}

func (m *MockPNGWrapper) Decode(r io.Reader) (image.Image, error) { return png.Decode(r) }

func TestGenerate(t *testing.T) {
	t.Run("invalid file name", func(t *testing.T) {
		g := generate.Generator{}

		err := g.Generate("invalid", "output.gif")
		assert.Error(t, err)
	})

	t.Run("open file error", func(t *testing.T) {
		g := generate.Generator{
			IOSWrapper: &MockOSWrapper{},
		}

		err := g.Generate("input.gif", "output.gif")
		assert.Error(t, err)
	})
}
