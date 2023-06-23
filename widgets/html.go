package widgets

type HtmlTagHtml struct {
	*Widget
}

func (receiver *HtmlTagHtml) Draw() {

}

func (receiver *HtmlTagHtml) Render() {

}

func SetWidgetPropertiesForHtmlTag() WidgetInterface {
	widget := new(HtmlTagHtml)
	return widget
}
