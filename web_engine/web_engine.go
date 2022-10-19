package web_engine

import (
	"gezgin_web_engine/css_scraper"
	"gezgin_web_engine/html_scraper"
)

func OpenWebEngine(fileUrl string) {
	//initialize drawer
	document := html_scraper.ScrapeHtmlFromFile(fileUrl)
	css_scraper.ExecuteCssScraper()
	css_scraper.ScrapeCssFromDocument(document)
	css_scraper.SetInheritCssProperties(document)
}

func DrawPage() {
	//draw document
	//set draw properties

}

func RenderPage() {
	//render document

}

func SendInput() {

}

func ExitWebEngine() {

}
