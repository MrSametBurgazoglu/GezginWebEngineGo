package DrawProperties

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/htmlParser/widget"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawHeaderFunction(widget *widget.Widget, renderer *sdl.Renderer) {
	if widget.CssProperties.Background != nil {
		drawerBackend.DrawBackground(widget, renderer)
	}
}

func RenderHeaderFunction(widget *widget.Widget, renderer *sdl.Renderer) {
}
