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

func SetFLexContainerWidthRowNoWrap(currentWidget widget.WidgetInterface) {
	totalWidth := currentWidget.GetLayout().ContentWidth
	childWidth := totalWidth / currentWidget.GetChildrenCount()
	for _, child := range currentWidget.GetChildren() {
		/*
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
		*/
		child.GetLayout().Width = childWidth
		child.GetLayout().ContentWidth = childWidth
	}
}

func SetFLexContainerWidthRowWrap(currentWidget widget.WidgetInterface) {
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
		child.GetLayout().Width = width
		child.GetLayout().ContentWidth = width
	}
}

func SetFLexContainerWidthRow(currentWidget widget.WidgetInterface) {
	SetWidthBlock(currentWidget, currentWidget.GetParent())
	switch currentWidget.GetStyleProperty().FlexWrap {
	case enums.CSS_FLEX_WRAP_NOWRAP:
		SetFLexContainerWidthRowNoWrap(currentWidget)
	case enums.CSS_FLEX_WRAP_WRAP:
		SetFLexContainerWidthRowWrap(currentWidget)
	case enums.CSS_FLEX_WRAP_EMPTY:
		SetFLexContainerWidthRowNoWrap(currentWidget)
	}
	/*
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
			child.GetLayout().Width = width
			child.GetLayout().ContentWidth = width
			totalWidth += width
		}

	*/
}

/*TOOO IF WRAP THEN ITS MEANS WE DON'T FIT THEM INTO ONE LINE WE USE THEIR WIDTHS*/
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
		child.GetLayout().ContentWidth = width
	}
}

func SetPositionFlex(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty().Parent.FlexDirection == enums.CSS_FLEX_DIRECTION_EMPTY || currentWidget.GetStyleProperty().Parent.FlexDirection == enums.CSS_FLEX_DIRECTION_ROW {
		SetPositionXFlex(currentWidget, parent, beforeCurrentWidget)
		SetPositionYFlex(currentWidget, parent, beforeCurrentWidget)
	} else {
		BlockSetPositionX(currentWidget, parent)
		BlockSetPositionY(currentWidget, parent, beforeCurrentWidget)
	}
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
