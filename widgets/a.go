package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleEngine/enums"
	"image"
)

type HtmlTagA struct {
	Widget
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

}

func (receiver *HtmlTagA) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForATag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagA)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_INLINE
	return widget
}
