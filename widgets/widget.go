package widgets

import (
	"gezgin_web_engine/CssParser"
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/HtmlParser/htmlVariables/standardHtmlTagVariables"
	"gezgin_web_engine/LayoutEngine"
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/StyleEngine/enums"
	"gezgin_web_engine/drawer/structs"
	"image"
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
	DrawProperties *structs.DrawProperties
	RenderWidget   func(*Widget, *image.RGBA)
	DrawWidget     func(*Widget, *image.RGBA)
	Children       []WidgetInterface
	Parent         WidgetInterface
	Draw           bool
	Rendered       bool
}

func (receiver *Widget) Initialize() {
	receiver.StyleProperty = new(StyleEngine.StyleProperty)
	receiver.StyleProperty.Initialize()
	receiver.DrawProperties = new(structs.DrawProperties)
	receiver.DrawProperties.Initialize()
	receiver.LayoutProperty = new(LayoutEngine.LayoutProperty)
	receiver.CopyFromHtmlElement(receiver.HtmlElement)

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

func (receiver *Widget) GetRect() *image.Rectangle {
	return receiver.DrawProperties.Rect
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
	receiver.HtmlElement = htmlElement
	receiver.ChildrenCount = htmlElement.ChildrenCount
	receiver.ChildrenIndex = htmlElement.ChildrenIndex
	receiver.ID = htmlElement.Attributes["id"]
	classes, found := htmlElement.Attributes["class"]
	if found {
		receiver.Classes = strings.Split(classes, " ")
	}
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

func (receiver *Widget) GetDrawProperties() *structs.DrawProperties {
	return receiver.DrawProperties
}

func (receiver *Widget) GetLayout() *LayoutEngine.LayoutProperty {
	return receiver.LayoutProperty
}

func (receiver *Widget) CalculateWidth() {

}

func (receiver *Widget) IsPreSetWidth() bool {
	return receiver.StyleProperty.Display == enums.CSS_DISPLAY_TYPE_BLOCK || receiver.StyleProperty.Display == enums.CSS_DISPLAY_TYPE_FLEX
}
