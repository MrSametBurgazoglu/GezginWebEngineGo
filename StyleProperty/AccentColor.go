package StyleProperty

import (
	"gezgin_web_engine/StyleProperty/structs"
)

func accentColorPropertySetValue(currentWidget *StyleProperty, value string) {
	if value == "inherit" {
		currentWidget.AccentColorInherit = true
	} else {
		currentWidget.AccentColorInherit = false
		if currentWidget.AccentColor == nil || currentWidget.AccentColorInherit {
			currentWidget.AccentColor = new(structs.ColorRGBA)
		}
		if value == "auto" || value == "initial" {
			currentWidget.AccentColor.SetColorByRGB(0, 0, 0)
		} else {
			currentWidget.AccentColor.SetColor(value)
		}
	}
}

func computeInheritAccentColor(dest, source *StyleProperty) {
	if dest.AccentColor == nil {
		dest.AccentColor = source.AccentColor
	}
}

func updateAccentColor(currentWidget, source *StyleProperty) {
	if source.AccentColorInherit {
		currentWidget.AccentColorInherit = true
		currentWidget.AccentColor = nil
	} else if source.AccentColor != nil {
		currentWidget.AccentColor = source.AccentColor
	}
}
