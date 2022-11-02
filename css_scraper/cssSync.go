package css_scraper

import "gezgin_web_engine/css_scraper/structs"
import "gezgin_web_engine/css_scraper/properties"

func computeInheritCssProperties(dest *structs.CssProperties, source *structs.CssProperties) {
	properties.UpdateBackground(dest, source)
	properties.UpdateColor(dest, source)
	/*
		updateAccentColor(dest, source)
		updateAlign(dest, source)
		updateAnimation(dest, source)
		updateBackdrop_filter(dest, source)
		updateBackface_visibility(dest, source)
		updateBackground(dest, source)
		updateBorder(dest, source)
		updateColor(dest, source)
		updateColumn(dest, source)
		updateFlex(dest, source)
		updateFont(dest, source)
		updateHeight(dest, source)
		updateMargin(dest, source)
		updateOpacity(dest, source)
		updateOutline(dest, source)
		updateOverflow(dest, source)
		updatePadding(dest, source)
		updatePosition(dest, source)
		updateResize(dest, source)
		updateText(dest, source)
		updateVisibility(dest, source)
		updateWidth(dest, source)
	*/
}

func updateCssProperties(dest *structs.CssProperties, source *structs.CssProperties) {
	/*
		updateAccentColor(dest, source)
		updateAlign(dest, source)
		updateAnimation(dest, source)
		updateBackdrop_filter(dest, source)
		updateBackface_visibility(dest, source)
		updateBackground(dest, source)
		updateBorder(dest, source)
		updateColor(dest, source)
		updateColumn(dest, source)
		updateFlex(dest, source)
		updateFont(dest, source)
		updateHeight(dest, source)
		updateMargin(dest, source)
		updateOpacity(dest, source)
		updateOutline(dest, source)
		updateOverflow(dest, source)
		updatePadding(dest, source)
		updatePosition(dest, source)
		updateResize(dest, source)
		updateText(dest, source)
		updateVisibility(dest, source)
		updateWidth(dest, source)
	*/
}
