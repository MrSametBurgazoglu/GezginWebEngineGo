package drawerBackend

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func GetImageTexture(renderer *sdl.Renderer, imagePath string, texture **sdl.Texture, rect *sdl.Rect) {
	image, err := img.LoadTexture(renderer, "exampleHtmlFiles/"+imagePath)
	if err != nil {
		println(err.Error())
	}
	_, _, w, h, imageQueryErr := image.Query()
	if imageQueryErr != nil {
		println(err.Error())
	}
	rect.X = 0
	rect.Y = 0
	rect.W = w
	rect.H = h

	*texture = image
}
