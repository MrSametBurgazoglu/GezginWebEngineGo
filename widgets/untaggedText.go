package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/LayoutEngine"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleEngine/enums"
	"gezgin_web_engine/drawer/Fonts"
	"gezgin_web_engine/drawer/structs"
	"image"
	"image/draw"
	"strings"
)

type UntaggedText struct {
	Widget
	Value string
}

func (receiver *UntaggedText) Draw(mainImage *image.RGBA) {
	if receiver.Parent.GetStyleProperty() != nil && receiver.Parent.GetStyleProperty().TextAlign == enums.CSS_TEXT_ALIGN_CENTER {
		middlePoint := receiver.Parent.GetLayout().XPosition + receiver.Parent.GetLayout().Width/2
		draw.Draw(mainImage, image.Rect(middlePoint-receiver.LayoutProperty.Width/2, receiver.LayoutProperty.YPosition, middlePoint-receiver.LayoutProperty.Width/2+receiver.LayoutProperty.Width, receiver.LayoutProperty.YPosition+receiver.LayoutProperty.Height), receiver.DrawProperties.Texture, image.Point{X: 0, Y: 0}, draw.Over)
	} else {
		draw.Draw(mainImage, image.Rect(receiver.LayoutProperty.XPosition, receiver.LayoutProperty.YPosition, receiver.LayoutProperty.XPosition+receiver.LayoutProperty.Width, receiver.LayoutProperty.YPosition+receiver.LayoutProperty.Height), receiver.DrawProperties.Texture, image.Point{X: 0, Y: 0}, draw.Over)
	}
}

func (receiver *UntaggedText) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {
	if receiver.GetParent().GetDrawProperties().Font == nil {
		if receiver.GetParent().GetStyleProperty().Font != nil {
			receiver.GetParent().GetDrawProperties().Font = Fonts.GetFont(receiver.GetParent().GetStyleProperty().Font.FontSizeValue)
		} else {
			receiver.GetParent().GetDrawProperties().Font = Fonts.GetFont(14)
		}
	}
	if currentWidth := int(receiver.GetParent().GetDrawProperties().Font.Size * float64(len(receiver.Value)) * 0.5); currentWidth > receiver.GetParent().GetLayout().Width {
		Lines, maxTextWidth := splitTextAndRenderByLines(receiver.Value, receiver.GetParent().GetLayout().Width, receiver.GetParent().GetDrawProperties().Font.Size)
		receiver.DrawProperties.Texture = image.NewRGBA(image.Rect(0, 0, maxTextWidth*3, 500)) // change this later
		height, width := Fonts.DrawText(receiver.GetParent().GetDrawProperties().Font, Lines, receiver.DrawProperties.Texture, receiver.GetParent().GetStyleProperty().Color)
		receiver.LayoutProperty.Height, receiver.LayoutProperty.Width = int(height), int(width)
	} else {
		//change this to calculated text
		receiver.DrawProperties.Texture = image.NewRGBA(image.Rect(0, 0, receiver.GetParent().GetLayout().Width, 500)) // change this later
		height, width := Fonts.DrawText(receiver.GetParent().GetDrawProperties().Font, []string{receiver.Value}, receiver.DrawProperties.Texture, receiver.GetParent().GetStyleProperty().Color)
		receiver.LayoutProperty.Height, receiver.LayoutProperty.Width = int(height), int(width)
	}
	//receiver.DrawProperties.W = int32(receiver.GetParent().GetDrawProperties().Font.Size * float64(len(receiver.Value)) * 0.5) // change this later and calculate text width

	if receiver.LayoutProperty.Width > receiver.GetParent().GetLayout().Width {
		Lines, _ := splitTextAndRenderByLines(receiver.Value, int(receiver.GetParent().GetLayout().Width), receiver.GetParent().GetDrawProperties().Font.Size)
		println(Lines, " Lines")
	}
}

func SetWidgetPropertiesForUntaggedText(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(UntaggedText)
	widget.HtmlElement = element
	widget.DrawProperties = new(structs.DrawProperties)
	widget.LayoutProperty = new(LayoutEngine.LayoutProperty)
	widget.DrawProperties.Initialize()
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

func splitTextAndRenderByLines(text string, maxWidth int, size float64) ([]string, int) {
	var Lines []string
	var err error
	var currentWidth, _, start, end int
	calculatedMaxWidth := 0
	length := len(text)
	start = 0
	end = length
	currentWidth = int(size * float64(length) * 0.55)
	if maxWidth <= 0 || !strings.Contains(text, " ") {
		return append(Lines, text), currentWidth
	}
	for start < length {
		currentWidth = int(size * float64(len(text[start:end])) * 0.55)
		for currentWidth > maxWidth {
			end = findLastSpace(text, end)
			newCurrentWidth := int(size * float64(len(text[start:end])) * 0.55)
			if newCurrentWidth > calculatedMaxWidth {
				calculatedMaxWidth = newCurrentWidth
			}
			if newCurrentWidth == currentWidth {
				break
			}
			currentWidth = newCurrentWidth
			if err != nil {
				panic(err)
			}
		}
		Lines = append(Lines, text[start:end])
		start = end + 1
		end = length
	}
	return Lines, calculatedMaxWidth
}

func (receiver *UntaggedText) IsPreSetWidth() bool {
	return false
}

func (receiver *UntaggedText) IsSetWidthSelf() bool {
	return false
}

func (receiver *UntaggedText) SetParent(parent WidgetInterface) {
	receiver.Parent = parent
	receiver.LayoutProperty.Parent = parent.GetLayout()
	receiver.Parent.GetLayout().Children = append(receiver.Parent.GetLayout().Children, receiver.LayoutProperty)
}
