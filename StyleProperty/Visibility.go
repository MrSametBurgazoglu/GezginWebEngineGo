package StyleProperty

import (
	"gezgin_web_engine/StyleProperty/enums"
)

func setVisibility(cssProperties *StyleProperty, value string) {
	if value == "hidden" {
		cssProperties.Visibility = enums.CSS_VISIBILITY_HIDDEN
	} else if value == "collapse" {
		cssProperties.Visibility = enums.CSS_VISIBILITY_COLLAPSE
	} else {
		cssProperties.Visibility = enums.CSS_VISIBILITY_VISIBLE
	}
}

func VisibilityPropertySetValue(cssProperties *StyleProperty, value string) {
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

func computeInheritVisibility(dest, source *StyleProperty) {
	if dest.Visibility == enums.CSS_VISIBILITY_EMPTY {
		dest.Visibility = source.Visibility
	}
}

func updateVisibility(cssProperties, source *StyleProperty) {
	if source.VisibilityInherit {
		cssProperties.VisibilityInherit = true
	} else if source.Visibility != enums.CSS_VISIBILITY_EMPTY {
		cssProperties.Visibility = source.Visibility
	}
}
