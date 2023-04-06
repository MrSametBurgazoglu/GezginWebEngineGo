package properties

import (
	"gezgin_web_engine/css_scraper/enums"
	"gezgin_web_engine/css_scraper/structs"
	"gezgin_web_engine/utils"
	"strconv"
	"strings"
)

const PositionStringCount = 5

var positionStrings = []string{
	"",
	"absolute",
	"fixed",
	"relative",
	"static",
	"sticky",
}

func setPosition(cssProperties *structs.CssProperties, value string) {
	index := utils.IndexFounder(positionStrings, value, PositionStringCount)
	if index != -1 {
		cssProperties.Position = enums.CssPositionType(index)
	} else {
		cssProperties.Position = enums.CSS_POSITION_TYPE_STATIC
	}
}

func setTop(cssProperties *structs.CssProperties, value string) {
	if strings.HasSuffix(value, "px") {
		top, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err != nil {
			// handle error
		}
		cssProperties.Top = uint(top)
		cssProperties.TopValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
	} else if strings.HasSuffix(value, "%") {
		top, err := strconv.Atoi(strings.TrimSuffix(value, "%"))
		if err != nil {
			// handle error
		}
		cssProperties.Top = uint(top)
		cssProperties.TopValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
	} else {
		// DEFAULT VALUE
		cssProperties.TopValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	}
}

func setBottom(cssProperties *structs.CssProperties, value string) {
	if strings.HasSuffix(value, "px") {
		bottom, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err != nil {
			// handle error
		}
		cssProperties.Bottom = uint(bottom)
		cssProperties.BottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
	} else if strings.HasSuffix(value, "%") {
		bottom, err := strconv.Atoi(strings.TrimSuffix(value, "%"))
		if err != nil {
			// handle error
		}
		cssProperties.Bottom = uint(bottom)
		cssProperties.BottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
	} else {
		// DEFAULT VALUE
		cssProperties.BottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	}
}

func setLeft(cssProperties *structs.CssProperties, value string) {
	if strings.HasSuffix(value, "px") {
		left, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err != nil {
			// handle error
		}
		cssProperties.Left = uint(left)
		cssProperties.LeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
	} else if strings.HasSuffix(value, "%") {
		left, err := strconv.Atoi(strings.TrimSuffix(value, "%"))
		if err != nil {
			// handle error
		}
		cssProperties.Left = uint(left)
		cssProperties.LeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
	} else {
		// DEFAULT VALUE
		cssProperties.LeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	}
}

func setRight(cssProperties *structs.CssProperties, value string) {
	if strings.HasSuffix(value, "px") {
		right, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err != nil {
			// handle error
		}
		cssProperties.Right = uint(right)
		cssProperties.RightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
	} else if strings.HasSuffix(value, "%") {
		right, err := strconv.Atoi(strings.TrimSuffix(value, "%"))
		if err != nil {
			// handle error
		}
		cssProperties.Right = uint(right)
		cssProperties.RightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
	} else {
		// DEFAULT VALUE
		cssProperties.RightValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	}
}

func PositionPropertySetValue(cssProperties *structs.CssProperties, value string) {
	if value == "inherit" {
		cssProperties.PositionInherit = true
	} else if value == "initial" {
		cssProperties.Position = enums.CSS_POSITION_TYPE_STATIC
	} else {
		setPosition(cssProperties, value)
	}
}

func TopPropertySetValue(cssProperties *structs.CssProperties, value string) {
	if value == "inherit" {
		cssProperties.TopInherit = true
	} else if value == "initial" {
		cssProperties.TopValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	} else {
		setTop(cssProperties, value)
	}
}

func BottomPropertySetValue(cssProperties *structs.CssProperties, value string) {
	if value == "inherit" {
		cssProperties.BottomInherit = true
	} else if value == "initial" {
		cssProperties.BottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	} else {
		setBottom(cssProperties, value)
	}
}

func LeftPropertySetValue(cssProperties *structs.CssProperties, value string) {
	if value == "inherit" {
		cssProperties.LeftInherit = true
	} else if value == "initial" {
		cssProperties.LeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	} else {
		setLeft(cssProperties, value)
	}
}

func RightPropertySetValue(cssProperties *structs.CssProperties, value string) {
	if value == "inherit" {
		cssProperties.RightInherit = true
	} else if value == "initial" {
		cssProperties.RightValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	} else {
		setRight(cssProperties, value)
	}
}

func computeInheritPosition(dest, source *structs.CssProperties) {
	if dest.PositionInherit {
		dest.Position = source.Position
	}
	if dest.LeftInherit {
		dest.Left = source.Left
	}
	if dest.RightInherit {
		dest.Right = source.Right
	}
	if dest.TopInherit {
		dest.Top = source.Top
	}
	if dest.BottomInherit {
		dest.Bottom = source.Bottom
	}
}

func UpdatePosition(cssProperties, source *structs.CssProperties) {
	if source.PositionInherit {
		cssProperties.PositionInherit = true
	} else if source.Position != enums.CSS_POSITION_TYPE_EMPTY {
		cssProperties.Position = source.Position
	}
	if source.TopInherit {
		cssProperties.TopInherit = true
	} else if source.TopValueType != enums.CSS_PROPERTY_VALUE_TYPE_EMPTY {
		cssProperties.Top = source.Top
		cssProperties.TopValueType = source.TopValueType
	}
	if source.BottomInherit {
		cssProperties.BottomInherit = true
	} else if source.BottomValueType != enums.CSS_PROPERTY_VALUE_TYPE_EMPTY {
		cssProperties.Bottom = source.Bottom
		cssProperties.BottomValueType = source.BottomValueType
	}
	if source.LeftInherit {
		cssProperties.LeftInherit = true
	} else if source.LeftValueType != enums.CSS_PROPERTY_VALUE_TYPE_EMPTY {
		cssProperties.Left = source.Left
		cssProperties.LeftValueType = source.LeftValueType
	}
	if source.RightInherit {
		cssProperties.RightInherit = true
	} else if source.RightValueType != enums.CSS_PROPERTY_VALUE_TYPE_EMPTY {
		cssProperties.Right = source.Right
		cssProperties.RightValueType = source.RightValueType
	}
}
