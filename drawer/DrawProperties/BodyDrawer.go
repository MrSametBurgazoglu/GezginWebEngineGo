package DrawProperties

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawBodyFunction(widget *tags.Widget, renderer *sdl.Renderer) {
	if widget.CssProperties.Background != nil {
		drawerBackend.DrawBody(widget, renderer)
	}
}

func RenderBodyFunction(widget *tags.Widget, renderer *sdl.Renderer) {
}
