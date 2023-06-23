package widgets

type HtmlTagLabel struct {
	*Widget
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

func (receiver *HtmlTagLabel) Draw() {

}

func (receiver *HtmlTagLabel) Render() {

}

func SetWidgetPropertiesForLabelTag() WidgetInterface {
	widget := new(HtmlTagLabel)
	return widget
}
