package drawerBackend

import (
	"image"
	"image/color"
	"image/draw"
)

func DrawBackground(red uint8, green uint8, blue uint8, alpha uint8, texture *image.RGBA) {
	bgColor := color.RGBA{R: red, G: green, B: blue, A: alpha}
	draw.Draw(texture, texture.Bounds(), &image.Uniform{C: bgColor}, image.Point{X: 0, Y: 0}, draw.Src)
}
