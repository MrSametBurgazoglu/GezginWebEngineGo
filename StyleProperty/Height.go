package StyleProperty

import (
	"gezgin_web_engine/StyleProperty/enums"
	"strconv"
	"strings"
)

func SetHeight(cssProperties *StyleProperty, value string) {
	if strings.HasSuffix(value, "px") {
		height, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err != nil {
			// handle error
		}
		cssProperties.Height = uint(height)
		cssProperties.HeightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
	} else if strings.HasSuffix(value, "%") {
		height, err := strconv.Atoi(strings.TrimSuffix(value, "%"))
		if err != nil {
			// handle error
		}
		cssProperties.Height = uint(height)
		cssProperties.HeightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
	} else if strings.HasSuffix(value, "rem") {
		height, err := strconv.Atoi(strings.TrimSuffix(value, "rem"))
		if err != nil {
			// handle error
		}
		cssProperties.Height = uint(height) * 16
		cssProperties.HeightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
	} else {
		// DEFAULT VALUE
		cssProperties.HeightValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	}
}

func SetMinHeight(cssProperties *StyleProperty, value string) {
	if strings.HasSuffix(value, "px") {
		height, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err != nil {
			// handle error
		}
		cssProperties.MinHeight = uint(height)
		cssProperties.MinHeightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
	} else if strings.HasSuffix(value, "%") {
		height, err := strconv.Atoi(strings.TrimSuffix(value, "%"))
		if err != nil {
			// handle error
		}
		cssProperties.MinHeight = uint(height)
		cssProperties.MinHeightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
	} else {
		// DEFAULT VALUE
		cssProperties.MinHeight = 0
	}
}

func SetMaxHeight(cssProperties *StyleProperty, value string) {
	if strings.HasSuffix(value, "px") {
		height, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err != nil {
			// handle error
		}
		cssProperties.MaxHeight = uint(height)
		cssProperties.MaxHeightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
	} else if strings.HasSuffix(value, "%") {
		height, err := strconv.Atoi(strings.TrimSuffix(value, "%"))
		if err != nil {
			// handle error
		}
		cssProperties.MaxHeight = uint(height)
		cssProperties.MaxHeightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
	} else {
		// DEFAULT VALUE
		cssProperties.MaxHeightValueType = enums.CSS_PROPERTY_VALUE_TYPE_NONE
	}
}

func HeightPropertySetValue(cssProperties *StyleProperty, value string) {
	if strings.Contains(value, "!important") {
		value = strings.ReplaceAll(value, "!important", "")
	}
	if value == "inherit" {
		cssProperties.HeightInherit = true
	} else {
		cssProperties.HeightInherit = false
		if value == "initial" {
			cssProperties.HeightValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
		} else {
			SetHeight(cssProperties, value)
		}
	}
}

func MinHeightPropertySetValue(cssProperties *StyleProperty, value string) {
	if value == "inherit" {
		cssProperties.MinHeightInherit = true
	} else {
		cssProperties.MinHeightInherit = false
		if value == "initial" {
			cssProperties.MinHeight = 0
			cssProperties.MinHeightValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
		} else {
			SetMinHeight(cssProperties, value)
		}
	}
}

func MaxHeightPropertySetValue(cssProperties *StyleProperty, value string) {
	if value == "inherit" {
		cssProperties.MaxHeightInherit = true
	} else {
		cssProperties.MaxHeightInherit = false
		if value == "initial" {
			cssProperties.MaxHeight = 0
			cssProperties.MaxHeightValueType = enums.CSS_PROPERTY_VALUE_TYPE_NONE
		} else {
			SetMaxHeight(cssProperties, value)
		}
	}
}

func ComputeInheritHeight(dest, source *StyleProperty) {
	if dest.HeightInherit {
		dest.Height = source.Height
		dest.HeightValueType = source.HeightValueType
	}
	if dest.MinHeightInherit {
		dest.MinHeight = source.MinHeight
		dest.MinHeightValueType = source.MinHeightValueType
	}
	if dest.MaxHeightInherit {
		dest.MaxHeight = source.MaxHeight
		dest.MaxHeightValueType = source.MaxHeightValueType
	}
}

func UpdateHeight(cssProperties, source *StyleProperty) {
	if source.HeightInherit {
		cssProperties.HeightInherit = true
	} else if cssProperties.HeightValueType != enums.CSS_PROPERTY_VALUE_TYPE_EMPTY {
		cssProperties.Height = source.Height
		cssProperties.HeightValueType = source.HeightValueType
	}
	if source.MaxHeightInherit {
		cssProperties.MaxHeightInherit = true
	} else if cssProperties.MaxHeightValueType != enums.CSS_PROPERTY_VALUE_TYPE_EMPTY {
		cssProperties.MaxHeight = source.MaxHeight
		cssProperties.MaxHeightValueType = source.MaxHeightValueType
	}
	if source.MinHeightInherit {
		cssProperties.MinHeightInherit = true
	} else if cssProperties.MinHeightValueType != enums.CSS_PROPERTY_VALUE_TYPE_EMPTY {
		cssProperties.MinHeight = source.MinHeight
		cssProperties.MinHeightValueType = source.MinHeightValueType
	}
}
