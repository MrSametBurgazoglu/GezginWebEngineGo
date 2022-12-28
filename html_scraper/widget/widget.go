package widget

import (
	"gezgin_web_engine/html_scraper/HtmlElementWidget"
	"gezgin_web_engine/html_scraper/HtmlTags"
	"gezgin_web_engine/html_scraper/htmlVariables/standardHtmlTagVariables"
)

// Widget TODO MAKE WIDGET NOT STRUCT, WIDGET WILL BE BASE INTERFACE FOR WIDGET STRUCTS IN BASE OPERATIONS
type BaseWidget struct {
	ChildrenCount int
	ChildrenIndex int
	HtmlTag       HtmlTags.HtmlTags
	*standardHtmlTagVariables.StandardHtmlTagVariables
	Children []HtmlElementWidget.HtmlElementWidgetInterface
	Parent   HtmlElementWidget.HtmlElementWidgetInterface
	Draw     bool
	Rendered bool
}

func (receiver *BaseWidget) GetChildrenCount() int {
	return receiver.ChildrenCount
}

func (receiver *BaseWidget) SetChildrenCount(value int) {
	receiver.ChildrenCount = value
}

func (receiver *BaseWidget) GetChildrenIndex() int {
	return receiver.ChildrenIndex
}

func (receiver *BaseWidget) SetChildrenIndex(value int) {
	receiver.ChildrenIndex = value
}

func (receiver *BaseWidget) GetChildren() []HtmlElementWidget.HtmlElementWidgetInterface {
	return receiver.Children
}

func (receiver *BaseWidget) AppendChild(widgetInterface HtmlElementWidget.HtmlElementWidgetInterface) {
	receiver.Children = append(receiver.Children, widgetInterface)
}

func (receiver *BaseWidget) RemoveChild(widgetInterface HtmlElementWidget.HtmlElementWidgetInterface) {
	receiver.Children = append(receiver.Children[:widgetInterface.GetChildrenIndex()], receiver.Children[widgetInterface.GetChildrenIndex()+1:]...)
}

func (receiver *BaseWidget) GetChild(index int) HtmlElementWidget.HtmlElementWidgetInterface {
	return receiver.Children[index]
}

func (receiver *BaseWidget) GetParent() HtmlElementWidget.HtmlElementWidgetInterface {
	return receiver.Parent
}

func (receiver *BaseWidget) SetParent(element HtmlElementWidget.HtmlElementWidgetInterface) {
	receiver.Parent = element
}

func (receiver *BaseWidget) GetHtmlTag() HtmlTags.HtmlTags {
	return receiver.HtmlTag
}

func (receiver *BaseWidget) SetHtmlTag(tag HtmlTags.HtmlTags) {
	receiver.HtmlTag = tag
}

func (receiver *BaseWidget) IsDrawable() bool {
	return receiver.Draw
}

func (receiver *BaseWidget) GetRendered() bool {
	return receiver.Rendered
}

func (receiver *BaseWidget) SetRendered(value bool) {
	receiver.Rendered = value
}
