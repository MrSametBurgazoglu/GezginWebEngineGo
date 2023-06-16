package tags

import "gezgin_web_engine/htmlParser/widget"

type HtmlTagLabel struct {
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

func SetWidgetPropertiesForLabelTag(widget *widget.Widget) {
	widget.WidgetProperties = new(HtmlTagLabel)
}
