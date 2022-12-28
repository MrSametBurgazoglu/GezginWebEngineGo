package drawer

import (
	"gezgin_web_engine/drawer/Fonts"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/drawer/calculator"
	"gezgin_web_engine/drawer/structs"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/slices"
)

func allChildsRendered(widget structs.DrawableWidget) bool {
	for _, child := range widget.GetChildren() {
		if child.GetRendered() == false {
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

func RenderDocument(document structs.DrawableWidget, renderer *sdl.Renderer) {
	widgetList := []structs.DrawableWidget{document}
	var edgeList []structs.DrawableWidget
	length := len(widgetList)
	keepGo := true
	for keepGo {
		keepGo = false
		for _, w := range widgetList {
			if w.GetChildrenCount() > 0 {
				for _, child := range w.GetChildren() {
					if child.IsDrawable() {
						widgetList = append(widgetList, child.(structs.DrawableWidget))
						child.SetRendered(false)
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
	length = len(widgetList)
	keepGo = true
	for keepGo {
		keepGo = false
		for _, w := range widgetList {
			if allChildsRendered(w) {
				w.RenderWidget(renderer)
				w.SetRendered(true)
			}
		}
		for _, w := range widgetList {
			if w.GetParent() != nil {
				widgetList = append(widgetList, w.GetParent().(structs.DrawableWidget))
				keepGo = true
			}
		}
		if keepGo {
			widgetList = widgetList[length:]
			length = len(widgetList)
		}
	}
}

func DrawDocument(document structs.DrawableWidget, renderer *sdl.Renderer) {
	widgetList := []structs.DrawableWidget{document}
	widgetIndexList := []int{0}
	//initialize document
	currentIndex := 0
	for widgetIndexList[0] != document.GetChildrenCount() {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].GetChildrenCount() {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			child, ok := widgetList[currentIndex].GetChild(widgetIndexList[currentIndex]).(structs.DrawableWidget)
			if ok {
				if child.GetChildrenCount() > 0 {
					widgetList = append(widgetList, child)
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := child
					currentWidget.DrawWidget(renderer)
					currentIndex++
				} else {
					currentWidget := child
					currentWidget.DrawWidget(renderer)
				}
			} else {
				widgetIndexList[currentIndex]++
			}
		}
	}
	/*widgetList := []*widget.Widget{document}
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
					currentWidget.WidgetElement.DrawWidget(currentWidget, renderer)
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					currentWidget.WidgetElement.DrawWidget(currentWidget, renderer)
				}
				widgetIndexList[currentIndex]++
			}
		}
	}

	*/
}

func setWHForWidget(widget structs.DrawableWidget) {
	width := calculator.CalculateWidthOfWidget(widget)
	height := calculator.CalculateHeightOfWidget(widget)
	widget.GetRect().W = int32(width)
	widget.GetRect().H = int32(height)
}

func setXYForWidget(widget structs.DrawableWidget) {
	posX := calculator.CalculateXPosOfWidget(widget)
	posY := calculator.CalculateYPosOfWidget(widget)
	widget.GetRect().X = posX
	widget.GetRect().Y = posY
}

func SetDrawPropertiesDocument(document structs.DrawableWidget, renderer *sdl.Renderer) {
	widgetList := []structs.DrawableWidget{document}
	var edgeList []structs.DrawableWidget
	length := len(widgetList)
	keepGo := true
	for keepGo {
		keepGo = false
		for _, w := range widgetList {
			if w.GetChildrenCount() > 0 {
				for _, child := range w.GetChildren() {
					if child.IsDrawable() {
						widgetList = append(widgetList, child.(structs.DrawableWidget))
						child.SetRendered(false)
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
	length = len(widgetList)
	keepGo = true
	for keepGo {
		keepGo = false
		for _, w := range widgetList {
			if allChildsRendered(w) {
				setWHForWidget(w)
				w.SetRendered(true)
			}
		}
		for _, w := range widgetList {
			if w.GetParent() != nil && slices.Contains(widgetList, w) {
				widgetList = append(widgetList, w.GetParent().(structs.DrawableWidget))
				keepGo = true
			}
		}
		if keepGo {
			widgetList = widgetList[length:]
			length = len(widgetList)
		}
	}
	widgetList = []structs.DrawableWidget{document}
	widgetIndexList := []int{0}
	//initialize document
	currentIndex := 0
	for widgetIndexList[0] != document.GetChildrenCount() {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].GetChildrenCount() {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			child, ok := widgetList[currentIndex].GetChild(widgetIndexList[currentIndex]).(structs.DrawableWidget)
			if ok {
				if child.GetChildrenCount() > 0 {
					widgetList = append(widgetList, child)
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := child
					setXYForWidget(currentWidget)
					currentIndex++
				} else {
					currentWidget := child
					setXYForWidget(currentWidget)
				}
			} else {
				widgetIndexList[currentIndex]++
			}
		}
	}
	/*
		widgetList = []structs.DrawableWidget{document}
		widgetIndexList := []int{0}
		currentIndex := 0
		for widgetIndexList[0] != document.GetChildrenCount() {
			if widgetIndexList[currentIndex] == widgetList[currentIndex].GetChildrenCount() {
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

	*/
}
