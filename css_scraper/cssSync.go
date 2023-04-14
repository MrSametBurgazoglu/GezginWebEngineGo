package css_scraper

import "gezgin_web_engine/css_scraper/structs"
import "gezgin_web_engine/css_scraper/properties"

func computeInheritCssProperties(dest *structs.CssProperties, source *structs.CssProperties) {
	properties.UpdateBackground(dest, source)
	properties.UpdateColor(dest, source)
}

func updateCssProperties(dest *structs.CssProperties, source *structs.CssProperties) {
	properties.UpdateBackground(dest, source)
	properties.UpdatePosition(dest, source)
}
