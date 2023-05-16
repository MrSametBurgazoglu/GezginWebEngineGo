package Fonts

import "github.com/veandco/go-sdl2/ttf"

var FontMap = make(map[int]*ttf.Font)

func GetFont(size int) *ttf.Font {
	if font, ok := FontMap[size]; ok {
		return font
	} else {
		if font, err := ttf.OpenFont("fonts/Sans.ttf", size); err != nil {
			panic(err)
		} else {
			FontMap[size] = font
			return font
		}
	}
	return nil
}
