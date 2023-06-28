package properties

import (
	"gezgin_web_engine/cssParser/structs"
)

func accentColorPropertySetValue(currentWidget *structs.CssProperties, value string) {
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

func computeInheritAccentColor(dest, source *structs.CssProperties) {
	if dest.AccentColor == nil {
		dest.AccentColor = source.AccentColor
	}
}

func updateAccentColor(currentWidget, source *structs.CssProperties) {
	if source.AccentColorInherit {
		currentWidget.AccentColorInherit = true
		currentWidget.AccentColor = nil
	} else if source.AccentColor != nil {
		currentWidget.AccentColor = source.AccentColor
	}
}