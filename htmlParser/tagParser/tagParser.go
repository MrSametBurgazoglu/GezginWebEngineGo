package tagParser

import (
	"gezgin_web_engine/cssParser"
	"gezgin_web_engine/htmlParser/htmlVariables"
	"gezgin_web_engine/htmlParser/widget"
	"gezgin_web_engine/javascript_interpreter"
	"gezgin_web_engine/utils"
	"strings"
	"sync"
)

type VarReaderInterface interface {
	VarReaderFunc(string, string)
}

type ContextReaderInterface interface {
	ContextReaderFunc(string)
}

func UntaggedTextClosed(widget *widget.Widget) {
	if widget.Parent.HtmlTag == htmlVariables.HTML_STYLE {
		cssParser.CreateCssPropertiesFromStyleTag(widget.Parent)
	} else if widget.Parent.HtmlTag == htmlVariables.HTML_SCRIPT {
		javascript_interpreter.ScriptElements = append(javascript_interpreter.ScriptElements, widget.Parent)
	}

}

func ParseParameters(widget *widget.Widget, parameters []string, group *sync.WaitGroup) {
	if len(parameters) > 1 {
		parameters = utils.MergeAttributes(parameters)
	}
	for _, s := range parameters[0:] {
		varName, varValue, found := strings.Cut(s, "=")
		if found {
			if isStandard := widget.SetStandardVariables(varName, varValue, widget); isStandard == false {
				var varReader, ok = widget.WidgetProperties.(VarReaderInterface)
				if ok {
					varReader.VarReaderFunc(varName, varValue)
				}
			}
		} else {
			if isStandard := widget.SetStandardContextVariables(s); isStandard == false {
				var contextReader, ok = widget.WidgetProperties.(ContextReaderInterface)
				if ok {
					contextReader.ContextReaderFunc(s)
				}
			}
		}
	}
	group.Done()
}

func ParseInsideOfTag(widget *widget.Widget, text string, group *sync.WaitGroup) bool {
	parameters := strings.Split(text, " ")
	result := htmlVariables.SetHtmlTag(parameters[0], widget)
	group.Add(1)
	go ParseParameters(widget, parameters, group)
	return result
}
