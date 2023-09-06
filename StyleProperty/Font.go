package StyleProperty

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/StyleProperty/structs"
	"gezgin_web_engine/utils"
	"strconv"
	"strings"
)

const FontSizeStringCount = 9

var fontSizeStrings = []string{
	"large",
	"larger",
	"medium",
	"small",
	"smaller",
	"x-large",
	"x-small",
	"xx-large",
	"xx-small",
}

func setFontSize(font *structs.Font, value string) {
	index := utils.IndexFounder(fontSizeStrings, value, FontSizeStringCount)
	if index != -1 {
		font.FontSizeValueType = enums.CssFontSizeType(index)
	} else {
		firstValue, err := strconv.Atoi(value)
		if err == nil {
			if strings.HasSuffix(value, "px") {
				font.FontSizeValueType = enums.CSS_FONT_SIZE_TYPE_LENGTH
				font.FontSizeValue = firstValue
			} else if strings.HasSuffix(value, "%") {
				font.FontSizeValueType = enums.CSS_FONT_SIZE_TYPE_PERCENTAGE
				font.FontSizeValue = firstValue
			}
		} else {
			font.FontSizeValueType = enums.CSS_FONT_SIZE_TYPE_MEDIUM
		}
	}
}

func FontSizePropertySetValue(properties *StyleProperty, value string) {
	if value == "inherit" {
		if !properties.FontInherit {
			if properties.Font == nil {
				properties.Font = new(structs.Font)
			}
			properties.Font.FontSizeInherit = true
		}
	} else {
		properties.FontInherit = false
		if properties.Font == nil {
			properties.Font = new(structs.Font)
		}
		if properties.FontInherit {
			// set inherited other font properties other than font-size
		}
		if value == "initial" {
			properties.Font.FontSizeValueType = enums.CSS_FONT_SIZE_TYPE_MEDIUM
		} else {
			setFontSize(properties.Font, value)
		}
	}
}
