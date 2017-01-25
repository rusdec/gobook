/**
Задание 3.4

Следуя подходу, использованному в примере с фигурами Лиссажу из раздела 1.7,
создайте веб-сервер, который вычисляетповерхности и возвращает клиенту
SVG-файл. Сервер должениспользовать в ответе заголовок ContentType наподобие
следующего:

w.Header().Set("ContentType", "image/svg+xml")

*/

package main

import (
	"fmt"
	"math"
	"errors"
	"net/http"
	"strconv"
)

const (
	widthDefault, heightDefault = 100.0, 100.0
	sizeMax = 800.0
	sizeMin = 100.0
	cells = 100
	xyrangeDefault = 40.0
	angle = math.Pi/6

	colorMinDefault = 0xffffff
	colorMaxDefault = 0xffffff
)

var (
		sin30, cos30 = math.Sin(angle), math.Cos(angle)
		color int64
		xyscale, xyrange, zscale float64
		colorMin, colorMax int64
		width, height float64
)

func parseFloat64(r *http.Request, s string) (float64, error) {
	return strconv.ParseFloat(r.URL.Query().Get(s), 64)
}

func parseInt64(r *http.Request, s string, base int) (int64, error) {
	return strconv.ParseInt(r.URL.Query().Get(s), base, 64)
}

func Polygon(w http.ResponseWriter, r *http.Request) {
	var err error

	xyrange, err = parseFloat64(r, "xyrange")
	if err != nil {
		xyrange = xyrangeDefault
	}
	colorMin, err = parseInt64(r, "colorMin", 16)
	if err != nil {
		colorMin = colorMinDefault
	}
	colorMax, err = parseInt64(r, "colorMax", 16)
	if err != nil {
		colorMax = colorMaxDefault
	}
	size, err := parseFloat64(r, "size")
	if err != nil || size < sizeMin || size > sizeMax {
		width = widthDefault
		height = heightDefault
	} else {
		width = size
		height = size
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; stroke-width: 0.7' width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#%06x'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
		}
		
	}
	fmt.Fprintf(w,"</svg>")
}

func corner(i, j int) (float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x,y)

	xyscale = float64(width)/2/xyrange
	zscale = height*0.4
	setColor(z)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	if math.IsNaN(sx) || math.IsNaN(sy) {
		return sx, sy, errors.New("nan")
	}
	 	
	return sx, sy, nil	
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	
	return math.Sin(r)/r
}

func setColor(z float64) {
	if z < 0 {
		color = colorMin
	} else {
		color = colorMax
	}
}
