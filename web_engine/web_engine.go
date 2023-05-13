package web_engine

import (
	"gezgin_web_engine/cssParser"
	"gezgin_web_engine/drawer"
	"gezgin_web_engine/htmlParser"
	"gezgin_web_engine/htmlParser/widget"
	"gezgin_web_engine/javascript_interpreter"
	"github.com/veandco/go-sdl2/sdl"
)

var document *widget.Widget

func GetDocument() *widget.Widget {
	return document
}

func OpenWebEngine(fileUrl string) {
	document = htmlParser.ParseHtmlFromFile(fileUrl)
	cssParser.WaitCssScrapingOperations()
	cssParser.ParseCssFromDocument(document)
	cssParser.SetInheritCssProperties(document)
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
