package javascript_interpreter

import (
	"gezgin_web_engine/html_scraper/htmlVariables"
	"gezgin_web_engine/html_scraper/tags"
	"gezgin_web_engine/html_scraper/widget"
	v8 "rogchap.com/v8go"
)

var currentIsolate *v8.Isolate
var globalDocument *widget.Widget

func getStyleElements(document *widget.Widget) []*widget.Widget {
	var elementList []*widget.Widget
	for _, child := range document.Children {
		if child.HtmlTag == htmlVariables.HTML_HTML {
			for _, htmlChild := range child.Children {
				if htmlChild.HtmlTag == htmlVariables.HTML_HEAD {
					for _, headChild := range htmlChild.Children {
						if headChild.HtmlTag == htmlVariables.HTML_SCRIPT {
							elementList = append(elementList, headChild)
						}
					}
					break
				}
			}
			break
		}
	}
	return elementList
}

func scrapeScriptElements(context *v8.Context, element *widget.Widget) {
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
	styleList := getStyleElements(document)
	for _, script := range styleList {
		scrapeScriptElements(ctx, script)
	}
}
