package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func BlockSetPosition(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) (int, int) {
	return BlockSetPositionX(currentWidget, parent), BlockSetPositionY(currentWidget, parent, beforeCurrentWidget)
}

func BlockSetPositionX(currentWidget, parent widget.WidgetInterface) int {
	position := 0
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			position = parent.GetLayout().ContentXPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			position = parent.GetLayout().XPosition
			if currentWidget.GetStyleProperty().Margin != nil {
				if currentWidget.GetStyleProperty().Margin.MarginLeftValueType == enums.CSS_PROPERTY_VALUE_TYPE_AUTO && currentWidget.GetStyleProperty().Margin.MarginRightValueType == enums.CSS_PROPERTY_VALUE_TYPE_AUTO {
					position += (currentWidget.GetLayout().Parent.Width - currentWidget.GetLayout().Width) / 2
				}
			}
			//position = parent.ContentXPosition
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
	currentWidget.GetLayout().XPosition = position
	return currentWidget.GetLayout().ContentXPosition
}

func BlockSetPositionY(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			return parent.GetLayout().XPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			if beforeCurrentWidget != nil {
				marginTop := 0
				if currentWidget.GetStyleProperty().Margin != nil {
					marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
				}
				return beforeCurrentWidget.GetLayout().YPosition + beforeCurrentWidget.GetLayout().Height + marginTop
			} else {
				marginTop := 0
				if currentWidget.GetStyleProperty().Margin != nil {
					marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
				}
				return parent.GetLayout().YPosition + marginTop
			}
		case enums.CSS_POSITION_TYPE_STATIC:
			marginTop := 0
			if currentWidget.GetStyleProperty().Margin != nil {
				marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
			}
			if beforeCurrentWidget != nil {
				return beforeCurrentWidget.GetLayout().YPosition + beforeCurrentWidget.GetLayout().Height + marginTop
			} else {
				return parent.GetLayout().YPosition + parent.GetLayout().Height + marginTop
			}
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			if currentWidget.GetStyleProperty().Top != 0 {
				return parent.GetLayout().YPosition + int(currentWidget.GetStyleProperty().Top)
			} else if currentWidget.GetStyleProperty().Bottom != 0 {
				return parent.GetLayout().YPosition + parent.GetLayout().Height - int(currentWidget.GetStyleProperty().Bottom)
			} else {
				return parent.GetLayout().YPosition + parent.GetLayout().Height
			}
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			if beforeCurrentWidget != nil {
				return beforeCurrentWidget.GetLayout().YPosition + beforeCurrentWidget.GetLayout().Height + int(currentWidget.GetStyleProperty().Top)
			} else {
				return parent.GetLayout().YPosition + int(currentWidget.GetStyleProperty().Top)
			}
		}
	} else {
		if beforeCurrentWidget == nil {
			return parent.GetLayout().YPosition
		} else {
			return beforeCurrentWidget.GetLayout().Height
		}
	}
	return 0
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
