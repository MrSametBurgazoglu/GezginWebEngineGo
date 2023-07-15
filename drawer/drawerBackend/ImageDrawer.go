package drawerBackend

import (
	"image"
	"image/draw"
	_ "image/png"
)

func GetImageTexture(imageSrc *image.Image, imageDest *image.RGBA, rect *image.Rectangle) {
	draw.Draw(imageDest, *rect, *imageSrc, image.Point{X: 0, Y: 0}, 0)
}
