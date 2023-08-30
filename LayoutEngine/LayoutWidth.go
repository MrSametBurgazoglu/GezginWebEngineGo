package LayoutEngine

import (
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/StyleEngine/enums"
)

func (receiver *LayoutProperty) SetWidth(parent *LayoutProperty, children []*LayoutProperty, styleProperty, parentStyleProperty *StyleEngine.StyleProperty) int {
	if styleProperty == nil {
		receiver.SetWidthInline(children, styleProperty)
	} else if styleProperty.Display == enums.CSS_DISPLAY_TYPE_FLEX {
		receiver.SetFLexContainerWidth()
	} else if styleProperty.Display == enums.CSS_DISPLAY_TYPE_BLOCK {
		receiver.SetWidthBlock(parent, styleProperty)
	} else if styleProperty.Display == enums.CSS_DISPLAY_TYPE_FLEX {
		receiver.SetWidthBlock(parent, styleProperty)
	} else {
		receiver.SetWidthInline(children, styleProperty)
	}
	return receiver.Width
}
