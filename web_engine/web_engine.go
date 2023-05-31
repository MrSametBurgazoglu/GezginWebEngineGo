package web_engine

import (
	"gezgin_web_engine/cssParser"
	"gezgin_web_engine/cssParser/tree"
	"gezgin_web_engine/drawer"
	"gezgin_web_engine/eventSystem"
	"gezgin_web_engine/htmlParser"
	"gezgin_web_engine/htmlParser/widget"
	"gezgin_web_engine/javascript_interpreter"
	"github.com/veandco/go-sdl2/sdl"
)

type WebTab struct {
	document         *widget.Widget
	cssPropertyLists *tree.CssStyleSheets
	EventMap         map[string][]eventSystem.InputWidget
	htmlParser       *htmlParser.HtmlParser
}

func (receiver *WebTab) OpenWebPageFromFile(fileUrl string) {
	receiver.document = htmlParser.CreateDocumentWidget()
	receiver.htmlParser.ParseHtmlFromFile(receiver.document, fileUrl)
	cssParser.WaitCssScrapingOperations()
	cssParser.ParseCssFromDocument(receiver.document)
	cssParser.SetInheritCssProperties(receiver.document)
	javascript_interpreter.InitializeJSInterpreter(receiver.document)
}

func (receiver *WebTab) DrawPage(renderer *sdl.Renderer) {
	drawer.DrawDocument(receiver.document, renderer)
}

func NewTab() *WebTab {
	newWebTab := new(WebTab)
	newWebTab.cssPropertyLists = new(tree.CssStyleSheets)
	return newWebTab
}

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

func InitDrawer(height, width int) {
	drawer.SetWindowSize(height, width)
}

func DrawPage(renderer *sdl.Renderer) {
	drawer.DrawDocument(document, renderer)
}

func RenderPage(renderer *sdl.Renderer) {
	drawer.SetDrawPropertiesForWidgets(document, renderer)
}

func SendInput() {

}

func ExitWebEngine() {

}
