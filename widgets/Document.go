package widgets

import "github.com/veandco/go-sdl2/sdl"

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

func (receiver *DocumentWidget) Draw() {
	//TODO implement me
	panic("implement me")
}

func (receiver *DocumentWidget) Render() {
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

func RenderDocument(document *tags.Widget, renderer *sdl.Renderer) {

}

func DrawDocument(document *tags.Widget, renderer *sdl.Renderer) {
	widgetList := []*tags.Widget{document}
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
