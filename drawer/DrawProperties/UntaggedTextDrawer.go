package DrawProperties

import (
	"gezgin_web_engine/HtmlParser/tags"
	"gezgin_web_engine/drawer/Fonts"
	"gezgin_web_engine/drawer/drawerBackend"
	tags2 "gezgin_web_engine/widgets"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"strings"
)

func findLastSpace(text string, last int) int {
	for i := last - 1; i > 0; i-- {
		if text[i] == ' ' {
			return i
		}
	}
	return last
}

func splitTextAndRenderByLines(text string, renderer *sdl.Renderer, font *ttf.Font, maxWidth int) string {
	println(text)
	var Lines []string
	var err error
	var currentWidth, _, start, end int
	length := len(text)
	start = 0
	end = length
	for start < length {
		currentWidth, _, err = font.SizeUTF8(text[start:end])
		for currentWidth > maxWidth {
			end = findLastSpace(text, end)
			currentWidth, _, err = font.SizeUTF8(text[start:end])
			if err != nil {
				panic(err)
			}
		}
		Lines = append(Lines, text[start:end])
		start = end + 1
		end = length
	}
	return strings.Join(Lines, "\n")
}

func DrawUntaggedTextFunction(widget *tags2.Widget, renderer *sdl.Renderer) {
	renderer.Copy(widget.DrawProperties.Texture, nil, &widget.DrawProperties.Rect)
}

func RenderUntaggedTextFunction(widget *tags2.Widget, renderer *sdl.Renderer) {
	drawText, ok := widget.WidgetProperties.(tags.UntaggedText)
	if widget.Parent.DrawProperties.Font == nil {
		if widget.Parent.CssProperties.Font != nil {
			widget.Parent.DrawProperties.Font = Fonts.GetFont(widget.Parent.CssProperties.Font.FontSizeValue)
		} else {
			widget.Parent.DrawProperties.Font = Fonts.GetFont(14)
		}
	}
	if ok {
		if currentWidth, _, _ := widget.Parent.DrawProperties.Font.SizeUTF8(drawText.Value); currentWidth > int(widget.Parent.DrawProperties.Rect.W) {
			Lines := splitTextAndRenderByLines(drawText.Value, renderer, widget.Parent.DrawProperties.Font, int(widget.Parent.DrawProperties.Rect.W))
			drawerBackend.GetTextTexture(
				renderer,
				Lines,
				widget.Parent.CssProperties.Color,
				widget.Parent.DrawProperties.Font,
				&widget.DrawProperties.Texture,
				&widget.DrawProperties.Rect,
			)
		} else {
			drawerBackend.GetTextTexture(
				renderer,
				drawText.Value,
				widget.Parent.CssProperties.Color,
				widget.Parent.DrawProperties.Font,
				&widget.DrawProperties.Texture,
				&widget.DrawProperties.Rect,
			)
		}
	}
	if widget.DrawProperties.Rect.W > widget.Parent.DrawProperties.Rect.W {
		println("bigger than parent")
		println(widget.DrawProperties.Rect.W)
		Lines := splitTextAndRenderByLines(drawText.Value, renderer, widget.Parent.DrawProperties.Font, int(widget.Parent.DrawProperties.Rect.W))
		println(Lines)
	}
}
