package StyleProperty

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/StyleProperty/structs"
	"strconv"
	"strings"
)

func setMarginTop(margin *structs.Margin, value string) {
	if strings.Contains(value, "!important") {
		value = strings.ReplaceAll(value, "!important", "")
	}
	if value == "auto" {
		margin.MarginTopValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	} else {
		intValue, err := strconv.Atoi(strings.TrimRight(value, "px"))
		if err == nil {
			margin.MarginTopValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			margin.MarginTop = intValue
		}
		percentValue, err := strconv.Atoi(strings.TrimRight(value, "%"))
		if err == nil {
			margin.MarginTopValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
			margin.MarginTop = percentValue
		}
		remValue, err := strconv.Atoi(strings.TrimRight(value, "rem"))
		if err == nil {
			margin.MarginTopValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
			margin.MarginTop = remValue * 16
		}
	}
}

func setMarginBottom(margin *structs.Margin, value string) {
	if strings.Contains(value, "!important") {
		value = strings.ReplaceAll(value, "!important", "")
	}
	if value == "auto" {
		margin.MarginBottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	} else {
		intValue, err := strconv.Atoi(strings.TrimRight(value, "px"))
		if err == nil {
			margin.MarginBottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			margin.MarginBottom = intValue
		}
		percentValue, err := strconv.Atoi(strings.TrimRight(value, "%"))
		if err == nil {
			margin.MarginBottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
			margin.MarginBottom = percentValue
		}
		remValue, err := strconv.Atoi(strings.TrimRight(value, "rem"))
		if err == nil {
			margin.MarginBottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
			margin.MarginBottom = remValue * 16
		}
	}
}

func setMarginLeft(margin *structs.Margin, value string) {
	if value == "auto" {
		margin.MarginLeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	} else {
		intValue, err := strconv.Atoi(strings.TrimRight(value, "px"))
		if err == nil {
			margin.MarginLeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			margin.MarginLeft = intValue
		}
		percentValue, err := strconv.Atoi(strings.TrimRight(value, "%"))
		if err == nil {
			margin.MarginLeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
			margin.MarginLeft = percentValue
		}
		remValue, err := strconv.Atoi(strings.TrimRight(value, "rem"))
		if err == nil {
			margin.MarginLeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
			margin.MarginLeft = remValue * 16
		}
	}
}

func setMarginRight(margin *structs.Margin, value string) {
	if value == "auto" {
		margin.MarginRightValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	} else {
		intValue, err := strconv.Atoi(strings.TrimRight(value, "px"))
		if err == nil {
			margin.MarginRightValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			margin.MarginRight = intValue
		}
		percentValue, err := strconv.Atoi(strings.TrimRight(value, "%"))
		if err == nil {
			margin.MarginRightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
			margin.MarginRight = percentValue
		}
		remValue, err := strconv.Atoi(strings.TrimRight(value, "rem"))
		if err == nil {
			margin.MarginRightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
			margin.MarginRight = remValue * 16
		}
	}
}

func setMargin(margin *structs.Margin, value string) {
	values := strings.Split(value, " ")
	switch len(values) {
	case 1:
		setMarginTop(margin, values[0])
		setMarginRight(margin, values[0])
		setMarginBottom(margin, values[0])
		setMarginLeft(margin, values[0])
	case 2:
		setMarginTop(margin, values[0])
		setMarginRight(margin, values[1])
		setMarginBottom(margin, values[0])
		setMarginLeft(margin, values[1])
	case 3:
		setMarginTop(margin, values[0])
		setMarginRight(margin, values[1])
		setMarginBottom(margin, values[2])
		setMarginLeft(margin, values[1])
	case 4:
		setMarginTop(margin, values[0])
		setMarginRight(margin, values[1])
		setMarginBottom(margin, values[2])
		setMarginLeft(margin, values[3])
	default:
		// Do nothing
	}
}

func MarginTopPropertySetValue(currentWidget *StyleProperty, value string) {
	if value == "inherit" {
		if !currentWidget.MarginInherit {
			if currentWidget.Margin == nil {
				currentWidget.Margin = new(structs.Margin)
			}
			currentWidget.Margin.MarginTopInherit = true
		}
	} else {
		if currentWidget.Margin == nil {
			currentWidget.Margin = new(structs.Margin)
		}
		if currentWidget.MarginInherit {
			currentWidget.Margin.MarginBottomInherit = true
			currentWidget.Margin.MarginLeftInherit = true
			currentWidget.Margin.MarginRightInherit = true
		}
		currentWidget.Margin.MarginTopInherit = false
		if value == "initial" {
			currentWidget.Margin.MarginTopValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			currentWidget.Margin.MarginTop = 0
		} else {
			setMarginTop(currentWidget.Margin, value)
		}
	}
}

func MarginBottomPropertySetValue(currentWidget *StyleProperty, value string) {
	if value == "inherit" {
		if !currentWidget.MarginInherit {
			if currentWidget.Margin == nil {
				currentWidget.Margin = new(structs.Margin)
			}
			currentWidget.Margin.MarginBottomInherit = true
		}
	} else {
		if currentWidget.Margin == nil {
			currentWidget.Margin = new(structs.Margin)
		}
		if currentWidget.MarginInherit {
			currentWidget.Margin.MarginTopInherit = true
			currentWidget.Margin.MarginLeftInherit = true
			currentWidget.Margin.MarginRightInherit = true
		}
		currentWidget.Margin.MarginBottomInherit = false
		if value == "initial" {
			currentWidget.Margin.MarginBottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			currentWidget.Margin.MarginBottom = 0
		} else {
			setMarginBottom(currentWidget.Margin, value)
		}
	}
}

func MarginLeftPropertyValue(currentWidget *StyleProperty, value string) {
	if strings.Contains(value, "!important") {
		value = strings.ReplaceAll(value, "!important", "")
	}
	if value == "inherit" {
		if !currentWidget.MarginInherit {
			if currentWidget.Margin == nil {
				currentWidget.Margin = new(structs.Margin)
			}
			currentWidget.Margin.MarginLeftInherit = true
		}
	} else {
		if currentWidget.Margin == nil {
			currentWidget.Margin = new(structs.Margin)
		}
		if currentWidget.MarginInherit {
			currentWidget.Margin.MarginTopInherit = true
			currentWidget.Margin.MarginBottomInherit = true
			currentWidget.Margin.MarginRightInherit = true
		}
		currentWidget.Margin.MarginLeftInherit = false
		if value == "initial" {
			currentWidget.Margin.MarginLeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			currentWidget.Margin.MarginLeft = 0
		} else {
			setMarginLeft(currentWidget.Margin, value)
		}
	}
}

func MarginRightPropertySetValue(currentWidget *StyleProperty, value string) {
	if strings.Contains(value, "!important") {
		value = strings.ReplaceAll(value, "!important", "")
	}
	if value == "inherit" {
		if !currentWidget.MarginInherit {
			if currentWidget.Margin == nil {
				currentWidget.Margin = new(structs.Margin)
			}
			currentWidget.Margin.MarginRightInherit = true
		}
	} else {
		if currentWidget.Margin == nil {
			currentWidget.Margin = new(structs.Margin)
		}
		if currentWidget.MarginInherit {
			currentWidget.Margin.MarginTopInherit = true
			currentWidget.Margin.MarginBottomInherit = true
			currentWidget.Margin.MarginLeftInherit = true
		}
		currentWidget.Margin.MarginRightInherit = false
		if value == "initial" {
			currentWidget.Margin.MarginRightValueType = enums.CSS_PROPERTY_VALUE_TYPE_LENGTH
			currentWidget.Margin.MarginRight = 0
		} else {
			setMarginRight(currentWidget.Margin, value)
		}
	}
}

func MarginPropertySetValue(currentWidget *StyleProperty, value string) {
	if value == "inherit" {
		currentWidget.MarginInherit = true
	} else {
		currentWidget.MarginInherit = false
		if currentWidget.Margin == nil {
			currentWidget.Margin = new(structs.Margin)
		}
		if value == "initial" {
			currentWidget.Margin.MarginTop = 0
			currentWidget.Margin.MarginBottom = 0
			currentWidget.Margin.MarginLeft = 0
			currentWidget.Margin.MarginRight = 0
			currentWidget.Margin.MarginTopValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
			currentWidget.Margin.MarginBottomValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
			currentWidget.Margin.MarginLeftValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
			currentWidget.Margin.MarginRightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
		} else {
			setMargin(currentWidget.Margin, value)
		}
	}
}
