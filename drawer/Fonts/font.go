package Fonts

import "github.com/veandco/go-sdl2/ttf"

var DefaultFont *ttf.Font

func InitializeFont() {
	if font, err := ttf.OpenFont("fonts/Sans.ttf", 14); err != nil {
		panic(err)
	} else {
		DefaultFont = font
	}
}
