package DrawProperties

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/html_scraper/widget"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawHtmlFunction(widget *widget.Widget, renderer *sdl.Renderer) {
	if widget.CssProperties.Background != nil {
		drawerBackend.DrawBackground(widget, renderer)
	}
}

func RenderHtmlFunction(widget *widget.Widget, renderer *sdl.Renderer) {
}
