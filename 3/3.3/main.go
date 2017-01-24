/**
## Задание 3.3

Окрасьте каждый многоугольник цветом, зависящим от его высоты, так, чтобы
пики были красными (#ff0000), а низины синими (#0000ff).

*/

package main

import (
	"fmt"
	"math"
	"errors"
)

const (
	width, height = 600, 300
	cells = 100
	xyrange = 30.0
	xyscale = width/2/xyrange
	zscale = height*0.4
	angle = math.Pi/6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>\n", width, height)
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
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
		
	}
	fmt.Printf("</svg>")
}

func corner(i, j int) (float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x,y)
	
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
