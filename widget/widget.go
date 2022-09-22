package widget

import (
	"gezgin_web_engine/css_scraper"
	"gezgin_web_engine/html_scraper"
	"gezgin_web_engine/html_scraper/tags"
)

type Widget struct {
	childrenCount         int
	childrenIndex         int
	htmlTag               tags.HtmlTags
	widgetProperties      any
	standardHtmlVariables html_scraper.StandardHtmlTagVariables
	cssProperties         css_scraper.CssProperties
	//drawProperties
	children []*Widget
	parent   *Widget
	draw     bool
}
