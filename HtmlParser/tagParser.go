package HtmlParser

import (
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

func ParseParameters(element *HtmlElement, parameters []string, group *sync.WaitGroup) {
	if len(parameters) > 1 {
		parameters = utils.MergeAttributes(parameters)
	}
	for _, s := range parameters[1:] {
		varName, varValue, found := strings.Cut(s, "=")
		if found {
			element.Attributes[varName] = varValue[1 : len(varValue)-1]
		} else {
			element.Attributes[varName] = ""
		}
	}
	group.Done()
}

func ParseInsideOfTag(element *HtmlElement, text string, group *sync.WaitGroup) bool {
	parameters := strings.Split(text, " ")
	htmlTag, endTag := FindHtmlTag(parameters[0])
	element.HtmlTag = htmlTag
	group.Add(1)
	go ParseParameters(element, parameters, group)
	return endTag
}
