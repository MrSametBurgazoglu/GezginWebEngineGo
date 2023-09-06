package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/StyleProperty/enums"
)

func (receiver *LayoutProperty) BlockSetPosition(parent, beforeCurrentWidget *LayoutProperty, styleProperty *StyleProperty.StyleProperty) (int, int) {
	return receiver.BlockSetPositionX(parent, styleProperty), receiver.BlockSetPositionY(parent, beforeCurrentWidget, styleProperty)
}

func (receiver *LayoutProperty) BlockSetPositionX(parent *LayoutProperty, styleProperty *StyleProperty.StyleProperty) int {
	position := 0
	if styleProperty != nil {
		switch styleProperty.Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			position = parent.ContentXPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			position = parent.XPosition
			if styleProperty.Margin != nil {
				if styleProperty.Margin.MarginLeftValueType == enums.CSS_PROPERTY_VALUE_TYPE_AUTO && styleProperty.Margin.MarginRightValueType == enums.CSS_PROPERTY_VALUE_TYPE_AUTO {
					position += (receiver.Parent.Width - receiver.Width) / 2
				}
			}
			//position = parent.ContentXPosition
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

func (receiver *LayoutProperty) BlockSetPositionY(parent, beforeCurrentWidget *LayoutProperty, styleProperty *StyleProperty.StyleProperty) int {
	if styleProperty != nil {
		switch styleProperty.Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			return parent.XPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			if beforeCurrentWidget != nil {
				marginTop := 0
				if styleProperty.Margin != nil {
					marginTop = styleProperty.Margin.MarginTop
				}
				return beforeCurrentWidget.YPosition + beforeCurrentWidget.Height + marginTop
			} else {
				marginTop := 0
				if styleProperty.Margin != nil {
					marginTop = styleProperty.Margin.MarginTop
				}
				return parent.YPosition + marginTop
			}
		case enums.CSS_POSITION_TYPE_STATIC:
			marginTop := 0
			if styleProperty.Margin != nil {
				marginTop = styleProperty.Margin.MarginTop
			}
			if beforeCurrentWidget != nil {
				return beforeCurrentWidget.YPosition + beforeCurrentWidget.Height + marginTop
			} else {
				return parent.YPosition + parent.Height + marginTop
			}
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

func (receiver *LayoutProperty) SetWidthBlock(parent *LayoutProperty, styleProperty *StyleProperty.StyleProperty) {
	width := parent.Width
	if styleProperty.MaxWidth > 0 && uint(width) > styleProperty.MaxWidth {
		width = int(styleProperty.MaxWidth)
	}
	if styleProperty.MinWidth > 0 && uint(width) < styleProperty.MinWidth {
		width = int(styleProperty.MinWidth)
	}
	receiver.Width = width
	if styleProperty != nil && styleProperty.Margin != nil {
		contentWidth := width - (styleProperty.Margin.MarginLeft + styleProperty.Margin.MarginRight)
		receiver.ContentWidth = contentWidth
	} else {
		receiver.ContentWidth = width
	}
}
