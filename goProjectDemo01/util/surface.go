package util

import (
	"bytes"
	"fmt"
	"math"
)

const (
	width, height = 1000, 500           // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func GetPicture(line string, ground string, w int, h int) []byte {
	if w == 0 {
		w = width
	}
	if h == 0 {
		h = height
	}
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "<!DocTYPE html>\n"+
		"<html lang = 'en'>\n"+
		"<head>\n"+
		"<meta charset='UTF-8'>\n"+
		"<meta name='viewport' content='width-device-width, initial-scale-1.0'>\n"+
		"<title>Text</title>\n"+
		"</head>\n"+
		"<body>\n")

	fmt.Fprintf(buffer, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", line, ground, w, h)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			var b bytes.Buffer
			if !math.IsNaN(ax) && !math.IsNaN(ay) {
				b.WriteString(fmt.Sprintf("%g,%g ", ax, ay))
			}
			bx, by := corner(i, j)
			if !math.IsNaN(bx) && !math.IsNaN(by) {
				b.WriteString(fmt.Sprintf("%g,%g ", bx, by))
			}
			cx, cy := corner(i, j+1)
			if !math.IsNaN(cx) && !math.IsNaN(cy) {
				b.WriteString(fmt.Sprintf("%g,%g ", cx, cy))
			}
			dx, dy := corner(i+1, j+1)
			if !math.IsNaN(dx) && !math.IsNaN(dy) {
				b.WriteString(fmt.Sprintf("%g,%g ", dx, dy))
			}

			fmt.Fprintf(buffer, "<polygon points='%s'/>\n",
				b.String())

		}
	}
	fmt.Fprintln(buffer, "</svg>\n")
	fmt.Fprintln(buffer, "</body>\n"+
		"</html>\n")
	return buffer.Bytes()
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if floatoverfloat(z) {
		return math.NaN(), math.NaN()
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func floatoverfloat(f float64) bool {
	return math.IsNaN(f) || f == math.MaxFloat64 || f == math.SmallestNonzeroFloat64
}
