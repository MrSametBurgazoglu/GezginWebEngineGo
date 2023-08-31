package LayoutEngine

import (
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/StyleEngine/enums"
)

func LookForWidth(layoutProperty *LayoutProperty) int {
	if len(layoutProperty.Children) == 0 {
		return layoutProperty.Width
	} else {
		maxWidth := 0
		for _, child := range layoutProperty.Children {
			currentWidth := LookForWidth(child)
			println(currentWidth, " asdadadasdsadsa")
			if currentWidth > maxWidth {
				maxWidth = currentWidth
			}
		}
		return maxWidth
	}
}

func (receiver *LayoutProperty) SetFLexContainerWidth(styleProperty *StyleEngine.StyleProperty) {
	receiver.SetWidthBlock(receiver.Parent, styleProperty)
	totalWidth := 0
	for i, child := range receiver.Children {
		width := LookForWidth(receiver)
		if styleProperty.Children[i].Width != 0 {
			switch styleProperty.Children[i].WidthValueType {
			case enums.CSS_PROPERTY_VALUE_TYPE_PIXEL:
				width = int(styleProperty.Children[i].Width)
			case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
				width = int(float64(receiver.Width) * (float64(styleProperty.Children[i].Width) / 100.0))
			}
		}
		//look width here
		if styleProperty.Children[i].MaxWidth != 0 && width > int(styleProperty.Children[i].MaxWidth) {
			width = int(styleProperty.Children[i].MaxWidth)
		}
		if width < int(styleProperty.Children[i].MinWidth) {
			width = int(styleProperty.Children[i].MinWidth)
		}
		child.Width = width
		totalWidth += width
	}
	if totalWidth > receiver.Width {
		receiver.Width = totalWidth
		receiver.ContentWidth = totalWidth
	}
}

func (receiver *LayoutProperty) SetPositionFlex(parent, beforeCurrentWidget *LayoutProperty, styleProperty *StyleEngine.StyleProperty) (int, int) {
	return receiver.SetPositionXFlex(parent, beforeCurrentWidget, styleProperty), receiver.SetPositionYFlex(parent, beforeCurrentWidget, styleProperty)
}

func (receiver *LayoutProperty) SetPositionXFlex(parent, beforeCurrentWidget *LayoutProperty, styleProperty *StyleEngine.StyleProperty) int {
	position := 0
	if styleProperty != nil {
		switch styleProperty.Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			position = parent.ContentXPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			x := 0
			if beforeCurrentWidget == nil {
				x = parent.XPosition
				if styleProperty.Margin != nil {
					x += receiver.MarginLeft
				}
				if styleProperty.Padding != nil {
					x += receiver.PaddingLeft
				}
			} else {
				x = beforeCurrentWidget.XPosition + beforeCurrentWidget.Width
				if styleProperty.Margin != nil {
					x += receiver.MarginLeft
				}
				if styleProperty.Padding != nil {
					x += receiver.PaddingLeft
				}
			}
			position = x

		case enums.CSS_POSITION_TYPE_STATIC:
			position = parent.ContentXPosition
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			if styleProperty.Left != 0 {
				position = parent.ContentXPosition + int(styleProperty.Left)
			} else if styleProperty.Right != 0 {
				position = parent.ContentWidth - int(styleProperty.Right)
			} else {
				position = parent.ContentXPosition
			}
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			if styleProperty.Left != 0 {
				position = parent.ContentXPosition + int(styleProperty.Left)
			} else if styleProperty.Right != 0 {
				position = parent.ContentWidth - int(styleProperty.Right)
			} else {
				position = parent.ContentXPosition
			}
		}
	} else {
		position = parent.ContentXPosition
	}
	receiver.ContentXPosition = position
	receiver.XPosition = position
	return receiver.ContentXPosition
}

func (receiver *LayoutProperty) SetPositionYFlex(parent, beforeCurrentWidget *LayoutProperty, styleProperty *StyleEngine.StyleProperty) int {
	if styleProperty != nil {
		switch styleProperty.Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			return parent.XPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			marginTop := 0
			if styleProperty.Margin != nil {
				marginTop = styleProperty.Margin.MarginTop
			}
			return parent.YPosition + marginTop

		case enums.CSS_POSITION_TYPE_STATIC:
			marginTop := 0
			if styleProperty.Margin != nil {
				marginTop = styleProperty.Margin.MarginTop
			}
			return parent.YPosition + parent.Height + marginTop

		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			if styleProperty.Top != 0 {
				return parent.YPosition + int(styleProperty.Top)
			} else if styleProperty.Bottom != 0 {
				return parent.YPosition + parent.Height - int(styleProperty.Bottom)
			} else {
				return parent.YPosition + parent.Height
			}
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			if beforeCurrentWidget != nil {
				return beforeCurrentWidget.YPosition + beforeCurrentWidget.Height + int(styleProperty.Top)
			} else {
				return parent.YPosition + int(styleProperty.Top)
			}
		}
	} else {
		if beforeCurrentWidget == nil {
			return parent.YPosition
		} else {
			return beforeCurrentWidget.Height
		}
	}
	return 0
}

func (receiver *LayoutProperty) SetWidthFlexChild(children []*LayoutProperty, styleProperty *StyleEngine.StyleProperty) {
	//you must set childrens width first
	if children != nil {
		width := 0
		for _, child := range children {
			width += child.Width
		}
		contentWidth := width
		if styleProperty != nil && styleProperty.Margin != nil {
			contentWidth = width - (styleProperty.Margin.MarginLeft + styleProperty.Margin.MarginRight)
		}
		receiver.Width = width
		receiver.ContentWidth = contentWidth
	}
}
