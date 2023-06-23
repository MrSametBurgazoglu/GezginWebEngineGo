package widgets

type HtmlTagBody struct {
	*Widget
}

func (receiver *HtmlTagBody) Draw() {

}

func (receiver *HtmlTagBody) Render() {

}

func SetWidgetPropertiesForBodyTag() WidgetInterface {
	widget := new(HtmlTagBody)
	return widget
}
