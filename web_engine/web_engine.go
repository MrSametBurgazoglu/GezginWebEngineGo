package web_engine

import (
	"gezgin_web_engine/css_scraper"
	"gezgin_web_engine/html_scraper"
)

func OpenWebEngine(fileUrl string) {
	//initialize drawer
	//initalize css scraper
	document := html_scraper.ScrapeHtmlFromFile(fileUrl)
	println(document)
	css_scraper.ExecuteCssScraper()
	css_scraper.ScrapeCssFromDocument(document)
}
