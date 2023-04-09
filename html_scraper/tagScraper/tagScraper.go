package tagScraper

import (
	"gezgin_web_engine/css_scraper"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"gezgin_web_engine/html_scraper/widget"
	"gezgin_web_engine/javascript_interpreter"
	"strings"
	"sync"
)

func UntaggedTextClosed(widget *widget.Widget) {
	if widget.Parent.HtmlTag == htmlVariables.HTML_STYLE {
		css_scraper.CreateCssPropertiesFromStyleTag(widget.Parent)
	} else if widget.Parent.HtmlTag == htmlVariables.HTML_SCRIPT {
		javascript_interpreter.ScriptElements = append(javascript_interpreter.ScriptElements, widget.Parent)
	}

}

func ScrapeParameters(widget *widget.Widget, parameters []string, group *sync.WaitGroup) {
	for _, s := range parameters[0:] {
		println(s)
		varName, varValue, found := strings.Cut(s, "=")
		if found {
			println(varName, varValue, "asdad")
			if isStandard := widget.StandardHtmlVariables.SetStandardVariables(varName, varValue); isStandard == false && widget.ContextReaderFunc != nil {
				widget.VarReaderFunc(widget, varName, varValue)
			}
		} else {
			if isStandard := widget.StandardHtmlVariables.SetStandardContextVariables(s); isStandard == false && widget.ContextReaderFunc != nil {
				widget.ContextReaderFunc(widget, s)
			}
		}
	}
	group.Done()
	println("scraping parameters finished", widget)
}

func ScrapeInsideOfTag(widget *widget.Widget, text string, group *sync.WaitGroup) bool {
	parameters := strings.Split(text, " ")
	result := htmlVariables.SetHtmlTag(parameters[0], widget)
	group.Add(1)
	go ScrapeParameters(widget, parameters, group)
	return result
}
