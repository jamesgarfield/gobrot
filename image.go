package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
)

//An image can be defined by its size and a func that specifies what color a pixel is at a given coordinate
type ImageGenerator interface {
	Size() Size
	ColorSpecifier
}

type ColorSpecifier interface {
	Color(x, y int) color.Color
}

type Size struct {
	X, Y int
}

type Pixel struct {
	X, Y  int
	Color color.Color
}

//write an image to disk at a given path
func saveToPNG(img image.Image, path string) (err error) {
	out, err := os.Create(path)
	if err != nil {
		return
	}

	err = png.Encode(out, img)
	return
}

//Generate an image serially
func generateImage(gen ImageGenerator) *image.RGBA {
	size := gen.Size()
	img := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			img.Set(x, y, gen.Color(x-size.X/2, y-size.Y/2))
		}
	}
	return img
}

//generate concurrent pipeline code (pixelpipe_pipeline.go)
//go:generate goast write impl --prefix=generated_ goast.net/x/pipeline
type PixelPipe <-chan Pixel

type Pixels []Pixel

//Generate an image by processing its pixels in parallel
func generateImageP(gen ImageGenerator) *image.RGBA {

	size := gen.Size()

	done := make(chan bool)
	pipe := generatePixels(done, size)

	workers := runtime.NumCPU() - 1

	img := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))

	//Fan-Out color generation work over a number of workers, collect all results to block until completion
	pipe.Fan(done, workers, func(p Pixel) Pixel {
		p.Color = gen.Color(p.X-size.X/2, p.Y-size.Y/2)
		img.Set(p.X, p.Y, p.Color)
		return p
	}).Collect(done, size.X*size.Y)

	return img
}

//Creates a pixel pipeline and fills it with one per pixel in the image
func generatePixels(done <-chan bool, s Size) PixelPipe {
	out := make(chan Pixel)
	go (func() {
		for x := 0; x < s.X; x++ {
			for y := 0; y < s.Y; y++ {
				select {
				case out <- Pixel{x, y, nil}:
				case <-done:
					return
				}
			}
		}
	})()
	return out
}
