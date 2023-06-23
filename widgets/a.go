package widgets

type HtmlTagA struct {
	*Widget
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

func (receiver *HtmlTagA) Draw() {

}

func (receiver *HtmlTagA) Render() {

}

func SetWidgetPropertiesForATag() WidgetInterface {
	widget := new(HtmlTagA)
	return widget
}
