package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	structs2 "gezgin_web_engine/StyleEngine/structs"
	"gezgin_web_engine/drawer/drawerBackend"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagP struct {
	Widget
}

func (receiver *HtmlTagP) Draw(renderer *sdl.Renderer) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, &receiver.DrawProperties.Rect, renderer)
	}
}

func (receiver *HtmlTagP) Render(renderer *sdl.Renderer, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForPTag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagP)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Font = new(structs2.Font)
	widget.StyleProperty.Font.FontSizeValue = 14
	widget.StyleProperty.Margin = new(structs2.Margin)
	widget.StyleProperty.Margin.MarginTop = 10
	return widget
}
