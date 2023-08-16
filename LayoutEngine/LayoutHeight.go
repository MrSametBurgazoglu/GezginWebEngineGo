package LayoutEngine

import "gezgin_web_engine/StyleEngine"

func (receiver *LayoutProperty) SetHeight(parent *LayoutProperty, children []*LayoutProperty, styleProperty *StyleEngine.StyleProperty) int {
	height := 0
	for _, child := range children {
		height += child.Height
	}

	contentHeight := height - (styleProperty.Margin.MarginTop + styleProperty.Margin.MarginBottom)
	receiver.Height = height
	receiver.ContentHeight = contentHeight
	return receiver.ContentHeight
}
