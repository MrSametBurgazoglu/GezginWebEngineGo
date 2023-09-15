package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func SetPosition(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) (int, int) {
	if currentWidget.GetStyleProperty() == nil {
		return InlineSetPosition(currentWidget, parent, beforeCurrentWidget)
	} else if currentWidget.GetStyleProperty().Parent.Display == enums.CSS_DISPLAY_TYPE_FLEX {
		return SetPositionFlex(currentWidget, parent, beforeCurrentWidget)
	} else if currentWidget.GetStyleProperty().Float != enums.CSS_FLOAT_EMPTY && currentWidget.GetStyleProperty().Float != enums.CSS_FLOAT_NONE {
		return SetPositionFloat(currentWidget, parent, beforeCurrentWidget)
	}
	switch currentWidget.GetStyleProperty().Display {
	case enums.CSS_DISPLAY_TYPE_BLOCK:
		return BlockSetPosition(currentWidget, parent, beforeCurrentWidget)
	case enums.CSS_DISPLAY_TYPE_INLINE:
		return InlineSetPosition(currentWidget, parent, beforeCurrentWidget)
	case enums.CSS_DISPLAY_TYPE_FLEX:
		return BlockSetPosition(currentWidget, parent, beforeCurrentWidget)
	case enums.CSS_DISPLAY_TYPE_INLINE_BLOCK:
		return InlineSetPosition(currentWidget, parent, beforeCurrentWidget)
	}
	return 0, 0
}
