package StyleProperty

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/utils"
	"strings"
)

const FloatStringCount = 5

var floatStrings = []string{
	"initial",
	"left",
	"none",
	"right",
}

func setFloatValue(cssProperties *StyleProperty, value string) {
	index := utils.IndexFounder(floatStrings, value, FloatStringCount)
	if index != -1 {
		cssProperties.Float = enums.CssFloatType(index)
	} else {
		cssProperties.Float = enums.CSS_FLOAT_EMPTY
	}
}

func FloatPropertySetValue(cssProperties *StyleProperty, value string) {
	if strings.Contains(value, "!important") {
		value = strings.ReplaceAll(value, "!important", "")
	}
	if value == "inherit" {
		cssProperties.FloatInherit = true
	} else {
		cssProperties.FloatInherit = false
		setFloatValue(cssProperties, value)
	}
}
