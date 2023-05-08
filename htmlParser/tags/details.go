package tags

import "gezgin_web_engine/htmlParser/widget"

type HtmlTagDetails struct {
	open bool
}

func (receiver *HtmlTagDetails) ContextReaderFunc(context string) {
	//receiver
	if context == "open" {
		receiver.open = true
	}
}

func SetWidgetPropertiesForDetailsTag(widget *widget.Widget) {
	widget.WidgetProperties = new(HtmlTagDetails)
}
