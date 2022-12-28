package tags

import (
	structs2 "gezgin_web_engine/drawer/structs"
	"gezgin_web_engine/html_scraper/widget"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlDocument struct {
	*widget.BaseWidget
	*structs2.DrawProperties
}

func (receiver *HtmlDocument) RenderWidget(renderer *sdl.Renderer) {
	return
}

func (receiver *HtmlDocument) DrawWidget(renderer *sdl.Renderer) {
	renderer.Copy(receiver.DrawProperties.Texture, nil, &receiver.DrawProperties.Rect)
}
