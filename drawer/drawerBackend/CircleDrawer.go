package drawerBackend

import (
	"golang.org/x/image/draw"
	"image/color"
)

func DrawCircle(img draw.Image, x0, y0, r int, c color.Color) {
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)

	for x > y {
		img.Set(x0+x, y0+y, c)
		img.Set(x0+y, y0+x, c)
		img.Set(x0-y, y0+x, c)
		img.Set(x0-x, y0+y, c)
		img.Set(x0-x, y0-y, c)
		img.Set(x0-y, y0-x, c)
		img.Set(x0+y, y0-x, c)
		img.Set(x0+x, y0-y, c)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
}

func DrawFilledCircle(img draw.Image, x0, y0, r int, c color.Color) {
	largestX := r
	centerX := x0
	centerY := y0
	for y := 0; y <= r; y++ {
		for x := largestX; x >= 0; x-- {
			if x*x+y*y <= r*r {
				DrawLine(img, centerX-x, centerX+x, centerY+y, centerY+y, c)
				DrawLine(img, centerX-x, centerX+x, centerY-y, centerY-y, c)
				largestX = x
				break
			}
		}
	}
}
