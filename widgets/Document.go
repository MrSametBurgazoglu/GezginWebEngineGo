package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/drawer/drawerBackend"
	"image"
)

type DocumentWidget struct {
	Widget
	ResourceManager *ResourceManager.ResourceManager
}

func allChildrenRendered(widget WidgetInterface) bool {
	for _, child := range widget.GetChildren() {
		if child.IsRender() == false {
			return false
		}
	}
	return true
}

func (receiver *DocumentWidget) DrawPage(mainImage *image.RGBA) {
	receiver.Draw(mainImage)
	widgetList := []WidgetInterface{receiver}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != receiver.ChildrenCount {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].GetChildrenCount() {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).GetChildrenCount() > 0 {
				widgetList = append(widgetList, widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]))
				widgetIndexList = append(widgetIndexList, 0)
				widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).Draw(mainImage)
				currentIndex++
			} else {
				widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).Draw(mainImage)
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func (receiver *DocumentWidget) RenderDocument(mainImage *image.RGBA) {
	widgetList := []WidgetInterface{receiver}
	var edgeList []WidgetInterface
	length := len(widgetList)
	keepGo := true
	for keepGo {
		keepGo = false
		for _, w := range widgetList {
			if w.GetChildrenCount() > 0 {
				for _, child := range w.GetChildren() {
					widgetList = append(widgetList, child)
					child.SetRender(false)
					keepGo = true
				}
			} else {
				edgeList = append(edgeList, w)
			}
		}
		if keepGo {
			widgetList = widgetList[length:]
			length = len(widgetList)
		}
	}
	widgetList = edgeList
	keepGo = true
	for keepGo {
		keepGo = false
		for _, w := range widgetList {
			if allChildrenRendered(w) {
				w.Render(mainImage, receiver.ResourceManager)
				w.SetRender(true)
			}
		}
		for _, w := range widgetList {
			if w.GetParent() != nil {
				widgetList = append(widgetList, w.GetParent())
				keepGo = true
			}
		}
		if keepGo {
			widgetList = widgetList[length:]
			length = len(widgetList)
		}
	}
}

// This function and sub functions will be rewritten
func (receiver *DocumentWidget) RenderPage(mainImage *image.RGBA) {
	receiver.DrawProperties.W = int32(ScreenProperties.WindowWidth)
	receiver.DrawProperties.H = int32(ScreenProperties.WindowHeight)
	SetWidthForBlockElements(receiver)
	receiver.RenderDocument(mainImage)
	SetWidthForInlineElements(receiver)
	receiver.SetHeightForElements()
	//SetHeightForInlineElements(receiver)
	//SetHeightForBlockElements(receiver)
	SetPositionOfElements(receiver)
}

func (receiver *DocumentWidget) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {
	//render body
}

func (receiver *DocumentWidget) Draw(mainImage *image.RGBA) {
	//draw.Draw(mainImage, mainImage.Bounds(), image.White, image.Point{X: 0, Y: 0}, draw.Over)
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, mainImage, receiver.DrawProperties)
	}
}

func SetWidthForWidget(widget WidgetInterface) {
	width := CalculateWidthOfWidget(widget)
	widget.GetDrawProperties().W = int32(width)
}
func SetHeightForWidget(widget WidgetInterface) {
	height := CalculateHeightOfWidget(widget)
	widget.GetDrawProperties().H = int32(height)
}

func SetXYForWidget(widget WidgetInterface) {
	posX := CalculateXPosOfWidget(widget)
	posY := CalculateYPosOfWidget(widget)
	widget.GetDrawProperties().X = posX
	widget.GetDrawProperties().Y = posY
}

func SetWidthForBlockElements(document WidgetInterface) {
	widgetList := []WidgetInterface{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != document.GetChildrenCount() {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].GetChildrenCount() {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).GetChildrenCount() > 0 {
				widgetList = append(widgetList, widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]))
				widgetIndexList = append(widgetIndexList, 0)
				currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
				if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_UNTAGGED_TEXT && HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_IMG {
					SetWidthForWidget(currentWidget)
				}
				currentIndex++

			} else {
				currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
				if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_UNTAGGED_TEXT && HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_IMG {
					SetWidthForWidget(currentWidget)
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func SetWidthForInlineElements(document WidgetInterface) {
	widgetList := []WidgetInterface{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != document.GetChildrenCount() {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].GetChildrenCount() {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).GetChildrenCount() > 0 {

				widgetList = append(widgetList, widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]))
				widgetIndexList = append(widgetIndexList, 0)
				currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
				if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_UNTAGGED_TEXT || HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_IMG {
					SetWidthForWidget(currentWidget)
				}
				currentIndex++

			} else {

				currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
				if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_UNTAGGED_TEXT || HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_IMG {
					SetWidthForWidget(currentWidget)
				}

				widgetIndexList[currentIndex]++
			}
		}
	}
}

func (receiver *DocumentWidget) SetHeightForElements() {
	widgetList := []WidgetInterface{receiver}
	var edgeList []WidgetInterface
	length := len(widgetList)
	keepGo := true
	for keepGo {
		keepGo = false
		for _, w := range widgetList {
			if w.GetChildrenCount() > 0 {
				for _, child := range w.GetChildren() {
					widgetList = append(widgetList, child)
					child.SetRender(false)
					keepGo = true
				}
			} else {
				edgeList = append(edgeList, w)
			}
		}
		if keepGo {
			widgetList = widgetList[length:]
			length = len(widgetList)
		}
	}
	widgetList = edgeList
	keepGo = true
	for keepGo {
		keepGo = false
		for _, w := range widgetList {
			if allChildrenRendered(w) {
				SetHeightForWidget(w)
				w.SetRender(true)
			}
		}
		for _, w := range widgetList {
			if w.GetParent() != receiver {
				widgetList = append(widgetList, w.GetParent())
				keepGo = true
			}
		}
		if keepGo {
			widgetList = widgetList[length:]
			length = len(widgetList)
		}
	}
}

func SetHeightForInlineElements(document WidgetInterface) {
	widgetList := []WidgetInterface{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != document.GetChildrenCount() {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].GetChildrenCount() {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).GetChildrenCount() > 0 {
				widgetList = append(widgetList, widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]))
				widgetIndexList = append(widgetIndexList, 0)
				currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
				if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_UNTAGGED_TEXT || HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_IMG {
					SetHeightForWidget(currentWidget)
				}
				currentIndex++

			} else {
				currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
				if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_UNTAGGED_TEXT || HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_IMG {
					SetHeightForWidget(currentWidget)
				}

				widgetIndexList[currentIndex]++
			}
		}
	}
}

func SetHeightForBlockElements(document WidgetInterface) {
	widgetList := []WidgetInterface{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != document.GetChildrenCount() {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].GetChildrenCount() {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).GetChildrenCount() > 0 {
				widgetList = append(widgetList, widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]))
				widgetIndexList = append(widgetIndexList, 0)
				currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
				if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_UNTAGGED_TEXT && HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_IMG {
					SetHeightForWidget(currentWidget)
				}
				currentIndex++
			} else {
				currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
				if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_UNTAGGED_TEXT && HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_IMG {
					SetHeightForWidget(currentWidget)
				}

				widgetIndexList[currentIndex]++
			}
		}
	}
}

func SetPositionOfElements(document WidgetInterface) {
	widgetList := []WidgetInterface{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != document.GetChildrenCount() {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].GetChildrenCount() {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).GetChildrenCount() > 0 {
				widgetList = append(widgetList, widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]))
				widgetIndexList = append(widgetIndexList, 0)
				currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
				SetXYForWidget(currentWidget)
				currentIndex++
			} else {
				currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
				SetXYForWidget(currentWidget)
				widgetIndexList[currentIndex]++
			}
		}
	}
}
