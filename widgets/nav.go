package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleEngine/enums"
	"gezgin_web_engine/drawer/drawerBackend"
	"image"
)

type HtmlTagNav struct {
	Widget
}

func (receiver *HtmlTagNav) Draw(mainImage *image.RGBA) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, mainImage, receiver.DrawProperties)
	}
}

func (receiver *HtmlTagNav) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForNavTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagNav)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_BLOCK
	return widget
}
