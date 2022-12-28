package tags

import (
	"gezgin_web_engine/drawer/drawerBackend"
	structs2 "gezgin_web_engine/drawer/structs"
	"gezgin_web_engine/html_scraper/widget"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlUntaggedText struct {
	*widget.BaseWidget
	Value string
	*structs2.DrawProperties
}

func (receiver *HtmlUntaggedText) ReadParameters(found bool, items ...string) {
	return
}

func (receiver *HtmlUntaggedText) RenderWidget(renderer *sdl.Renderer) {
	parent := receiver.GetParent().(structs2.DrawableWidget)
	drawerBackend.GetTextTexture(
		renderer,
		receiver.Value,
		parent.GetCssProperties().Color,
		parent.GetFont(),
		&receiver.DrawProperties.Texture,
		&receiver.DrawProperties.Rect,
	)
}

func (receiver *HtmlUntaggedText) DrawWidget(renderer *sdl.Renderer) {
	renderer.Copy(receiver.DrawProperties.Texture, nil, &receiver.DrawProperties.Rect)
}
