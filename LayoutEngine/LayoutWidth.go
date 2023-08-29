package LayoutEngine

import "gezgin_web_engine/StyleEngine"

func (receiver *LayoutProperty) SetWidth(parent *LayoutProperty, children []*LayoutProperty, styleProperty *StyleEngine.StyleProperty) int {
	if receiver.Display == "block" {
		receiver.SetWidthBlock(parent, styleProperty)
	} else {
		receiver.SetWidthInline(children, styleProperty)
	}
	return receiver.Width
}

func (receiver *LayoutProperty) SetWidthBlock(parent *LayoutProperty, styleProperty *StyleEngine.StyleProperty) {
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

func (receiver *LayoutProperty) SetWidthInline(children []*LayoutProperty, styleProperty *StyleEngine.StyleProperty) {
	if children == nil {
		receiver.Width = 0
		receiver.ContentWidth = 0
	} else {
		width := 0
		for _, child := range children {
			width += child.Width
		}
		contentWidth := width - (styleProperty.Margin.MarginLeft + styleProperty.Margin.MarginRight)
		receiver.Width = width
		receiver.ContentWidth = contentWidth
	}
}
