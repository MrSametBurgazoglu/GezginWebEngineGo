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

func (tag *StandardHtmlTagVariables) setStandardVariables(name string, value string) (free int) {
	fmt.Println("heyyo")
	return 0
}

func (tag *StandardHtmlTagVariables) setStandardContextVariables(context string) {
	fmt.Println("heyyo")
}
