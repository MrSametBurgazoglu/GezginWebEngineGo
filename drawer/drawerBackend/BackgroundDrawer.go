package drawerBackend

import (
	"gezgin_web_engine/htmlParser/widget"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawBackground(widget *widget.Widget, renderer *sdl.Renderer) {
	if widget.CssProperties.Background.BackgroundColor != nil {
		alpha, red, green, blue := widget.CssProperties.Background.BackgroundColor.GetColorByRGBA()
		renderer.SetDrawColor(red, green, blue, alpha)
		renderer.FillRect(&widget.DrawProperties.Rect)
	}
}
