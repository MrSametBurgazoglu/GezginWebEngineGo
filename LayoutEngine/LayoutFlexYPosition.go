package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func FlexSetPositionYSticky(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	return parent.GetLayout().ContentYPosition
}

func FlexSetPositionYStatic(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	x := 0
	if beforeCurrentWidget == nil {
		x = parent.GetLayout().XPosition
		if currentWidget.GetStyleProperty().Margin != nil {
			if currentWidget.GetStyleProperty().Margin.MarginLeftValueType == enums.CSS_PROPERTY_VALUE_TYPE_AUTO && currentWidget.GetStyleProperty().Margin.MarginRightValueType == enums.CSS_PROPERTY_VALUE_TYPE_AUTO {
				x += (currentWidget.GetLayout().Parent.Width - currentWidget.GetLayout().Width) / 2
			} else {
				x += currentWidget.GetLayout().MarginLeft
			}
		}
		if currentWidget.GetStyleProperty().Padding != nil {
			x += currentWidget.GetLayout().PaddingLeft
		}
	} else {
		x = beforeCurrentWidget.GetLayout().XPosition + beforeCurrentWidget.GetLayout().Width
		if currentWidget.GetStyleProperty().Margin != nil {
			x += currentWidget.GetLayout().MarginLeft
		}
		if currentWidget.GetStyleProperty().Padding != nil {
			x += currentWidget.GetLayout().PaddingLeft
		}
	}
	return x
	/*
		if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Margin != nil {
			CalculateTopMargin(currentWidget, false)
			CalculateBottomMargin(currentWidget, false)
		}
		if beforeCurrentWidget != nil {
			return beforeCurrentWidget.GetLayout().YPosition + beforeCurrentWidget.GetLayout().GetTotalHeight() + currentWidget.GetLayout().MarginTop
		} else {
			return parent.GetLayout().YPosition + currentWidget.GetLayout().MarginTop
		}

	*/
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

func FlexSetPositionYRelative(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	position := 0
	if beforeCurrentWidget != nil {
		position = beforeCurrentWidget.GetLayout().YPosition + beforeCurrentWidget.GetLayout().Height + int(currentWidget.GetStyleProperty().Top)
	} else {
		position = parent.GetLayout().YPosition + int(currentWidget.GetStyleProperty().Top)
	}
	return position
}
