package tags

import (
	"gezgin_web_engine/css_scraper/structs"
	structs2 "gezgin_web_engine/drawer/structs"
	"gezgin_web_engine/html_scraper/HtmlElementWidget"
	"gezgin_web_engine/html_scraper/widget"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagA struct {
	*widget.BaseWidget
	download string
	href     string
	hrefLang string
	*structs2.DrawProperties
	*structs.CssProperties
}

func (receiver *HtmlTagA) RenderWidget(renderer *sdl.Renderer) {
}

func (receiver *HtmlTagA) DrawWidget(renderer *sdl.Renderer) {
}

func (receiver *HtmlTagA) VarReader(variableName, variableValue string) {
	switch variableName {
	case "download":
		receiver.download = variableValue
	case "href":
		receiver.href = variableValue
	case "hrefLang":
		receiver.hrefLang = variableValue
	}
}

func SetWidgetPropertiesForATag() (newWidget HtmlElementWidget.HtmlElementWidgetInterface) {
	newWidget = &HtmlTagA{DrawProperties: new(structs2.DrawProperties), BaseWidget: &widget.BaseWidget{Draw: true}, CssProperties: new(structs.CssProperties)}
	return
}
