package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func BlockSetPositionYSticky(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	return parent.GetLayout().ContentYPosition
}

func BlockSetPositionYStatic(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Margin != nil {
		CalculateTopMargin(currentWidget, false)
		CalculateBottomMargin(currentWidget, false)
	}
	if beforeCurrentWidget != nil {
		return beforeCurrentWidget.GetLayout().YPosition + beforeCurrentWidget.GetLayout().GetTotalHeight() + currentWidget.GetLayout().MarginTop
	} else {
		return parent.GetLayout().YPosition + currentWidget.GetLayout().MarginTop
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

func BlockSetPositionYRelative(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	position := 0
	if beforeCurrentWidget != nil {
		position = beforeCurrentWidget.GetLayout().YPosition + beforeCurrentWidget.GetLayout().Height + int(currentWidget.GetStyleProperty().Top)
	} else {
		position = parent.GetLayout().YPosition + int(currentWidget.GetStyleProperty().Top)
	}
	return position
}

func BlockSetPositionY(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	position := 0
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			position = BlockSetPositionYSticky(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_EMPTY:
			position = BlockSetPositionYStatic(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_STATIC:
			position = BlockSetPositionYStatic(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			position = BlockSetPositionYAbsolute(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_FIXED:
			position = BlockSetPositionYFixed(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_RELATIVE:
			position = BlockSetPositionYRelative(currentWidget, parent, beforeCurrentWidget)
		}
	} else {
		if beforeCurrentWidget == nil {
			position = parent.GetLayout().YPosition
		} else {
			position = beforeCurrentWidget.GetLayout().Height
		}
	}
	currentWidget.GetLayout().YPosition = position
	return position
}
