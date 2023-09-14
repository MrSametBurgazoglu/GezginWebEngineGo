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

func SetFLexContainerWidth(currentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_EMPTY || currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_ROW {
		SetFLexContainerWidthRow(currentWidget)
	} else {
		SetFLexContainerWidthColumn(currentWidget)
	}
}

func SetFLexContainerWidthRow(currentWidget widget.WidgetInterface) {
	SetWidthBlock(currentWidget, currentWidget.GetParent())
	totalWidth := 0
	for i, child := range currentWidget.GetChildren() {
		width := LookForWidth(child.GetLayout())
		if currentWidget.GetStyleProperty().Children != nil && currentWidget.GetStyleProperty().Children[i].Width != 0 {
			switch currentWidget.GetStyleProperty().Children[i].WidthValueType {
			case enums.CSS_PROPERTY_VALUE_TYPE_PIXEL:
				width = int(currentWidget.GetStyleProperty().Children[i].Width)
			case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
				width = int(float64(currentWidget.GetLayout().Width) * (float64(currentWidget.GetStyleProperty().Children[i].Width) / 100.0))
			}
			//look width here
			if currentWidget.GetStyleProperty().Children[i].MaxWidth != 0 && width > int(currentWidget.GetStyleProperty().Children[i].MaxWidth) {
				width = int(currentWidget.GetStyleProperty().Children[i].MaxWidth)
			}
			if width < int(currentWidget.GetStyleProperty().Children[i].MinWidth) {
				width = int(currentWidget.GetStyleProperty().Children[i].MinWidth)
			}
		}
		child.GetLayout().Width = width
		totalWidth += width
	}
	if totalWidth > currentWidget.GetLayout().Width {
		currentWidget.GetLayout().Width = totalWidth
		currentWidget.GetLayout().ContentWidth = totalWidth
	}
}

func SetFLexContainerWidthColumn(currentWidget widget.WidgetInterface) {
	SetWidthBlock(currentWidget, currentWidget.GetParent())
	for i, child := range currentWidget.GetChildren() {
		width := currentWidget.GetLayout().Width
		if currentWidget.GetStyleProperty().Children[i].Width != 0 {
			switch currentWidget.GetStyleProperty().Children[i].WidthValueType {
			case enums.CSS_PROPERTY_VALUE_TYPE_PIXEL:
				width = int(currentWidget.GetStyleProperty().Children[i].Width)
			case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
				width = int(float64(currentWidget.GetLayout().Width) * (float64(currentWidget.GetStyleProperty().Children[i].Width) / 100.0))
			}
		}
		//look width here
		if currentWidget.GetStyleProperty().Children[i].MaxWidth != 0 && width > int(currentWidget.GetStyleProperty().Children[i].MaxWidth) {
			width = int(currentWidget.GetStyleProperty().Children[i].MaxWidth)
		}
		if width < int(currentWidget.GetStyleProperty().Children[i].MinWidth) {
			width = int(currentWidget.GetStyleProperty().Children[i].MinWidth)
		}
		child.GetLayout().Width = width
	}
}

func SetPositionFlex(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) (int, int) {
	if currentWidget.GetStyleProperty().Parent.FlexDirection == enums.CSS_FLEX_DIRECTION_EMPTY || currentWidget.GetStyleProperty().Parent.FlexDirection == enums.CSS_FLEX_DIRECTION_ROW {
		return SetPositionXFlex(currentWidget, parent, beforeCurrentWidget), SetPositionYFlex(currentWidget, parent, beforeCurrentWidget)
	} else {
		return BlockSetPositionX(currentWidget, parent), BlockSetPositionY(currentWidget, parent, beforeCurrentWidget)
	}
}

func SetPositionYFlex(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	position := 0
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			position = FlexSetPositionYSticky(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_EMPTY:
			println("hello")
			println(currentWidget.GetParent().GetStyleProperty().FlexDirection)
			//position = FlexSetPositionYStatic(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_STATIC:
			position = FlexSetPositionYStatic(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			position = FlexSetPositionYAbsolute(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_FIXED:
			position = FlexSetPositionYFixed(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_RELATIVE:
			position = FlexSetPositionYRelative(currentWidget, parent, beforeCurrentWidget)
		}
	} else {
		if beforeCurrentWidget == nil {
			position = parent.GetLayout().YPosition
		} else {
			position = beforeCurrentWidget.GetLayout().YPosition
		}
	}
	return position
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
