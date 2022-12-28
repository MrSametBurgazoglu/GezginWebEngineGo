package css_scraper

import (
	"gezgin_web_engine/css_scraper/structs"
	"gezgin_web_engine/html_scraper/HtmlTags"
)

type ICssProperties interface {
	GetCssProperties() *structs.CssProperties
	GetClass() []string
	GetHtmlTag() HtmlTags.HtmlTags
	GetID() string
	GetStyle() string
	GetChild(int) any
	GetChildrenCount() int
	IsDrawable() bool
	GetParent() ICssProperties
}
