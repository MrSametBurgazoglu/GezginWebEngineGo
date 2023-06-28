package widgets

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagHtml struct {
	Widget
}

func (receiver *HtmlTagHtml) Draw(renderer *sdl.Renderer) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, &receiver.DrawProperties.Rect, renderer)
	}
}

func (receiver *HtmlTagHtml) Render(renderer *sdl.Renderer) {

}

func SetWidgetPropertiesForHtmlTag() WidgetInterface {
	widget := new(HtmlTagHtml)
	return widget
}
