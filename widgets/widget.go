package widgets

import (
	"gezgin_web_engine/CssParser"
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/HtmlParser/htmlVariables/standardHtmlTagVariables"
	"gezgin_web_engine/LayoutEngine"
	"gezgin_web_engine/StyleEngine"
	structs2 "gezgin_web_engine/drawer/structs"
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

type Widget struct {
	ID             string
	Classes        []string
	ChildrenCount  int
	ChildrenIndex  int
	HtmlElement    *HtmlParser.HtmlElement
	StyleProperty  *StyleEngine.StyleProperty
	LayoutProperty *LayoutEngine.LayoutProperty
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

func (receiver *Widget) Initialize() {
	receiver.StyleProperty = new(StyleEngine.StyleProperty)
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

func (receiver *Widget) GetChildrenByIndex(index int) WidgetInterface {
	return receiver.Children[index]
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

func (receiver *Widget) CopyFromHtmlElement(htmlElement *HtmlParser.HtmlElement) {
	receiver.ChildrenCount = htmlElement.ChildrenCount
	receiver.ChildrenIndex = htmlElement.ChildrenIndex
	receiver.ID = htmlElement.Attributes["id"]
	receiver.Classes = strings.Split(htmlElement.Attributes["class"], " ")
	styleText, found := htmlElement.Attributes["style"]
	if found {
		receiver.StyleRules = CssParser.ParseCssFromInlineStyle(styleText)
	}
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

func (receiver *Widget) GetDrawProperties() *structs2.DrawProperties {
	return receiver.DrawProperties
}
