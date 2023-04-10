package tags

import "gezgin_web_engine/html_scraper/widget"

type HtmlTagA struct {
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

func SetWidgetPropertiesForATag(widget *widget.Widget) {
	widget.WidgetProperties = new(HtmlTagA)
	widget.HaveAttrAsVar = true
}
