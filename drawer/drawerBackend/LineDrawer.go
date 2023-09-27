package drawerBackend

import (
	"image/color"
	"image/draw"
)

func DrawLine(image draw.Image, x0, x1, y0, y1 int, c color.Color) {
	if x1 == x0 {
		return
	}
	eq := float64(y1-y0) / float64(x1-x0)
	j := float64(y0)
	for i := x0; i < x1; i++ {
		image.Set(i, int(j), c)
		j += eq
	}
}
