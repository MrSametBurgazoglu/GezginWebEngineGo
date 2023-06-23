package DrawProperties

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawPFunction(widget *tags.Widget, renderer *sdl.Renderer) {
	if widget.CssProperties.Background != nil {
		drawerBackend.DrawBackground(widget, renderer)
	}
}

func RenderPFunction(widget *tags.Widget, renderer *sdl.Renderer) {
}
