package structs

import "gezgin_web_engine/cssParser/enums"

type Margin struct {
	MarginTopInherit    bool
	MarginBottomInherit bool
	MarginLeftInherit   bool
	MarginRightInherit  bool

	MarginTop    int
	MarginBottom int
	MarginLeft   int
	MarginRight  int

	MarginTopValueType    enums.CssPropertyValueType
	MarginBottomValueType enums.CssPropertyValueType
	MarginLeftValueType   enums.CssPropertyValueType
	MarginRightValueType  enums.CssPropertyValueType
}
