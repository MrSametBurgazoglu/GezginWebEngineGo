package tagScraper

import (
	"gezgin_web_engine/css_scraper/tree"
	"gezgin_web_engine/html_scraper/HtmlElementWidget"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"strings"
)

func ScrapeInsideOfTag(widget HtmlElementWidget.HtmlElementWidgetInterface, text string) bool {
	parameters := strings.Split(text, " ")
	result := htmlVariables.SetHtmlTag(parameters[0], widget)
	if widget.GetHtmlTag() == htmlVariables.HTML_STYLE {
		tree.CssStyleTagList = append(tree.CssStyleTagList, widget)
	}
	for _, s := range parameters[1:] {
		println(s)
		varName, varValue, found := strings.Cut(s, "=")
		if found {
			if widget.SetStandardVariables(varName, varValue) == false {
				VarReader, ok := widget.(VarReaderInterface)
				if ok {
					VarReader.VarReaderFunc(varName, varValue)
				}
			}
		} else {
			if widget.SetStandardContextVariables(s) == false { //remove later
				ContextReader, ok := widget.(ContextReaderInterface)
				if ok {
					ContextReader.ContextReaderFunc(s)
				}
			}
		}
	}
	return result
}
