package HtmlParser

import (
	"strings"
	"sync"
)

type HtmlParser struct {
	workerGroup string
}

func (receiver *HtmlParser) ParseHtmlFromFile(document *HtmlElement, dat []byte, nodes chan *HtmlElement) {
	currentElement := document

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
				newHtmlElement := HtmlElement{
					HtmlTag:       HTML_UNTAGGED_TEXT,
					Text:          data[seek : seek+start],
					Parent:        currentElement,
					ChildrenIndex: currentElement.ChildrenCount,
					//draw and render function
				}
				currentElement.ChildrenCount++
				currentElement.Children = append(currentElement.Children, &newHtmlElement)
				nodes <- currentElement
			}
			if data[seek+start+1] == '/' {
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
				nodes <- &newElement
				if ParseInsideOfTag(currentElement, data[seek+start+1:seek+start+end], &wg) {
					currentElement = currentElement.Parent
				}
			}
			seek += start + end + 1
		}
	}
	wg.Wait()
	close(nodes)
}

func CreateDocumentElement() (document *HtmlElement) {
	document = new(HtmlElement)
	document.HtmlTag = HTML_DOCUMENT
	return
}