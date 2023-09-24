package LayoutEngine

import (
	"gezgin_web_engine/widget"
)

func FlexSetPositionYSticky(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	return parent.GetLayout().ContentYPosition
}

func FlexSetPositionYStatic(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Margin != nil {
		CalculateTopMargin(currentWidget, false)
		CalculateBottomMargin(currentWidget, false)
	}
	currentWidget.GetLayout().YPosition = parent.GetLayout().ContentYPosition + currentWidget.GetLayout().MarginTop
	currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
}

func FlexSetPositionYAbsolute(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	position := 0
	if currentWidget.GetStyleProperty().Top != 0 {
		position = parent.GetLayout().YPosition + int(currentWidget.GetStyleProperty().Top)
	} else if currentWidget.GetStyleProperty().Bottom != 0 {
		position = parent.GetLayout().YPosition + parent.GetLayout().Height - int(currentWidget.GetStyleProperty().Bottom)
	} else {
		position = parent.GetLayout().YPosition + parent.GetLayout().Height
	}
	return position
}

func FlexSetPositionYFixed(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	return 0
}

func FlexSetPositionYRelative(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	FlexSetPositionYStatic(currentWidget, parent, beforeCurrentWidget)
	currentWidget.GetLayout().YPosition += int(currentWidget.GetStyleProperty().Top)
	currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
}
