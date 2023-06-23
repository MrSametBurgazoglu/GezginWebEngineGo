package widgets

type HtmlTagStyle struct {
	media string
	Type  string
}

func (receiver *HtmlTagStyle) VarReaderFunc(variableName string, variableValue string) {
	switch variableName {
	case "media":
		receiver.media = variableValue
	case "type":
		receiver.Type = variableValue
	}
}

func SetWidgetPropertiesForStyleTag(widget *Widget) {
	widget.WidgetProperties = new(HtmlTagStyle)
}
