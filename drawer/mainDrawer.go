package drawer

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/drawer/calculator"
	"gezgin_web_engine/widgets"
)

func SetWindowSize(height int, width int) {
	ScreenProperties.SetWindowSize(height, width)
}

/*
func RenderDocument(document widgets.WidgetInterface, renderer *sdl.Renderer) {
	widgetList := []widgets.WidgetInterface{document}
	var edgeList []widgets.WidgetInterface
	length := len(widgetList)
	keepGo := true
	for keepGo {
		keepGo = false
		for _, w := range widgetList {
			if w.GetChildrenCount() > 0 {
				for _, child := range w.Children {
					if child.Draw {
						widgetList = append(widgetList, child)
						child.Rendered = false
					}
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
				w.RenderWidget(w, renderer)
				w.Rendered = true
			}
		}
		for _, w := range widgetList {
			if w.Parent != nil {
				widgetList = append(widgetList, w.Parent)
				keepGo = true
			}
		}
		if keepGo {
			widgetList = widgetList[length:]
			length = len(widgetList)
		}
	}
}

func DrawDocument(document widgets.WidgetInterface, renderer *sdl.Renderer) {
	widgetList := []widgets.WidgetInterface{document}
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
				if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).Draw {
					widgetList = append(widgetList, widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]))
					widgetIndexList = append(widgetIndexList, 0)
					widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).DrawWidget(widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]), renderer)
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).Draw {
					widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).DrawWidget(widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]), renderer)
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}
*/
// TEMPORARY WIDTH AND HEIGHT FUNCTIONS FOR MAKING TEXT RESPONSIVE
func SetWidthForWidget(widget widgets.WidgetInterface) {
	width := calculator.CalculateWidthOfWidget(widget)
	widget.GetDrawProperties().Rect.W = int32(width)
}
func SetHeightForWidget(widget widgets.WidgetInterface) {
	height := calculator.CalculateHeightOfWidget(widget)
	widget.GetDrawProperties().Rect.H = int32(height)
}

func SetXYForWidget(widget widgets.WidgetInterface) {
	posX := calculator.CalculateXPosOfWidget(widget)
	posY := calculator.CalculateYPosOfWidget(widget)
	widget.GetDrawProperties().Rect.X = posX
	widget.GetDrawProperties().Rect.Y = posY
}

func SetWidthForBlockElements(document widgets.WidgetInterface) {
	widgetList := []widgets.WidgetInterface{document}
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
				if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).IsDraw() {
					widgetList = append(widgetList, widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]))
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
					if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_UNTAGGED_TEXT && HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_IMG {
						SetWidthForWidget(currentWidget)
					}
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).IsDraw() {
					currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
					if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_UNTAGGED_TEXT && HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_IMG {
						SetWidthForWidget(currentWidget)
					}
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func SetWidthForInlineElements(document widgets.WidgetInterface) {
	widgetList := []widgets.WidgetInterface{document}
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
				if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).IsDraw() {
					widgetList = append(widgetList, widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]))
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
					if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_UNTAGGED_TEXT || HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_IMG {
						SetWidthForWidget(currentWidget)
					}
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).IsDraw() {
					currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
					if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_UNTAGGED_TEXT || HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_IMG {
						SetWidthForWidget(currentWidget)
					}
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func SetHeightForInlineElements(document widgets.WidgetInterface) {
	widgetList := []widgets.WidgetInterface{document}
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
				if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).IsDraw() {
					widgetList = append(widgetList, widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]))
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
					if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_UNTAGGED_TEXT || HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_IMG {
						SetHeightForInlineElements(currentWidget)
					}
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).IsDraw() {
					currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
					if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_UNTAGGED_TEXT || HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_IMG {
						SetHeightForInlineElements(currentWidget)
					}
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func SetHeightForBlockElements(document widgets.WidgetInterface) {
	widgetList := []widgets.WidgetInterface{document}
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
				if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).IsDraw() {
					widgetList = append(widgetList, widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]))
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
					if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_UNTAGGED_TEXT && HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_IMG {
						SetHeightForWidget(currentWidget)
					}
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).IsDraw() {
					currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
					if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_UNTAGGED_TEXT && HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) != HtmlParser.HTML_IMG {
						SetHeightForWidget(currentWidget)
					}
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

/*
	func SetDrawPropertiesForWidgets(document widgets.WidgetInterface, renderer *sdl.Renderer) {
		document.DrawProperties.Rect.W = int32(ScreenProperties.WindowWidth)
		document.DrawProperties.Rect.H = int32(ScreenProperties.WindowHeight)
		setWidthForBlockElements(document)
		RenderDocument(document, renderer)
		setWidthForInlineElements(document)
		setHeightForInlineElements(document)
		setHeightForBlockElements(document)
		SetPositionOfElements(document)
	}
*/
func SetPositionOfElements(document widgets.WidgetInterface) {
	widgetList := []widgets.WidgetInterface{document}
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
				if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).IsDraw() {
					widgetList = append(widgetList, widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]))
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
					SetXYForWidget(currentWidget)
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex]).IsDraw() {
					currentWidget := widgetList[currentIndex].GetChildrenByIndex(widgetIndexList[currentIndex])
					SetXYForWidget(currentWidget)
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}
