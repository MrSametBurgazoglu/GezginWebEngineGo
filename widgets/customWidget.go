package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
	"image"
)

type CustomWidget struct {
	widget.Widget
}

func (receiver *CustomWidget) Draw(mainImage *image.RGBA) {
	receiver.DrawBackground(mainImage)
}

func (receiver *CustomWidget) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetCustomWidgetProperties(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) widget.WidgetInterface {
	widget := new(CustomWidget)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_BLOCK
	return widget
}
