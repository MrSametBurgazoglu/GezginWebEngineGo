package structs

import (
	"gezgin_web_engine/css_scraper/structs"
	"gezgin_web_engine/html_scraper/HtmlElementWidget"
	"gezgin_web_engine/html_scraper/HtmlTags"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type BaseDrawableWidget interface {
	HtmlElementWidget.HtmlElementWidgetInterface
	RenderWidget(*sdl.Renderer)
	DrawWidget(*sdl.Renderer)
	GetRect() *sdl.Rect
	GetFont() *ttf.Font
	GetHtmlTag() HtmlTags.HtmlTags
	GetRendered() bool
}

type DrawableWidget interface {
	BaseDrawableWidget
	GetCssProperties() *structs.CssProperties
}

type UntaggedTextDrawableWidget interface {
	BaseDrawableWidget
}
