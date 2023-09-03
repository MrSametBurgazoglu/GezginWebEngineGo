package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleEngine/enums"
	"gezgin_web_engine/drawer/drawerBackend"
	"image"
)

type HtmlTagStrong struct {
	Widget
}

func (receiver *HtmlTagStrong) Draw(mainImage *image.RGBA) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, mainImage, receiver.DrawProperties)
	}
}

func (receiver *HtmlTagStrong) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForStrongTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagStrong)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_INLINE
	return widget
}
