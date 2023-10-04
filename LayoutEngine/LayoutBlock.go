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
	if currentWidget.GetStyleProperty().MaxWidth > 0 && uint(width) > currentWidget.GetStyleProperty().MaxWidth {
		width = int(currentWidget.GetStyleProperty().MaxWidth)
	}
	if currentWidget.GetStyleProperty().MinWidth > 0 && uint(width) < currentWidget.GetStyleProperty().MinWidth {
		width = int(currentWidget.GetStyleProperty().MinWidth)
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
