package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleEngine/enums"
	"gezgin_web_engine/drawer/drawerBackend"
	"image"
)

type HtmlTagFooter struct {
	Widget
}

func (receiver *HtmlTagFooter) Draw(mainImage *image.RGBA) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, mainImage, receiver.DrawProperties)
	}
}

func (receiver *HtmlTagFooter) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForFooterTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagFooter)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_BLOCK
	return widget
}
