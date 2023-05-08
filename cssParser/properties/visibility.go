package properties

import (
	"gezgin_web_engine/cssParser/enums"
	"gezgin_web_engine/cssParser/structs"
)

func setVisibility(cssProperties *structs.CssProperties, value string) {
	if value == "hidden" {
		cssProperties.Visibility = enums.CSS_VISIBILITY_HIDDEN
	} else if value == "collapse" {
		cssProperties.Visibility = enums.CSS_VISIBILITY_COLLAPSE
	} else {
		cssProperties.Visibility = enums.CSS_VISIBILITY_VISIBLE
	}
}

func VisibilityPropertySetValue(cssProperties *structs.CssProperties, value string) {
	if value == "inherit" {
		cssProperties.VisibilityInherit = true
	} else {
		cssProperties.VisibilityInherit = false
		if value == "initial" {
			cssProperties.Visibility = enums.CSS_VISIBILITY_VISIBLE
		} else {
			setVisibility(cssProperties, value)
		}
	}
}

func computeInheritVisibility(dest, source *structs.CssProperties) {
	if dest.Visibility == enums.CSS_VISIBILITY_EMPTY {
		dest.Visibility = source.Visibility
	}
}

func updateVisibility(cssProperties, source *structs.CssProperties) {
	if source.VisibilityInherit {
		cssProperties.VisibilityInherit = true
	} else if source.Visibility != enums.CSS_VISIBILITY_EMPTY {
		cssProperties.Visibility = source.Visibility
	}
}
