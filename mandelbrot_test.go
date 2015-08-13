package main

import (
	"testing"
)

func Test_mandelbrot(t *testing.T) {
	brot := []complex128{complex(0, 0), complex(-1, 0), complex(0, .5), complex(.1, -.2), complex(-2, 0)}
	for _, c := range brot {
		if escaped, _, _ := mandelbrot(c, 50, 1000); escaped {
			t.Errorf("%+v within mandelbrot, should not escape", c)
		}
	}

	notbrot := []complex128{complex(1, 1), complex(-1, -1), complex(0, 1.5), complex(0, -1.5)}
	for _, c := range notbrot {
		if escaped, _, _ := mandelbrot(c, 50, 1000); !escaped {
			t.Errorf("%+v not within mandelbrot, should escape", c)
		}
	}
}
