package widgets

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagBody struct {
	*Widget
}

func (receiver *HtmlTagBody) Draw(renderer *sdl.Renderer) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, &receiver.DrawProperties.Rect, renderer)
	}
}

func (receiver *HtmlTagBody) Render(renderer *sdl.Renderer) {

}

func SetWidgetPropertiesForBodyTag() WidgetInterface {
	widget := new(HtmlTagBody)
	return widget
}
