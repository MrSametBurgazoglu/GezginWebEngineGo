package widgets

import (
	"gezgin_web_engine/drawer"
	"gezgin_web_engine/drawer/ScreenProperties"
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

func (receiver *DocumentWidget) Draw(renderer *sdl.Renderer) {
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

func (receiver *DocumentWidget) Render(renderer *sdl.Renderer) {
	receiver.DrawProperties.Rect.W = int32(ScreenProperties.WindowWidth)
	receiver.DrawProperties.Rect.H = int32(ScreenProperties.WindowHeight)
	drawer.SetWidthForBlockElements(receiver)
	receiver.RenderDocument(renderer)
	drawer.SetWidthForInlineElements(receiver)
	drawer.SetHeightForInlineElements(receiver)
	drawer.SetHeightForBlockElements(receiver)
	drawer.SetPositionOfElements(receiver)
	/*
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
		}*/
}
