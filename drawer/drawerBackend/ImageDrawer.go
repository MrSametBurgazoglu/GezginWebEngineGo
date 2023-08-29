package drawerBackend

import (
	"gezgin_web_engine/LayoutEngine"
	"image"
	"image/draw"
	_ "image/png"
)

func GetImageTexture(imageSrc *image.Image, imageDest *image.RGBA, layoutProperty *LayoutEngine.LayoutProperty) {
	draw.Draw(imageDest, imageDest.Bounds(), *imageSrc, image.Point{X: 0, Y: 0}, draw.Src)
	layoutProperty.Width = imageDest.Rect.Max.X
	layoutProperty.Height = imageDest.Rect.Max.Y
}
