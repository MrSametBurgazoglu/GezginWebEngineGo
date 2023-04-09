package html_scraper

import (
	"fmt"
	"gezgin_web_engine/drawer/DrawProperties"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/drawer/structs"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"gezgin_web_engine/html_scraper/tagScraper"
	"gezgin_web_engine/html_scraper/tags"
	"gezgin_web_engine/html_scraper/widget"
	"os"
	"strings"
	"sync"
)

func FreeHtmlTree() {
	fmt.Println("heyyo")
}

func ScrapeHtmlFromFile(fileUrl string) *widget.Widget {
	dat, err := os.ReadFile(fileUrl)
	if err != nil {
		panic(err)
	}
	documentWidget := widget.Widget{ChildrenCount: 0, ChildrenIndex: 0, HtmlTag: htmlVariables.HTML_DOCUMENT, Draw: true}
	documentWidget.DrawProperties = new(structs.DrawProperties)
	documentWidget.DrawProperties.Rect.X = 0
	documentWidget.DrawProperties.Rect.Y = 0
	documentWidget.DrawProperties.Rect.W = int32(ScreenProperties.WindowWidth)
	currentWidget := &documentWidget

	var wg sync.WaitGroup

	data := string(dat)
	dataLength := len(data)
	seek := 0
	for seek < dataLength {
		if data[seek] == ' ' || data[seek] == '\n' {
			seek += 1
		} else {
			result := strings.Index(data[seek:], "<")
			if result > 0 {
				//make untagged text to strip
				newWidget := widget.Widget{
					HtmlTag:          htmlVariables.HTML_UNTAGGED_TEXT,
					WidgetProperties: tags.UntaggedText{Value: data[seek : seek+result]},
					Parent:           currentWidget,
					ChildrenIndex:    currentWidget.ChildrenCount,
					RenderWidget:     DrawProperties.RenderUntaggedTextFunction,
					DrawWidget:       DrawProperties.DrawUntaggedTextFunction,
					Draw:             true,
					DrawProperties:   new(structs.DrawProperties),
					//draw and render function
				}
				currentWidget.ChildrenCount++
				currentWidget.Children = append(currentWidget.Children, &newWidget)
				tagScraper.UntaggedTextClosed(&newWidget)
				println("untagged text", data[seek:seek+result])
			}
			result2 := strings.Index(data[seek+result:], ">")
			if data[seek+result+1] == '/' {
				println("tag ended", data[seek+result+1:seek+result+result2])
				currentWidget = currentWidget.Parent
			} else {
				newWidget := widget.Widget{
					Parent:        currentWidget,
					ChildrenIndex: currentWidget.ChildrenCount,
					//draw and render function
				}
				currentWidget.ChildrenCount++
				currentWidget.Children = append(currentWidget.Children, &newWidget)
				println("inside of tag", data[seek+result+1:seek+result+result2])
				currentWidget = &newWidget
				if tagScraper.ScrapeInsideOfTag(currentWidget, data[seek+result+1:seek+result+result2], &wg) {
					currentWidget = currentWidget.Parent
				}
			}
			seek += result + result2 + 1
		}
	}
	wg.Wait()
	return &documentWidget
}
