package StyleProperty

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/StyleProperty/structs"
	"strconv"
	"strings"
)

func setPaddingTop(padding *structs.Padding, value string) {
	if value == "auto" {
		padding.PaddingTopValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	} else {
		intValue, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err == nil {
			padding.PaddingTopValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			padding.PaddingTop = float64(intValue)
			return
		}
		floatValue, err := strconv.ParseFloat(strings.TrimSuffix(value, "%"), 64)
		if err == nil {
			padding.PaddingTopValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
			padding.PaddingTop = floatValue
			return
		}
		floatValue2, err := strconv.ParseFloat(strings.TrimSuffix(value, "rem"), 64)
		if err == nil {
			padding.PaddingTopValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			padding.PaddingTop = floatValue2 * 16
			return
		}
	}
}

func setPaddingBottom(padding *structs.Padding, value string) {
	if value == "auto" {
		padding.PaddingBottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	} else {
		intValue, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err == nil {
			padding.PaddingBottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			padding.PaddingBottom = float64(intValue)
			return
		}
		floatValue, err := strconv.ParseFloat(strings.TrimSuffix(value, "%"), 64)
		if err == nil {
			padding.PaddingBottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
			padding.PaddingBottom = floatValue
			return
		}
		floatValue2, err := strconv.ParseFloat(strings.TrimSuffix(value, "rem"), 64)
		if err == nil {
			padding.PaddingBottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			padding.PaddingBottom = floatValue2 * 16
			return
		}
	}
}

func setPaddingLeft(padding *structs.Padding, value string) {
	if value == "auto" {
		padding.PaddingLeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	} else {
		intValue, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err == nil {
			padding.PaddingLeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			padding.PaddingLeft = float64(intValue)
			return
		}
		floatValue, err := strconv.ParseFloat(strings.TrimSuffix(value, "%"), 64)
		if err == nil {
			padding.PaddingLeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
			padding.PaddingLeft = floatValue
			return
		}
		floatValue2, err := strconv.ParseFloat(strings.TrimSuffix(value, "rem"), 64)
		if err == nil {
			padding.PaddingLeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			padding.PaddingLeft = floatValue2 * 16
			return
		}
	}
}

func setPaddingRight(padding *structs.Padding, value string) {
	if value == "auto" {
		padding.PaddingRightValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	} else {
		intValue, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err == nil {
			padding.PaddingRightValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			padding.PaddingRight = float64(intValue)
			return
		}
		floatValue, err := strconv.ParseFloat(strings.TrimSuffix(value, "%"), 64)
		if err == nil {
			padding.PaddingRightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
			padding.PaddingRight = floatValue
			return
		}
		floatValue2, err := strconv.ParseFloat(strings.TrimSuffix(value, "rem"), 64)
		if err == nil {
			padding.PaddingRightValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			padding.PaddingRight = floatValue2 * 16
			return
		}
	}
}

func setPadding(padding *structs.Padding, value string) {
	if strings.Contains(value, "!important") {
		value = strings.ReplaceAll(value, "!important", "")
	}
	values := strings.Split(value, " ")
	switch len(values) {
	case 1:
		setPaddingTop(padding, values[0])
		setPaddingRight(padding, values[0])
		setPaddingBottom(padding, values[0])
		setPaddingLeft(padding, values[0])
	case 2:
		setPaddingTop(padding, values[0])
		setPaddingRight(padding, values[1])
		setPaddingBottom(padding, values[0])
		setPaddingLeft(padding, values[1])
	case 3:
		setPaddingTop(padding, values[0])
		setPaddingRight(padding, values[1])
		setPaddingBottom(padding, values[2])
		setPaddingLeft(padding, values[1])
	case 4:
		setPaddingTop(padding, values[0])
		setPaddingRight(padding, values[1])
		setPaddingBottom(padding, values[2])
		setPaddingLeft(padding, values[3])
	default:
		break
	}
}

func PaddingTopPropertySetValue(currentWidget *StyleProperty, value string) {
	if value == "inherit" {
		if !currentWidget.PaddingInherit {
			if currentWidget.Padding == nil {
				currentWidget.Padding = new(structs.Padding)
			}
			currentWidget.Padding.PaddingTopInherit = true
		}
	} else {
		if currentWidget.Padding == nil {
			currentWidget.Padding = new(structs.Padding)
		}
		if currentWidget.PaddingInherit {
			currentWidget.Padding.PaddingBottomInherit = true
			currentWidget.Padding.PaddingLeftInherit = true
			currentWidget.Padding.PaddingRightInherit = true
		}
		currentWidget.Padding.PaddingTopInherit = false
		if value == "initial" {
			currentWidget.Padding.PaddingTopValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			currentWidget.Padding.PaddingTop = 0
		} else {
			setPaddingTop(currentWidget.Padding, value)
		}
	}
}

func PaddingBottomPropertySetValue(currentWidget *StyleProperty, value string) {
	if value == "inherit" {
		if !currentWidget.PaddingInherit {
			if currentWidget.Padding == nil {
				currentWidget.Padding = new(structs.Padding)
			}
			currentWidget.Padding.PaddingBottomInherit = true
		}
	} else {
		if currentWidget.Padding == nil {
			currentWidget.Padding = new(structs.Padding)
		}
		if currentWidget.PaddingInherit {
			currentWidget.Padding.PaddingTopInherit = true
			currentWidget.Padding.PaddingLeftInherit = true
			currentWidget.Padding.PaddingRightInherit = true
		}
		currentWidget.Padding.PaddingBottomInherit = false
		if value == "initial" {
			currentWidget.Padding.PaddingBottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			currentWidget.Padding.PaddingBottom = 0
		} else {
			setPaddingBottom(currentWidget.Padding, value)
		}
	}
}

func PaddingLeftPropertyValue(currentWidget *StyleProperty, value string) {
	if value == "inherit" {
		if !currentWidget.PaddingInherit {
			if currentWidget.Padding == nil {
				currentWidget.Padding = new(structs.Padding)
			}
			currentWidget.Padding.PaddingLeftInherit = true
		}
	} else {
		if currentWidget.Padding == nil {
			currentWidget.Padding = new(structs.Padding)
		}
		if currentWidget.PaddingInherit {
			currentWidget.Padding.PaddingTopInherit = true
			currentWidget.Padding.PaddingBottomInherit = true
			currentWidget.Padding.PaddingRightInherit = true
		}
		currentWidget.Padding.PaddingLeftInherit = false
		if value == "initial" {
			currentWidget.Padding.PaddingLeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			currentWidget.Padding.PaddingLeft = 0
		} else {
			setPaddingLeft(currentWidget.Padding, value)
		}
	}
}

func PaddingRightPropertySetValue(currentWidget *StyleProperty, value string) {
	if value == "inherit" {
		if !currentWidget.PaddingInherit {
			if currentWidget.Padding == nil {
				currentWidget.Padding = new(structs.Padding)
			}
			currentWidget.Padding.PaddingRightInherit = true
		}
	} else {
		if currentWidget.Padding == nil {
			currentWidget.Padding = new(structs.Padding)
		}
		if currentWidget.PaddingInherit {
			currentWidget.Padding.PaddingTopInherit = true
			currentWidget.Padding.PaddingBottomInherit = true
			currentWidget.Padding.PaddingLeftInherit = true
		}
		currentWidget.Padding.PaddingRightInherit = false
		if value == "initial" {
			currentWidget.Padding.PaddingRightValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			currentWidget.Padding.PaddingRight = 0
		} else {
			setPaddingRight(currentWidget.Padding, value)
		}
	}
}

func PaddingPropertySetValue(currentWidget *StyleProperty, value string) {
	if value == "inherit" {
		currentWidget.PaddingInherit = true
	} else {
		currentWidget.PaddingInherit = false
		if currentWidget.Padding == nil {
			currentWidget.Padding = new(structs.Padding)
		}
		if value == "initial" {
			currentWidget.Padding.PaddingTop = 0
			currentWidget.Padding.PaddingBottom = 0
			currentWidget.Padding.PaddingLeft = 0
			currentWidget.Padding.PaddingRight = 0
			currentWidget.Padding.PaddingTopValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
			currentWidget.Padding.PaddingBottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
			currentWidget.Padding.PaddingLeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
			currentWidget.Padding.PaddingRightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
		} else {
			setPadding(currentWidget.Padding, value)
		}
	}
}
