package drawerBackend

import (
	"gezgin_web_engine/StyleEngine/structs"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func GetTextTexture(renderer *sdl.Renderer, text string, color *structs.ColorRGBA, font *ttf.Font, texture **sdl.Texture, rect *sdl.Rect) {
	alpha, red, green, blue := color.GetColorByRGBA()
	textColor := sdl.Color{R: red, G: green, B: blue, A: alpha}
	surface, err := font.RenderUTF8BlendedWrapped(text, textColor, 0)
	if err != nil {
		return
	}
	*texture, err = renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return
	}
	_, _, width, height, _ := (*texture).Query()
	surface.Free()
	rect.X = 0
	rect.Y = 0
	rect.W = width
	rect.H = height
}
