package LayoutEngine

import (
	"gezgin_web_engine/widget"
)

func SetHeight(currentWidget widget.WidgetInterface) int {
	if len(currentWidget.GetChildren()) > 0 {
		height := 0
		for _, child := range currentWidget.GetChildren() {
			height += child.GetLayout().Height
		}
		if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Margin != nil {
			contentHeight := height - (currentWidget.GetStyleProperty().Margin.MarginTop + currentWidget.GetStyleProperty().Margin.MarginBottom)
			currentWidget.GetLayout().ContentHeight = contentHeight
		} else {
			currentWidget.GetLayout().ContentHeight = height
		}
		currentWidget.GetLayout().Height = height
		return currentWidget.GetLayout().Height
	} else {
		return currentWidget.GetLayout().Height
	}
}
