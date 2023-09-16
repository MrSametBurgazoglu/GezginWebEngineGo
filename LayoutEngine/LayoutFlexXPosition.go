package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func FlexSetPositionXSticky(currentWidget, parent widget.WidgetInterface) int {
	return parent.GetLayout().ContentXPosition
}

func FlexSetPositionXStatic(currentWidget, beforeCurrentWidget, parent widget.WidgetInterface) int {
	position := parent.GetLayout().ContentXPosition
	x := 0
	if beforeCurrentWidget == nil {
		x = parent.GetLayout().XPosition
		if currentWidget.GetStyleProperty().Margin != nil {
			CalculateLeftMargin(currentWidget, true)
			CalculateRightMargin(currentWidget, true)
		}
		x += currentWidget.GetLayout().MarginLeft
		println(currentWidget.GetLayout().MarginLeft)
		/*
			if currentWidget.GetStyleProperty().Margin != nil {
				if currentWidget.GetStyleProperty().Margin.MarginLeftValueType == enums.CSS_PROPERTY_VALUE_TYPE_AUTO && currentWidget.GetStyleProperty().Margin.MarginRightValueType == enums.CSS_PROPERTY_VALUE_TYPE_AUTO {
					marginLeft := (currentWidget.GetLayout().Parent.Width - currentWidget.GetLayout().Width) / 2
					x += marginLeft
					currentWidget.GetLayout().MarginLeft = marginLeft
					currentWidget.GetLayout().MarginRight = marginLeft
				} else {
					x += currentWidget.GetLayout().MarginLeft
				}
			}
		*/
		if currentWidget.GetStyleProperty().Padding != nil {
			x += currentWidget.GetLayout().PaddingLeft
		}
	} else {
		x = beforeCurrentWidget.GetLayout().XPosition + beforeCurrentWidget.GetLayout().Width
		if currentWidget.GetStyleProperty().Margin != nil {
			CalculateLeftMargin(currentWidget, false)
			CalculateRightMargin(currentWidget, false)
			x += currentWidget.GetLayout().MarginLeft
		}
		if currentWidget.GetStyleProperty().Padding != nil {
			x += currentWidget.GetLayout().PaddingLeft
		}
	}
	position = x
	return position
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

func SetPositionXFlex(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	position := 0
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			position = FlexSetPositionXSticky(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_EMPTY:
			position = FlexSetPositionXStatic(currentWidget, beforeCurrentWidget, parent)
		case enums.CSS_POSITION_TYPE_STATIC:
			position = FlexSetPositionXStatic(currentWidget, beforeCurrentWidget, parent)
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			position = FlexSetPositionXAbsolute(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_FIXED:
			position = FlexSetPositionXFixed(currentWidget, parent)
		case enums.CSS_POSITION_TYPE_RELATIVE:
			position = FlexSetPositionXRelative(currentWidget, parent)
		}
	} else {
		position = parent.GetLayout().ContentXPosition
	}
	currentWidget.GetLayout().ContentXPosition = position
	currentWidget.GetLayout().XPosition = position
	return currentWidget.GetLayout().ContentXPosition
}
