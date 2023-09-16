package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func GetTotalChildrenHeight(currentWidget widget.WidgetInterface) int {
	if currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_BLOCK || currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_INLINE_BLOCK ||
		(currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_FLEX && currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_COLUMN) {
		height := 0
		for _, child := range currentWidget.GetChildren() {
			height += child.GetLayout().Height
		}
		return height
	} else if currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_INLINE ||
		(currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_FLEX && (currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_ROW || currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_EMPTY)) {
		currentHeight := 0
		for _, child := range currentWidget.GetChildren() {
			if child.GetLayout().Height > currentHeight {
				currentHeight = child.GetLayout().Height
			}
		}
		return currentHeight
	} else {
		return 0
	}
}

/*TODO ADD STYLE PROPERTY HEIGHT VALUE TO CALCULATE HEIGHT*/
func SetHeight(currentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Height != 0 {
		height := 0
		switch currentWidget.GetStyleProperty().HeightValueType {
		case enums.CSS_PROPERTY_VALUE_TYPE_PIXEL:
			height = int(currentWidget.GetStyleProperty().Height)
		case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
			height = currentWidget.GetLayout().Parent.GetPresetHeight() * int(currentWidget.GetStyleProperty().Height) / 100
		}
		currentWidget.GetLayout().ContentHeight = height
		if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Margin != nil {
			totalHeight := height + (currentWidget.GetStyleProperty().Margin.MarginTop + currentWidget.GetStyleProperty().Margin.MarginBottom)
			currentWidget.GetLayout().Height = totalHeight
		} else {
			currentWidget.GetLayout().Height = height
		}
	} else if len(currentWidget.GetChildren()) > 0 {
		height := GetTotalChildrenHeight(currentWidget)
		currentWidget.GetLayout().ContentHeight = height
		if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Margin != nil {
			totalHeight := height + (currentWidget.GetStyleProperty().Margin.MarginTop + currentWidget.GetStyleProperty().Margin.MarginBottom)
			currentWidget.GetLayout().Height = totalHeight
		} else {
			currentWidget.GetLayout().Height = height
		}
	}
}
