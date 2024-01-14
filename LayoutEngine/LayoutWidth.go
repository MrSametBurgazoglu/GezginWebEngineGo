package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func SetWidth(currentWidget widget.WidgetInterface) {
	/*
		if classes := currentWidget.GetClasses(); len(classes) > 0 && classes[0] == "p-4" {
			print("hey")
		}
		if classes := currentWidget.GetClasses(); len(classes) == 4 && classes[2] == "g-lg-5" {
			print("hey")
		}
		if currentWidget.GetHtmlTag() == int(HtmlParser.HTML_FORM) || currentWidget.GetHtmlTag() == int(HtmlParser.HTML_IMG) {
			println("hey")
		}

	*/
	if currentWidget.GetStyleProperty() == nil {
		SetWidthInline(currentWidget, currentWidget.GetStyleProperty())
	} else if currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_NONE {
		currentWidget.GetLayout().Width = 0
		currentWidget.GetLayout().ContentWidth = 0
	} else if currentWidget.GetStyleProperty().Float != enums.CSS_FLOAT_EMPTY && currentWidget.GetStyleProperty().Float != enums.CSS_FLOAT_NONE {
		SetFloatWidth(currentWidget)
	} else if currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_FLEX {
		SetFLexContainerWidth(currentWidget)
	} else if currentWidget.GetStyleProperty().Parent.Display == enums.CSS_DISPLAY_TYPE_FLEX && currentWidget.GetStyleProperty().Display != enums.CSS_DISPLAY_TYPE_FLEX && currentWidget.GetStyleProperty().Display != enums.CSS_DISPLAY_TYPE_INLINE_FLEX {
		//println("heyyo")
	} else if currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_BLOCK {
		SetWidthBlock(currentWidget, currentWidget.GetParent())
	} else if currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_INLINE_BLOCK {
		SetWidthInlineBlock(currentWidget, currentWidget.GetStyleProperty())
	} else if currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_INLINE_FLEX {
		SetInlineFLexContainerWidth(currentWidget)
	} else {
		SetWidthInline(currentWidget, currentWidget.GetStyleProperty())
	}
}
