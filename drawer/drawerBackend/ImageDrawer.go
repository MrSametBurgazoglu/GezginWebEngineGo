package drawerBackend

import (
	"gezgin_web_engine/drawer/structs"
	"image"
	"image/draw"
	_ "image/png"
)

func GetImageTexture(imageSrc *image.Image, imageDest *image.RGBA, drawProperties *structs.DrawProperties) {
	draw.Draw(imageDest, imageDest.Bounds(), *imageSrc, image.Point{X: 0, Y: 0}, draw.Src)
	drawProperties.W = int32(imageDest.Rect.Max.X)
	drawProperties.H = int32(imageDest.Rect.Max.Y)
}
