package tags

import (
	"gezgin_web_engine/html_scraper/widget"
)

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

func SetWidgetPropertiesForStyleTag(widget *widget.Widget) {
	widget.WidgetProperties = new(HtmlTagStyle)
	widget.HaveAttrAsVar = true
}
