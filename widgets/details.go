package widgets

import "github.com/veandco/go-sdl2/sdl"

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

func (receiver *HtmlTagDetails) Draw(renderer *sdl.Renderer) {

}

func (receiver *HtmlTagDetails) Render(renderer *sdl.Renderer) {

}

func SetWidgetPropertiesForDetailsTag() WidgetInterface {
	widget := new(HtmlTagDetails)
	return widget
}
