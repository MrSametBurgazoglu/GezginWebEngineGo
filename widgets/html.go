package widgets

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagHtml struct {
	*Widget
}

func (receiver *HtmlTagHtml) Draw(renderer *sdl.Renderer) {
	if receiver.GetStyleProperty().Background != nil {
		drawerBackend.DrawBackground(receiver, renderer)
	}
}

func (receiver *HtmlTagHtml) Render(renderer *sdl.Renderer) {

}

func SetWidgetPropertiesForHtmlTag() WidgetInterface {
	widget := new(HtmlTagHtml)
	return widget
}
