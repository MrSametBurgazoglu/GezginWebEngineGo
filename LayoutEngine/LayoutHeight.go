package LayoutEngine

import "gezgin_web_engine/StyleEngine"

func (receiver *LayoutProperty) SetHeight(parent *LayoutProperty, children []*LayoutProperty, styleProperty *StyleEngine.StyleProperty) int {
	if children != nil {
		height := 0
		for _, child := range children {
			height += child.Height
		}
		if styleProperty != nil && styleProperty.Margin != nil {
			contentHeight := height - (styleProperty.Margin.MarginTop + styleProperty.Margin.MarginBottom)
			receiver.ContentHeight = contentHeight
		} else {
			receiver.ContentHeight = height
		}
		receiver.Height = height
		return receiver.Height
	} else {
		return receiver.Height
	}
}
