package web_engine

import (
	"gezgin_web_engine/css_scraper"
	"gezgin_web_engine/drawer"
	"gezgin_web_engine/html_scraper"
	"gezgin_web_engine/html_scraper/widget"
	"gezgin_web_engine/javascript_interpreter"
	"github.com/veandco/go-sdl2/sdl"
)

var document *widget.Widget

func GetDocument() *widget.Widget {
	return document
}

func OpenWebEngine(fileUrl string) {
	//initialize drawer
	document = html_scraper.ScrapeHtmlFromFile(fileUrl)
	css_scraper.WaitCssScrapingOperations()
	css_scraper.ScrapeCssFromDocument(document)
	css_scraper.SetInheritCssProperties(document)
	javascript_interpreter.InitializeJSInterpreter(document)
}

func InitDrawer() {
	drawer.LoadDefaultFont()
	drawer.SetWindowSize(600, 800)
}

func DrawPage(renderer *sdl.Renderer) {
	drawer.SetDrawPropertiesDocument(document)
	drawer.DrawDocument(document, renderer)
}

func RenderPage(renderer *sdl.Renderer) {
	drawer.RenderDocument(document, renderer)

}

func SendInput() {

}

func ExitWebEngine() {

}
