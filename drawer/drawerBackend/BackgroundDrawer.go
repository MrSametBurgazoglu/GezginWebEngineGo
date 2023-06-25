package drawerBackend

import (
	"gezgin_web_engine/widgets"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawBackground(widget widgets.WidgetInterface, renderer *sdl.Renderer) {
	if widget.GetStyleProperty().Background.BackgroundColor != nil {
		alpha, red, green, blue := widget.GetStyleProperty().Background.BackgroundColor.GetColorByRGBA()
		renderer.SetDrawColor(red, green, blue, alpha)
		renderer.FillRect(&widget.GetDrawProperties().Rect)
	}
}
