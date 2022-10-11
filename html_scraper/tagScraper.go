package html_scraper

import "strings"

func ScrapeInsideofTag(widget *Widget, text string) {
	parameters := strings.Split(text, " ")
	widget.HtmlTag.setHtmlTag(parameters[0], widget)
}
