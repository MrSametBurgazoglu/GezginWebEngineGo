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
	width := parent.GetLayout().ContentWidth
	if currentWidget.GetStyleProperty().Width != 0 {
		width = currentWidget.GetLayout().GetWidthFromStyleProperty()
	}
	currentWidget.GetLayout().Width = width
	contentWidth := width
	if currentWidget.GetStyleProperty() != nil {
		if currentWidget.GetStyleProperty().Margin != nil {
			contentWidth -= currentWidget.GetLayout().MarginLeft + currentWidget.GetLayout().MarginRight
		}
	}
	currentWidget.GetLayout().ContentWidth = contentWidth
}
