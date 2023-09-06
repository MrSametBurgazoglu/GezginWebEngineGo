package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/StyleProperty/structs"
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/widget"
	"image"
)

type HtmlTagHR struct {
	widget.Widget
}

func (receiver *HtmlTagHR) Draw(mainImage *image.RGBA) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, mainImage, receiver.LayoutProperty)
	}
}

func (receiver *HtmlTagHR) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForHRTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) widget.WidgetInterface {
	widget := new(HtmlTagHR)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_BLOCK
	widget.StyleProperty.Margin = new(structs.Margin)
	widget.StyleProperty.Margin.MarginTop = 1
	widget.StyleProperty.Margin.MarginTopValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
	widget.StyleProperty.Margin.MarginBottom = 1
	widget.StyleProperty.Margin.MarginBottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
	return widget
}
