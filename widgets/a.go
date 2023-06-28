package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagA struct {
	Widget
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

func (receiver *HtmlTagA) Draw(renderer *sdl.Renderer) {

}

func (receiver *HtmlTagA) Render(renderer *sdl.Renderer) {

}

func SetWidgetPropertiesForATag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagA)
	widget.HtmlElement = element
	return widget
}
