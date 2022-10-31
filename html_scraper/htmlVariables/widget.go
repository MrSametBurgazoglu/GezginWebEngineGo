package htmlVariables

import (
	"gezgin_web_engine/css_scraper/structs"
	structs2 "gezgin_web_engine/drawer/structs"
	"github.com/veandco/go-sdl2/sdl"
)

type Widget struct {
	ChildrenCount         int
	ChildrenIndex         int
	HtmlTag               HtmlTags
	WidgetProperties      any
	StandardHtmlVariables StandardHtmlTagVariables
	CssProperties         *structs.CssProperties
	DrawProperties        *structs2.DrawProperties
	VarReaderFunc         func(*Widget, string, string)
	ContextReaderFunc     func(*Widget, string)
	RenderWidget          func(*Widget, *sdl.Renderer)
	DrawWidget            func(*Widget, *sdl.Renderer)
	Children              []*Widget
	Parent                *Widget
	Draw                  bool
}
