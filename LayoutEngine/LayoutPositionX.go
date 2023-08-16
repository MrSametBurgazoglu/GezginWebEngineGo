package LayoutEngine

import (
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/StyleEngine/enums"
)

func (receiver *LayoutProperty) SetPositionX(parent *LayoutProperty, styleProperty *StyleEngine.StyleProperty) int {
	position := 0
	if styleProperty != nil {
		switch styleProperty.Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			position = parent.ContentXPosition
		case enums.CSS_POSITION_TYPE_EMPTY:
			position = parent.ContentXPosition
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
	position = 0
	receiver.ContentXPosition = position
	return receiver.ContentXPosition
}
