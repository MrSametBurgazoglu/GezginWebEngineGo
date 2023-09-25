package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func GetTotalChildrenHeight(currentWidget widget.WidgetInterface) int {
	if currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_INLINE ||
		(currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_FLEX && (currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_ROW || currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_EMPTY && currentWidget.GetStyleProperty().FlexWrap != enums.CSS_FLEX_WRAP_WRAP)) {
		currentHeight := 0
		for _, child := range currentWidget.GetChildren() {
			if child.GetLayout().Height > currentHeight {
				currentHeight = child.GetLayout().Height
			}
		}
		return currentHeight
	} else if currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_FLEX && currentWidget.GetStyleProperty().FlexWrap == enums.CSS_FLEX_WRAP_WRAP && (currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_ROW || currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_EMPTY) {
		parentWidth := currentWidget.GetLayout().ContentWidth
		currentWidth := 0
		maxHeight := 0
		currentHeight := 0
		for _, widgetInterface := range currentWidget.GetChildren() {
			currentWidth += widgetInterface.GetLayout().Width
			if currentWidth >= parentWidth {
				currentHeight += maxHeight
				maxHeight = 0
				currentWidth = 0
			}
			if widgetInterface.GetLayout().Height > maxHeight {
				maxHeight = widgetInterface.GetLayout().Height
			}
		}
		currentHeight += maxHeight
		return currentHeight
	} else {
		lastChild := currentWidget.GetChildrenByIndex(currentWidget.GetChildrenCount() - 1)
		height := lastChild.GetLayout().Height
		for _, widgetInterface := range currentWidget.GetChildren()[1:] {
			if lastChild.GetStyleProperty() != nil && (lastChild.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_BLOCK || lastChild.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_INLINE_BLOCK) {
				height += widgetInterface.GetLayout().Height
			} else if lastChild.GetStyleProperty() == nil || widgetInterface.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_INLINE ||
				(widgetInterface.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_FLEX && (widgetInterface.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_ROW || currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_EMPTY)) {
				if widgetInterface.GetLayout().Height > lastChild.GetLayout().Height {
					height += widgetInterface.GetLayout().Height - lastChild.GetLayout().Height
				}
			}
		}
		return height
	}
	/*
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

	*/
}

/*TODO ADD STYLE PROPERTY HEIGHT VALUE TO CALCULATE HEIGHT*/
func SetHeight(currentWidget widget.WidgetInterface) {
	if classes := currentWidget.GetClasses(); len(classes) == 2 && classes[1] == "shadow-sm" {
		println("hey")
	}
	if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_NONE {
		currentWidget.GetLayout().Height = 0
		currentWidget.GetLayout().ContentHeight = 0
	} else if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Height != 0 {
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
