package tagScraper

import (
	"gezgin_web_engine/css_scraper/tree"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"gezgin_web_engine/html_scraper/widget"
	"strings"
)

func ScrapeInsideOfTag(widget *widget.Widget, text string) bool {
	parameters := strings.Split(text, " ")
	result := htmlVariables.SetHtmlTag(parameters[0], widget)
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
