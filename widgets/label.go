package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/widget"
	"image"
)

type HtmlTagLabel struct {
	widget.Widget
	for_ string
	form string
}

func (receiver *HtmlTagLabel) VarReaderFunc(variableName string, variableValue string) {
	//receiver
	if variableName == "form" {
		receiver.form = variableValue
	} else if variableName == "for" {
		receiver.for_ = variableValue
	}
}

func (receiver *HtmlTagLabel) Draw(mainImage *image.RGBA) {

}

func (receiver *HtmlTagLabel) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForLabelTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) widget.WidgetInterface {
	widget := new(HtmlTagLabel)
	widget.HtmlElement = element
	widget.Initialize()
	return widget
}
