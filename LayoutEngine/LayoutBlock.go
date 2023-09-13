package LayoutEngine

import (
	"gezgin_web_engine/widget"
)

func BlockSetPosition(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) (int, int) {
	return BlockSetPositionX(currentWidget, parent), BlockSetPositionY(currentWidget, parent, beforeCurrentWidget)
}

func SetWidthBlock(currentWidget, parent widget.WidgetInterface) {
	width := parent.GetLayout().Width
	if currentWidget.GetStyleProperty().MaxWidth > 0 && uint(width) > currentWidget.GetStyleProperty().MaxWidth {
		width = int(currentWidget.GetStyleProperty().MaxWidth)
	}
	if currentWidget.GetStyleProperty().MinWidth > 0 && uint(width) < currentWidget.GetStyleProperty().MinWidth {
		width = int(currentWidget.GetStyleProperty().MinWidth)
	}
	currentWidget.GetLayout().Width = width
	if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Margin != nil {
		contentWidth := width - (currentWidget.GetStyleProperty().Margin.MarginLeft + currentWidget.GetStyleProperty().Margin.MarginRight)
		currentWidget.GetLayout().ContentWidth = contentWidth
	} else {
		currentWidget.GetLayout().ContentWidth = width
	}
}
