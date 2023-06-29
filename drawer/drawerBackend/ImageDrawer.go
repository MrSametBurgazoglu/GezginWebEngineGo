package drawerBackend

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func GetImageTexture(renderer *sdl.Renderer, imageSrc *sdl.RWops, texture **sdl.Texture, rect *sdl.Rect) {
	image, err := img.LoadTextureRW(renderer, imageSrc, true)
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
