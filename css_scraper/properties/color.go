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
