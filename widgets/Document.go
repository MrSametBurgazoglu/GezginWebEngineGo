package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/LayoutEngine"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/drawer/drawerBackend"
	"image"
	"sync"
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
	println("rendering finished")
}

// This function and sub functions will be rewritten
func (receiver *DocumentWidget) RenderPage(mainImage *image.RGBA) {
	receiver.LayoutProperty.Width = ScreenProperties.WindowWidth
	receiver.LayoutProperty.Height = ScreenProperties.WindowHeight
	receiver.SetWidthForBlockElements()
	println("width for block elements finished")
	receiver.RenderDocument(mainImage)
	println("widget rendered")
	SetWidthForInlineElements(receiver)
	println("width for inline elements finished")
	receiver.SetHeightForElements()
	//SetHeightForInlineElements(receiver)
	//SetHeightForBlockElements(receiver)
	SetPositionOfElements(receiver)
}

func (receiver *DocumentWidget) LayoutBlockWidgets() {

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
	layout := widget.GetLayout()
	var layoutList []*LayoutEngine.LayoutProperty
	for _, widgetInterface := range widget.GetChildren() {
		layoutList = append(layoutList, widgetInterface.GetLayout())
	}

	layout.SetWidth(widget.GetParent().GetLayout(), layoutList, widget.GetStyleProperty())
	widget.GetLayout().Width = layout.Width
}
func SetHeightForWidget(widget WidgetInterface) {
	layout := widget.GetLayout()
	var layoutList []*LayoutEngine.LayoutProperty
	for _, widgetInterface := range widget.GetChildren() {
		layoutList = append(layoutList, widgetInterface.GetLayout())
	}

	layout.SetHeight(widget.GetParent().GetLayout(), layoutList, widget.GetStyleProperty())
	widget.GetLayout().Height = layout.Height
}

func SetXYForWidget(widget WidgetInterface) {
	layout := widget.GetLayout()
	var layoutList []*LayoutEngine.LayoutProperty
	for _, widgetInterface := range widget.GetChildren() {
		layoutList = append(layoutList, widgetInterface.GetLayout())
	}

	x := layout.SetPositionX(widget.GetParent().GetLayout(), widget.GetStyleProperty())
	layout.XPosition = x
	if widget.GetChildrenIndex() == 0 {
		y := layout.SetPositionY(widget.GetParent().GetLayout(), nil, widget.GetStyleProperty())
		layout.YPosition = y
	} else {
		y := layout.SetPositionY(widget.GetParent().GetLayout(), widget.GetParent().GetChildrenByIndex(widget.GetChildrenIndex()-1).GetLayout(), widget.GetStyleProperty())
		layout.YPosition = y
	}
	widget.GetLayout().XPosition = layout.XPosition
	widget.GetLayout().YPosition = layout.YPosition
}

func (receiver *DocumentWidget) SetWidthForBlockElements() {
	var wg sync.WaitGroup
	for _, child := range receiver.Children {
		wg.Add(1)
		println("added")
		receiver.SetWidthOfWidget(child, &wg)
	}
	wg.Wait()
}

func (receiver *DocumentWidget) SetWidthOfWidget(widget WidgetInterface, group *sync.WaitGroup) { //TODO html tag must be string and can be custom
	SetWidthForWidget(widget)
	println("widget width setted", widget.GetHtmlTag())
	for _, child := range widget.GetChildren() {
		if child.IsBlockElement() {
			group.Add(1)
			println("added")
			go receiver.SetWidthOfWidget(child, group)
		}
	}
	group.Done()
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
			if allChildrenRendered(w) && !w.IsBlockElement() {
				SetHeightForWidget(w)
				w.SetRender(true)
			}
		}
		for _, w := range widgetList {
			if !w.GetParent().IsBlockElement() {
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
