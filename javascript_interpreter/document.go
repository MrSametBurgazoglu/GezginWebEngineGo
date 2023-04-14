package javascript_interpreter

import v8 "rogchap.com/v8go"

func setDocumentTemplate(iso *v8.Isolate, context *v8.Context, template *v8.ObjectTemplate) {
	getElementByID := v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		elementObj, _ := Element.NewInstance(context)
		elementObj.Set("filter_value", info.Args()[0])
		elementObj.Set("filter_type", "id")
		return elementObj.Value
	})
	template.Set("getElementById", getElementByID)
}
