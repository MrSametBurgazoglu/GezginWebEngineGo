package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func BlockSetPositionXSticky(currentWidget, parent widget.WidgetInterface) int {
	return parent.GetLayout().ContentXPosition
}

func BlockSetPositionXStatic(currentWidget, parent widget.WidgetInterface) int {
	position := parent.GetLayout().ContentXPosition
	if currentWidget.GetStyleProperty().Margin != nil {
		CalculateLeftMargin(currentWidget, true)
		CalculateRightMargin(currentWidget, true)
		position += currentWidget.GetLayout().MarginLeft
	}
	return position
}

func BlockSetPositionXAbsolute(currentWidget, parent widget.WidgetInterface) int {
	position := 0
	if currentWidget.GetStyleProperty().Left != 0 {
		position = parent.GetLayout().ContentXPosition + int(currentWidget.GetStyleProperty().Left)
	} else if currentWidget.GetStyleProperty().Right != 0 {
		position = parent.GetLayout().ContentWidth - int(currentWidget.GetStyleProperty().Right)
	} else {
		position = parent.GetLayout().ContentXPosition
	}
	return position
}

func BlockSetPositionXFixed(currentWidget, parent widget.WidgetInterface) int {
	return 0
}

func BlockSetPositionXRelative(currentWidget, parent widget.WidgetInterface) int {
	position := 0
	if currentWidget.GetStyleProperty().Left != 0 {
		position = parent.GetLayout().ContentXPosition + int(currentWidget.GetStyleProperty().Left)
	} else if currentWidget.GetStyleProperty().Right != 0 {
		position = parent.GetLayout().ContentWidth - int(currentWidget.GetStyleProperty().Right)
	} else {
		position = parent.GetLayout().ContentXPosition
	}
	return position
}

func BlockSetPositionX(currentWidget, parent widget.WidgetInterface) int {
	position := 0
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			position = BlockSetPositionXSticky(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_EMPTY:
			position = BlockSetPositionXStatic(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_STATIC:
			position = BlockSetPositionXStatic(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			position = BlockSetPositionXAbsolute(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_FIXED:
			position = BlockSetPositionXFixed(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_RELATIVE:
			position = BlockSetPositionXRelative(currentWidget, parent)
		}
	} else {
		position = parent.GetLayout().ContentXPosition
	}
	currentWidget.GetLayout().ContentXPosition = position
	currentWidget.GetLayout().XPosition = position
	return currentWidget.GetLayout().ContentXPosition
}
