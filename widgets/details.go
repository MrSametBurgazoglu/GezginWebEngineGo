package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"github.com/veandco/go-sdl2/sdl"
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

func (receiver *HtmlTagDetails) Draw(renderer *sdl.Renderer) {

}

func (receiver *HtmlTagDetails) Render(renderer *sdl.Renderer, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForDetailsTag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagDetails)
	widget.HtmlElement = element
	return widget
}
