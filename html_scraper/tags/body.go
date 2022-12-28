package tags

import (
	"gezgin_web_engine/css_scraper/structs"
	structs2 "gezgin_web_engine/drawer/structs"
	"gezgin_web_engine/html_scraper/HtmlElementWidget"
	"gezgin_web_engine/html_scraper/widget"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagBody struct {
	*widget.BaseWidget
	*structs2.DrawProperties
	*structs.CssProperties
}

func (receiver *HtmlTagBody) RenderWidget(renderer *sdl.Renderer) {
}

func (receiver *HtmlTagBody) DrawWidget(renderer *sdl.Renderer) {
}

func SetWidgetPropertiesForBodyTag() (newWidget HtmlElementWidget.HtmlElementWidgetInterface) {
	newWidget = &HtmlTagBody{DrawProperties: new(structs2.DrawProperties), BaseWidget: &widget.BaseWidget{Draw: true}, CssProperties: new(structs.CssProperties)}
	return
}
