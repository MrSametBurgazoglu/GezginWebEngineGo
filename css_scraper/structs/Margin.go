package structs

import "gezgin_web_engine/css_scraper/enums"

type Margin struct {
	MarginTopValueType    enums.CssPropertyValueType
	MarginTop             int
	MarginBottomValueType enums.CssPropertyValueType
	MarginBottom          int
	MarginLeftValueType   enums.CssPropertyValueType
	MarginLeft            int
	MarginRightValueType  enums.CssPropertyValueType
	MarginRight           int
	MarginTopInherit      bool
	MarginBottomInherit   bool
	MarginLeftInherit     bool
	MarginRightInherit    bool
}
