package properties

import (
	"gezgin_web_engine/cssParser/structs"
	"strings"
)

func ColorPropertySetValue(properties *structs.CssProperties, value string) {
	if value == "inherit" {
		properties.ColorInherit = true
	} else {
		properties.ColorInherit = false
		if properties.Color == nil {
			properties.Color = new(structs.ColorRGBA)
		}
		if strings.Compare(value, "auto") == 0 || strings.Compare(value, "initial") == 0 {
			properties.Color.SetColorByRGB(0, 0, 0)
		} else {
			properties.Color.SetColor(value)
		}
	}
}

func UpdateColor(properties *structs.CssProperties, source *structs.CssProperties) {
	if source.ColorInherit && properties.Color == nil {
		properties.ColorInherit = true
		properties.Color = nil
	} else if source.Color != nil && properties.Color == nil {
		properties.Color = source.Color
	}
}
