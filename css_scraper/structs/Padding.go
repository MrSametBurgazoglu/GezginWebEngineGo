package structs

import "gezgin_web_engine/css_scraper/enums"

type Padding struct {
	PaddingTopInherit    bool
	PaddingBottomInherit bool
	PaddingLeftInherit   bool
	PaddingRightInherit  bool

	PaddingTopValueType    enums.CssPropertyValueType
	PaddingBottomValueType enums.CssPropertyValueType
	PaddingLeftValueType   enums.CssPropertyValueType
	PaddingRightValueType  enums.CssPropertyValueType

	PaddingTop    int
	PaddingBottom int
	PaddingLeft   int
	PaddingRight  int
}
