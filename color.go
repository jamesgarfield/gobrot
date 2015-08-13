package main

import (
	"image/color"
	"math"
)

//Produce a color based on a percentage input (0-1)
type ColorScale interface {
	Color(float64) color.Color
}

type Palette []color.RGBA

func (p Palette) Color(x float64) color.Color {
	n := x * float64(len(p))
	c1 := p[int(math.Floor(n))]
	c2 := p[int(math.Min(math.Ceil(n), float64(len(p)-1)))]
	_, t := math.Modf(n)
	return blendRGBA(c1, c2, t)
}

//Blends two RGBA colors using linear interpolation
func blendRGBA(c1, c2 color.RGBA, t float64) color.RGBA {
	return color.RGBA{
		uint8(lerp(float64(c1.R), float64(c2.R), t)),
		uint8(lerp(float64(c1.G), float64(c2.G), t)),
		uint8(lerp(float64(c1.B), float64(c2.B), t)),
		uint8(lerp(float64(c1.A), float64(c2.A), t)),
	}
}

//linear interpolation between a & b at p
func lerp(a, b, p float64) float64 {
	return (1-p)*a + p*b
}
