package structs

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//maybe we can make a struct for every widget

type DrawProperties struct {
	font              *ttf.Font
	rect              sdl.Rect
	texture           *sdl.Texture
	backgroundTexture *sdl.Texture
}
