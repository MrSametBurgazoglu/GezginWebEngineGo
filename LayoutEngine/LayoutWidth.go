package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/StyleProperty/enums"
)

func (receiver *LayoutProperty) SetWidth(parent *LayoutProperty, children []*LayoutProperty, styleProperty, parentStyleProperty *StyleProperty.StyleProperty) int {
	if styleProperty == nil {
		receiver.SetWidthInline(children, styleProperty)
	} else if styleProperty.Display == enums.CSS_DISPLAY_TYPE_FLEX {
		receiver.SetFLexContainerWidth(styleProperty)
	} else if styleProperty.Parent.Display == enums.CSS_DISPLAY_TYPE_FLEX {
		println("heyyo")
	} else if styleProperty.Display == enums.CSS_DISPLAY_TYPE_BLOCK {
		receiver.SetWidthBlock(parent, styleProperty)
	} else if styleProperty.Display == enums.CSS_DISPLAY_TYPE_FLEX {
		receiver.SetWidthBlock(parent, styleProperty)
	} else {
		receiver.SetWidthInline(children, styleProperty)
	}
	return receiver.Width
}
