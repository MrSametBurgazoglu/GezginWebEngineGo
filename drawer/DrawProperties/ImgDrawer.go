package DrawProperties

import (
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/htmlParser/tags"
	tags2 "gezgin_web_engine/widgets"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawImgFunction(widget *tags2.Widget, renderer *sdl.Renderer) {
	renderer.Copy(widget.DrawProperties.Texture, nil, &widget.DrawProperties.Rect)
}

func RenderImgFunction(widget *tags2.Widget, renderer *sdl.Renderer) {
	drawImg, ok := widget.WidgetProperties.(*tags.HtmlTagImg)
	if ok {
		drawerBackend.GetImageTexture(
			renderer,
			drawImg.Src,
			&widget.DrawProperties.Texture,
			&widget.DrawProperties.Rect,
		)
	}
}
