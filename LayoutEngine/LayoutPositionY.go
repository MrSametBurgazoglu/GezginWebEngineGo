package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func SetPositionY(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
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
