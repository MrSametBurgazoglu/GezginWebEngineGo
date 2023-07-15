package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/drawer/drawerBackend"
	"image"
)

type HtmlTagDiv struct {
	Widget
}

func (receiver *HtmlTagDiv) Draw(mainImage *image.RGBA) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, mainImage)
	}
}

func (receiver *HtmlTagDiv) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForDivTag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagDiv)
	widget.HtmlElement = element
	widget.Initialize()
	return widget
}
