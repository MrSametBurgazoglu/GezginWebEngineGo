package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func InlineSetPosition(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	InlineSetPositionX(currentWidget, parent, beforeCurrentWidget)
	InlineSetPositionY(currentWidget, parent, beforeCurrentWidget)
}

func InlineSetPositionXStatic(currentWidget widget.WidgetInterface) {
	if currentWidget.GetParent().GetStyleProperty() != nil && currentWidget.GetParent().GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_BLOCK && currentWidget.GetParent().GetStyleProperty().TextAlign == enums.CSS_TEXT_ALIGN_CENTER {
		childrenTotalWidth := 0
		currentParent := currentWidget.GetParent()
		for _, widgetInterface := range currentParent.GetChildren() {
			if widgetInterface.GetStyleProperty() == nil || widgetInterface.GetStyleProperty().Display != enums.CSS_DISPLAY_TYPE_BLOCK {
				childrenTotalWidth += widgetInterface.GetLayout().GetTotalWidth()
			} else {
				childrenTotalWidth = 0
			}
		}
		startPoint := currentParent.GetLayout().ContentXPosition + currentParent.GetLayout().ContentWidth/2 + childrenTotalWidth/2
		currentStartPoint := startPoint
		for _, widgetInterface := range currentParent.GetChildren()[currentWidget.GetChildrenIndex():] {
			if widgetInterface.GetStyleProperty() == nil || widgetInterface.GetStyleProperty().Display != enums.CSS_DISPLAY_TYPE_BLOCK {
				currentStartPoint -= widgetInterface.GetLayout().GetTotalWidth()
			} else {
				currentStartPoint = startPoint
			}
		}
		currentWidget.GetLayout().XPosition = currentStartPoint
		currentWidget.GetLayout().ContentXPosition = currentWidget.GetLayout().XPosition + currentWidget.GetLayout().PaddingLeft
	} else {
		if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Margin != nil {
			CalculateLeftMargin(currentWidget, true)
			CalculateRightMargin(currentWidget, true)
		}
		if currentWidget.GetStyleProperty() != nil && currentWidget.GetStyleProperty().Padding != nil {
			currentWidget.GetLayout().PaddingLeft = int(currentWidget.GetStyleProperty().Padding.PaddingLeft)
			currentWidget.GetLayout().PaddingRight = int(currentWidget.GetStyleProperty().Padding.PaddingRight)
			currentWidget.GetLayout().PaddingTop = int(currentWidget.GetStyleProperty().Padding.PaddingTop)
			currentWidget.GetLayout().PaddingBottom = int(currentWidget.GetStyleProperty().Padding.PaddingBottom)
		}
		currentWidget.GetLayout().XPosition = currentWidget.GetParent().GetLayout().ContentXPosition + currentWidget.GetLayout().MarginLeft
		currentWidget.GetLayout().ContentXPosition = currentWidget.GetLayout().XPosition
	}
}

func InlineSetPositionYStatic(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if beforeCurrentWidget == nil || beforeCurrentWidget.GetStyleProperty() == nil {
		marginTop := 0
		if currentWidget.GetStyleProperty().Margin != nil {
			marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
		}
		currentWidget.GetLayout().YPosition = parent.GetLayout().YPosition + marginTop
		currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
	} else if beforeCurrentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_INLINE || beforeCurrentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_INLINE_FLEX {
		marginTop := 0
		if currentWidget.GetStyleProperty().Margin != nil {
			marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
		}
		currentWidget.GetLayout().YPosition = beforeCurrentWidget.GetLayout().YPosition + marginTop
		currentWidget.GetLayout().ContentYPosition = beforeCurrentWidget.GetLayout().YPosition
	} else if beforeCurrentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_BLOCK {
		marginTop := 0
		if currentWidget.GetStyleProperty().Margin != nil {
			marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
		}
		currentWidget.GetLayout().YPosition = beforeCurrentWidget.GetLayout().YPosition + beforeCurrentWidget.GetLayout().Height + marginTop
		currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
	}
}

func InlineSetPositionX(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			currentWidget.GetLayout().XPosition = parent.GetLayout().ContentXPosition
			currentWidget.GetLayout().ContentXPosition = parent.GetLayout().ContentXPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			if beforeCurrentWidget != nil && beforeCurrentWidget.GetStyleProperty() != nil && beforeCurrentWidget.GetStyleProperty().Display != enums.CSS_DISPLAY_TYPE_BLOCK {
				currentWidget.GetLayout().XPosition = beforeCurrentWidget.GetLayout().XPosition + beforeCurrentWidget.GetLayout().Width
				currentWidget.GetLayout().ContentXPosition = currentWidget.GetLayout().XPosition
			} else {
				InlineSetPositionXStatic(currentWidget)
			}
		case enums.CSS_POSITION_TYPE_STATIC:
			if beforeCurrentWidget != nil {
				currentWidget.GetLayout().XPosition = beforeCurrentWidget.GetLayout().XPosition + beforeCurrentWidget.GetLayout().Width
				currentWidget.GetLayout().ContentXPosition = currentWidget.GetLayout().XPosition
			} else {
				InlineSetPositionXStatic(currentWidget)
			}
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			break
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			break
		}
	} else {
		if beforeCurrentWidget != nil {
			currentWidget.GetLayout().XPosition = beforeCurrentWidget.GetLayout().XPosition + beforeCurrentWidget.GetLayout().Width
			currentWidget.GetLayout().ContentXPosition = currentWidget.GetLayout().XPosition
		} else {
			InlineSetPositionXStatic(currentWidget)
		}
	}
}

func InlineSetPositionY(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			println("not implemented yet")
		case enums.CSS_POSITION_TYPE_EMPTY:
			InlineSetPositionYStatic(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_STATIC:
			InlineSetPositionYStatic(currentWidget, parent, beforeCurrentWidget)
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			println("not implemented yet")
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			println("not implemented yet")
		}
	} else {
		if beforeCurrentWidget == nil {
			currentWidget.GetLayout().YPosition = parent.GetLayout().ContentYPosition
			currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
		} else {
			if beforeCurrentWidget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_INLINE {
				currentWidget.GetLayout().YPosition = parent.GetLayout().YPosition
				currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
			} else {
				currentWidget.GetLayout().YPosition = beforeCurrentWidget.GetLayout().Height + beforeCurrentWidget.GetLayout().YPosition
				currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
			}
		}
	}
}

func SetWidthInline(currentWidget widget.WidgetInterface, styleProperty *StyleProperty.StyleProperty) {
	if len(currentWidget.GetChildren()) > 0 {
		width := 0
		for _, child := range currentWidget.GetChildren() {
			width += child.GetLayout().GetTotalWidth()
		}
		if styleProperty != nil && styleProperty.Padding != nil {
			currentWidget.GetLayout().PaddingLeft = int(styleProperty.Padding.PaddingLeft)
			currentWidget.GetLayout().PaddingRight = int(styleProperty.Padding.PaddingRight)
			width += currentWidget.GetLayout().PaddingLeft + currentWidget.GetLayout().PaddingRight
		}
		currentWidget.GetLayout().Width = width
		currentWidget.GetLayout().ContentWidth = width
	}
}
