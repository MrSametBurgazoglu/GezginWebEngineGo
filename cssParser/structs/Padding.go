package structs

import "gezgin_web_engine/cssParser/enums"

type Padding struct {
	PaddingTopInherit    bool
	PaddingBottomInherit bool
	PaddingLeftInherit   bool
	PaddingRightInherit  bool

	PaddingTop    int
	PaddingBottom int
	PaddingLeft   int
	PaddingRight  int

	PaddingTopValueType    enums.CssPropertyValueType
	PaddingBottomValueType enums.CssPropertyValueType
	PaddingLeftValueType   enums.CssPropertyValueType
	PaddingRightValueType  enums.CssPropertyValueType
}
