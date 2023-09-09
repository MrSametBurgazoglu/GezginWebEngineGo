package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func SetFloatWidth(currentWidget widget.WidgetInterface) {
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

func SetPositionFloat(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) (int, int) {
	return SetPositionXFloat(currentWidget, parent, beforeCurrentWidget), SetPositionYFloat(currentWidget, parent, beforeCurrentWidget)
}

func SetPositionXFloat(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	position := 0
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			position = parent.GetLayout().ContentXPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			x := 0
			if beforeCurrentWidget == nil {
				if currentWidget.GetStyleProperty().Float == enums.CSS_FLOAT_LEFT {
					x = parent.GetLayout().XPosition
					if currentWidget.GetStyleProperty().Margin != nil {
						x += currentWidget.GetLayout().MarginLeft
					}
					if currentWidget.GetStyleProperty().Padding != nil {
						x += currentWidget.GetLayout().PaddingLeft
					}
				} else {
					x = parent.GetLayout().XPosition + parent.GetLayout().Width
					if currentWidget.GetStyleProperty().Margin != nil {
						x -= currentWidget.GetLayout().MarginRight
					}
					if currentWidget.GetStyleProperty().Padding != nil {
						x -= currentWidget.GetLayout().PaddingRight
					}
				}
			} else {
				index := currentWidget.GetChildrenIndex()
				if currentWidget.GetStyleProperty().Float == enums.CSS_FLOAT_LEFT {
					x = parent.GetLayout().XPosition
				} else {
					x = parent.GetLayout().XPosition + parent.GetLayout().Width - currentWidget.GetLayout().Width
				}
				for i := index - 1; i >= 0; i-- {
					if currentWidget.GetParent().GetChildrenByIndex(i).GetStyleProperty().Float == currentWidget.GetStyleProperty().Float {
						if currentWidget.GetStyleProperty().Float == enums.CSS_FLOAT_LEFT {
							x = currentWidget.GetParent().GetChildrenByIndex(i).GetLayout().XPosition + currentWidget.GetParent().GetChildrenByIndex(i).GetLayout().Width
						} else {
							x = currentWidget.GetParent().GetChildrenByIndex(i).GetLayout().XPosition - currentWidget.GetLayout().Width
						}
						break
					}
				}
			}
			position = x

		case enums.CSS_POSITION_TYPE_STATIC:
			position = parent.GetLayout().ContentXPosition
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
		position = parent.GetLayout().ContentXPosition
	}
	currentWidget.GetLayout().ContentXPosition = position
	currentWidget.GetLayout().XPosition = position
	return currentWidget.GetLayout().ContentXPosition
}

func SetPositionYFloat(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) int {
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			return parent.GetLayout().XPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			marginTop := 0
			if currentWidget.GetStyleProperty().Margin != nil {
				marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
			}
			return parent.GetLayout().YPosition + marginTop

		case enums.CSS_POSITION_TYPE_STATIC:
			marginTop := 0
			if currentWidget.GetStyleProperty().Margin != nil {
				marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
			}
			return parent.GetLayout().YPosition + marginTop

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
		} else {
			return beforeCurrentWidget.GetLayout().YPosition
		}
	}
	return 0
}

func SetWidthFloatChild(currentWidget widget.WidgetInterface, styleProperty *StyleProperty.StyleProperty) {
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
