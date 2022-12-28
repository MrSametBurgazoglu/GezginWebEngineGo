package drawerBackend

import (
	"gezgin_web_engine/css_scraper/structs"
	structs2 "gezgin_web_engine/drawer/structs"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawBackground(cssProperties *structs.CssProperties, drawProperties *structs2.DrawProperties, renderer *sdl.Renderer) {
	if cssProperties.Background.BackgroundColor != nil {
		alpha, red, green, blue := cssProperties.Background.BackgroundColor.GetColorByRGBA()
		renderer.SetDrawColor(red, green, blue, alpha)
		renderer.FillRect(&drawProperties.Rect)
	}
}
