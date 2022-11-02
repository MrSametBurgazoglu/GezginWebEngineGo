package DrawProperties

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/html_scraper/tags"
	"gezgin_web_engine/html_scraper/widget"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawUntaggedTextFunction(widget *widget.Widget, renderer *sdl.Renderer) {
	renderer.Copy(widget.DrawProperties.Texture, nil, &widget.DrawProperties.Rect)
	println("çizdim ya la")
}

func RenderUntaggedTextFunction(widget *widget.Widget, renderer *sdl.Renderer) {
	println("renderladım ya")
	drawText, ok := widget.WidgetProperties.(tags.UntaggedText)
	if ok {
		drawerBackend.GetTextTexture(
			renderer,
			drawText.Value,
			widget.Parent.CssProperties.Color,
			widget.Parent.DrawProperties.Font,
			&widget.DrawProperties.Texture,
			&widget.DrawProperties.Rect,
		)
		widget.DrawProperties.Rect.X = 0
		widget.DrawProperties.Rect.Y = 0
		println(widget.DrawProperties.Rect.H)
		println(widget.DrawProperties.Rect.W)
		println(drawText.Value)
		println("render Untagged text")
		//widget.DrawProperties.Rect.X = CalculateXPos(widget)
		//widget.DrawProperties.Rect.Y = CalculateYPos(widget)
	}

}
