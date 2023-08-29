package TaskManager

import (
	"gezgin_web_engine/CssParser"
	"gezgin_web_engine/FileManager"
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/JavascriptHandler"
	"gezgin_web_engine/NetworkManager"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/StyleEngine/enums"
	"gezgin_web_engine/StyleEngine/structs"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/eventSystem"
	"gezgin_web_engine/widgets"
	"github.com/gammazero/workerpool"
	"image"
	"image/draw"
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
	NetworkManager   *NetworkManager.NetworkManager
	ResourceManager  *ResourceManager.ResourceManager
	WebView          *image.RGBA
	HtmlWidget       widgets.WidgetInterface
	BodyWidget       widgets.WidgetInterface
}

func (receiver *TaskManager) Initialize() {
	receiver.WorkerPool = workerpool.New(runtime.NumCPU() - 1)
	receiver.htmlParser = new(HtmlParser.HtmlParser)
	receiver.cssParser = new(CssParser.CssParser)
	receiver.styleEngine = new(StyleEngine.StyleEngine)
	receiver.styleEngine.Initialize()
	receiver.javascriptEngine = new(JavascriptHandler.JavascriptEngine)
	receiver.NetworkManager = new(NetworkManager.NetworkManager)
	receiver.NetworkManager.Initialize()
	receiver.ResourceManager = new(ResourceManager.ResourceManager)
	receiver.ResourceManager.Initialize()
	receiver.ResourceManager.NetworkManager = receiver.NetworkManager
	receiver.WebView = image.NewRGBA(image.Rect(0, 0, ScreenProperties.WindowWidth, ScreenProperties.WindowHeight))
	draw.Draw(receiver.WebView, receiver.WebView.Bounds(), image.White, image.Point{Y: 0, X: 0}, draw.Src)
}

func (receiver *TaskManager) CreateFromFile(fileUrl string) {
	receiver.ResourceManager.Online = false
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
	receiver.ExecuteScripts()
}

func (receiver *TaskManager) CreateFromWeb(webUrl string) {
	receiver.ResourceManager.Online = true
	dat := receiver.NetworkManager.GetPage(webUrl)
	receiver.HtmlDocument = HtmlParser.CreateDocumentElement()
	nodes := make(chan *HtmlParser.HtmlElement)
	go receiver.htmlParser.ParseHtmlFromFile(receiver.HtmlDocument, dat, nodes)
	count := 0
	for node := range nodes {
		count += 1
		println(node.HtmlTag, "html tag")
		if node.HtmlTag == HtmlParser.HTML_SCRIPT {
			receiver.HandleScriptTag(node)
		} else if node.HtmlTag == HtmlParser.HTML_STYLE {
			styleSheet := receiver.styleEngine.CreateCssSheet(false)
			receiver.styleEngine.WorkerPool.Submit(func() { receiver.HandleStyleTag(node, styleSheet) }) //maybe worker pool
		} else if node.HtmlTag == HtmlParser.HTML_IMG {
			println("img tag", node.Attributes["src"])
			if src := node.Attributes["src"]; src != "" {
				receiver.HandleWebImgResource(src)
			}
		} else if node.HtmlTag == HtmlParser.HTML_IMG {
			if src := node.Attributes["src"]; src != "" {
				receiver.HandleWebImgResource(src)
			}
		} else if node.HtmlTag == HtmlParser.HTML_LINK {
			if node.Attributes["rel"] != "" && node.Attributes["href"] != "" {
				switch node.Attributes["rel"] {
				case "stylesheet":
					receiver.HandleWebLinkStyleSheet(node.Attributes["href"])
				}
			}
		}
	}
	receiver.styleEngine.WorkerPool.StopWait()
	receiver.CreateWidgetTree()
	receiver.SetStylePropertiesOfDocument()
	receiver.SetInheritStylePropertiesOfDocument()
	receiver.ExecuteScripts()
}

func (receiver *TaskManager) HandleWebImgResource(url string) {
	receiver.ResourceManager.CreateResourceFromWeb(url)
}

func (receiver *TaskManager) HandleWebLinkStyleSheet(url string) {
	styleSheet := receiver.styleEngine.CreateCssSheet(true)
	receiver.styleEngine.WorkerPool.Submit(func() {
		dat := receiver.NetworkManager.Get(url)
		styleTag := HtmlParser.HtmlElement{HtmlTag: HtmlParser.HTML_STYLE}
		untaggedText := HtmlParser.HtmlElement{HtmlTag: HtmlParser.HTML_UNTAGGED_TEXT, Text: string(dat)}
		styleTag.Children = append(styleTag.Children, &untaggedText)
		receiver.HandleStyleTag(&styleTag, styleSheet)
	})
}

func (receiver *TaskManager) HandleStyleTag(htmlElement *HtmlParser.HtmlElement, styleSheet *StyleEngine.StyleSheet) {
	result := receiver.cssParser.ParseCssFromStyleTag(htmlElement, htmlElement.Children[0].GetText())
	receiver.styleEngine.CreateStyleRules(styleSheet, result)
}

func (receiver *TaskManager) HandleScriptTag(scriptElement *HtmlParser.HtmlElement) {
	//give style element to v8 engine
	receiver.javascriptEngine.AppendScript(scriptElement.Children[0].GetText())
}

func (receiver *TaskManager) ExecuteScripts() {
	//give style element to v8 engine
	receiver.javascriptEngine.ExecuteScripts()
}

func (receiver *TaskManager) CreateWidgetTree() {
	receiver.DocumentWidget = new(widgets.DocumentWidget)
	element := receiver.FindBody()
	receiver.DocumentWidget.HtmlElement = element
	receiver.DocumentWidget.Initialize()
	receiver.DocumentWidget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_BLOCK
	receiver.DocumentWidget.ResourceManager = receiver.ResourceManager
	receiver.DocumentWidget.StyleProperty.Color = new(structs.ColorRGBA)
	receiver.DocumentWidget.StyleProperty.Color.SetColorByRGB(0, 0, 0)
	var wg sync.WaitGroup
	wg.Add(1)
	receiver.CreateWidgetForTree(receiver.DocumentWidget, element, &wg)
	wg.Wait()
}

func (receiver *TaskManager) CreateWidgetForTree(parentWidget widgets.WidgetInterface, parentHtmlElement *HtmlParser.HtmlElement, group *sync.WaitGroup) {
	for _, child := range parentHtmlElement.Children {
		function := widgets.WidgetFunctions[child.HtmlTag] // function will always return even if not drawen html elements
		newWidget := function(child, receiver)             // but function return value can be nil because not drawen html elements don't exist in widget tree
		newWidget.SetParent(parentWidget)
		parentWidget.AppendChild(newWidget)
		group.Add(1)
		go receiver.CreateWidgetForTree(newWidget, child, group)
	}
	group.Done()
}

// we will not need this function because body widget can set now body of document
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
	receiver.styleEngine.InitializeRoot()
	var wg sync.WaitGroup
	wg.Add(1)
	receiver.styleEngine.Root.InheritVariables(receiver.DocumentWidget.GetStyleProperty())
	receiver.SetStylePropertiesOfWidget(receiver.DocumentWidget, &wg)
	wg.Wait()
}

func (receiver *TaskManager) SetStylePropertiesOfWidget(widget widgets.WidgetInterface, group *sync.WaitGroup) { //TODO html tag must be string and can be custom
	widget.GetStyleProperty().ApplyCssRules(receiver.styleEngine, widget.GetID(), widget.GetClasses(), widget.GetHtmlTag(), widget.GetStyleRules())
	for _, child := range widget.GetChildren() {
		if child.GetHtmlTag() != 106 { //untagged text shouldn't have style property
			group.Add(1)
			widget.GetStyleProperty().InheritVariables(child.GetStyleProperty())
			go receiver.SetStylePropertiesOfWidget(child, group)
		}
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

func (receiver *TaskManager) Draw() {
	receiver.DocumentWidget.DrawPage(receiver.WebView)

}

func (receiver *TaskManager) Render() {
	receiver.DocumentWidget.RenderPage(receiver.WebView)
}

func (receiver *TaskManager) IsRendered() bool {
	return receiver.DocumentWidget.Rendered
}

func (receiver *TaskManager) SetRendered(rendered bool) {
	receiver.DocumentWidget.Rendered = rendered
}

func (receiver *TaskManager) SetHtmlElement(widget widgets.WidgetInterface) {
	receiver.HtmlWidget = widget
}

func (receiver *TaskManager) SetBodyElement(widget widgets.WidgetInterface) {
	receiver.BodyWidget = widget
}
