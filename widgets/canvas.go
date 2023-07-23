package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"image"
	"strconv"
)

type HtmlTagCanvas struct {
	Widget
	width  int
	height int
}

func (receiver *HtmlTagCanvas) VarReaderFunc(variableName string, variableValue string) {
	//receiver
	switch variableName {
	case "width":
		receiver.width, _ = strconv.Atoi(variableValue)
	case "height":
		receiver.height, _ = strconv.Atoi(variableValue)
	}
}

func (receiver *HtmlTagCanvas) Draw(mainImage *image.RGBA) {

}

func (receiver *HtmlTagCanvas) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForCanvasTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagCanvas)
	widget.HtmlElement = element
	return widget
}
