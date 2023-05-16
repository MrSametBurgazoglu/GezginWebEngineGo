package drawer

import (
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/drawer/calculator"
	"gezgin_web_engine/htmlParser/htmlVariables"
	"gezgin_web_engine/htmlParser/widget"
	"github.com/veandco/go-sdl2/sdl"
)

func allChildrenRendered(widget *widget.Widget) bool {
	for _, child := range widget.Children {
		if child.Rendered == false {
			return false
		}
	}
	return true
}

func SetWindowSize(height int, width int) {
	ScreenProperties.SetWindowSize(height, width)
}

func RenderDocument(document *widget.Widget, renderer *sdl.Renderer) {
	widgetList := []*widget.Widget{document}
	var edgeList []*widget.Widget
	length := len(widgetList)
	keepGo := true
	for keepGo {
		keepGo = false
		for _, w := range widgetList {
			if w.ChildrenCount > 0 {
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

func DrawDocument(document *widget.Widget, renderer *sdl.Renderer) {
	widgetList := []*widget.Widget{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != document.ChildrenCount {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].ChildrenCount {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].ChildrenCount > 0 {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					widgetList = append(widgetList, widgetList[currentIndex].Children[widgetIndexList[currentIndex]])
					widgetIndexList = append(widgetIndexList, 0)
					widgetList[currentIndex].Children[widgetIndexList[currentIndex]].DrawWidget(widgetList[currentIndex].Children[widgetIndexList[currentIndex]], renderer)
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					widgetList[currentIndex].Children[widgetIndexList[currentIndex]].DrawWidget(widgetList[currentIndex].Children[widgetIndexList[currentIndex]], renderer)
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

// TEMPORARY WIDTH AND HEIGHT FUNCTIONS FOR MAKING TEXT RESPONSIVE
func setWidthForWidget(widget *widget.Widget) {
	width := calculator.CalculateWidthOfWidget(widget)
	widget.DrawProperties.Rect.W = int32(width)
}
func setHeightForWidget(widget *widget.Widget) {
	height := calculator.CalculateHeightOfWidget(widget)
	widget.DrawProperties.Rect.H = int32(height)
}

func setXYForWidget(widget *widget.Widget) {
	posX := calculator.CalculateXPosOfWidget(widget)
	posY := calculator.CalculateYPosOfWidget(widget)
	widget.DrawProperties.Rect.X = posX
	widget.DrawProperties.Rect.Y = posY
}

func setWidthForBlockElements(document *widget.Widget) {
	widgetList := []*widget.Widget{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != document.ChildrenCount {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].ChildrenCount {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].ChildrenCount > 0 {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					widgetList = append(widgetList, widgetList[currentIndex].Children[widgetIndexList[currentIndex]])
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					if currentWidget.HtmlTag != htmlVariables.HTML_UNTAGGED_TEXT && currentWidget.HtmlTag != htmlVariables.HTML_IMG {
						setWidthForWidget(currentWidget)
					}
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					if currentWidget.HtmlTag != htmlVariables.HTML_UNTAGGED_TEXT && currentWidget.HtmlTag != htmlVariables.HTML_IMG {
						setWidthForWidget(currentWidget)
					}
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func setWidthForInlineElements(document *widget.Widget) {
	widgetList := []*widget.Widget{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != document.ChildrenCount {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].ChildrenCount {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].ChildrenCount > 0 {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					widgetList = append(widgetList, widgetList[currentIndex].Children[widgetIndexList[currentIndex]])
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					if currentWidget.HtmlTag == htmlVariables.HTML_UNTAGGED_TEXT || currentWidget.HtmlTag == htmlVariables.HTML_IMG {
						setWidthForWidget(currentWidget)
					}
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					if currentWidget.HtmlTag == htmlVariables.HTML_UNTAGGED_TEXT || currentWidget.HtmlTag == htmlVariables.HTML_IMG {
						setWidthForWidget(currentWidget)
					}
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func setHeightForInlineElements(document *widget.Widget) {
	widgetList := []*widget.Widget{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != document.ChildrenCount {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].ChildrenCount {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].ChildrenCount > 0 {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					widgetList = append(widgetList, widgetList[currentIndex].Children[widgetIndexList[currentIndex]])
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					if currentWidget.HtmlTag == htmlVariables.HTML_UNTAGGED_TEXT || currentWidget.HtmlTag == htmlVariables.HTML_IMG {
						setHeightForInlineElements(currentWidget)
					}
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					if currentWidget.HtmlTag == htmlVariables.HTML_UNTAGGED_TEXT || currentWidget.HtmlTag == htmlVariables.HTML_IMG {
						setHeightForInlineElements(currentWidget)
					}
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func setHeightForBlockElements(document *widget.Widget) {
	widgetList := []*widget.Widget{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != document.ChildrenCount {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].ChildrenCount {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].ChildrenCount > 0 {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					widgetList = append(widgetList, widgetList[currentIndex].Children[widgetIndexList[currentIndex]])
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					if currentWidget.HtmlTag != htmlVariables.HTML_UNTAGGED_TEXT && currentWidget.HtmlTag != htmlVariables.HTML_IMG {
						setHeightForWidget(currentWidget)
					}
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					if currentWidget.HtmlTag != htmlVariables.HTML_UNTAGGED_TEXT && currentWidget.HtmlTag != htmlVariables.HTML_IMG {
						setHeightForWidget(currentWidget)
					}
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func SetDrawPropertiesForWidgets(document *widget.Widget, renderer *sdl.Renderer) {
	document.DrawProperties.Rect.W = int32(ScreenProperties.WindowWidth)
	document.DrawProperties.Rect.H = int32(ScreenProperties.WindowHeight)
	setWidthForBlockElements(document)
	RenderDocument(document, renderer)
	setWidthForInlineElements(document)
	setHeightForInlineElements(document)
	setHeightForBlockElements(document)
	SetPositionOfElements(document)
}

func SetPositionOfElements(document *widget.Widget) {
	widgetList := []*widget.Widget{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != document.ChildrenCount {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].ChildrenCount {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].ChildrenCount > 0 {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					widgetList = append(widgetList, widgetList[currentIndex].Children[widgetIndexList[currentIndex]])
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					setXYForWidget(currentWidget)
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					setXYForWidget(currentWidget)
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}
