package LayoutProperty

import (
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/StyleProperty/enums"
)

type LayoutProperty struct {
	Parent           *LayoutProperty
	Children         []*LayoutProperty
	StyleProperty    *StyleProperty.StyleProperty
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

func (receiver *LayoutProperty) GetPresetHeight() int {
	if receiver.Height > 0 {
		return receiver.Height
	} else if receiver.StyleProperty.HeightValueType == enums.CSS_PROPERTY_VALUE_TYPE_PIXEL {
		return int(receiver.StyleProperty.Height)
	} else if receiver.StyleProperty.HeightValueType == enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE {
		return receiver.Parent.GetPresetHeight() * int(receiver.StyleProperty.Height) / 100
	}
	return 0
}

func (receiver *LayoutProperty) GetTotalWidth() int {
	return receiver.MarginLeft + receiver.PaddingLeft + receiver.ContentWidth + receiver.PaddingRight + receiver.MarginRight
}

func (receiver *LayoutProperty) GetTotalHeight() int {
	return receiver.MarginTop + receiver.PaddingTop + receiver.ContentHeight + receiver.PaddingBottom + receiver.MarginBottom
}

func (receiver *LayoutProperty) GetTotalContentHeight() int {
	return receiver.PaddingTop + receiver.ContentHeight + receiver.PaddingBottom
}
