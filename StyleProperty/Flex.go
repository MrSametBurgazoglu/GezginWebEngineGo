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
		cssProperties.FlexInherit = true
	} else {
		cssProperties.FlexInherit = false
		if value == "initial" {
			cssProperties.FlexDirection = 0
		} else {
			setFlexDirectionValue(cssProperties, value)
		}
	}
}
