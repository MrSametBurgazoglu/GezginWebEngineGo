package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/drawer/Fonts"
	"gezgin_web_engine/drawer/structs"
	"image"
	"image/draw"
)

type UntaggedText struct {
	Widget
	Value string
}

func (receiver *UntaggedText) Draw(mainImage *image.RGBA) {
	draw.Draw(mainImage, *receiver.DrawProperties.Rect, receiver.DrawProperties.Texture, image.Point{X: 0, Y: 0}, draw.Src)
}

func (receiver *UntaggedText) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {
	if receiver.GetParent().GetDrawProperties().Font == nil {
		if receiver.GetParent().GetStyleProperty().Font != nil {
			receiver.GetParent().GetDrawProperties().Font = Fonts.GetFont(receiver.GetParent().GetStyleProperty().Font.FontSizeValue)
		} else {
			receiver.GetParent().GetDrawProperties().Font = Fonts.GetFont(14)
		}
	}
	if currentWidth := int(receiver.GetParent().GetDrawProperties().Font.Size * float64(len(receiver.Value)) * 1.2); currentWidth > int(receiver.GetParent().GetDrawProperties().W) {
		Lines := splitTextAndRenderByLines(receiver.Value, int(receiver.GetParent().GetDrawProperties().W), receiver.GetParent().GetDrawProperties().Font.Size)
		Fonts.DrawText(receiver.GetParent().GetDrawProperties().Font, Lines, receiver.DrawProperties.Texture)
	} else {
		Fonts.DrawText(receiver.GetParent().GetDrawProperties().Font, []string{receiver.Value}, receiver.DrawProperties.Texture)
	}

	if receiver.GetDrawProperties().W > receiver.GetParent().GetDrawProperties().W {
		println("bigger than parent")
		println(receiver.GetDrawProperties().W)
		Lines := splitTextAndRenderByLines(receiver.Value, int(receiver.GetParent().GetDrawProperties().W), receiver.GetParent().GetDrawProperties().Font.Size)
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

func splitTextAndRenderByLines(text string, maxWidth int, size float64) []string {
	println(text)
	var Lines []string
	var err error
	var currentWidth, _, start, end int
	length := len(text)
	start = 0
	end = length
	for start < length {
		currentWidth = int(size * float64(len(text[start:end])) * 1.2)
		for currentWidth > maxWidth {
			end = findLastSpace(text, end)
			currentWidth = int(size * float64(len(text[start:end])) * 1.2)
			if err != nil {
				panic(err)
			}
		}
		Lines = append(Lines, text[start:end])
		start = end + 1
		end = length
	}
	return Lines
}
