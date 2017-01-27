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
	width, height = 2100, 2100
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	
	for py := 0; py <= height; py++ {
		y	:= float64(py) / height * (ymax - ymin) + ymin
		y1	:= float64(py+1) / height * (ymax - ymin) + ymin
		for px := 0; px <= width; px++ {
			x	:= float64(px) / width * (xmax - xmin) + xmin
			x1	:= float64(px+1) / width * (xmax - xmin) + xmin
			img.Set(px, py, superSampling(x, y, x1, y1))
			
		}
	}

	png.Encode(os.Stdout, img)
}

func superSampling(x, y, x1, y1 float64) color.Color{
	colors := [3]uint16{}
	z := complex(x, y)
	for i, k := range mandelbrot(z) {
		colors[i] += k
	}
	z = complex(x1, y)
	for i, k := range mandelbrot(z) {
		colors[i] += k
	}
	z = complex(x, y1)
	for i, k := range mandelbrot(z) {
		colors[i] += k
	}
	z = complex(x1, y1)
	for i, k := range mandelbrot(z) {
		colors[i] += k
	}
	for i, k := range colors {
		colors[i] = k/4
	}
	
	return color.YCbCr{uint8(colors[0]), uint8(colors[1]), uint8(colors[2])}
}

func mandelbrot(z complex128) ([3]uint16) {
	const (
		iterations = 200
		contrast = 15
	)
	
	var v complex128
	var colorMax uint16 = 128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return [3]uint16{16, colorMax, colorMax}
		}
	}
	
	return [3]uint16{161,118,10}
}
