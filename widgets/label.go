package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagLabel struct {
	Widget
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

func (receiver *HtmlTagLabel) Draw(renderer *sdl.Renderer) {

}

func (receiver *HtmlTagLabel) Render(renderer *sdl.Renderer) {

}

func SetWidgetPropertiesForLabelTag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagLabel)
	widget.HtmlElement = element
	widget.Initialize()
	return widget
}
