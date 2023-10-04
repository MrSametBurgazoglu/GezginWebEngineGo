package drawerBackend

import (
	"gezgin_web_engine/LayoutProperty"
	"image"
	"image/color"
	"image/draw"
)

func DrawBackground(red uint8, green uint8, blue uint8, alpha uint8, texture *image.RGBA, layoutProperty *LayoutProperty.LayoutProperty) {
	if alpha == 0 {
		return
	}
	bgColor := color.RGBA{R: red, G: green, B: blue, A: 50}
	//println(bgColor.A, " color")
	//println("color:", red, green, blue, "pos", layoutProperty.XPosition, layoutProperty.YPosition, layoutProperty.Width, layoutProperty.Height)
	draw.Draw(texture, image.Rect(layoutProperty.XPosition, layoutProperty.YPosition, layoutProperty.XPosition+layoutProperty.ContentWidth, layoutProperty.YPosition+layoutProperty.ContentHeight), &image.Uniform{C: bgColor}, image.Point{X: 0, Y: 0}, draw.Src)
}
