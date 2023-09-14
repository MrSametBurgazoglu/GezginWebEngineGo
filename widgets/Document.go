package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/LayoutEngine"
	"gezgin_web_engine/LayoutProperty"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/widget"
	"image"
	"sync"
)

type DocumentWidget struct {
	widget.Widget
	ResourceManager *ResourceManager.ResourceManager
}

func allChildrenRendered(widget widget.WidgetInterface) bool {
	for _, child := range widget.GetChildren() {
		if child.IsRender() == false {
			return false
		}
	}
	return true
}

func (receiver *DocumentWidget) DrawPage(mainImage *image.RGBA) {
	receiver.Draw(mainImage)
	widgetList := []widget.WidgetInterface{receiver}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != widgetList[0].GetChildrenCount() {
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
	widgetList := []widget.WidgetInterface{receiver}
	var edgeList []widget.WidgetInterface
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
	receiver.LayoutProperty.Width = ScreenProperties.WindowWidth
	//receiver.LayoutProperty.Height = ScreenProperties.WindowHeight
	//TODO FLEX ITEM MUST NEED TO CALCULATED WIDTH OF UNTAGGED TEXT BUT
	receiver.RenderDocument(mainImage)
	receiver.SetWidthForBlockElements()
	receiver.RenderDocument(mainImage)
	SetWidthForInlineElements(receiver)
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
		drawerBackend.DrawBackground(red, green, blue, alpha, mainImage, receiver.LayoutProperty)
	}
}

func SetWidthForWidget(widget widget.WidgetInterface) {
	layout := widget.GetLayout()
	var layoutList []*LayoutProperty.LayoutProperty
	for _, widgetInterface := range widget.GetChildren() {
		layoutList = append(layoutList, widgetInterface.GetLayout())
	}

	LayoutEngine.SetWidth(widget)
	widget.GetLayout().Width = layout.Width
}
func SetHeightForWidget(widget widget.WidgetInterface) {
	var layoutList []*LayoutProperty.LayoutProperty
	for _, widgetInterface := range widget.GetChildren() {
		layoutList = append(layoutList, widgetInterface.GetLayout())
	}
	LayoutEngine.SetHeight(widget)
}

func SetXYForWidget(currentWidget widget.WidgetInterface) {
	layout := currentWidget.GetLayout()
	var layoutList []*LayoutProperty.LayoutProperty
	for _, widgetInterface := range currentWidget.GetChildren() {
		layoutList = append(layoutList, widgetInterface.GetLayout())
	}

	var beforeCurrentWidget widget.WidgetInterface
	if currentWidget.GetChildrenIndex() > 0 {
		beforeCurrentWidget = currentWidget.GetParent().GetChildrenByIndex(currentWidget.GetChildrenIndex() - 1)
	}

	x, y := LayoutEngine.SetPosition(currentWidget, currentWidget.GetParent(), beforeCurrentWidget)

	if untaggedText, ok := currentWidget.(*UntaggedText); ok {
		println("x:", x, "y:", y, "value:", untaggedText.Value, "w:", layout.Width, "h:", layout.Height)
	}

	if HtmlParser.HtmlTags(currentWidget.GetHtmlTag()) == HtmlParser.HTML_PRE {
		println("heyyo")
	}

	currentWidget.GetLayout().XPosition = x
	currentWidget.GetLayout().YPosition = y
}

func (receiver *DocumentWidget) SetWidthForBlockElements() {
	var wg sync.WaitGroup
	for _, child := range receiver.Children {
		wg.Add(1)
		receiver.SetWidthOfWidget(child, &wg)
	}
	wg.Wait()
}

func (receiver *DocumentWidget) SetWidthOfWidget(widget widget.WidgetInterface, group *sync.WaitGroup) { //TODO html tag must be string and can be custom
	SetWidthForWidget(widget)
	for _, child := range widget.GetChildren() {
		if child.IsPreSetWidth() {
			group.Add(1)
			go receiver.SetWidthOfWidget(child, group)
		} else if child.IsSetWidthSelf() { //we need to look at here
			group.Add(1)
			go receiver.SetWidthOfWidget(child, group)
		}
	}
	group.Done()
}

func SetWidthForBlockElements(document widget.WidgetInterface) {
	widgetList := []widget.WidgetInterface{document}
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

func SetWidthForInlineElements(document widget.WidgetInterface) {
	widgetList := []widget.WidgetInterface{document}
	var edgeList []widget.WidgetInterface
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
			if allChildrenRendered(w) && !w.IsPreSetWidth() {
				SetWidthForWidget(w)
				w.SetRender(true)
			}
		}
		for _, w := range widgetList {
			//TODO FLEX CONTAINER'S CHILDREN MUST BE ALSO FLEX
			//THE FIX THAT WE PUT HERE IS CAN BE WRONG BUT IF A BLOCK ELEMENT WIDTH COULDN'T BE 0
			if !w.GetParent().IsPreSetWidth() {
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

/*TODO SIZE AND POSITIONS SHOULDN'T BE SET THEY MUST BE GET AND ITS MUST CALCULATED WHEN GETTING*/
func (receiver *DocumentWidget) SetHeightForElements() {
	widgetList := []widget.WidgetInterface{receiver}
	var edgeList []widget.WidgetInterface
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
	SetHeightForWidget(receiver)
}

func SetPositionOfElements(document widget.WidgetInterface) {
	widgetList := []widget.WidgetInterface{document}
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
