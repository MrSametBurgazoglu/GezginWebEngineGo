package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/widget"
	"image"
)

type HtmlTagA struct {
	widget.Widget
	download string
	href     string
	hrefLang string
}

func (receiver *HtmlTagA) VarReaderFunc(variableName string, variableValue string) {
	//receiver
	switch variableName {
	case "download":
		receiver.download = variableValue
	case "href":
		receiver.href = variableValue
	case "hrefLang":
		receiver.hrefLang = variableValue
	}
}

func (receiver *HtmlTagA) Draw(mainImage *image.RGBA) {
	if receiver.GetStyleProperty().Background != nil {
		/*TODO WHY WE DRAW BACKGROUND IF IT SAME COLOR WITH PARENT*/
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, mainImage, receiver.LayoutProperty)
	}
}

func (receiver *HtmlTagA) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForATag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) widget.WidgetInterface {
	widget := new(HtmlTagA)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_INLINE
	return widget
}
