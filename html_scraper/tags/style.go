package tags

import (
	"gezgin_web_engine/html_scraper/widget"
)

type HtmlTagStyle struct {
	media string
	Type  string
}

func ChooseVariableForStyleTag(widget *widget.Widget, varName string, varValue string) {
	styleTag, ok := widget.WidgetProperties.(HtmlTagStyle)
	if ok {
		switch varName {
		case "media":
			styleTag.media = varValue
		case "type":
			styleTag.Type = varValue
		}
	}
}

func SetWidgetPropertiesForStyleTag(widget *widget.Widget) {
	widget.WidgetProperties = HtmlTagStyle{}
	widget.VarReaderFunc = ChooseVariableForStyleTag
}
