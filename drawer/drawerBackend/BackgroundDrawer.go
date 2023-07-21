package drawerBackend

import (
	"gezgin_web_engine/drawer/structs"
	"image"
	"image/color"
	"image/draw"
)

func DrawBackground(red uint8, green uint8, blue uint8, alpha uint8, texture *image.RGBA, drawProperties *structs.DrawProperties) {
	bgColor := color.RGBA{R: red, G: green, B: blue, A: alpha}
	draw.Draw(texture, image.Rect(int(drawProperties.X), int(drawProperties.Y), int(drawProperties.X+drawProperties.W), int(drawProperties.Y+drawProperties.H)), &image.Uniform{C: bgColor}, image.Point{X: 0, Y: 0}, draw.Src)
}
