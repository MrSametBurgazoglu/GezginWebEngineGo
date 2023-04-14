package drawer

import (
	"gezgin_web_engine/drawer/Fonts"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/drawer/calculator"
	"gezgin_web_engine/html_scraper/widget"
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

func LoadDefaultFont() {
	Fonts.InitializeFont()
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

func setWHForWidget(widget *widget.Widget, channel chan *widget.Widget) {
	width := calculator.CalculateWidthOfWidget(widget)
	height := calculator.CalculateHeightOfWidget(widget)
	widget.DrawProperties.Rect.W = int32(width)
	widget.DrawProperties.Rect.H = int32(height)
	widget.Rendered = true
	channel <- widget
}

func setXYForWidget(widget *widget.Widget) {
	posX := calculator.CalculateXPosOfWidget(widget)
	posY := calculator.CalculateYPosOfWidget(widget)
	widget.DrawProperties.Rect.X = posX
	widget.DrawProperties.Rect.Y = posY
}

func SetDrawPropertiesDocument(document *widget.Widget) {
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
	channel := make(chan *widget.Widget)
	for keepGo {
		keepGo = false
		widgetCount := 0
		for _, w := range widgetList {
			if allChildrenRendered(w) {
				widgetCount += 1
				go setWHForWidget(w, channel)
			}
		}
		for widgetCount > 0 {
			currentWidget := <-channel
			if currentWidget.Parent != nil {
				widgetList = append(widgetList, currentWidget.Parent)
				keepGo = true
			}
			widgetCount--
		}
		if keepGo {
			widgetList = widgetList[length:]
			length = len(widgetList)
		}
	}
	widgetList = []*widget.Widget{document}
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
