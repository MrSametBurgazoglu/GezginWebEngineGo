package tags

import "gezgin_web_engine/html_scraper/widget"

type HtmlTagButton struct {
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
		switch variableValue {
		case "button":
			receiver.buttonType = BUTTON_TYPE_BUTTON
		case "reset":
			receiver.buttonType = BUTTON_TYPE_RESET
		case "submit":
			receiver.buttonType = BUTTON_TYPE_SUBMIT
		}
	case "form":
		receiver.form = variableValue
	case "formaction":
		receiver.formAction = variableValue
	case "formenctype":
		switch variableValue {
		case "text/plain":
			receiver.formEncType = FORM_ENC_TYPE_TEXT
		case "multipart/form-data":
			receiver.formEncType = FORM_ENC_TYPE_MULTIPART
		case "application/x-www-form-urlencoded":
			receiver.formEncType = FORM_ENC_TYPE_APPLICATION
		}
	case "formmethod":
		switch variableValue {
		case "get":
			receiver.formMethod = FORM_METHOD_GET
		case "post":
			receiver.formMethod = FORM_METHOD_POST
		}
	case "formtarget":
		switch variableValue {
		case "_blank":
			receiver.formTarget = FORM_TARGET_BLANK
		case "_self":
			receiver.formTarget = FORM_TARGET_SELF
		case "_parent":
			receiver.formTarget = FORM_TARGET_PARENT
		case "_top":
			receiver.formTarget = FORM_TARGET_TOP
		default:
			receiver.formTarget = FORM_TARGET_CUSTOM
		}
	case "name":
		receiver.name = variableValue
	case "value":
		receiver.value = variableValue
	}
}

func SetWidgetPropertiesForButtonTag(widget *widget.Widget) {
	widget.WidgetProperties = new(HtmlTagButton)
	widget.HaveAttrAsVar = true
	widget.HaveAttrAsContext = true
}
