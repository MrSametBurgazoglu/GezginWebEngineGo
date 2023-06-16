package htmlParser

import (
	"gezgin_web_engine/htmlParser/htmlVariables/standardHtmlTagVariables"
)

type HtmlElement struct {
	ChildrenCount    int
	ChildrenIndex    int
	HtmlTag          HtmlTags
	WidgetProperties any
	Attributes       map[string]string
	standardHtmlTagVariables.StandardHtmlTagVariables
	Children []*HtmlElement
	Parent   *HtmlElement
}
