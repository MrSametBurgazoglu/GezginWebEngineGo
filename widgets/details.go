package widgets

type HtmlTagDetails struct {
	*Widget
	open bool
}

func (receiver *HtmlTagDetails) ContextReaderFunc(context string) {
	//receiver
	if context == "open" {
		receiver.open = true
	}
}

func (receiver *HtmlTagDetails) Draw() {

}

func (receiver *HtmlTagDetails) Render() {

}

func SetWidgetPropertiesForDetailsTag() WidgetInterface {
	widget := new(HtmlTagDetails)
	return widget
}
