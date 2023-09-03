package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleEngine/enums"
	"gezgin_web_engine/drawer/drawerBackend"
	"image"
)

type HtmlTagSVG struct {
	Widget
}

func (receiver *HtmlTagSVG) Draw(mainImage *image.RGBA) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, mainImage, receiver.DrawProperties)
	}
}

func (receiver *HtmlTagSVG) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForSVGTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagSVG)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_INLINE
	return widget
}
