package main

import (
	"flag"
	"image/color"
	"math/rand"
	"runtime"
	"time"
)

var (
	BLACK  = color.RGBA{0, 0, 0, 255}
	WHITE  = color.RGBA{255, 255, 255, 255}
	RED    = color.RGBA{255, 0, 0, 255}
	GREEN  = color.RGBA{0, 255, 0, 255}
	BLUE   = color.RGBA{0, 0, 255, 255}
	ORANGE = color.RGBA{255, 115, 0, 255}
	PURPLE = color.RGBA{255, 0, 255, 255}
)

var (
	path        *string
	center_x    *float64
	center_y    *float64
	zoom        *float64
	width       *int
	height      *int
	iter        *int
	limit       *float64
	interesting *bool
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(int64(time.Now().Nanosecond()))

	path = flag.String("path", "./brot.png", "Output filename")
	center_x = flag.Float64("x", -.5, "X center on mandelbrot plane")
	center_y = flag.Float64("y", 0, "Y center on mandelbrot plane")
	zoom = flag.Float64("zoom", 1, "Zoom Multiplier")
	width = flag.Int("width", 500, "Width of final image")
	height = flag.Int("height", 500, "Height of final iage")
	iter = flag.Int("iter", 150, "Escape analysis iterations")
	limit = flag.Float64("limit", 1000, "Limit for escape analysis")
	interesting = flag.Bool("interesting", false, "Use preselected interesting parameters")
}

func main() {

	flag.Parse()

	//make something pretty
	if *interesting {
		*center_x = -.8525
		*center_y = -.20995
		*zoom = 128000
		*iter = 1000
		*limit = 2000
		*width = 1000
		*height = 1000
	}

	img := generateImageP(MandelbrotSet{
		MandelbrotView{
			Center: complex(*center_x, *center_y),
			Zoom:   *zoom * 100,
			Limit:  *limit,
			Iter:   *iter,
			Size:   Size{*width, *height},
		},
		Palette{BLUE, WHITE, ORANGE, WHITE, PURPLE},
	})

	if err := saveToPNG(img, *path); err != nil {
		panic(err)
	}

}
