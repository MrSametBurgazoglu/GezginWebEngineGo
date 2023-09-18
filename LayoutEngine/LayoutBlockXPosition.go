package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func BlockSetPositionXSticky(currentWidget, parent widget.WidgetInterface) {
}

func BlockSetPositionXStatic(currentWidget, parent widget.WidgetInterface) {
	position := parent.GetLayout().ContentXPosition
	if currentWidget.GetStyleProperty().Margin != nil {
		CalculateLeftMargin(currentWidget, true)
		CalculateRightMargin(currentWidget, true)
		position += currentWidget.GetLayout().MarginLeft
	}
	currentWidget.GetLayout().XPosition = position
	if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Padding != nil {
		currentWidget.GetLayout().PaddingLeft = int(currentWidget.GetStyleProperty().Padding.PaddingLeft)
		currentWidget.GetLayout().PaddingRight = int(currentWidget.GetStyleProperty().Padding.PaddingRight)
		position += currentWidget.GetLayout().PaddingLeft
	}
	currentWidget.GetLayout().ContentXPosition = position
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

func BlockSetPositionX(currentWidget, parent widget.WidgetInterface) {
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			BlockSetPositionXSticky(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_EMPTY:
			BlockSetPositionXStatic(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_STATIC:
			BlockSetPositionXStatic(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			BlockSetPositionXAbsolute(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_FIXED:
			BlockSetPositionXFixed(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_RELATIVE:
			BlockSetPositionXRelative(currentWidget, parent)
		}
	}
}
