package widgets

type HtmlTagDiv struct {
	*Widget
}

func (receiver *HtmlTagDiv) Draw() {

}

func (receiver *HtmlTagDiv) Render() {

}

func SetWidgetPropertiesForDivTag() WidgetInterface {
	widget := new(HtmlTagDiv)
	return widget
}
