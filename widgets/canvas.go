package widgets

import (
	"strconv"
)

type HtmlTagCanvas struct {
	*Widget
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

func (receiver *HtmlTagCanvas) Draw() {

}

func (receiver *HtmlTagCanvas) Render() {

}

func SetWidgetPropertiesForCanvasTag() WidgetInterface {
	widget := new(HtmlTagCanvas)
	return widget
}
