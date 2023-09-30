package widget

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/LayoutProperty"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/drawer/structs"
	"image"
)

type WidgetInterface interface {
	GetID() string
	GetClasses() []string
	GetAttributes() map[string]string
	GetHtmlTag() int
	GetHtmlName() string
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
	GetStyleProperty() *StyleProperty.StyleProperty
	GetDrawProperties() *structs.DrawProperties
	GetLayout() *LayoutProperty.LayoutProperty
	Draw(rgba *image.RGBA)
	Render(rgba *image.RGBA, resourceManager *ResourceManager.ResourceManager)
	IsPreSetWidth() bool
	IsSetWidthSelf() bool
	GetIsNotDrawChildren() bool
}
