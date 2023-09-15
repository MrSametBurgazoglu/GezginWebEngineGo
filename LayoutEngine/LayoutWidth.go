package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func SetWidth(currentWidget widget.WidgetInterface) int {
	if currentWidget.GetStyleProperty() == nil {
		SetWidthInline(currentWidget, currentWidget.GetStyleProperty())
	} else if currentWidget.GetStyleProperty().Float != enums.CSS_FLOAT_EMPTY && currentWidget.GetStyleProperty().Float != enums.CSS_FLOAT_NONE {
		SetFloatWidth(currentWidget)
	} else if currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_FLEX {
		SetFLexContainerWidth(currentWidget)
	} else if currentWidget.GetStyleProperty().Parent.Display == enums.CSS_DISPLAY_TYPE_FLEX {
		println("heyyo")
	} else if currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_BLOCK {
		SetWidthBlock(currentWidget, currentWidget.GetParent())
	} else if currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_FLEX {
		SetWidthBlock(currentWidget, currentWidget.GetParent())
	} else if currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_INLINE_BLOCK {
		SetWidthInline(currentWidget, currentWidget.GetStyleProperty())
	} else {
		SetWidthInline(currentWidget, currentWidget.GetStyleProperty())
	}
	return currentWidget.GetLayout().Width
}
