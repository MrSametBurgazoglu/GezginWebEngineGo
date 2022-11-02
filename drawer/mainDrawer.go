package drawer

import (
	"gezgin_web_engine/drawer/Fonts"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/html_scraper/widget"
	"github.com/veandco/go-sdl2/sdl"
)

func SetWindowSize(height int, width int) {
	ScreenProperties.SetWindowSize(height, width)
}

func LoadDefaultFont() {
	Fonts.InitializeFont()
}

func RenderDocument(document *widget.Widget, renderer *sdl.Renderer) {
	widgetList := []*widget.Widget{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != document.ChildrenCount {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].ChildrenCount {
			currentIndex--
			widgetIndexList = widgetIndexList[:currentIndex+1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].ChildrenCount > 0 {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					widgetList = append(widgetList, widgetList[currentIndex].Children[widgetIndexList[currentIndex]])
					widgetIndexList = append(widgetIndexList, 0)
					widgetList[currentIndex].Children[widgetIndexList[currentIndex]].RenderWidget(widgetList[currentIndex].Children[widgetIndexList[currentIndex]], renderer)
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				println(widgetList[currentIndex].Draw, widgetList[currentIndex].HtmlTag)
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					println(widgetList[currentIndex].HtmlTag)
					widgetList[currentIndex].Children[widgetIndexList[currentIndex]].RenderWidget(widgetList[currentIndex].Children[widgetIndexList[currentIndex]], renderer)
				}
				widgetIndexList[currentIndex]++
			}
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
			widgetIndexList = widgetIndexList[:currentIndex+1]
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
				println(widgetList[currentIndex].Draw, widgetList[currentIndex].HtmlTag)
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					println(widgetList[currentIndex].HtmlTag)
					widgetList[currentIndex].Children[widgetIndexList[currentIndex]].DrawWidget(widgetList[currentIndex].Children[widgetIndexList[currentIndex]], renderer)
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}
