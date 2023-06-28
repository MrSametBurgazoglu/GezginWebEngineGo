package widgets

import (
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type HtmlTagCanvas struct {
	Widget
	width  int
	height int
}

func (receiver *HtmlTagCanvas) VarReaderFunc(variableName string, variableValue string) {
	//receiver
	switch variableName {
	case "width":
		receiver.width, _ = strconv.Atoi(variableValue)
	case "height":
		receiver.height, _ = strconv.Atoi(variableValue)
	}
}

func (receiver *HtmlTagCanvas) Draw(renderer *sdl.Renderer) {

}

func (receiver *HtmlTagCanvas) Render(renderer *sdl.Renderer) {

}

func SetWidgetPropertiesForCanvasTag() WidgetInterface {
	widget := new(HtmlTagCanvas)
	return widget
}
