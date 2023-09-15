package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func InlineSetPosition(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) (int, int) {
	return InlineSetPositionX(currentWidget, parent, beforeCurrentWidget), InlineSetPositionY(currentWidget, parent, beforeCurrentWidget)
}

func InlineSetPositionXStatic(currentWidget widget.WidgetInterface) int {
	if currentWidget.GetParent().GetStyleProperty() != nil && currentWidget.GetParent().GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_BLOCK && currentWidget.GetParent().GetStyleProperty().TextAlign == enums.CSS_TEXT_ALIGN_CENTER {
		childrenTotalWidth := 0
		currentParent := currentWidget.GetParent()
		for _, widgetInterface := range currentParent.GetChildren() {
			if widgetInterface.GetStyleProperty() == nil || widgetInterface.GetStyleProperty().Display != enums.CSS_DISPLAY_TYPE_BLOCK {
				childrenTotalWidth += widgetInterface.GetLayout().Width
			} else {
				childrenTotalWidth = 0
			}
		}
		startPoint := currentParent.GetLayout().XPosition + currentParent.GetLayout().Width/2 + childrenTotalWidth/2
		currentStartPoint := startPoint
		for _, widgetInterface := range currentParent.GetChildren()[currentWidget.GetChildrenIndex():] {
			if widgetInterface.GetStyleProperty() == nil || widgetInterface.GetStyleProperty().Display != enums.CSS_DISPLAY_TYPE_BLOCK {
				currentStartPoint -= widgetInterface.GetLayout().Width
			} else {
				currentStartPoint = startPoint
			}
		}
		return currentStartPoint
	}
	return currentWidget.GetParent().GetLayout().ContentXPosition
}

func InlineSetPositionX(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	position := 0
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			position = parent.GetLayout().ContentXPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			if beforeCurrentWidget != nil {
				position = beforeCurrentWidget.GetLayout().XPosition + beforeCurrentWidget.GetLayout().Width
			} else {
				position = InlineSetPositionXStatic(currentWidget)
			}
		case enums.CSS_POSITION_TYPE_STATIC:
			if beforeCurrentWidget != nil {
				position = beforeCurrentWidget.GetLayout().XPosition + beforeCurrentWidget.GetLayout().Width
			} else {
				position = InlineSetPositionXStatic(currentWidget)
			}
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			if currentWidget.GetStyleProperty().Left != 0 {
				position = parent.GetLayout().ContentXPosition + int(currentWidget.GetStyleProperty().Left)
			} else if currentWidget.GetStyleProperty().Right != 0 {
				position = parent.GetLayout().ContentWidth - int(currentWidget.GetStyleProperty().Right)
			} else {
				position = parent.GetLayout().ContentXPosition
			}
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			if currentWidget.GetStyleProperty().Left != 0 {
				position = parent.GetLayout().ContentXPosition + int(currentWidget.GetStyleProperty().Left)
			} else if currentWidget.GetStyleProperty().Right != 0 {
				position = parent.GetLayout().ContentWidth - int(currentWidget.GetStyleProperty().Right)
			} else {
				position = parent.GetLayout().ContentXPosition
			}
		}
	} else {
		if beforeCurrentWidget != nil {
			position = beforeCurrentWidget.GetLayout().XPosition + beforeCurrentWidget.GetLayout().Width
		} else {
			position = InlineSetPositionXStatic(currentWidget)
		}
	}
	currentWidget.GetLayout().ContentXPosition = position
	currentWidget.GetLayout().XPosition = position
	return currentWidget.GetLayout().ContentXPosition
}

func InlineSetPositionY(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			return parent.GetLayout().XPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			if beforeCurrentWidget == nil || beforeCurrentWidget.GetStyleProperty() == nil {
				marginTop := 0
				if currentWidget.GetStyleProperty().Margin != nil {
					marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
				}
				return parent.GetLayout().YPosition + marginTop
			} else if beforeCurrentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_INLINE {
				marginTop := 0
				if currentWidget.GetStyleProperty().Margin != nil {
					marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
				}
				return parent.GetLayout().YPosition + marginTop
			} else if beforeCurrentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_BLOCK {
				marginTop := 0
				if currentWidget.GetStyleProperty().Margin != nil {
					marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
				}
				return beforeCurrentWidget.GetLayout().YPosition + beforeCurrentWidget.GetLayout().Height + marginTop
			}
		case enums.CSS_POSITION_TYPE_STATIC:
			marginTop := 0
			if currentWidget.GetStyleProperty().Margin != nil {
				marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
			}
			if beforeCurrentWidget != nil {
				return beforeCurrentWidget.GetLayout().YPosition + beforeCurrentWidget.GetLayout().Height + marginTop
			} else {
				return parent.GetLayout().YPosition + parent.GetLayout().Height + marginTop
			}
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			if currentWidget.GetStyleProperty().Top != 0 {
				return parent.GetLayout().YPosition + int(currentWidget.GetStyleProperty().Top)
			} else if currentWidget.GetStyleProperty().Bottom != 0 {
				return parent.GetLayout().YPosition + parent.GetLayout().Height - int(currentWidget.GetStyleProperty().Bottom)
			} else {
				return parent.GetLayout().YPosition + parent.GetLayout().Height
			}
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			if beforeCurrentWidget != nil {
				return beforeCurrentWidget.GetLayout().YPosition + beforeCurrentWidget.GetLayout().Height + int(currentWidget.GetStyleProperty().Top)
			} else {
				return parent.GetLayout().YPosition + int(currentWidget.GetStyleProperty().Top)
			}
		}
	} else {
		if beforeCurrentWidget == nil {
			return parent.GetLayout().YPosition
		} else { //TODO look before current widget if it is block then before height + y
			if beforeCurrentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_INLINE {
				return parent.GetLayout().YPosition
			} else {
				return beforeCurrentWidget.GetLayout().Height + beforeCurrentWidget.GetLayout().YPosition
			}
		}
	}
	return 0
}

func SetWidthInline(currentWidget widget.WidgetInterface, styleProperty *StyleProperty.StyleProperty) {
	if len(currentWidget.GetChildren()) > 0 {
		width := 0
		for _, child := range currentWidget.GetChildren() {
			width += child.GetLayout().Width
		}
		currentWidget.GetLayout().Width = width
		currentWidget.GetLayout().ContentWidth = width
		if styleProperty != nil && styleProperty.Margin != nil {
			contentWidth := width - (styleProperty.Margin.MarginLeft + styleProperty.Margin.MarginRight)
			currentWidget.GetLayout().ContentWidth = contentWidth
		}
	}
}
