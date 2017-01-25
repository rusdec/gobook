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
	width, height = 600, 300
	cells = 100
	xyrangeDefault = 40.0
	zscale = height*0.4
	angle = math.Pi/6

	min = 50
	max = 500	
	colorMin = 0x0000ff
	colorMax = 0xff0000
)

var (
		sin30, cos30 = math.Sin(angle), math.Cos(angle)
		color int
		xyscale, xyrange float64
)

func Polygon(w http.ResponseWriter, r *http.Request) {
	val, err := strconv.ParseFloat(r.URL.Query().Get("xyrange"), 64)
	if err != nil {
		xyrange = xyrangeDefault
	} else {
		xyrange = float64(val)
	}
	xyscale = width/2/xyrange
	
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
