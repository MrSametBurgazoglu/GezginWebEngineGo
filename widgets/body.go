package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/drawer/drawerBackend"
	"image"
)

type HtmlTagBody struct {
	Widget
}

func (receiver *HtmlTagBody) Draw(mainImage *image.RGBA) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, mainImage, receiver.DrawProperties)
	}
}

func (receiver *HtmlTagBody) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForBodyTag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagBody)
	widget.HtmlElement = element
	widget.Initialize()
	return widget
}
