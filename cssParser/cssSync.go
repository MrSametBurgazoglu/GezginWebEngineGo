package cssParser

import "gezgin_web_engine/cssParser/structs"
import "gezgin_web_engine/cssParser/properties"

func computeInheritCssProperties(dest *structs.CssProperties, source *structs.CssProperties) {
	properties.UpdateBackground(dest, source)
	properties.UpdateColor(dest, source)
}

func updateCssProperties(dest *structs.CssProperties, source *structs.CssProperties) {
	properties.UpdateBackground(dest, source)
	properties.UpdatePosition(dest, source)
	properties.UpdateColor(dest, source)
}
