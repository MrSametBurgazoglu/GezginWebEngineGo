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

type HtmlTagP struct {
	widget.Widget
}

func (receiver *HtmlTagP) Draw(mainImage *image.RGBA) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, mainImage, receiver.LayoutProperty)
	}
}

func (receiver *HtmlTagP) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForPTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) widget.WidgetInterface {
	widget := new(HtmlTagP)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_BLOCK
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 14
	widget.StyleProperty.Margin = new(structs.Margin)
	widget.StyleProperty.Margin.MarginTop = 10
	return widget
}
