package structs

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//maybe we can make a struct for every widget

type DrawProperties struct {
	Font              *ttf.Font
	Rect              sdl.Rect
	Texture           *sdl.Texture
	BackgroundTexture *sdl.Texture
}
