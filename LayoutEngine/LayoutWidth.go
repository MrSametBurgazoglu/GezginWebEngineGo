package LayoutEngine

import (
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/StyleEngine/enums"
)

func (receiver *LayoutProperty) SetWidth(parent *LayoutProperty, children []*LayoutProperty, styleProperty *StyleEngine.StyleProperty) int {
	if styleProperty == nil {
		receiver.SetWidthInline(children, styleProperty)
	} else if styleProperty.Display == enums.CSS_DISPLAY_TYPE_BLOCK {
		receiver.SetWidthBlock(parent, styleProperty)
	} else if styleProperty.Display == enums.CSS_DISPLAY_TYPE_FLEX {
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
	if children != nil {
		width := 0
		for _, child := range children {
			width += child.Width
		}
		contentWidth := width - (styleProperty.Margin.MarginLeft + styleProperty.Margin.MarginRight)
		receiver.Width = width
		receiver.ContentWidth = contentWidth
	}
}
