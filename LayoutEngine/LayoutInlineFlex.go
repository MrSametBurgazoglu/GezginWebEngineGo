package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func SetInlineFLexContainerWidth(currentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_EMPTY || currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_ROW {
		SetFlexRowContainerChildrenSizeAndPosition(currentWidget)
	} else {
		SetFlexColumnContainerChildrenSizeAndPosition(currentWidget)
	}
	SetWidthInline(currentWidget, currentWidget.GetStyleProperty())
}
