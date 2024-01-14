package LayoutEngine

import (
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
	if totalWidth > currentWidget.GetLayout().Width {
		currentWidget.GetLayout().Width = totalWidth
		currentWidget.GetLayout().ContentWidth = totalWidth
	}
}

func SetPositionFloat(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	SetPositionXFloat(currentWidget, parent, beforeCurrentWidget)
	SetPositionYFloat(currentWidget, parent, beforeCurrentWidget)
}

func SetPositionXFloat(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			currentWidget.GetLayout().XPosition = parent.GetLayout().ContentXPosition
			currentWidget.GetLayout().ContentXPosition = currentWidget.GetLayout().XPosition
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
			currentWidget.GetLayout().XPosition = x
			currentWidget.GetLayout().ContentXPosition = currentWidget.GetLayout().XPosition
		case enums.CSS_POSITION_TYPE_STATIC:
			currentWidget.GetLayout().XPosition = parent.GetLayout().ContentXPosition
			currentWidget.GetLayout().ContentXPosition = currentWidget.GetLayout().XPosition
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			//println("not implemented yet")
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			//println("not implemented yet")
		}
	} else {
		currentWidget.GetLayout().XPosition = parent.GetLayout().ContentXPosition
		currentWidget.GetLayout().ContentXPosition = currentWidget.GetLayout().XPosition
	}
}

func SetPositionYFloat(currentWidget, parent, beforeCurrentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			currentWidget.GetLayout().YPosition = parent.GetLayout().ContentYPosition
			currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			marginTop := 0
			if currentWidget.GetStyleProperty().Margin != nil {
				marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
			}
			currentWidget.GetLayout().YPosition = parent.GetLayout().YPosition + marginTop
			currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
		case enums.CSS_POSITION_TYPE_STATIC:
			marginTop := 0
			if currentWidget.GetStyleProperty().Margin != nil {
				marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
			}
			currentWidget.GetLayout().YPosition = parent.GetLayout().YPosition + marginTop
			currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			//println("implemented yet")
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			//println("implemented yet")
		}
	} else {
		if beforeCurrentWidget == nil {
			currentWidget.GetLayout().YPosition = parent.GetLayout().ContentYPosition
			currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
		} else {
			currentWidget.GetLayout().YPosition = beforeCurrentWidget.GetLayout().YPosition
			currentWidget.GetLayout().ContentYPosition = currentWidget.GetLayout().YPosition
		}
	}
}
