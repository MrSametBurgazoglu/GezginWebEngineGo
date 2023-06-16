package tags

import (
	"gezgin_web_engine/htmlParser/widget"
	"strconv"
)

type HtmlTagCanvas struct {
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

func SetWidgetPropertiesForCanvasTag(widget *widget.Widget) {
	widget.WidgetProperties = new(HtmlTagCanvas)
}
