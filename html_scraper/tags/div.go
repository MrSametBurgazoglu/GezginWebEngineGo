package tags

import (
	"gezgin_web_engine/css_scraper/structs"
	structs2 "gezgin_web_engine/drawer/structs"
	"gezgin_web_engine/html_scraper/HtmlElementWidget"
	"gezgin_web_engine/html_scraper/widget"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagDiv struct {
	*widget.BaseWidget
	*structs2.DrawProperties
	*structs.CssProperties
}

func (receiver *HtmlTagDiv) RenderWidget(renderer *sdl.Renderer) {
}

func (receiver *HtmlTagDiv) DrawWidget(renderer *sdl.Renderer) {
}

func SetWidgetPropertiesForDivTag() (newWidget HtmlElementWidget.HtmlElementWidgetInterface) {
	newWidget = &HtmlTagDiv{DrawProperties: new(structs2.DrawProperties), BaseWidget: &widget.BaseWidget{Draw: true}, CssProperties: new(structs.CssProperties)}
	return
}
