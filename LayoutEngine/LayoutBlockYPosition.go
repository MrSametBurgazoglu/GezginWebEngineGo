package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func BlockSetPositionYSticky(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	return parent.GetLayout().ContentYPosition
}

func BlockSetPositionYStatic(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Margin != nil {
		CalculateTopMargin(currentWidget, false)
		CalculateBottomMargin(currentWidget, false)
	}
	if beforeCurrentWidget != nil {
		totalHeight := beforeCurrentWidget.GetLayout().GetTotalHeight()
		if totalHeight < 0 {
			totalHeight = 0
		}
		currentWidget.GetLayout().YPosition = beforeCurrentWidget.GetLayout().YPosition + totalHeight
		currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition + currentWidget.GetLayout().MarginTop
	} else {
		currentWidget.GetLayout().YPosition = parent.GetLayout().YPosition
		currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition + currentWidget.GetLayout().MarginTop
	}
}

func BlockSetPositionYAbsolute(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
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

func BlockSetPositionYFixed(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	return 0
}

func BlockSetPositionYRelative(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	position := 0
	if beforeCurrentWidget != nil {
		position = beforeCurrentWidget.GetLayout().YPosition + beforeCurrentWidget.GetLayout().Height + int(currentWidget.GetStyleProperty().Top)
	} else {
		position = parent.GetLayout().YPosition + int(currentWidget.GetStyleProperty().Top)
	}
	currentWidget.GetLayout().YPosition = position
	currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
}

func BlockSetPositionY(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			BlockSetPositionYSticky(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_EMPTY:
			BlockSetPositionYStatic(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_STATIC:
			BlockSetPositionYStatic(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			BlockSetPositionYAbsolute(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_FIXED:
			BlockSetPositionYFixed(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_RELATIVE:
			BlockSetPositionYRelative(currentWidget, parent, beforeCurrentWidget)
		}
	} else {
		if beforeCurrentWidget == nil {
			currentWidget.GetLayout().YPosition = parent.GetLayout().ContentYPosition
		} else {
			currentWidget.GetLayout().YPosition = beforeCurrentWidget.GetLayout().ContentYPosition + beforeCurrentWidget.GetLayout().Height
		}
	}
}
