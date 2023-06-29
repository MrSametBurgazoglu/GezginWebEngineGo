package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
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

func (receiver *HtmlTagDiv) Render(renderer *sdl.Renderer, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForDivTag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagDiv)
	widget.HtmlElement = element
	widget.Initialize()
	return widget
}
