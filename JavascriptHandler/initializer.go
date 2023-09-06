package JavascriptHandler

import (
	"gezgin_web_engine/widget"
	v8 "rogchap.com/v8go"
)

type JavascriptEngine struct {
	DocumentWidget widget.WidgetInterface
	ScriptTexts    []string
	GlobalContext  *v8.Context
}

func (receiver *JavascriptEngine) parseScript(scriptText string) {
	_, err := receiver.GlobalContext.RunScript(scriptText, "script.js")
	if err != nil {
		println(err.Error(), "script error")
		return
	}
}

func (receiver *JavascriptEngine) ExecuteScripts() {
	for _, text := range receiver.ScriptTexts {
		receiver.parseScript(text)
	}
}

func (receiver *JavascriptEngine) AppendScript(scriptText string) {
	receiver.ScriptTexts = append(receiver.ScriptTexts, scriptText)
}

func (receiver *JavascriptEngine) InitializeJSInterpreter() {
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
	receiver.GlobalContext = ctx

}
