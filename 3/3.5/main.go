/**
## Задание 3.5

Реализуйте полноцветное множество Мандельброта с использованием функции
image.NewRGBA и типа color.RGBA или color.YCbCr

*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height = 2000, 2000
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	
	for py := 0; py <= height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px <= width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
			
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		contrast = 15
	)
	
	var v complex128
	var colorMax uint8 = 128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{colorMax - contrast*n, colorMax - contrast*n, colorMax - contrast*n}
		}
	}
	
	return color.YCbCr{200,78,90}
}
