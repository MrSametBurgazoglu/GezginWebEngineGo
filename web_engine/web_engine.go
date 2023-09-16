package web_engine

import (
	"gezgin_web_engine/TaskManager"
	"gezgin_web_engine/drawer"
	"image"
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

func (receiver *WebTab) DrawPage() {
	receiver.taskManager.Draw()
}

func (receiver *WebTab) RenderPage() {
	receiver.taskManager.Render()
}

func (receiver *WebTab) IsRendered() bool {
	return receiver.taskManager.IsRendered()
}

func (receiver *WebTab) SetRendered(rendered bool) {
	receiver.taskManager.SetRendered(rendered)
}

func (receiver *WebTab) GetWebView() *image.RGBA {
	return receiver.taskManager.WebView.Image
}

func NewTab() *WebTab {
	newWebTab := new(WebTab)
	newWebTab.taskManager = new(TaskManager.TaskManager)
	newWebTab.taskManager.Initialize()
	return newWebTab
}

func InitDrawer(width, height int) {
	drawer.SetWindowSize(width, height)
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

	func DrawPage(mainImage *image.RGBA) {
		drawer.DrawDocument(document, renderer)
	}

	func RenderPage(mainImage *image.RGBA) {
		drawer.SetDrawPropertiesForWidgets(document, renderer)
	}
*/
func SendInput() {

}

func ExitWebEngine() {

}
