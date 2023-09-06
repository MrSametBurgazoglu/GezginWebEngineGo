package StyleProperty

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/utils"
	"strings"
)

const TextAlignStringCount = 4

var textAlignStrings = []string{
	"center",
	"justify",
	"left",
	"right",
}

func setTextAlignValue(cssProperties *StyleProperty, value string) {
	if strings.Contains(value, "!") { //TEMPORARY SOLUTION TO PASS IMPORTANT
		value = value[:strings.Index(value, "!")]
	}
	index := utils.IndexFounder(textAlignStrings, value, TextAlignStringCount)
	if index != -1 {
		cssProperties.TextAlign = enums.CssTextAlignType(index + 1)
	} else {
		cssProperties.TextAlign = enums.CSS_TEXT_ALIGN_LEFT
	}
}

func TextAlignPropertySetValue(cssProperties *StyleProperty, value string) {
	//TODO FIX IMPORTANT PROBLEM
	if value == "inherit" {
		cssProperties.TextAlignInherit = true
	} else {
		cssProperties.TextAlignInherit = false
		if value == "initial" {
			cssProperties.TextAlign = enums.CSS_TEXT_ALIGN_LEFT
		} else {
			setTextAlignValue(cssProperties, value)
		}
	}
}

func UpdateText(properties *StyleProperty, source *StyleProperty) {
	if source.TextAlignInherit && properties.TextAlign == enums.CSS_TEXT_ALIGN_EMPTY {
		properties.TextAlignInherit = true
		properties.TextAlign = enums.CSS_TEXT_ALIGN_EMPTY
	} else if source.TextAlign != enums.CSS_TEXT_ALIGN_EMPTY && properties.TextAlign == enums.CSS_TEXT_ALIGN_EMPTY {
		properties.TextAlign = source.TextAlign
	}
}
