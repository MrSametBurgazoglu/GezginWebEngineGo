package properties

import (
	"gezgin_web_engine/css_scraper/structs"
)

func ColorPropertySetValue(properties *structs.CssProperties, value string) {
	if value == "inherit" {
		properties.ColorInherit = true
	} else {
		properties.ColorInherit = false
		/*
			if properties.Color == nil {
				properties.Color = new()
			}
		*/
	}
}

func UpdateColor(properties *structs.CssProperties, source *structs.CssProperties) {
	if source.ColorInherit {
		properties.ColorInherit = true
		properties.Color = nil
	} else if source.Color != nil {
		properties.Color = source.Color
	}
}
