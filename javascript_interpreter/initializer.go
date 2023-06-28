package javascript_interpreter

import (
	"gezgin_web_engine/htmlParser/tags"
	"gezgin_web_engine/htmlParser/widget"
	v8 "rogchap.com/v8go"
)

var ScriptElements []*widget.Widget

var globalDocument *widget.Widget

func parseScriptElements(context *v8.Context, element *widget.Widget) {
	scriptWidget, ok := element.Children[0].WidgetProperties.(tags.UntaggedText)
	if ok {
		_, err := context.RunScript(scriptWidget.Value, "script.js")
		if err != nil {
			println(err.Error(), "script error")
			return
		}
	}
}

func InitializeJSInterpreter(document *widget.Widget) {
	globalDocument = document
	iso := v8.NewIsolate()
	//defer iso.Dispose()
	global := v8.NewObjectTemplate(iso)
	createTimeoutFunc(iso, global)
	ctx := v8.NewContext(iso, global)
	documentTemplate := v8.NewObjectTemplate(iso)
	setDocumentTemplate(iso, ctx, documentTemplate)
	documentObj, _ := documentTemplate.NewInstance(ctx)
	ctx.Global().Set("document", documentObj)
	CreateElementTemplate(iso)
	for _, script := range ScriptElements {
		parseScriptElements(ctx, script)
	}
}