package html_scraper

import "fmt"

type StandardHtmlTagVariables struct {
	contentEditable bool
	draggable       bool
	hidden          bool
	accessKey       string
	//dir
	class []string
	id    string
	style string
}

func (tag *StandardHtmlTagVariables) setStandardVariables(name string, value string) bool {
	switch name {
	case "id":
		tag.id = value
	case "class":
		tag.class = append(tag.class, value)
	case "style":
		tag.style = value
	case "contenteditable":
		tag.contentEditable = value != "false"
	case "draggable":
		tag.draggable = value != "false"
	default:
		return false
	}
	fmt.Println("heyyo")
	return true
}

func (tag *StandardHtmlTagVariables) setStandardContextVariables(context string) bool {
	switch context {
	case "contenteditable":
		tag.contentEditable = true
	default:
		return false
	}
	return true
}
