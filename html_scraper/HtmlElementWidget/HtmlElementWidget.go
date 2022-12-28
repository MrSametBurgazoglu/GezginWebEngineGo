package HtmlElementWidget

import (
	"gezgin_web_engine/html_scraper/HtmlTags"
)

type HtmlElementWidgetInterface interface {
	GetChildrenCount() int
	SetChildrenCount(int)
	GetChildrenIndex() int
	SetChildrenIndex(int)
	GetChildren() []HtmlElementWidgetInterface
	AppendChild(HtmlElementWidgetInterface)
	RemoveChild(HtmlElementWidgetInterface)
	GetChild(int) HtmlElementWidgetInterface
	GetHtmlTag() HtmlTags.HtmlTags
	SetHtmlTag(HtmlTags.HtmlTags)
	SetParent(HtmlElementWidgetInterface)
	GetParent() HtmlElementWidgetInterface
	SetStandardVariables(string, string) bool
	SetStandardContextVariables(string) bool
	IsDrawable() bool
	GetRendered() bool
	SetRendered(bool)
	//GetCssProperties() *structs.CssProperties
}
