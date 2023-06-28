package widgets

import (
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type HtmlTagTextArea struct {
	autofocus   bool
	disabled    bool
	readonly    bool
	required    bool
	dirname     string
	form        string
	name        string
	placeholder string
	rows        int
	cols        int
	maxlength   int
	wrap        TextAreaWrap
}

func (receiver *HtmlTagTextArea) ContextReaderFunc(context string) {
	//receiver
	switch context {
	case "autofocus":
		receiver.autofocus = true
	case "disabled":
		receiver.disabled = true
	case "readonly":
		receiver.readonly = true
	case "required":
		receiver.required = true
	}
}

func (receiver *HtmlTagTextArea) VarReaderFunc(variableName string, variableValue string) {
	//receiver
	switch variableName {
	case "col":
		receiver.cols, _ = strconv.Atoi(variableValue)
	case "dirname":
		receiver.form = variableValue
	case "form":
		receiver.form = variableValue
	case "maxlength":
		receiver.maxlength, _ = strconv.Atoi(variableValue)
	case "name":
		receiver.name = variableValue
	case "placeholder":
		receiver.placeholder = variableValue
	case "rows":
		receiver.rows, _ = strconv.Atoi(variableValue)
	case "wrap":
		receiver.wrap.Set(variableValue)
	}
}

func (receiver *HtmlTagTextArea) Draw(renderer *sdl.Renderer) {

}

func (receiver *HtmlTagTextArea) Render(renderer *sdl.Renderer) {

}

func SetWidgetPropertiesForTextAreaTag() *HtmlTagTextArea {
	widget := new(HtmlTagTextArea)
	return widget
}
