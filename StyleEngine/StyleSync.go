package StyleEngine

import StyleProperty2 "gezgin_web_engine/StyleProperty"

func computeInheritCssProperties(dest *StyleProperty2.StyleProperty, source *StyleProperty2.StyleProperty) {
	StyleProperty2.UpdateBackground(dest, source)
	StyleProperty2.UpdateColor(dest, source)
}

func updateCssProperties(dest *StyleProperty2.StyleProperty, source *StyleProperty2.StyleProperty) {
	StyleProperty2.UpdateBackground(dest, source)
	StyleProperty2.UpdatePosition(dest, source)
	StyleProperty2.UpdateColor(dest, source)
}
