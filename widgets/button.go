package widgets

import "github.com/veandco/go-sdl2/sdl"

type HtmlTagButton struct {
	Widget
	autoFocus      bool
	disabled       bool
	formNovalidate bool
	form           string
	formAction     string
	name           string
	value          string
	formEncType    FormEncType
	formMethod     FormMethodType
	formTarget     FormTargetType
	buttonType     ButtonType
}

func (receiver *HtmlTagButton) ContextReaderFunc(context string) {
	//receiver
	switch context {
	case "autofocus":
		receiver.autoFocus = true
	case "disabled":
		receiver.disabled = true
	case "formnovalidate":
		receiver.formNovalidate = true
	}
}

func (receiver *HtmlTagButton) VarReaderFunc(variableName string, variableValue string) {
	//receiver
	switch variableName {
	case "type":
		receiver.buttonType.Set(variableValue)
	case "form":
		receiver.form = variableValue
	case "formaction":
		receiver.formAction = variableValue
	case "formenctype":
		receiver.formEncType.Set(variableValue)
	case "formmethod":
		receiver.formMethod.Set(variableValue)
	case "formtarget":
		receiver.formTarget.Set(variableValue)
	case "name":
		receiver.name = variableValue
	case "value":
		receiver.value = variableValue
	}
}

func (receiver *HtmlTagButton) Draw(renderer *sdl.Renderer) {

}

func (receiver *HtmlTagButton) Render(renderer *sdl.Renderer) {

}

func SetWidgetPropertiesForButtonTag() WidgetInterface {
	widget := new(HtmlTagButton)
	return widget
}
