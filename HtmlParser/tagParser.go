package HtmlParser

import (
	"gezgin_web_engine/utils"
	"strings"
)

type VarReaderInterface interface {
	VarReaderFunc(string, string)
}

type ContextReaderInterface interface {
	ContextReaderFunc(string)
}

func ParseParameters(element *HtmlElement, parameters []string) {
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
}

func ParseInsideOfTag(element *HtmlElement, text string) (bool, bool, string) {
	parameters := strings.Fields(text)
	htmlTag, endTag, notParseInside := FindHtmlTag(parameters[0])
	element.HtmlTag = htmlTag
	//group.Add(1)
	ParseParameters(element, parameters)
	return endTag, notParseInside, parameters[0]
}
