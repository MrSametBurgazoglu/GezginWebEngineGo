package TaskManager

import (
	"gezgin_web_engine/FileManager"
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/cssParser"
	"gezgin_web_engine/eventSystem"
	"gezgin_web_engine/htmlParser"
	"gezgin_web_engine/widgets"
	"github.com/gammazero/workerpool"
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
	"sync"
)

type TaskType uint8

type Task interface {
	ExecuteTask()
}

type TaskManager struct {
	WorkerPool     *workerpool.WorkerPool
	HtmlDocument   *htmlParser.HtmlElement
	EventMap       map[string][]eventSystem.InputWidget
	htmlParser     *htmlParser.HtmlParser
	cssParser      *cssParser.CssParser
	styleEngine    *StyleEngine.StyleEngine
	DocumentWidget *widgets.DocumentWidget
}

func (receiver *TaskManager) Initialize() {
	receiver.WorkerPool = workerpool.New(runtime.NumCPU() - 1)
	receiver.styleEngine.WorkerPool = workerpool.New(runtime.NumCPU() - 1)
	//receiver.styleEngine.Run()
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
			styleSheet := receiver.styleEngine.CreateCssSheet(false)
			receiver.styleEngine.WorkerPool.Submit(func() { receiver.HandleStyleTag(node, styleSheet) }) //maybe worker pool
		}
	}
	receiver.styleEngine.WorkerPool.StopWait()
	receiver.CreateWidgetTree()
	receiver.SetStylePropertiesOfDocument()
	receiver.SetInheritStylePropertiesOfDocument()
	/*
		cssParser.ParseCssFromDocument(receiver.Document)
		cssParser.SetInheritCssProperties(receiver.Document)
		javascript_interpreter.InitializeJSInterpreter(receiver.Document)
	*/
}

func (receiver *TaskManager) HandleStyleTag(htmlElement *htmlParser.HtmlElement, styleSheet *StyleEngine.StyleSheet) {
	result := receiver.cssParser.ParseCssFromStyleTag(htmlElement, htmlElement.Children[0].GetText())
	receiver.styleEngine.CreateStyleRules(styleSheet, result)
}

func (receiver *TaskManager) HandleScriptTag(styleElement *htmlParser.HtmlElement) {
	//give style element to v8 engine
}

func (receiver *TaskManager) CreateWidgetTree() {
	receiver.DocumentWidget = new(widgets.DocumentWidget)
	element := receiver.FindBody()
	var wg *sync.WaitGroup
	wg.Add(1)
	receiver.CreateWidgetForTree(receiver.DocumentWidget, element, wg)
	wg.Wait()
}

func (receiver *TaskManager) CreateWidgetForTree(parentWidget widgets.WidgetInterface, parentHtmlElement *htmlParser.HtmlElement, group *sync.WaitGroup) {
	for _, child := range parentHtmlElement.Children {
		newWidget := widgets.WidgetFunctions[child.HtmlTag]()
		newWidget.CopyFromHtmlElement(child)
		newWidget.SetParent(parentWidget)
		parentWidget.AppendChild(newWidget)
		group.Add(1)
		go receiver.CreateWidgetForTree(newWidget, child, group)
	}
	group.Done()
}

func (receiver *TaskManager) FindBody() *htmlParser.HtmlElement {
	elementList := []*htmlParser.HtmlElement{receiver.HtmlDocument}
	length := len(elementList)
	keepGo := true
	for keepGo {
		keepGo = false
		for _, w := range elementList {
			if w.HtmlTag == htmlParser.HTML_BODY {
				return w
			} else if w.ChildrenCount > 0 {
				for _, child := range w.Children {
					elementList = append(elementList, child)
					keepGo = true
				}
			}
		}
		if keepGo {
			elementList = elementList[length:]
			length = len(elementList)
		}
	}
	return nil
}

func (receiver *TaskManager) SetStylePropertiesOfDocument() {
	var wg *sync.WaitGroup
	wg.Add(1)
	receiver.SetStylePropertiesOfWidget(receiver.DocumentWidget, wg)
	wg.Wait()
}

func (receiver *TaskManager) SetStylePropertiesOfWidget(widget widgets.WidgetInterface, group *sync.WaitGroup) {
	widget.GetStyleProperty().ApplyCssRules(receiver.styleEngine, widget.GetID(), widget.GetClasses(), widget.GetHtmlTag(), widget.GetStyleRules())
	for _, child := range widget.GetChildren() {
		group.Add(1)
		go receiver.SetStylePropertiesOfWidget(child, group)
	}
	group.Done()
}

func (receiver *TaskManager) SetInheritStylePropertiesOfDocument() {
	var wg *sync.WaitGroup
	wg.Add(1)
	receiver.SetInheritStylePropertiesOfWidget(receiver.DocumentWidget, wg)
	wg.Wait()
}

func (receiver *TaskManager) SetInheritStylePropertiesOfWidget(widget widgets.WidgetInterface, group *sync.WaitGroup) {
	widget.GetStyleProperty().ApplyCssRules(receiver.styleEngine, widget.GetID(), widget.GetClasses(), widget.GetHtmlTag(), widget.GetStyleRules())
	for _, child := range widget.GetChildren() {
		child.GetStyleProperty().SetInheritStyleProperties(widget.GetStyleProperty())
		group.Add(1)
		go receiver.SetStylePropertiesOfWidget(child, group)
	}
	group.Done()
}

func (receiver *TaskManager) Draw(renderer *sdl.Renderer) {

}

func (receiver *TaskManager) Render(renderer *sdl.Renderer) {

}
