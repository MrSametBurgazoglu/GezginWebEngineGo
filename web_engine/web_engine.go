package web_engine

import (
	"gezgin_web_engine/TaskManager"
	"gezgin_web_engine/drawer"
	"github.com/veandco/go-sdl2/sdl"
)

type WebTab struct {
	taskManager *TaskManager.TaskManager
}

func (receiver *WebTab) OpenWebPageFromFile(fileUrl string) {
	receiver.taskManager.CreateFromFile(fileUrl)
}

func (receiver *WebTab) OpenWebPageFromWeb(Url string) {
	receiver.taskManager.CreateFromWeb(Url)
}

func (receiver *WebTab) DrawPage(renderer *sdl.Renderer) {
	receiver.taskManager.Draw(renderer)
}

func (receiver *WebTab) RenderPage(renderer *sdl.Renderer) {
	receiver.taskManager.Render(renderer)
}

func (receiver *WebTab) IsRendered() bool {
	return receiver.taskManager.IsRendered()
}

func (receiver *WebTab) SetRendered(rendered bool) {
	receiver.taskManager.SetRendered(rendered)
}

func NewTab() *WebTab {
	newWebTab := new(WebTab)
	newWebTab.taskManager = new(TaskManager.TaskManager)
	newWebTab.taskManager.Initialize()
	return newWebTab
}

func InitDrawer(height, width int) {
	drawer.SetWindowSize(height, width)
}

/*
var document *tags.Widget

	func GetDocument() *tags.Widget {
		return document
	}

	func OpenWebEngine(fileUrl string) {
		document = HtmlParser.ParseHtmlFromFile(fileUrl)
		CssParser.WaitCssScrapingOperations()
		CssParser.ParseCssFromDocument(document)
		CssParser.SetInheritCssProperties(document)
		JavascriptHandler.InitializeJSInterpreter(document)
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
*/
func SendInput() {

}

func ExitWebEngine() {

}
