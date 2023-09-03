package HtmlParser

import (
	"strings"
	"sync"
)

type HtmlParser struct {
	workerGroup string
}

func (receiver *HtmlParser) CreateUntaggedHtmlText(text string, currentElement *HtmlElement) {
	newHtmlElement := HtmlElement{
		HtmlTag:       HTML_UNTAGGED_TEXT,
		Text:          text,
		Parent:        currentElement,
		ChildrenIndex: currentElement.ChildrenCount,
		//draw and render function
	}
	currentElement.ChildrenCount++
	currentElement.Children = append(currentElement.Children, &newHtmlElement)
}

func (receiver *HtmlParser) ParseHtmlFromFile(document *HtmlElement, dat []byte, nodes chan *HtmlElement) {
	/*TODO DO NOT PARSE INSIDE SCRIPT TAG*/
	currentElement := document

	var wg sync.WaitGroup

	data := string(dat)
	dataLength := len(data)
	seek := 0
	notParseInside := false
	tagName := ""
	endTag := false
	for seek < dataLength {
		if data[seek] == ' ' || data[seek] == '\n' || data[seek] == '\r' {
			seek += 1
		} else {
			if notParseInside {
				index := strings.Index(data[seek:], "</"+tagName+">")
				receiver.CreateUntaggedHtmlText(data[seek:seek+index], currentElement)
				seek += index
				notParseInside = false
			}
			start := strings.Index(data[seek:], "<")
			end := strings.Index(data[seek+start:], ">")
			if start > 0 {
				receiver.CreateUntaggedHtmlText(data[seek:seek+start], currentElement)
				nodes <- currentElement
			}
			if data[seek:seek+4] == "<!--" {
				endCommentTag := strings.Index(data[seek:], "-->")
				seek += endCommentTag + 4
				println("there is comment")
				continue
			}
			if data[seek+start+1] == '/' {
				nodes <- currentElement
				if currentElement.Name != data[seek+start+2:seek+start+end] {
					println("something wrong")
					for currentElement.Name != data[seek+start+2:seek+start+end] {
						println(currentElement.Name, " -> ", currentElement.Parent.Name)
						currentElement = currentElement.Parent
						nodes <- currentElement
					}
				}
				println(currentElement.Name, " -> ", currentElement.Parent.Name)
				currentElement = currentElement.Parent
			} else {
				newElement := HtmlElement{
					Parent:        currentElement,
					ChildrenIndex: currentElement.ChildrenCount,
					//draw and render function
				}
				currentElement.ChildrenCount++
				currentElement.Children = append(currentElement.Children, &newElement)
				newElement.Attributes = make(map[string]string)
				currentElement = &newElement
				if data[seek+start+end-1] == '/' {
					end -= 2
					endTag, notParseInside, tagName = ParseInsideOfTag(currentElement, data[seek+start+1:seek+start+end])
					nodes <- currentElement
					currentElement = currentElement.Parent
					println(currentElement.Name, " -> ", currentElement.Parent.Name)
				} else {
					endTag, notParseInside, tagName = ParseInsideOfTag(currentElement, data[seek+start+1:seek+start+end])
					newElement.Name = tagName
					if endTag {
						if currentElement.Parent != nil {
							println(currentElement.Name, " -> ", currentElement.Parent.Name)
						}
						nodes <- currentElement
						currentElement = currentElement.Parent
					}
				}

				//nodes <- &newElement
			}
			seek += start + end + 1
		}
	}
	wg.Wait()
	println("waited")
	close(nodes)
}

func CreateDocumentElement() (document *HtmlElement) {
	document = new(HtmlElement)
	document.HtmlTag = HTML_DOCUMENT
	return
}
