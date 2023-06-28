package drawerBackend

import (
	"github.com/veandco/go-sdl2/sdl"
)

func DrawBackground(red uint8, green uint8, blue uint8, alpha uint8, rect *sdl.Rect, renderer *sdl.Renderer) {
	renderer.SetDrawColor(red, green, blue, alpha)
	renderer.FillRect(rect)
}
