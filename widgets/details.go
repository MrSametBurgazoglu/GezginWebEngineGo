package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"image"
)

type HtmlTagDetails struct {
	Widget
	open bool
}

func (receiver *HtmlTagDetails) ContextReaderFunc(context string) {
	//receiver
	if context == "open" {
		receiver.open = true
	}
}

func (receiver *HtmlTagDetails) Draw(mainImage *image.RGBA) {

}

func (receiver *HtmlTagDetails) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForDetailsTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagDetails)
	widget.HtmlElement = element
	return widget
}
