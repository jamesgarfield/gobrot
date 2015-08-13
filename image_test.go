package main

import (
	"testing"
)

//base case, full unzoomed view
var baseBrot = MandelbrotSet{
	MandelbrotView{
		Center: complex(0, 0),
		Zoom:   100,
		Limit:  1000,
		Iter:   500,
		Size:   Size{100, 100},
	},
	Palette{BLUE, WHITE, ORANGE, WHITE, PURPLE},
}

//worst case scenario, all rendered points within mandelbrot
var worstBrot MandelbrotSet = MandelbrotSet{
	MandelbrotView{
		Center: complex(0, 0),
		Zoom:   5000,
		Limit:  5000,
		Iter:   5000,
		Size:   Size{100, 100},
	},
	Palette{BLUE, WHITE, ORANGE, WHITE, PURPLE},
}

func Benchmark_ImgGen_Base(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generateImage(baseBrot)
	}
}

func Benchmark_ImgGenP_Base(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generateImageP(baseBrot)
	}
}

func Benchmark_ImgGen_Worst(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generateImage(worstBrot)
	}
}

func Benchmark_ImgGenP_Worst(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generateImageP(worstBrot)
	}
}
