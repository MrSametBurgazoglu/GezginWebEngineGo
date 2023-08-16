package LayoutEngine

import "gezgin_web_engine/StyleEngine"

func (receiver *LayoutProperty) SetWidth(parent *LayoutProperty, children []*LayoutProperty, styleProperty *StyleEngine.StyleProperty) int {
	if receiver.Display == "block" {
		receiver.SetWidthBlock(parent, styleProperty)
	} else {
		receiver.SetWidthInline(children, styleProperty)
	}
	return receiver.ContentWidth
}

func (receiver *LayoutProperty) SetWidthBlock(parent *LayoutProperty, styleProperty *StyleEngine.StyleProperty) {
	width := parent.ContentWidth
	if uint(width) > styleProperty.MaxWidth {
		width = int(styleProperty.MaxWidth)
	}
	if uint(width) < styleProperty.MinWidth {
		width = int(styleProperty.MinWidth)
	}
	contentWidth := width - (styleProperty.Margin.MarginLeft + styleProperty.Margin.MarginRight)
	receiver.Width = width
	receiver.ContentWidth = contentWidth
}

func (receiver *LayoutProperty) SetWidthInline(children []*LayoutProperty, styleProperty *StyleEngine.StyleProperty) {
	width := 0
	for _, child := range children {
		width += child.Width
	}
	contentWidth := width - (styleProperty.Margin.MarginLeft + styleProperty.Margin.MarginRight)
	receiver.Width = width
	receiver.ContentWidth = contentWidth
}
