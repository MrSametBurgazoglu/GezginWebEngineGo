package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
	"image"
)

type HtmlTagButton struct {
	widget.Widget
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

func (receiver *HtmlTagButton) Draw(mainImage *image.RGBA) {

}

func (receiver *HtmlTagButton) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForButtonTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) widget.WidgetInterface {
	widget := new(HtmlTagButton)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_INLINE
	return widget
}
