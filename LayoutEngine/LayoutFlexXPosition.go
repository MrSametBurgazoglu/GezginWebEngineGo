package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func FlexSetPositionXSticky(currentWidget, parent widget.WidgetInterface) {
}

func FlexSetPositionXStatic(currentWidget, beforeCurrentWidget, parent widget.WidgetInterface) {
	if beforeCurrentWidget == nil {
		x := parent.GetLayout().XPosition
		if currentWidget.GetStyleProperty().Margin != nil {
			CalculateLeftMargin(currentWidget, true)
			CalculateRightMargin(currentWidget, true)
		}
		x += currentWidget.GetLayout().MarginLeft
		currentWidget.GetLayout().XPosition = x
		if currentWidget.GetStyleProperty().Padding != nil {
			x += currentWidget.GetLayout().PaddingLeft
		}
		currentWidget.GetLayout().ContentXPosition = x
	} else {
		x := beforeCurrentWidget.GetLayout().XPosition + beforeCurrentWidget.GetLayout().Width
		if currentWidget.GetStyleProperty().Margin != nil {
			CalculateLeftMargin(currentWidget, false)
			CalculateRightMargin(currentWidget, false)
			x += currentWidget.GetLayout().MarginLeft
		}
		currentWidget.GetLayout().XPosition = x
		if currentWidget.GetStyleProperty().Padding != nil {
			x += currentWidget.GetLayout().PaddingLeft
		}
		currentWidget.GetLayout().ContentXPosition = x
	}
}

func FlexSetPositionXAbsolute(currentWidget, parent widget.WidgetInterface) int {
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

func FlexSetPositionXFixed(currentWidget, parent widget.WidgetInterface) int {
	return 0
}

func FlexSetPositionXRelative(currentWidget, parent widget.WidgetInterface) int {
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

func SetPositionXFlex(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			FlexSetPositionXSticky(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_EMPTY:
			FlexSetPositionXStatic(currentWidget, beforeCurrentWidget, parent)
		case enums.CSS_POSITION_TYPE_STATIC:
			FlexSetPositionXStatic(currentWidget, beforeCurrentWidget, parent)
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			FlexSetPositionXAbsolute(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_FIXED:
			FlexSetPositionXFixed(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_RELATIVE:
			FlexSetPositionXRelative(currentWidget, parent)
		}
	} else {
		currentWidget.GetLayout().XPosition = parent.GetLayout().ContentXPosition
		currentWidget.GetLayout().ContentXPosition = parent.GetLayout().ContentXPosition
	}
}
