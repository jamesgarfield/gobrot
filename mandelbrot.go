package main

import (
	"image/color"
	"math/cmplx"
)

type MandelbrotSet struct {
	MandelbrotView
	ColorScale
}

func (m MandelbrotSet) Size() Size {
	return m.MandelbrotView.Size
}

func (m MandelbrotSet) Color(x, y int) color.Color {
	c := scale(x, y, m.Center, m.Zoom)
	escaped, at, _ := mandelbrot(c, m.Iter, m.Limit)
	if escaped {
		return m.ColorScale.Color(float64(at) / float64(m.Iter))
	}
	return color.Black
}

type MandelbrotView struct {
	Center      complex128
	Zoom, Limit float64
	Iter        int
	Size
}

//Applies the mandelbrot escape analysis on a given complex coordinate.
func mandelbrot(c complex128, iter int, limit float64) (escaped bool, at int, magnitute float64) {
	z := complex(0, 0)
	for at = 0; at < iter; at++ {
		z = z*z + c
		magnitute = cmplx.Abs(z)
		if magnitute > limit {
			escaped = true
			return
		}
	}
	return
}

//Transform an x,y pixel coordinate to a complex coordinate based on a given center and zoomc
func scale(x, y int, center complex128, zoom float64) complex128 {
	return center + complex(float64(x)/zoom, float64(y)/zoom)
}
