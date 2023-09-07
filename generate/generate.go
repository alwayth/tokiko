package generate

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"io"
	"os"
	"regexp"
)

type IOSWrapper interface {
	Open(name string) (*os.File, error)
	Create(name string) (*os.File, error)
}

type OSWrapper struct{}

func (m *OSWrapper) Open(name string) (*os.File, error)   { return os.Open(name) }
func (m *OSWrapper) Create(name string) (*os.File, error) { return os.Create(name) }

type GIFWrapper struct{}

func (m *GIFWrapper) EncodeAll(w io.Writer, g *gif.GIF) error { return gif.EncodeAll(w, g) }

type PNGWrapper struct{}

func (m *PNGWrapper) Decode(r io.Reader) (image.Image, error) { return png.Decode(r) }

type Generator struct {
	IOSWrapper IOSWrapper
}

func NewGenerator() *Generator {
	return &Generator{
		IOSWrapper: &OSWrapper{},
	}
}

func (g *Generator) Generate(
	inputFile,
	outputFile string,
) error {

	if err := g.validate(inputFile); err != nil {
		return err
	}

	f, err := g.IOSWrapper.Open(inputFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	decoded, err := gif.DecodeAll(f)
	if err != nil {
		panic(err)
	}

	lgtmImg := g.readLGTMImg()
	src := g.composite(decoded, lgtmImg)

	newFile, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := newFile.Close(); err != nil {
			panic(err)
		}
	}()

	err = gif.EncodeAll(newFile, &src)
	if err != nil {
		panic(err)
	}

	return nil
}

func (g *Generator) validate(
	name string,
) error {
	r := regexp.MustCompile(`\.(gif|GIF)$`)
	if !r.MatchString(name) {
		return errors.New(fmt.Sprintf("tokiko: %s is not gif", name))
	}

	return nil
}

func (g *Generator) readLGTMImg() image.Image {
	f, err := os.Open("assets/lgtm.png")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	return img
}

func (g *Generator) composite(
	src *gif.GIF,
	img image.Image,
) gif.GIF {
	out := gif.GIF{
		Delay:     src.Delay,
		LoopCount: src.LoopCount,
		Disposal:  src.Disposal,
		Config: image.Config{
			Width:  src.Config.Width,
			Height: src.Config.Height,
		},
	}

	x := src.Config.Width/2 - img.Bounds().Dx()/2
	y := src.Config.Height/2 - img.Bounds().Dy()/2

	for _, frame := range src.Image {
		draw.Draw(frame, frame.Bounds(), img, image.Point{-x, -y}, draw.Over)
		out.Image = append(out.Image, frame)
	}

	return out
}

func (g *Generator) createGIFFile(
	src *gif.GIF,
	outputFilePath string,
) {
	newFile, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := newFile.Close(); err != nil {
			panic(err)
		}
	}()

	err = gif.EncodeAll(newFile, src)
	if err != nil {
		panic(err)
	}
}
