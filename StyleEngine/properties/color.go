package properties

import (
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/StyleEngine/structs"
)

func ColorPropertySetValue(properties *StyleEngine.StyleProperty, value string) {
	if value == "inherit" {
		properties.ColorInherit = true
	} else {
		properties.ColorInherit = false
		if properties.Color == nil {
			properties.Color = new(structs.ColorRGBA)
		}
		if value == "auto" || value == "initial" {
			properties.Color.SetColorByRGB(0, 0, 0)
		} else {
			properties.Color.SetColor(value)
		}
	}
}

func UpdateColor(properties *StyleEngine.StyleProperty, source *StyleEngine.StyleProperty) {
	if source.ColorInherit && properties.Color == nil {
		properties.ColorInherit = true
		properties.Color = nil
	} else if source.Color != nil && properties.Color == nil {
		properties.Color = source.Color
	}
}
