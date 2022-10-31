package web_engine

import (
	"gezgin_web_engine/css_scraper"
	"gezgin_web_engine/drawer"
	"gezgin_web_engine/html_scraper"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"github.com/veandco/go-sdl2/sdl"
)

var document *htmlVariables.Widget

func OpenWebEngine(fileUrl string) {
	//initialize drawer
	document = html_scraper.ScrapeHtmlFromFile(fileUrl)
	css_scraper.ExecuteCssScraper()
	css_scraper.ScrapeCssFromDocument(document)
	css_scraper.SetInheritCssProperties(document)
}

func DrawPage() {
	//draw document
	//set draw properties

}

func RenderPage(renderer *sdl.Renderer) {
	drawer.RenderDocument(document, renderer)

}

func SendInput() {

}

func ExitWebEngine() {

}
