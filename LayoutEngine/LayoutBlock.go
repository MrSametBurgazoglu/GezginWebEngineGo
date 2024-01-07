package LayoutEngine

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/widget"
)

func BlockSetPosition(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if currentWidget.GetHtmlTag() == int(HtmlParser.HTML_UL) {
		println("hey")
	}
	BlockSetPositionX(currentWidget, parent)
	BlockSetPositionY(currentWidget, parent, beforeCurrentWidget)
}

func SetWidthBlock(currentWidget, parent widget.WidgetInterface) {
	if currentWidget.GetHtmlTag() == int(HtmlParser.HTML_INPUT) || currentWidget.GetHtmlTag() == int(HtmlParser.HTML_SVG) {
		println("hey")
	}
	width := parent.GetLayout().ContentWidth
	if currentWidget.GetStyleProperty().Width != 0 {
		width = currentWidget.GetLayout().GetWidthFromStyleProperty()
	}
	currentWidget.GetLayout().ContentWidth = width
	if currentWidget.GetStyleProperty() != nil {
		if currentWidget.GetStyleProperty().Margin != nil {
			CalculateLeftMargin(currentWidget, true)
			CalculateRightMargin(currentWidget, true)
			width += currentWidget.GetLayout().MarginLeft + currentWidget.GetLayout().MarginRight
		}
		if currentWidget.GetStyleProperty().Padding != nil {
			width += int(currentWidget.GetStyleProperty().Padding.PaddingLeft + currentWidget.GetStyleProperty().Padding.PaddingRight)
			currentWidget.GetLayout().ContentWidth -= int(currentWidget.GetStyleProperty().Padding.PaddingLeft + currentWidget.GetStyleProperty().Padding.PaddingRight)
		}
	}
	currentWidget.GetLayout().Width = width
}
