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

type CssProperties interface {
	CreateCssPropertiesFromStyleTag(*widget.Widget)
}

type JavascriptInterpreter interface {
	CreateScriptElements()
}

type HtmlParser struct {
	workerGroup           string
	cssProperties         CssProperties
	javascriptInterpreter JavascriptInterpreter
}

func (receiver *HtmlParser) ParseHtmlFromFile(documentWidget *widget.Widget, dat []byte, nodes chan *widget.Widget) {
	currentWidget := documentWidget

	var wg sync.WaitGroup

	data := string(dat)
	dataLength := len(data)
	seek := 0
	for seek < dataLength {
		if data[seek] == ' ' || data[seek] == '\n' {
			seek += 1
		} else {
			start := strings.Index(data[seek:], "<")
			end := strings.Index(data[seek+start:], ">")
			if start > 0 {
				//make untagged text to strip
				newWidget := widget.Widget{
					HtmlTag:          htmlVariables.HTML_UNTAGGED_TEXT,
					WidgetProperties: tags.UntaggedText{Value: data[seek : seek+start]},
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
				nodes <- &newWidget
			}
			if data[seek+start+1] == '/' {
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
				nodes <- &newWidget
				if tagParser.ParseInsideOfTag(currentWidget, data[seek+start+1:seek+start+end], &wg) {
					currentWidget = currentWidget.Parent
				}
			}
			seek += start + end + 1
		}
	}
	wg.Wait()
}

func CreateDocumentWidget() (documentWidget *widget.Widget) {
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
	documentWidget := CreateDocumentWidget()
	currentWidget := documentWidget

	var wg sync.WaitGroup

	data := string(dat)
	dataLength := len(data)
	seek := 0
	for seek < dataLength {
		if data[seek] == ' ' || data[seek] == '\n' {
			seek += 1
		} else {
			start := strings.Index(data[seek:], "<")
			end := strings.Index(data[seek+start:], ">")
			if start > 0 {
				//make untagged text to strip
				newWidget := widget.Widget{
					HtmlTag:          htmlVariables.HTML_UNTAGGED_TEXT,
					WidgetProperties: tags.UntaggedText{Value: data[seek : seek+start]},
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
			if data[seek+start+1] == '/' {
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
				if tagParser.ParseInsideOfTag(currentWidget, data[seek+start+1:seek+start+end], &wg) {
					currentWidget = currentWidget.Parent
				}
			}
			seek += start + end + 1
		}
	}
	wg.Wait()
	return documentWidget
}
