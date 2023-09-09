package drawerBackend

import (
	"gezgin_web_engine/LayoutProperty"
	"image"
	"image/draw"
	_ "image/png"
)

func GetImageTexture(imageSrc *image.Image, imageDest *image.RGBA, layoutProperty *LayoutProperty.LayoutProperty) {
	draw.Draw(imageDest, imageDest.Bounds(), *imageSrc, image.Point{X: 0, Y: 0}, draw.Src)
	layoutProperty.Width = imageDest.Rect.Max.X
	layoutProperty.Height = imageDest.Rect.Max.Y
}
