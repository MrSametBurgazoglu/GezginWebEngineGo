package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func SetInlineFLexContainerWidth(currentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_EMPTY || currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_ROW {
		SetFlexRowContainerChildrenWidth(currentWidget)
	} else {
		SetFlexColumnContainerChildrenWidth(currentWidget)
	}
	SetWidthInline(currentWidget, currentWidget.GetStyleProperty())
}
