package TaskManager

import (
	"gezgin_web_engine/CssParser"
	"gezgin_web_engine/FileManager"
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/JavascriptHandler"
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/eventSystem"
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
	WorkerPool       *workerpool.WorkerPool
	HtmlDocument     *HtmlParser.HtmlElement
	EventMap         map[string][]eventSystem.InputWidget
	htmlParser       *HtmlParser.HtmlParser
	cssParser        *CssParser.CssParser
	styleEngine      *StyleEngine.StyleEngine
	javascriptEngine *JavascriptHandler.JavascriptEngine
	DocumentWidget   *widgets.DocumentWidget
}

func (receiver *TaskManager) Initialize() {
	receiver.WorkerPool = workerpool.New(runtime.NumCPU() - 1)
	receiver.htmlParser = new(HtmlParser.HtmlParser)
	receiver.cssParser = new(CssParser.CssParser)
	receiver.styleEngine = new(StyleEngine.StyleEngine)
	receiver.styleEngine.WorkerPool = workerpool.New(runtime.NumCPU() - 1)
	receiver.javascriptEngine = new(JavascriptHandler.JavascriptEngine)
	//receiver.styleEngine.Run()
}

func (receiver *TaskManager) CreateFromFile(fileUrl string) {
	dat := FileManager.LoadFile(fileUrl)
	receiver.HtmlDocument = HtmlParser.CreateDocumentElement()
	nodes := make(chan *HtmlParser.HtmlElement)
	go receiver.htmlParser.ParseHtmlFromFile(receiver.HtmlDocument, dat, nodes)
	for node := range nodes {
		if node.HtmlTag == HtmlParser.HTML_SCRIPT {
			receiver.HandleScriptTag(node)
		} else if node.HtmlTag == HtmlParser.HTML_STYLE {
			styleSheet := receiver.styleEngine.CreateCssSheet(false)
			receiver.styleEngine.WorkerPool.Submit(func() { receiver.HandleStyleTag(node, styleSheet) }) //maybe worker pool
		}
	}
	receiver.styleEngine.WorkerPool.StopWait()
	receiver.CreateWidgetTree()
	receiver.SetStylePropertiesOfDocument()
	receiver.SetInheritStylePropertiesOfDocument()
	/*
		CssParser.ParseCssFromDocument(receiver.Document)
		CssParser.SetInheritCssProperties(receiver.Document)
		JavascriptHandler.InitializeJSInterpreter(receiver.Document)
	*/
}

func (receiver *TaskManager) HandleStyleTag(htmlElement *HtmlParser.HtmlElement, styleSheet *StyleEngine.StyleSheet) {
	result := receiver.cssParser.ParseCssFromStyleTag(htmlElement, htmlElement.Children[0].GetText())
	receiver.styleEngine.CreateStyleRules(styleSheet, result)
}

func (receiver *TaskManager) HandleScriptTag(scriptElement *HtmlParser.HtmlElement) {
	//give style element to v8 engine
	receiver.javascriptEngine.AppendScript(scriptElement.Children[0].GetText())
}

func (receiver *TaskManager) ExecuteScripts(scriptElement *HtmlParser.HtmlElement) {
	//give style element to v8 engine
	receiver.javascriptEngine.ExecuteScripts()
}

func (receiver *TaskManager) CreateWidgetTree() {
	receiver.DocumentWidget = new(widgets.DocumentWidget)
	element := receiver.FindBody()
	receiver.DocumentWidget.HtmlElement = element
	receiver.DocumentWidget.Initialize()
	var wg sync.WaitGroup
	wg.Add(1)
	receiver.CreateWidgetForTree(receiver.DocumentWidget, element, &wg)
	wg.Wait()
}

func (receiver *TaskManager) CreateWidgetForTree(parentWidget widgets.WidgetInterface, parentHtmlElement *HtmlParser.HtmlElement, group *sync.WaitGroup) {
	for _, child := range parentHtmlElement.Children {
		function := widgets.WidgetFunctions[child.HtmlTag]
		newWidget := function(child)
		newWidget.SetParent(parentWidget)
		parentWidget.AppendChild(newWidget)
		group.Add(1)
		go receiver.CreateWidgetForTree(newWidget, child, group)
	}
	group.Done()
}

func (receiver *TaskManager) FindBody() *HtmlParser.HtmlElement {
	elementList := []*HtmlParser.HtmlElement{receiver.HtmlDocument}
	length := len(elementList)
	keepGo := true
	for keepGo {
		keepGo = false
		for _, w := range elementList {
			if w.HtmlTag == HtmlParser.HTML_BODY {
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
	var wg sync.WaitGroup
	wg.Add(1)
	receiver.SetStylePropertiesOfWidget(receiver.DocumentWidget, &wg)
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
	var wg sync.WaitGroup
	wg.Add(1)
	receiver.SetInheritStylePropertiesOfWidget(receiver.DocumentWidget, &wg)
	wg.Wait()
}

func (receiver *TaskManager) SetInheritStylePropertiesOfWidget(widget widgets.WidgetInterface, group *sync.WaitGroup) {
	//widget.GetStyleProperty().ApplyCssRules(receiver.styleEngine, widget.GetID(), widget.GetClasses(), widget.GetHtmlTag(), widget.GetStyleRules())
	for _, child := range widget.GetChildren() {
		if child.GetStyleProperty() != nil {
			child.GetStyleProperty().SetInheritStyleProperties(widget.GetStyleProperty())
			group.Add(1)
			go receiver.SetInheritStylePropertiesOfWidget(child, group)
		}
	}
	group.Done()
}

func (receiver *TaskManager) Draw(renderer *sdl.Renderer) {
	receiver.DocumentWidget.DrawPage(renderer)

}

func (receiver *TaskManager) Render(renderer *sdl.Renderer) {
	receiver.DocumentWidget.RenderPage(renderer)
}

func (receiver *TaskManager) IsRendered() bool {
	return receiver.DocumentWidget.Rendered
}

func (receiver *TaskManager) SetRendered(rendered bool) {
	receiver.DocumentWidget.Rendered = rendered
}
