package LayoutEngine

import (
	"gezgin_web_engine/LayoutProperty"
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func LookForWidth(layoutProperty *LayoutProperty.LayoutProperty) int {
	if len(layoutProperty.Children) == 0 {
		return layoutProperty.Width
	} else {
		maxWidth := 0
		for _, child := range layoutProperty.Children {
			currentWidth := LookForWidth(child)
			if currentWidth > maxWidth {
				maxWidth = currentWidth
			}
		}
		return maxWidth
	}
}

func LookForHeight(layoutProperty *LayoutProperty.LayoutProperty) int {
	if len(layoutProperty.Children) == 0 {
		return layoutProperty.Height
	} else {
		maxHeight := 0
		for _, child := range layoutProperty.Children {
			currentHeight := LookForHeight(child)
			if currentHeight > maxHeight {
				maxHeight = currentHeight
			}
		}
		return maxHeight
	}
}

func SetFLexContainerWidth(currentWidget widget.WidgetInterface) {
	SetWidthBlock(currentWidget, currentWidget.GetParent())
	if currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_EMPTY || currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_ROW {
		SetFlexRowContainerChildrenWidth(currentWidget)
	} else {
		SetFlexColumnContainerChildrenWidth(currentWidget)
	}
}

func GetTotalWidthOfChildren(currentWidget widget.WidgetInterface) (int, []int) {
	totalChildrenWidth := 0
	var widths []int
	for i, child := range currentWidget.GetChildren() {
		width := LookForWidth(child.GetLayout())
		if currentWidget.GetStyleProperty().Children != nil && currentWidget.GetStyleProperty().Children[i].Width != 0 {
			switch currentWidget.GetStyleProperty().Children[i].WidthValueType {
			case enums.CSS_PROPERTY_VALUE_TYPE_PIXEL:
				width = int(currentWidget.GetStyleProperty().Children[i].Width)
			case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
				width = int(float64(currentWidget.GetLayout().Width) * (float64(currentWidget.GetStyleProperty().Children[i].Width) / 100.0))
			}
			if currentWidget.GetStyleProperty().Children[i].MaxWidth != 0 && width > int(currentWidget.GetStyleProperty().Children[i].MaxWidth) {
				width = int(currentWidget.GetStyleProperty().Children[i].MaxWidth)
			}
			if width < int(currentWidget.GetStyleProperty().Children[i].MinWidth) {
				width = int(currentWidget.GetStyleProperty().Children[i].MinWidth)
			}
		}
		widths = append(widths, width)
		totalChildrenWidth += width
		//child.GetLayout().Width = childWidth
		//child.GetLayout().ContentWidth = childWidth
	}
	return totalChildrenWidth, widths
}

func SetFlexContainerHeight(currentWidget widget.WidgetInterface) {
	parentWidth := currentWidget.GetLayout().ContentWidth
	currentWidth := 0
	maxHeight := 0
	totalHeight := 0
	for _, child := range currentWidget.GetLayout().Children {
		if child.Height > maxHeight {
			maxHeight = child.Height
		}
		if currentWidth > parentWidth {
			totalHeight += maxHeight
			currentWidth = 0
			maxHeight = 0
		}
	}
}

func SetPositionFlex(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	BlockSetPosition(currentWidget, parent, beforeCurrentWidget)
	SetFlexContainerChildrenPosition(currentWidget)
}

func SetPositionYFlex(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			FlexSetPositionYSticky(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_EMPTY:
			FlexSetPositionYStatic(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_STATIC:
			FlexSetPositionYStatic(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			FlexSetPositionYAbsolute(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_FIXED:
			FlexSetPositionYFixed(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_RELATIVE:
			FlexSetPositionYRelative(currentWidget, parent, beforeCurrentWidget)
		}
	} else {
		if beforeCurrentWidget == nil {
			currentWidget.GetLayout().YPosition = parent.GetLayout().ContentYPosition
			currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
		} else {
			currentWidget.GetLayout().YPosition = beforeCurrentWidget.GetLayout().ContentYPosition
			currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
		}
	}
}

func SetWidthFlexChild(currentWidget widget.WidgetInterface, styleProperty *StyleProperty.StyleProperty) {
	//you must set childrens width first
	if len(currentWidget.GetChildren()) > 0 {
		width := 0
		for _, child := range currentWidget.GetChildren() {
			width += child.GetLayout().Width
		}
		contentWidth := width
		if styleProperty != nil && styleProperty.Margin != nil {
			contentWidth = width - (styleProperty.Margin.MarginLeft + styleProperty.Margin.MarginRight)
		}
		currentWidget.GetLayout().Width = width
		currentWidget.GetLayout().ContentWidth = contentWidth
	}
}
