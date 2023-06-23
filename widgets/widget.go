package widgets

import (
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/cssParser"
	structs2 "gezgin_web_engine/drawer/structs"
	"gezgin_web_engine/htmlParser"
	"gezgin_web_engine/htmlParser/htmlVariables/standardHtmlTagVariables"
	"github.com/veandco/go-sdl2/sdl"
	"strings"
)

type HtmlTagsInterface interface {
	SetHtmlTag() int
}

/*Widget
htmlElement
StyleProperty
DrawProperties
Events
*/

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
	AppendChild(WidgetInterface)
	GetParent() WidgetInterface
	SetParent(WidgetInterface)
	IsDraw() bool
	SetDraw(draw bool)
	IsRender() bool
	SetRender(render bool)
	CopyFromHtmlElement(htmlElement *htmlParser.HtmlElement)
	GetStyleProperty() *StyleEngine.StyleProperty
	Draw()
	Render()
}

type Widget struct {
	ID            string
	Classes       []string
	ChildrenCount int
	ChildrenIndex int
	HtmlElement   *htmlParser.HtmlElement
	StyleProperty *StyleEngine.StyleProperty
	standardHtmlTagVariables.StandardHtmlTagVariables
	StyleRules     map[string]string
	DrawProperties *structs2.DrawProperties
	RenderWidget   func(*Widget, *sdl.Renderer)
	DrawWidget     func(*Widget, *sdl.Renderer)
	Children       []WidgetInterface
	Parent         WidgetInterface
	Draw           bool
	Rendered       bool
}

func (receiver *Widget) SetChildrenCount(count int) {
	receiver.ChildrenCount = count
}

func (receiver *Widget) GetChildrenCount() int {
	return receiver.ChildrenCount
}

func (receiver *Widget) SetChildrenIndex(index int) {
	receiver.ChildrenIndex = index
}

func (receiver *Widget) GetChildrenIndex() int {
	return receiver.ChildrenIndex
}

func (receiver *Widget) GetChildren() []WidgetInterface {
	return receiver.Children
}

func (receiver *Widget) GetParent() WidgetInterface {
	return receiver.Parent
}

func (receiver *Widget) SetParent(widget WidgetInterface) {
	receiver.Parent = widget
}

func (receiver *Widget) AppendChild(child WidgetInterface) {
	receiver.Children = append(receiver.Children, child)
}

func (receiver *Widget) GetRect() *sdl.Rect {
	return &receiver.DrawProperties.Rect
}

func (receiver *Widget) IsDraw() bool {
	return receiver.Draw
}

func (receiver *Widget) SetDraw(draw bool) {
	receiver.Draw = draw
}

func (receiver *Widget) IsRender() bool {
	return receiver.Rendered
}

func (receiver *Widget) SetRender(render bool) {
	receiver.Rendered = render
}

func (receiver *Widget) CopyFromHtmlElement(htmlElement *htmlParser.HtmlElement) {
	receiver.ChildrenCount = htmlElement.ChildrenCount
	receiver.ChildrenIndex = htmlElement.ChildrenIndex
	receiver.ID = htmlElement.Attributes["id"]
	receiver.Classes = strings.Split(htmlElement.Attributes["class"], " ")
	receiver.StyleRules = cssParser.ParseCssFromInlineStyle(htmlElement.Attributes["style"])
}

func (receiver *Widget) GetStyleProperty() *StyleEngine.StyleProperty {
	return receiver.StyleProperty
}

func (receiver *Widget) GetID() string {
	return receiver.ID
}

func (receiver *Widget) GetClasses() []string {
	return receiver.Classes
}

func (receiver *Widget) GetHtmlTag() int {
	return int(receiver.HtmlElement.HtmlTag)
}

func (receiver *Widget) GetStyleRules() map[string]string {
	return receiver.StyleRules
}
