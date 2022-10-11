package html_scraper

import (
	"fmt"
	"gezgin_web_engine/html_scraper/tags"
	"os"
	"strings"
)

func FreeHtmlTree() {
	fmt.Println("heyyo")
}

func ScrapeHtmlFromFile(fileUrl string) *Widget {
	dat, err := os.ReadFile(fileUrl)
	if err != nil {
		panic(err)
	}
	documentWidget := Widget{ChildrenCount: 0, ChildrenIndex: 0, HtmlTag: HTML_DOCUMENT}
	currentWidget := &documentWidget
	data := string(dat)
	dataLength := len(data)
	seek := 0
	println(dataLength)
	for seek < dataLength {
		println("seek:", seek)
		println(string(data[seek]))
		if data[seek] == ' ' || data[seek] == '\n' {
			seek += 1
		} else {
			result := strings.Index(data[seek:], "<")
			println("result:", result)
			if result > 0 {
				//make untagged text to strip
				newWidget := Widget{
					HtmlTag:          HTML_UNTAGGED_TEXT,
					WidgetProperties: tags.UntaggedText{Value: data[seek : seek+result]},
					Parent:           currentWidget,
					ChildrenIndex:    currentWidget.ChildrenCount,
					//draw and render function
				}
				currentWidget.ChildrenCount++
				currentWidget.Children = append(currentWidget.Children, newWidget)
				println("untagged text", data[seek:seek+result])
				seek += result
			}
			result2 := strings.Index(data[seek+result:], ">")
			if data[seek+1] == '/' {
				println("tag ended", data[seek+1:seek+result+result2])
				currentWidget = currentWidget.Parent
			} else {
				newWidget := Widget{
					Parent:        currentWidget,
					ChildrenIndex: currentWidget.ChildrenCount,
					//draw and render function
				}
				currentWidget.ChildrenCount++
				currentWidget.Children = append(currentWidget.Children, newWidget)
				println("inside of tag", data[seek+result+1:seek+result2])
				currentWidget = &newWidget
				ScrapeInsideOfTag(currentWidget, data[seek+result+1:seek+result2])
			}
			seek += result + result2 + 1
		}
	}
	return &documentWidget
}
