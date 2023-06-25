package widgets

import (
	structs2 "gezgin_web_engine/StyleEngine/structs"
	"gezgin_web_engine/drawer/drawerBackend"
	"github.com/veandco/go-sdl2/sdl"
)

type HtmlTagP struct {
	*Widget
}

func (receiver *HtmlTagP) Draw(renderer *sdl.Renderer) {
	if receiver.GetStyleProperty().Background != nil {
		drawerBackend.DrawBackground(receiver, renderer)
	}
}

func (receiver *HtmlTagP) Render(renderer *sdl.Renderer) {

}

func SetWidgetPropertiesForPTag() WidgetInterface {
	widget := new(HtmlTagP)
	widget.StyleProperty.Font = new(structs2.Font)
	widget.StyleProperty.Font.FontSizeValue = 14
	widget.StyleProperty.Margin = new(structs2.Margin)
	widget.StyleProperty.Margin.MarginTop = 10
	return widget
}
