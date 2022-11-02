package DrawProperties

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/html_scraper/widget"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawBodyFunction(widget *widget.Widget, renderer *sdl.Renderer) {
	if widget.CssProperties.Background != nil {
		drawerBackend.DrawBackground(widget, renderer)
	}
}

func RenderBodyFunction(widget *widget.Widget, renderer *sdl.Renderer) {
}