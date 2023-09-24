package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func SetPosition(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty() == nil {
		InlineSetPosition(currentWidget, parent, beforeCurrentWidget)
	} else if currentWidget.GetStyleProperty().Parent.Display == enums.CSS_DISPLAY_TYPE_FLEX || currentWidget.GetStyleProperty().Parent.Display == enums.CSS_DISPLAY_TYPE_INLINE_FLEX {
		SetPositionFlex(currentWidget, parent, beforeCurrentWidget)
	} else if currentWidget.GetStyleProperty().Float != enums.CSS_FLOAT_EMPTY && currentWidget.GetStyleProperty().Float != enums.CSS_FLOAT_NONE {
		SetPositionFloat(currentWidget, parent, beforeCurrentWidget)
	} else {
		switch currentWidget.GetStyleProperty().Display {
		case enums.CSS_DISPLAY_TYPE_BLOCK:
			BlockSetPosition(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_DISPLAY_TYPE_INLINE:
			InlineSetPosition(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_DISPLAY_TYPE_FLEX:
			BlockSetPosition(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_DISPLAY_TYPE_INLINE_BLOCK:
			InlineSetPosition(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_DISPLAY_TYPE_INLINE_FLEX:
			InlineSetPosition(currentWidget, parent, beforeCurrentWidget)
		}
	}
}
