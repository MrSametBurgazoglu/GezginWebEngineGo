package TaskManager

import (
	"gezgin_web_engine/FileManager"
	"gezgin_web_engine/cssParser"
	"gezgin_web_engine/cssParser/tree"
	"gezgin_web_engine/eventSystem"
	"gezgin_web_engine/htmlParser"
	"github.com/gammazero/workerpool"
)

type TaskType uint8

type Task interface {
	ExecuteTask()
}

type TaskManager struct {
	WorkerPool       *workerpool.WorkerPool
	HtmlDocument     *htmlParser.HtmlElement
	cssPropertyLists *tree.CssStyleSheets
	EventMap         map[string][]eventSystem.InputWidget
	htmlParser       *htmlParser.HtmlParser
}

func (receiver *TaskManager) Initialize() {
	receiver.WorkerPool = workerpool.New(10)
}

func (receiver *TaskManager) CreateFromFile(fileUrl string) {
	dat := FileManager.LoadFile(fileUrl)
	receiver.HtmlDocument = htmlParser.CreateDocumentWidget()
	nodes := make(chan *htmlParser.HtmlElement)
	receiver.htmlParser.ParseHtmlFromFile(receiver.HtmlDocument, dat, nodes)
	for node := range nodes {
		if node.HtmlTag == htmlParser.HTML_SCRIPT {
			//to js interpreter
		} else if node.HtmlTag == htmlParser.HTML_STYLE {
			cssParser.CreateCssPropertiesFromStyleTag(node)
		}
	}
	cssParser.WaitCssScrapingOperations()
	/*
		cssParser.ParseCssFromDocument(receiver.Document)
		cssParser.SetInheritCssProperties(receiver.Document)
		javascript_interpreter.InitializeJSInterpreter(receiver.Document)
	*/
}

const (
	LOAD_FILE TaskType = iota
)
