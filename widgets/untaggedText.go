package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/drawer/Fonts"
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/drawer/structs"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"strings"
)

type UntaggedText struct {
	Widget
	Value string
}

func (receiver *UntaggedText) Draw(renderer *sdl.Renderer) {
	renderer.Copy(receiver.DrawProperties.Texture, nil, &receiver.DrawProperties.Rect)
}

func (receiver *UntaggedText) Render(renderer *sdl.Renderer, resourceManager *ResourceManager.ResourceManager) {
	if receiver.GetParent().GetDrawProperties().Font == nil {
		if receiver.GetParent().GetStyleProperty().Font != nil {
			receiver.GetParent().GetDrawProperties().Font = Fonts.GetFont(receiver.GetParent().GetStyleProperty().Font.FontSizeValue)
		} else {
			receiver.GetParent().GetDrawProperties().Font = Fonts.GetFont(14)
		}
	}
	if currentWidth, _, _ := receiver.GetParent().GetDrawProperties().Font.SizeUTF8(receiver.Value); currentWidth > int(receiver.GetParent().GetDrawProperties().Rect.W) {
		Lines := splitTextAndRenderByLines(receiver.Value, renderer, receiver.GetParent().GetDrawProperties().Font, int(receiver.GetParent().GetDrawProperties().Rect.W))
		drawerBackend.GetTextTexture(
			renderer,
			Lines,
			receiver.GetParent().GetStyleProperty().Color,
			receiver.GetParent().GetDrawProperties().Font,
			&receiver.GetDrawProperties().Texture,
			&receiver.GetDrawProperties().Rect,
		)
	} else {
		drawerBackend.GetTextTexture(
			renderer,
			receiver.Value,
			receiver.GetParent().GetStyleProperty().Color,
			receiver.GetParent().GetDrawProperties().Font,
			&receiver.GetDrawProperties().Texture,
			&receiver.GetDrawProperties().Rect,
		)
	}

	if receiver.GetDrawProperties().Rect.W > receiver.GetParent().GetDrawProperties().Rect.W {
		println("bigger than parent")
		println(receiver.GetDrawProperties().Rect.W)
		Lines := splitTextAndRenderByLines(receiver.Value, renderer, receiver.GetParent().GetDrawProperties().Font, int(receiver.GetParent().GetDrawProperties().Rect.W))
		println(Lines)
	}
}

func SetWidgetPropertiesForUntaggedText(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(UntaggedText)
	widget.HtmlElement = element
	widget.DrawProperties = new(structs.DrawProperties)
	widget.Value = element.Text
	return widget
}

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
