package TaskManager

import (
	"gezgin_web_engine/FileManager"
	"gezgin_web_engine/cssParser"
	"gezgin_web_engine/cssParser/tree"
	"gezgin_web_engine/eventSystem"
	"gezgin_web_engine/htmlParser"
	"gezgin_web_engine/htmlParser/htmlVariables"
	"gezgin_web_engine/htmlParser/widget"
	"gezgin_web_engine/javascript_interpreter"
	"github.com/gammazero/workerpool"
)

type TaskType uint8

type Task interface {
	ExecuteTask()
}

type TaskManager struct {
	WorkerPool       *workerpool.WorkerPool
	Document         *widget.Widget
	cssPropertyLists *tree.CssStyleSheets
	EventMap         map[string][]eventSystem.InputWidget
	htmlParser       *htmlParser.HtmlParser
}

func (receiver *TaskManager) Initialize() {
	receiver.WorkerPool = workerpool.New(10)
}

func (receiver *TaskManager) CreateFromFile(fileUrl string) {
	dat := FileManager.LoadFile(fileUrl)
	receiver.Document = htmlParser.CreateDocumentWidget()
	nodes := make(chan *widget.Widget)
	receiver.htmlParser.ParseHtmlFromFile(receiver.Document, dat, nodes)
	for node := range nodes {
		if node.HtmlTag == htmlVariables.HTML_SCRIPT {
			//to js interpreter
		} else if node.HtmlTag == htmlVariables.HTML_STYLE {
			//to css interpreter
		}
	}
	cssParser.WaitCssScrapingOperations()
	cssParser.ParseCssFromDocument(receiver.Document)
	cssParser.SetInheritCssProperties(receiver.Document)
	javascript_interpreter.InitializeJSInterpreter(receiver.Document)
}

const (
	LOAD_FILE TaskType = iota
)
