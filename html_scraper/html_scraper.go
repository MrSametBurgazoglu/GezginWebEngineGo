package html_scraper

import (
	"fmt"
	"gezgin_web_engine/drawer/structs"
	"gezgin_web_engine/html_scraper/HtmlElementWidget"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"gezgin_web_engine/html_scraper/tagScraper"
	"gezgin_web_engine/html_scraper/tags"
	"gezgin_web_engine/html_scraper/widget"
	"os"
	"strings"
)

func FreeHtmlTree() {
	fmt.Println("heyyo")
}

func ScrapeHtmlFromFile(fileUrl string) HtmlElementWidget.HtmlElementWidgetInterface {
	dat, err := os.ReadFile(fileUrl)
	if err != nil {
		panic(err)
	}
	documentWidget := tags.HtmlDocument{BaseWidget: &widget.BaseWidget{Draw: true, HtmlTag: htmlVariables.HTML_DOCUMENT}, DrawProperties: new(structs.DrawProperties)}
	var currentWidget HtmlElementWidget.HtmlElementWidgetInterface = &documentWidget
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
				newWidget := tags.HtmlUntaggedText{
					DrawProperties: new(structs.DrawProperties),
					BaseWidget: &widget.BaseWidget{
						Draw:          true,
						HtmlTag:       htmlVariables.HTML_UNTAGGED_TEXT,
						Parent:        currentWidget,
						ChildrenIndex: currentWidget.GetChildrenCount(),
					},
					Value: data[seek : seek+result],
					//draw and render function
				}
				currentWidget.SetChildrenCount(currentWidget.GetChildrenCount() + 1)
				currentWidget.AppendChild(&newWidget)
				println("untagged text", data[seek:seek+result])
			}
			result2 := strings.Index(data[seek+result:], ">")
			if data[seek+result+1] == '/' {
				println("tag ended", data[seek+result+1:seek+result+result2])
				currentWidget = currentWidget.GetParent()
			} else {
				var newWidget HtmlElementWidget.HtmlElementWidgetInterface
				/*newWidget := widget.Widget{
					Parent:        currentWidget,
					ChildrenIndex: currentWidget.ChildrenCount,
					//draw and render function
				}

				currentWidget.ChildrenCount++
				currentWidget.Children = append(currentWidget.Children, &newWidget)
				println("inside of tag", data[seek+result+1:seek+result+result2])
				currentWidget = &newWidget
				*/
				var resultEnd = tagScraper.ScrapeInsideOfTag(newWidget, data[seek+result+1:seek+result+result2])
				newWidget.SetParent(currentWidget)
				newWidget.SetChildrenIndex(currentWidget.GetChildrenCount())
				currentWidget.SetChildrenCount(currentWidget.GetChildrenCount() + 1)
				currentWidget.AppendChild(newWidget)
				currentWidget = newWidget
				if resultEnd {
					currentWidget = currentWidget.GetParent()
				}

			}
			seek += result + result2 + 1
		}
	}
	return &documentWidget
}
