package widgets

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagBody struct {
	*Widget
}

func (receiver *HtmlTagBody) Draw(renderer *sdl.Renderer) {
	if receiver.StyleProperty.Background != nil {
		drawerBackend.DrawBody(receiver, renderer)
	}
}

func (receiver *HtmlTagBody) Render(renderer *sdl.Renderer) {

}

func SetWidgetPropertiesForBodyTag() WidgetInterface {
	widget := new(HtmlTagBody)
	return widget
}
