package javascript_interpreter

import (
	v8 "rogchap.com/v8go"
	"strings"
	"time"
)

func executeAfter(functionName string, timeMs int, ctx *v8.Context) {
	time.Sleep(time.Duration(timeMs) * time.Millisecond)
	//time.Sleep(10000 * time.Millisecond)
	_, err := ctx.RunScript(functionName+"()", "script.js")
	if err != nil {
		println("script error", err.Error())
		return
	}
}

func createTimeoutFunc(iso *v8.Isolate, global *v8.ObjectTemplate) {
	setTimeout := v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		functionString := info.Args()[0].String()
		functionStringNameStart := strings.Index(functionString, " ")
		functionStringNameEnd := strings.Index(functionString, "(")
		functionName := functionString[functionStringNameStart+1 : functionStringNameEnd]
		go executeAfter(functionName, int(info.Args()[1].Integer()), info.Context())
		return nil // you can return a value back to the JS caller if required
	})
	err := global.Set("setTimeout", setTimeout)
	if err != nil {
		println(err.Error(), "setTimeout not implemted")
	}
}
