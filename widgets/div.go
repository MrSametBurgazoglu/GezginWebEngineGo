package widgets

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagDiv struct {
	Widget
}

func (receiver *HtmlTagDiv) Draw(renderer *sdl.Renderer) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, &receiver.DrawProperties.Rect, renderer)
	}
}

func (receiver *HtmlTagDiv) Render(renderer *sdl.Renderer) {

}

func SetWidgetPropertiesForDivTag() WidgetInterface {
	widget := new(HtmlTagDiv)
	widget.Initialize()
	return widget
}
