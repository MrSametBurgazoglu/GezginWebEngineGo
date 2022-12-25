package widget

import (
	"gezgin_web_engine/css_scraper/structs"
	structs2 "gezgin_web_engine/drawer/structs"
	"gezgin_web_engine/html_scraper/HtmlTags"
	"gezgin_web_engine/html_scraper/htmlVariables/standardHtmlTagVariables"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagsInterface interface {
	SetHtmlTag() int
}

type Widget struct {
	ChildrenCount         int
	ChildrenIndex         int
	HtmlTag               HtmlTags.HtmlTags
	WidgetProperties      any
	StandardHtmlVariables standardHtmlTagVariables.StandardHtmlTagVariables
	CssProperties         *structs.CssProperties
	DrawProperties        *structs2.DrawProperties
	VarReaderFunc         func(*Widget, string, string)
	ContextReaderFunc     func(*Widget, string)
	RenderWidget          func(*Widget, *sdl.Renderer)
	DrawWidget            func(*Widget, *sdl.Renderer)
	Children              []*Widget
	Parent                *Widget
	Draw                  bool
	Rendered              bool
}
