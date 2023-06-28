package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/StyleEngine/structs"
	"gezgin_web_engine/drawer/drawerBackend"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagHeader struct {
	Widget
}

func (receiver *HtmlTagHeader) Draw(renderer *sdl.Renderer) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, &receiver.DrawProperties.Rect, renderer)
	}
}

func (receiver *HtmlTagHeader) Render(renderer *sdl.Renderer) {

}

func SetWidgetPropertiesForH1Tag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 24
	return widget
}

func SetWidgetPropertiesForH2Tag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 20
	return widget
}

func SetWidgetPropertiesForH3Tag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.HtmlElement = element
	widget.StyleProperty = new(StyleEngine.StyleProperty)
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 12
	return widget
}

func SetWidgetPropertiesForH4Tag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.HtmlElement = element
	widget.Initialize()
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 10
	return widget
}

func SetWidgetPropertiesForH5Tag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.HtmlElement = element
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 8
	return widget
}

func SetWidgetPropertiesForH6Tag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.HtmlElement = element
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 6
	return widget
}