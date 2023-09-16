package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func SetPositionX(currentWidget, parent widget.WidgetInterface) int {
	position := 0
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			position = parent.GetLayout().ContentXPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			position = parent.GetLayout().ContentXPosition
		case enums.CSS_POSITION_TYPE_STATIC:
			position = parent.GetLayout().ContentXPosition
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			if currentWidget.GetStyleProperty().Left != 0 {
				position = parent.GetLayout().ContentXPosition + int(currentWidget.GetStyleProperty().Left)
			} else if currentWidget.GetStyleProperty().Right != 0 {
				position = parent.GetLayout().ContentWidth - int(currentWidget.GetStyleProperty().Right)
			} else {
				position = parent.GetLayout().ContentXPosition
			}
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			if currentWidget.GetStyleProperty().Left != 0 {
				position = parent.GetLayout().ContentXPosition + int(currentWidget.GetStyleProperty().Left)
			} else if currentWidget.GetStyleProperty().Right != 0 {
				position = parent.GetLayout().ContentWidth - int(currentWidget.GetStyleProperty().Right)
			} else {
				position = parent.GetLayout().ContentXPosition
			}
		}
	} else {
		position = parent.GetLayout().ContentXPosition
	}
	currentWidget.GetLayout().ContentXPosition = position
	return currentWidget.GetLayout().ContentXPosition
}
