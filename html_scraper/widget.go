package html_scraper

import (
	"gezgin_web_engine/css_scraper"
)

type Widget struct {
	ChildrenCount         int
	ChildrenIndex         int
	HtmlTag               HtmlTags
	WidgetProperties      any
	StandardHtmlVariables StandardHtmlTagVariables
	CssProperties         css_scraper.CssProperties
	//drawProperties
	VarReaderFunc     func(Widget, string, string)
	ContextReaderFunc func(Widget, string)
	Children          []Widget
	Parent            *Widget
	Draw              bool
}
