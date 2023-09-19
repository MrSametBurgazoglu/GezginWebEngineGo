package StyleProperty

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/utils"
	"strings"
)

const FlexDirectionStringCount = 5

var flexDirectionStrings = []string{
	"column",
	"colum-reverse",
	"row",
	"row-reverse",
}

const FlexWrapStringCount = 3

var flexWrapStrings = []string{
	"no-wrap",
	"wrap",
	"wrap-reverse",
}

func setFlexDirectionValue(cssProperties *StyleProperty, value string) {
	index := utils.IndexFounder(flexDirectionStrings, value, FlexDirectionStringCount)
	if index != -1 {
		cssProperties.FlexDirection = enums.CssFlexDirectionType(index + 1)
	} else {
		cssProperties.FlexDirection = enums.CSS_FLEX_DIRECTION_ROW
	}
}

func FlexDirectionPropertySetValue(cssProperties *StyleProperty, value string) {
	if strings.Contains(value, "!important") {
		value = strings.ReplaceAll(value, "!important", "")
	}
	if value == "inherit" {
		cssProperties.FlexDirectionInherit = true
	} else {
		cssProperties.FlexDirectionInherit = false
		if value == "initial" {
			cssProperties.FlexDirection = 0
		} else {
			setFlexDirectionValue(cssProperties, value)
		}
	}
}

func setFlexWrapValue(cssProperties *StyleProperty, value string) {
	index := utils.IndexFounder(flexWrapStrings, value, FlexWrapStringCount)
	if index != -1 {
		cssProperties.FlexWrap = enums.CssFlexWrapType(index + 1)
	} else {
		cssProperties.FlexWrap = enums.CSS_FLEX_WRAP_NOWRAP
	}
}

func FlexWrapPropertySetValue(cssProperties *StyleProperty, value string) {
	if strings.Contains(value, "!important") {
		value = strings.ReplaceAll(value, "!important", "")
	}
	if value == "inherit" {
		cssProperties.FlexInherit = true
	} else {
		cssProperties.FlexInherit = false
		if value == "initial" {
			cssProperties.FlexWrap = 0
		} else {
			setFlexWrapValue(cssProperties, value)
		}
	}
}
