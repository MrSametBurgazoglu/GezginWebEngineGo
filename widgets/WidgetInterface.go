package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/drawer/structs"
	"image"
)

type WidgetInterface interface {
	GetID() string
	GetClasses() []string
	GetHtmlTag() int
	GetStyleRules() map[string]string
	SetChildrenCount(int)
	GetChildrenCount() int
	SetChildrenIndex(int)
	GetChildrenIndex() int
	GetChildren() []WidgetInterface
	GetChildrenByIndex(int) WidgetInterface
	AppendChild(WidgetInterface)
	GetParent() WidgetInterface
	SetParent(WidgetInterface)
	IsDraw() bool
	SetDraw(draw bool)
	IsRender() bool
	SetRender(render bool)
	CopyFromHtmlElement(htmlElement *HtmlParser.HtmlElement)
	GetStyleProperty() *StyleEngine.StyleProperty
	GetDrawProperties() *structs.DrawProperties
	Draw(rgba *image.RGBA)
	Render(rgba *image.RGBA, resourceManager *ResourceManager.ResourceManager)
	IsBlockElement() bool
}
