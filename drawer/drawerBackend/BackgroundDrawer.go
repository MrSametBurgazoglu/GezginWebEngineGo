package drawerBackend

import (
	"gezgin_web_engine/html_scraper/htmlVariables"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawBackground(widget *htmlVariables.Widget, renderer *sdl.Renderer) {
	if widget.CssProperties.Background.BackgroundColor != nil {
		alpha, red, green, blue := widget.CssProperties.Background.BackgroundColor.GetColorByRGBA()
		renderer.SetDrawColor(red, green, blue, alpha)
		renderer.FillRect(&widget.DrawProperties.Rect)
	}
}
