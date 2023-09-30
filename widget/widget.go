package widget

import (
	"gezgin_web_engine/CssParser"
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/HtmlParser/htmlVariables/standardHtmlTagVariables"
	"gezgin_web_engine/LayoutProperty"
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/StyleProperty/enums"
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
	StyleProperty  *StyleProperty.StyleProperty
	LayoutProperty *LayoutProperty.LayoutProperty
	standardHtmlTagVariables.StandardHtmlTagVariables
	StyleRules        map[string]string
	DrawProperties    *structs.DrawProperties
	RenderWidget      func(*Widget, *image.RGBA)
	DrawWidget        func(*Widget, *image.RGBA)
	Children          []WidgetInterface
	Parent            WidgetInterface
	Draw              bool
	Rendered          bool
	IsNotDrawChildren bool
}

func (receiver *Widget) Initialize() {
	receiver.StyleProperty = new(StyleProperty.StyleProperty)
	receiver.StyleProperty.Initialize()
	receiver.DrawProperties = new(structs.DrawProperties)
	receiver.DrawProperties.Initialize()
	receiver.LayoutProperty = new(LayoutProperty.LayoutProperty)
	receiver.CopyFromHtmlElement(receiver.HtmlElement)
	receiver.LayoutProperty.StyleProperty = receiver.StyleProperty
}

func (receiver *Widget) SetChildrenCount(count int) {
	receiver.ChildrenCount = count
}

func (receiver *Widget) GetIsNotDrawChildren() bool {
	return receiver.IsNotDrawChildren
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
	if receiver.LayoutProperty != nil {
		receiver.LayoutProperty.Parent = widget.GetLayout()
		receiver.Parent.GetLayout().Children = append(receiver.Parent.GetLayout().Children, receiver.LayoutProperty)
	}
	if receiver.StyleProperty != nil {
		receiver.StyleProperty.Parent = widget.GetStyleProperty()
		receiver.Parent.GetStyleProperty().Children = append(receiver.Parent.GetStyleProperty().Children, receiver.StyleProperty)
	}
}

func (receiver *Widget) AppendChild(child WidgetInterface) {
	receiver.Children = append(receiver.Children, child)
	child.SetChildrenIndex(receiver.ChildrenCount)
	receiver.ChildrenCount += 1
}

func (receiver *Widget) GetRect() *image.Rectangle {
	return receiver.DrawProperties.Rect
}

func (receiver *Widget) IsDraw() bool {
	if receiver.GetStyleProperty() != nil && receiver.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_NONE {
		return false
	} else {
		return true
	}
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
	//receiver.ChildrenCount = htmlElement.ChildrenCount
	//receiver.ChildrenIndex = htmlElement.ChildrenIndex
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

func (receiver *Widget) GetStyleProperty() *StyleProperty.StyleProperty {
	return receiver.StyleProperty
}

func (receiver *Widget) GetID() string {
	return receiver.ID
}

func (receiver *Widget) GetClasses() []string {
	return receiver.Classes
}

func (receiver *Widget) GetAttributes() map[string]string {
	return receiver.HtmlElement.Attributes
}

func (receiver *Widget) GetHtmlTag() int {
	return int(receiver.HtmlElement.HtmlTag)
}

func (receiver *Widget) GetHtmlName() string {
	return receiver.HtmlElement.Name
}

func (receiver *Widget) GetStyleRules() map[string]string {
	return receiver.StyleRules
}

func (receiver *Widget) GetDrawProperties() *structs.DrawProperties {
	return receiver.DrawProperties
}

func (receiver *Widget) GetLayout() *LayoutProperty.LayoutProperty {
	return receiver.LayoutProperty
}

func (receiver *Widget) CalculateWidth() {

}

func (receiver *Widget) IsPreSetWidth() bool {
	return receiver.StyleProperty != nil && receiver.StyleProperty.Display == enums.CSS_DISPLAY_TYPE_BLOCK
}

func (receiver *Widget) IsSetWidthSelf() bool {
	return receiver.StyleProperty != nil && receiver.StyleProperty.Display == enums.CSS_DISPLAY_TYPE_FLEX || (receiver.StyleProperty.Float == enums.CSS_FLOAT_LEFT || receiver.StyleProperty.Float == enums.CSS_FLOAT_RIGHT)
}
