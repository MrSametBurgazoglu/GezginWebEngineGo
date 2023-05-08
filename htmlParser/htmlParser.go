package htmlParser

import (
	"gezgin_web_engine/drawer/DrawProperties"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/drawer/structs"
	"gezgin_web_engine/htmlParser/htmlVariables"
	"gezgin_web_engine/htmlParser/tagParser"
	"gezgin_web_engine/htmlParser/tags"
	"gezgin_web_engine/htmlParser/widget"
	"os"
	"strings"
	"sync"
)

func createDocumentWidget() (documentWidget *widget.Widget) {
	documentWidget = new(widget.Widget)
	documentWidget.HtmlTag = htmlVariables.HTML_DOCUMENT
	documentWidget.DrawProperties = new(structs.DrawProperties)
	documentWidget.DrawProperties.Rect.X = 0
	documentWidget.DrawProperties.Rect.Y = 0
	documentWidget.DrawProperties.Rect.W = int32(ScreenProperties.WindowWidth)
	documentWidget.Draw = true
	return
}

func ParseHtmlFromFile(fileUrl string) *widget.Widget {
	dat, err := os.ReadFile(fileUrl)
	if err != nil {
		panic(err)
	}
	documentWidget := createDocumentWidget()
	currentWidget := documentWidget

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
				tagParser.UntaggedTextClosed(&newWidget)
			}
			result2 := strings.Index(data[seek+result:], ">")
			if data[seek+result+1] == '/' {
				currentWidget = currentWidget.Parent
			} else {
				newWidget := widget.Widget{
					Parent:        currentWidget,
					ChildrenIndex: currentWidget.ChildrenCount,
					//draw and render function
				}
				currentWidget.ChildrenCount++
				currentWidget.Children = append(currentWidget.Children, &newWidget)
				currentWidget = &newWidget
				if tagParser.ParseInsideOfTag(currentWidget, data[seek+result+1:seek+result+result2], &wg) {
					currentWidget = currentWidget.Parent
				}
			}
			seek += result + result2 + 1
		}
	}
	wg.Wait()
	return documentWidget
}
