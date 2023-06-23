package drawerBackend

import (
	"gezgin_web_engine/drawer/ScreenProperties"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawBody(widget *tags.Widget, renderer *sdl.Renderer) {
	if widget.CssProperties.Background.BackgroundColor != nil {
		alpha, red, green, blue := widget.CssProperties.Background.BackgroundColor.GetColorByRGBA()
		bodyRect := sdl.Rect{X: 0, Y: 0, W: int32(ScreenProperties.WindowWidth), H: int32(ScreenProperties.WindowHeight)}
		renderer.SetDrawColor(red, green, blue, alpha)
		renderer.FillRect(&bodyRect)
	}
}
