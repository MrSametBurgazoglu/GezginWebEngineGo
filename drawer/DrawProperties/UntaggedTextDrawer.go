package DrawProperties

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/drawer/structs"
	"gezgin_web_engine/html_scraper"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawUntaggedTextFunction(widget *htmlVariables.Widget, renderer *sdl.Renderer) {
	renderer.Copy(widget.DrawProperties.Texture, nil, &widget.DrawProperties.Rect)
}

func RenderUntaggedTextFunction(widget *htmlVariables.Widget, renderer *sdl.Renderer) {
	if widget.DrawProperties == nil {
		widget.DrawProperties = new(structs.DrawProperties)
	}
	drawText, ok := widget.WidgetProperties.(html_scraper.UntaggedText)
	if ok {
		drawerBackend.GetTextTexture(
			renderer,
			drawText.Value,
			widget.CssProperties.Color,
			widget.Parent.DrawProperties.Font,
			widget.DrawProperties.Texture,
			&widget.DrawProperties.Rect,
		)
		//widget.DrawProperties.Rect.X = CalculateXPos(widget)
		//widget.DrawProperties.Rect.Y = CalculateYPos(widget)
	}

}
