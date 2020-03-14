// Package surface computes an SVG rendering of a 3-d surface function.
package surface

import (
	"fmt"
	"io"
	"math"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// Render returns the SVG string that renders a 3-d surface function
func Render(w io.Writer, color string, width float64, height float64) {
	//there must be a cleaner way to nil check arguments passed to a function
	if color == "" {
		color = "grey"
	}

	if width == 0 {
		width = 600
	}

	if height == 0 {
		height = 320
	}
	// ------------------------------------------------------------------------

	xyscale := float64(width / 2 / xyrange) // pixels per x or y unit
	zscale := float64(height * 0.4)         // pixels per z unit

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%f' height='%f'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(xyscale, zscale, width, height, i+1, j)
			bx, by := corner(xyscale, zscale, width, height, i, j)
			cx, cy := corner(xyscale, zscale, width, height, i, j+1)
			dx, dy := corner(xyscale, zscale, width, height, i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(xyscale, zscale, width, height float64, i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
