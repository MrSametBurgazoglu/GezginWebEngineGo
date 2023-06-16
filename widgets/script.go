package tags

import "gezgin_web_engine/htmlParser/widget"

type HtmlTagScript struct {
	async          bool
	defer_         bool
	noModule       bool
	integrity      string
	src            string
	type_          string
	crossOrigin    CrossOriginType
	referrerPolicy ReferrerPolicyType
}

func (receiver *HtmlTagScript) ContextReaderFunc(context string) {
	//receiver
	switch context {
	case "async":
		receiver.async = true
	case "defer":
		receiver.defer_ = true
	}
}

func (receiver *HtmlTagScript) VarReaderFunc(variableName string, variableValue string) {
	//receiver
	switch variableName {
	case "crossorigin":
		receiver.crossOrigin.Set(variableValue)
	case "integrity":
		receiver.integrity = variableValue
	case "nomodule":
		if variableValue == "True" {
			receiver.noModule = true
		} else if variableValue == "False" {
			receiver.noModule = false
		}
	case "referrerpolicy":
		receiver.referrerPolicy.Set(variableValue)
	case "src":
		receiver.src = variableValue
	case "type":
		receiver.type_ = variableValue
	}
}

func SetWidgetPropertiesForScriptTag(widget *widget.Widget) {
	widget.WidgetProperties = new(HtmlTagLink)
}
