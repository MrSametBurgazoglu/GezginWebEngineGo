package StyleEngine

func computeInheritCssProperties(dest *StyleProperty, source *StyleProperty) {
	UpdateBackground(dest, source)
	UpdateColor(dest, source)
}

func updateCssProperties(dest *StyleProperty, source *StyleProperty) {
	UpdateBackground(dest, source)
	UpdatePosition(dest, source)
	UpdateColor(dest, source)
}
