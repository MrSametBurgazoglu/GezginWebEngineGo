package DrawProperties

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawDivFunction(widget *tags.Widget, renderer *sdl.Renderer) {
	if widget.CssProperties.Background != nil {
		drawerBackend.DrawBackground(widget, renderer)
	}
}

func RenderDivFunction(widget *tags.Widget, renderer *sdl.Renderer) {
}
