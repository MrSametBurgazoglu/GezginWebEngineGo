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
			if start != 0 {
				println("heyy")
			}
			end := strings.Index(data[seek+start:], ">")
			if start > 0 {
				receiver.CreateUntaggedHtmlText(data[seek:seek+start], currentElement)
				nodes <- currentElement
			}
			if data[seek+start+1] == '/' {
				nodes <- currentElement
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
				endTag, notParseInside, tagName = ParseInsideOfTag(currentElement, data[seek+start+1:seek+start+end])
				newElement.Name = tagName
				println(data[seek : seek+50])
				println("htmlTag:", tagName)
				if strings.Contains(tagName, "ytd") {
					println("this is strange")
				}
				if endTag {
					currentElement = currentElement.Parent
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
