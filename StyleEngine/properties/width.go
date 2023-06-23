package properties

import (
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/StyleEngine/enums"
	"strconv"
	"strings"
)

func SetWidth(cssProperties *StyleEngine.StyleProperty, value string) {
	if strings.HasSuffix(value, "px") {
		width, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err != nil {
			// handle error
		}
		cssProperties.Width = uint(width)
		cssProperties.WidthValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
	} else if strings.HasSuffix(value, "%") {
		width, err := strconv.Atoi(strings.TrimSuffix(value, "%"))
		if err != nil {
			// handle error
		}
		cssProperties.Width = uint(width)
		cssProperties.WidthValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
	} else {
		// DEFAULT VALUE
		cssProperties.WidthValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
	}
}

func SetMinWidth(cssProperties *StyleEngine.StyleProperty, value string) {
	if strings.HasSuffix(value, "px") {
		width, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err != nil {
			// handle error
		}
		cssProperties.MinWidth = uint(width)
		cssProperties.MinWidthValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
	} else if strings.HasSuffix(value, "%") {
		width, err := strconv.Atoi(strings.TrimSuffix(value, "%"))
		if err != nil {
			// handle error
		}
		cssProperties.MinWidth = uint(width)
		cssProperties.MinWidthValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
	} else {
		// DEFAULT VALUE
		cssProperties.MinWidth = 0
	}
}

func SetMaxWidth(cssProperties *StyleEngine.StyleProperty, value string) {
	if strings.HasSuffix(value, "px") {
		width, err := strconv.Atoi(strings.TrimSuffix(value, "px"))
		if err != nil {
			// handle error
		}
		cssProperties.MaxWidth = uint(width)
		cssProperties.MaxWidthValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
	} else if strings.HasSuffix(value, "%") {
		width, err := strconv.Atoi(strings.TrimSuffix(value, "%"))
		if err != nil {
			// handle error
		}
		cssProperties.MaxWidth = uint(width)
		cssProperties.MaxWidthValueType = enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
	} else {
		// DEFAULT VALUE
		cssProperties.MaxWidthValueType = enums.CSS_PROPERTY_VALUE_TYPE_NONE
	}
}

func WidthPropertySetValue(cssProperties *StyleEngine.StyleProperty, value string) {
	if value == "inherit" {
		cssProperties.WidthInherit = true
	} else {
		cssProperties.WidthInherit = false
		if value == "initial" {
			cssProperties.WidthValueType = enums.CSS_PROPERTY_VALUE_TYPE_AUTO
		} else {
			SetWidth(cssProperties, value)
		}
	}
}

func MinWidthPropertySetValue(cssProperties *StyleEngine.StyleProperty, value string) {
	if value == "inherit" {
		cssProperties.MinWidthInherit = true
	} else {
		cssProperties.MinWidthInherit = false
		if value == "initial" {
			cssProperties.MinWidth = 0
			cssProperties.MinWidthValueType = enums.CSS_PROPERTY_VALUE_TYPE_PIXEL
		} else {
			SetMinWidth(cssProperties, value)
		}
	}
}

func MaxWidthPropertySetValue(cssProperties *StyleEngine.StyleProperty, value string) {
	if value == "inherit" {
		cssProperties.MaxWidthInherit = true
	} else {
		cssProperties.MaxWidthInherit = false
		if value == "initial" {
			cssProperties.MaxWidth = 0
			cssProperties.MaxWidthValueType = enums.CSS_PROPERTY_VALUE_TYPE_NONE
		} else {
			SetMaxWidth(cssProperties, value)
		}
	}
}

func ComputeInheritWidth(dest, source *StyleEngine.StyleProperty) {
	if dest.WidthInherit {
		dest.Width = source.Width
		dest.WidthValueType = source.WidthValueType
	}
	if dest.MinWidthInherit {
		dest.MinWidth = source.MinWidth
		dest.MinWidthValueType = source.MinWidthValueType
	}
	if dest.MaxWidthInherit {
		dest.MaxWidth = source.MaxWidth
		dest.MaxWidthValueType = source.MaxWidthValueType
	}
}

func UpdateWidth(cssProperties, source *StyleEngine.StyleProperty) {
	if source.WidthInherit {
		cssProperties.WidthInherit = true
	} else if cssProperties.WidthValueType != enums.CSS_PROPERTY_VALUE_TYPE_EMPTY {
		cssProperties.Width = source.Width
		cssProperties.WidthValueType = source.WidthValueType
	}
	if source.MaxWidthInherit {
		cssProperties.MaxWidthInherit = true
	} else if cssProperties.MaxWidthValueType != enums.CSS_PROPERTY_VALUE_TYPE_EMPTY {
		cssProperties.MaxWidth = source.MaxWidth
		cssProperties.MaxWidthValueType = source.MaxWidthValueType
	}
	if source.MinWidthInherit {
		cssProperties.MinWidthInherit = true
	} else if cssProperties.MinWidthValueType != enums.CSS_PROPERTY_VALUE_TYPE_EMPTY {
		cssProperties.MinWidth = source.MinWidth
		cssProperties.MinWidthValueType = source.MinWidthValueType
	}
}
