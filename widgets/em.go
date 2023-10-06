package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
	"image"
)

type HtmlTagEM struct {
	widget.Widget
}

func (receiver *HtmlTagEM) Draw(mainImage *image.RGBA) {
	receiver.DrawBackground(mainImage)
}

func (receiver *HtmlTagEM) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForEMTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) widget.WidgetInterface {
	widget := new(HtmlTagEM)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_INLINE
	return widget
}
