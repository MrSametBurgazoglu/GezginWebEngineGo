package StyleEngine

import (
	"gezgin_web_engine/StyleEngine/enums"
	"gezgin_web_engine/utils"
)

const DisplayStringCount = 5

var displayStrings = []string{
	"block",
	"inline",
	"contents",
	"flex",
	"hello",
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
