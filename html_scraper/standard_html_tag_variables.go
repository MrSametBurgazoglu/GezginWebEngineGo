package html_scraper

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
