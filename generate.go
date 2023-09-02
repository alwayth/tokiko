package main

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"os"
	"regexp"
)

type Generator struct {
	input, output string
}

func NewGenerator(
	input, output string,
) *Generator {
	return &Generator{
		input:  input,
		output: output,
	}
}

func (g *Generator) Generate() error {
	r := regexp.MustCompile(`\.(gif|GIF)$`)
	if !r.MatchString(g.input) {
		return errors.New(fmt.Sprintf("tokiko: %s is not gif", g.input))
	}

	f, err := os.Open(g.input)
	if err != nil {
		panic(err)
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
	g.composite(decoded, lgtmImg)

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

func (g *Generator) composite(src *gif.GIF, img image.Image) {
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

	newFile, err := os.Create(g.output)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := newFile.Close(); err != nil {
			panic(err)
		}
	}()

	err = gif.EncodeAll(newFile, &out)
	if err != nil {
		panic(err)
	}
}
