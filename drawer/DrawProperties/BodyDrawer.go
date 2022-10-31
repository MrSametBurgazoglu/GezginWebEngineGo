package DrawProperties

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawBodyFunction(widget *htmlVariables.Widget, renderer *sdl.Renderer) {
	if widget.CssProperties.Background != nil {
		drawerBackend.DrawBackground(widget, renderer)
	}
	for _, child := range widget.Children {
		child.DrawWidget(child, renderer)
	}
}

func RenderBodyFunction(widget *htmlVariables.Widget, renderer *sdl.Renderer) {
	for _, child := range widget.Children {
		child.RenderWidget(child, renderer)
	}
}
