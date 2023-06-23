package StyleEngine

import (
	properties2 "gezgin_web_engine/StyleEngine/properties"
)

func computeInheritCssProperties(dest *StyleProperty, source *StyleProperty) {
	properties2.UpdateBackground(dest, source)
	properties2.UpdateColor(dest, source)
}

func updateCssProperties(dest *StyleProperty, source *StyleProperty) {
	properties2.UpdateBackground(dest, source)
	properties2.UpdatePosition(dest, source)
	properties2.UpdateColor(dest, source)
}
