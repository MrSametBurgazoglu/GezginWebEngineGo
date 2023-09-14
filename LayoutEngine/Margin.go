package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func GetAlignedChildren(currentWidget widget.WidgetInterface, containerAlignSelf bool) []widget.WidgetInterface {
	var widgetList []widget.WidgetInterface
	if containerAlignSelf {
		widgetList = append(widgetList, currentWidget)
	} else {
		widgetList = append(widgetList, currentWidget.GetParent().GetChildren()...)
	}
	return widgetList
}

func GetHorizontalMarginAutoCountAndRemainWidth(currentWidget widget.WidgetInterface, widgetList []widget.WidgetInterface) int {
	parentContentWidth := currentWidget.GetParent().GetLayout().ContentWidth
	autoCount := 0
	for _, childWidgetInterface := range widgetList {
		parentContentWidth -= childWidgetInterface.GetLayout().GetTotalWidth()
		if childWidgetInterface.GetStyleProperty().Margin.MarginLeftValueType == enums.CSS_PROPERTY_VALUE_TYPE_AUTO {
			autoCount += 1
		}
		if childWidgetInterface.GetStyleProperty().Margin.MarginRightValueType == enums.CSS_PROPERTY_VALUE_TYPE_AUTO {
			autoCount += 1
		}
	}
	return parentContentWidth / autoCount
}

func GetVerticalMarginAutoCountAndRemainHeight(currentWidget widget.WidgetInterface, widgetList []widget.WidgetInterface) int {
	parentContentHeight := currentWidget.GetParent().GetLayout().ContentHeight
	autoCount := 0
	for _, childWidgetInterface := range widgetList {
		parentContentHeight -= childWidgetInterface.GetLayout().GetTotalContentHeight()
		if childWidgetInterface.GetStyleProperty().Margin != nil && childWidgetInterface.GetStyleProperty().Margin.MarginTopValueType == enums.CSS_PROPERTY_VALUE_TYPE_AUTO {
			autoCount += 1
		}
		if childWidgetInterface.GetStyleProperty().Margin != nil && childWidgetInterface.GetStyleProperty().Margin.MarginBottomValueType == enums.CSS_PROPERTY_VALUE_TYPE_AUTO {
			autoCount += 1
		}
	}
	return parentContentHeight / autoCount
}

func CalculateLeftMargin(currentWidget widget.WidgetInterface, containerAlignSelf bool) {
	if currentWidget.GetStyleProperty().Margin.MarginLeftValueType != enums.CSS_PROPERTY_VALUE_TYPE_AUTO {
		currentWidget.GetLayout().MarginLeft = currentWidget.GetStyleProperty().Margin.MarginLeft
	} else {
		widgetList := GetAlignedChildren(currentWidget, containerAlignSelf)
		currentWidget.GetLayout().MarginLeft = GetHorizontalMarginAutoCountAndRemainWidth(currentWidget, widgetList)
	}
}

func CalculateRightMargin(currentWidget widget.WidgetInterface, containerAlignSelf bool) {
	if currentWidget.GetStyleProperty().Margin.MarginRightValueType != enums.CSS_PROPERTY_VALUE_TYPE_AUTO {
		currentWidget.GetLayout().MarginRight = currentWidget.GetStyleProperty().Margin.MarginRight
	} else {
		widgetList := GetAlignedChildren(currentWidget, containerAlignSelf)
		currentWidget.GetLayout().MarginRight = GetHorizontalMarginAutoCountAndRemainWidth(currentWidget, widgetList)
	}
}

func CalculateTopMargin(currentWidget widget.WidgetInterface, containerAlignSelf bool) {
	if currentWidget.GetStyleProperty().Margin.MarginTopValueType != enums.CSS_PROPERTY_VALUE_TYPE_AUTO {
		currentWidget.GetLayout().MarginTop = currentWidget.GetStyleProperty().Margin.MarginTop
	} else {
		widgetList := GetAlignedChildren(currentWidget, containerAlignSelf)
		currentWidget.GetLayout().MarginTop = GetVerticalMarginAutoCountAndRemainHeight(currentWidget, widgetList)
	}
}

func CalculateBottomMargin(currentWidget widget.WidgetInterface, containerAlignSelf bool) {
	if currentWidget.GetStyleProperty().Margin.MarginBottomValueType != enums.CSS_PROPERTY_VALUE_TYPE_AUTO {
		currentWidget.GetLayout().MarginBottom = currentWidget.GetStyleProperty().Margin.MarginBottom
	} else {
		widgetList := GetAlignedChildren(currentWidget, containerAlignSelf)
		currentWidget.GetLayout().MarginBottom = GetVerticalMarginAutoCountAndRemainHeight(currentWidget, widgetList)
	}
}
