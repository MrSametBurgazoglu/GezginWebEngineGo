package StyleProperty

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/utils"
)

const DisplayStringCount = 7

var displayStrings = []string{
	"block",
	"contents",
	"flex",
	"grid",
	"inline",
	"inline-block",
	"inline-flex",
}

func setDisplayValue(cssProperties *StyleProperty, value string) {
	index := utils.IndexFounder(displayStrings, value, DisplayStringCount)
	if index != -1 {
		cssProperties.Display = enums.CssDisplayType(index)
	} else {
		cssProperties.Display = enums.CSS_DISPLAY_TYPE_BLOCK
	}
}

func DisplayPropertySetValue(cssProperties *StyleProperty, value string) {
	if value == "inherit" {
		cssProperties.MaxHeightInherit = true
	} else {
		cssProperties.MaxHeightInherit = false
		if value == "initial" {
			cssProperties.MaxHeight = 0
			cssProperties.MaxHeightValueType = enums.CSS_PROPERTY_VALUE_TYPE_NONE
		} else {
			setDisplayValue(cssProperties, value)
		}
	}
}
