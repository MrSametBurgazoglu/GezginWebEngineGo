package LayoutEngine

import (
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/StyleEngine/enums"
)

type LayoutProperty struct {
	Parent           *LayoutProperty
	Children         []*LayoutProperty
	Position         string
	XPosition        int
	YPosition        int
	ContentXPosition int
	ContentYPosition int
	Width            int
	Height           int
	ContentWidth     int
	ContentHeight    int
	PaddingLeft      int
	PaddingRight     int
	PaddingTop       int
	PaddingBottom    int
	MarginLeft       int
	MarginRight      int
	MarginTop        int
	MarginBottom     int
}

func (receiver *LayoutProperty) SetPosition(parent, beforeCurrentWidget *LayoutProperty, styleProperty *StyleEngine.StyleProperty) (int, int) {
	if styleProperty == nil {
		return receiver.InlineSetPosition(receiver.Parent, beforeCurrentWidget, styleProperty)
	} else if styleProperty.Parent.Display == enums.CSS_DISPLAY_TYPE_FLEX {
		return receiver.SetPositionFlex(parent, beforeCurrentWidget, styleProperty)
	}
	switch styleProperty.Display {
	case enums.CSS_DISPLAY_TYPE_BLOCK:
		return receiver.BlockSetPosition(receiver.Parent, beforeCurrentWidget, styleProperty)
	case enums.CSS_DISPLAY_TYPE_INLINE:
		return receiver.InlineSetPosition(receiver.Parent, beforeCurrentWidget, styleProperty)
	case enums.CSS_DISPLAY_TYPE_FLEX:
		return receiver.BlockSetPosition(receiver.Parent, beforeCurrentWidget, styleProperty)
	}
	return 0, 0
}
