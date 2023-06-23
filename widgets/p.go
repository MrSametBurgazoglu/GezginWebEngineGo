package widgets

import (
	structs2 "gezgin_web_engine/StyleEngine/structs"
)

type HtmlTagP struct {
	*Widget
}

func (receiver *HtmlTagP) Draw() {

}

func (receiver *HtmlTagP) Render() {

}

func SetWidgetPropertiesForPTag() WidgetInterface {
	widget := new(HtmlTagP)
	widget.StyleProperty.Font = new(structs2.Font)
	widget.StyleProperty.Font.FontSizeValue = 14
	widget.StyleProperty.Margin = new(structs2.Margin)
	widget.StyleProperty.Margin.MarginTop = 10
	return widget
}
