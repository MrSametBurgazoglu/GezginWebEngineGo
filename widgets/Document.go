package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/drawer/drawerBackend"
	"github.com/veandco/go-sdl2/sdl"
)

type DocumentWidget struct {
	Widget
}

func allChildrenRendered(widget WidgetInterface) bool {
	for _, child := range widget.GetChildren() {
		if child.IsRender() == false {
			return false
		}
	}
	return true
}

func (receiver *DocumentWidget) DrawPage(renderer *sdl.Renderer) {
	receiver.Draw(renderer)
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
				widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).Draw(renderer)
				currentIndex++
			} else {
				widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).Draw(renderer)
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func (receiver *DocumentWidget) RenderDocument(renderer *sdl.Renderer) {
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
				w.Render(renderer)
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

func (receiver *DocumentWidget) RenderPage(renderer *sdl.Renderer) {
	receiver.DrawProperties.Rect.W = int32(ScreenProperties.WindowWidth)
	receiver.DrawProperties.Rect.H = int32(ScreenProperties.WindowHeight)
	SetWidthForBlockElements(receiver)
	receiver.RenderDocument(renderer)
	SetWidthForInlineElements(receiver)
	SetHeightForInlineElements(receiver)
	SetHeightForBlockElements(receiver)
	SetPositionOfElements(receiver)
}

func (receiver *DocumentWidget) Render(renderer *sdl.Renderer) {
	//render body
}

func (receiver *DocumentWidget) Draw(renderer *sdl.Renderer) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, &receiver.DrawProperties.Rect, renderer)
	}
}

func SetWidthForWidget(widget WidgetInterface) {
	width := CalculateWidthOfWidget(widget)
	widget.GetDrawProperties().Rect.W = int32(width)
}
func SetHeightForWidget(widget WidgetInterface) {
	height := CalculateHeightOfWidget(widget)
	widget.GetDrawProperties().Rect.H = int32(height)
}

func SetXYForWidget(widget WidgetInterface) {
	posX := CalculateXPosOfWidget(widget)
	posY := CalculateYPosOfWidget(widget)
	widget.GetDrawProperties().Rect.X = posX
	widget.GetDrawProperties().Rect.Y = posY
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
				if currentWidget.GetHtmlTag() != 106 {
					SetXYForWidget(currentWidget)
				}
				currentIndex++
			} else {
				currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
				if currentWidget.GetHtmlTag() != 106 {
					SetXYForWidget(currentWidget)
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}
