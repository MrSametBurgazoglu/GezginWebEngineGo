package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/StyleProperty/enums"
)

func (receiver *LayoutProperty) SetPositionY(parent, beforeCurrentWidget *LayoutProperty, styleProperty *StyleProperty.StyleProperty) int {
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
