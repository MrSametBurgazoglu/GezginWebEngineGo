package drawer

import (
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"github.com/veandco/go-sdl2/sdl"
)

func SetWindowSize(height int, width int) {
	ScreenProperties.SetWindowSize(height, width)
}

func RenderDocument(document *htmlVariables.Widget, renderer *sdl.Renderer) {
	widgetList := []*htmlVariables.Widget{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	widgetCount := 0
	for widgetIndexList[0] != document.ChildrenCount {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].ChildrenCount {
			currentIndex--
			widgetCount--
			widgetIndexList = widgetIndexList[:widgetCount]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].ChildrenCount > 0 {
				widgetCount++
				widgetList = append(widgetList, widgetList[currentIndex].Children[widgetIndexList[currentIndex]])
				widgetIndexList[widgetCount-1] = 0
				currentIndex++
				if widgetList[currentIndex].Draw {
					widgetList[currentIndex].RenderWidget(widgetList[currentIndex], renderer)
				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					widgetList[currentIndex].RenderWidget(widgetList[currentIndex], renderer)
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}
