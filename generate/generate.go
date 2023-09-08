package generate

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

func Generate(
	inputFile,
	outputFile string,
) error {
	if err := validate(inputFile); err != nil {
		return err
	}

	f, err := os.Open(inputFile)
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

	lgtmImg := readLGTMImg()
	src := composite(decoded, lgtmImg)

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

func validate(
	name string,
) error {
	r := regexp.MustCompile(`\.(gif|GIF)$`)
	if !r.MatchString(name) {
		return errors.New(fmt.Sprintf("tokiko: %s is not gif", name))
	}

	return nil
}

func readLGTMImg() image.Image {
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

func composite(
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
		draw.Draw(frame, frame.Bounds(), img, image.Point{X: -x, Y: -y}, draw.Over)
		out.Image = append(out.Image, frame)
	}

	return out
}
