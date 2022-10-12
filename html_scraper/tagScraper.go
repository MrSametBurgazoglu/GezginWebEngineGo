package html_scraper

import (
	"gezgin_web_engine/css_scraper/tree"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"strings"
)

func ScrapeInsideOfTag(widget *htmlVariables.Widget, text string) bool {
	parameters := strings.Split(text, " ")
	result := widget.HtmlTag.SetHtmlTag(parameters[0], widget)
	if widget.HtmlTag == htmlVariables.HTML_STYLE {
		tree.CssStyleTagList = append(tree.CssStyleTagList, widget)
	}
	for _, s := range parameters[0:] {
		varName, varValue, found := strings.Cut(s, "=")
		if found {
			if isStandard := widget.StandardHtmlVariables.SetStandardVariables(varName, varValue); isStandard == false && widget.ContextReaderFunc != nil {
				widget.VarReaderFunc(widget, varName, varValue)
			}
		} else {
			if isStandard := widget.StandardHtmlVariables.SetStandardContextVariables(s); isStandard == false && widget.ContextReaderFunc != nil {
				widget.ContextReaderFunc(widget, s)
			}
		}
	}
	return result
}
