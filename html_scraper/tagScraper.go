package html_scraper

import "strings"

func ScrapeInsideOfTag(widget *Widget, text string) bool {
	parameters := strings.Split(text, " ")
	result := widget.HtmlTag.setHtmlTag(parameters[0], widget)
	for _, s := range parameters[0:] {
		varName, varValue, found := strings.Cut(s, "=")
		if found {
			if isStandard := widget.StandardHtmlVariables.setStandardVariables(varName, varValue); isStandard == false && widget.ContextReaderFunc != nil {
				widget.VarReaderFunc(widget, varName, varValue)
			}
		} else {
			if isStandard := widget.StandardHtmlVariables.setStandardContextVariables(s); isStandard == false && widget.ContextReaderFunc != nil {
				widget.ContextReaderFunc(widget, s)
			}
		}
	}
	return result
}
