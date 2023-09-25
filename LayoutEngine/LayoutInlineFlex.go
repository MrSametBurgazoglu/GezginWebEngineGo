package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func SetInlineFLexContainerWidth(currentWidget widget.WidgetInterface) {
	width := LookForWidth(currentWidget.GetLayout())
	currentWidget.GetLayout().Width = width
	currentWidget.GetLayout().ContentWidth = width
	if currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_EMPTY || currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_ROW {
		SetFlexRowContainerChildrenWidth(currentWidget)
	} else {
		SetFlexColumnContainerChildrenWidth(currentWidget)
	}
	//SetWidthInline(currentWidget, currentWidget.GetStyleProperty())
}

func SetPositionInlineFlex(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if !(currentWidget.GetStyleProperty().Parent.Display == enums.CSS_DISPLAY_TYPE_FLEX || currentWidget.GetStyleProperty().Parent.Display == enums.CSS_DISPLAY_TYPE_INLINE_FLEX) {
		InlineSetPosition(currentWidget, parent, beforeCurrentWidget)
	}
	SetFlexContainerChildrenPosition(currentWidget)
}
