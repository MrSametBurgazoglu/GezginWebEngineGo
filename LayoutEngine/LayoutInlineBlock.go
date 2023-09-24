package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/widget"
)

func SetWidthInlineBlock(currentWidget widget.WidgetInterface, styleProperty *StyleProperty.StyleProperty) {
	if len(currentWidget.GetChildren()) > 0 {
		width := 0
		for _, child := range currentWidget.GetChildren() {
			width += child.GetLayout().GetTotalWidth()
		}
		if styleProperty != nil && styleProperty.Padding != nil {
			currentWidget.GetLayout().PaddingLeft = int(styleProperty.Padding.PaddingLeft)
			currentWidget.GetLayout().PaddingRight = int(styleProperty.Padding.PaddingRight)
			width += currentWidget.GetLayout().PaddingLeft + currentWidget.GetLayout().PaddingRight
		}
		currentWidget.GetLayout().Width = width
		currentWidget.GetLayout().ContentWidth = width
	}
}
