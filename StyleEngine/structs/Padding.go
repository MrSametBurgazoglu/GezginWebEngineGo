package structs

import (
	"gezgin_web_engine/StyleEngine/enums"
)

type Padding struct {
	PaddingTopInherit    bool
	PaddingBottomInherit bool
	PaddingLeftInherit   bool
	PaddingRightInherit  bool

	PaddingTop    float64
	PaddingBottom float64
	PaddingLeft   float64
	PaddingRight  float64

	PaddingTopValueType    enums.CssPropertyValueType
	PaddingBottomValueType enums.CssPropertyValueType
	PaddingLeftValueType   enums.CssPropertyValueType
	PaddingRightValueType  enums.CssPropertyValueType
}
