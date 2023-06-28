package CssParser

import (
	"gezgin_web_engine/GlobalTypes"
	"gezgin_web_engine/utils"
	"strings"
	"sync"
)

type CssParser struct {
	wg sync.WaitGroup
}

type Result struct {
	ruleCount          int
	CssStyleSheetRules []GlobalTypes.CssRuleInterface
}

func (receiver *Result) GetRuleCount() int {
	return receiver.ruleCount
}

func (receiver *Result) GetRuleByIndex(index int) GlobalTypes.CssRuleInterface {
	return receiver.CssStyleSheetRules[index]
}

type StyleElement interface {
}

func (receiver *CssParser) ParseCssFromStyleTag(styleElement StyleElement, styleText string) (result *Result) {
	result = new(Result)
	newCssStyleSheet := new(CssStyleSheet)
	styleText = utils.RemoveCharsFromString(styleText)
	seek := 0
	index := 0
	index = strings.Index(styleText[seek:], "{")
	for index != -1 { //maybe go routine for every cssText
		newCssRule := new(CssRule)
		index2 := strings.Index(styleText[seek:], "}")
		selectors := styleText[seek : seek+index]
		cssText := styleText[seek+index+1 : seek+index2]
		seek += index2 + 1
		newCssRule.SetStyleSheet(newCssStyleSheet)
		newCssRule.SetCssSelectors(selectors)
		newCssRule.SetCssDeclarationBlock(cssText)
		result.CssStyleSheetRules = append(result.CssStyleSheetRules, newCssRule)
		index = strings.Index(styleText[seek:], "{")
	}
	return
}

func ParseCssFromInlineStyle(cssText string) (m map[string]string) {
	m = make(map[string]string)
	declarations := strings.Split(cssText, ";")
	for _, declaration := range declarations {
		list := strings.Split(declaration, ":")
		m[list[0]] = list[1]
	}
	return
}

/*
BU style_engine'e gidicek

	func getCssWidget(selector string, channel chan *structs.StyleProperty) {
		var cssWidget *structs.StyleProperty
		switch selector[0] {
		case '#':
			cssWidget = tree.GetCssPropertiesByID(selector[0:])
			if cssWidget == nil {
				cssWidget = tree.CreateNewCssPropertiesByID(selector[0:])
			}
		case '.':
			cssWidget = tree.GetCssRulesByClass(selector[1:])
			if cssWidget == nil {
				cssWidget = tree.CreateNewCssPropertiesByClass(selector[1:])
			}
		default:
			tag := HtmlParser.GetElementTag(selector[0:])
			cssWidget = tree.GetCssRulesByElement(tag)
			if cssWidget == nil {
				cssWidget = tree.CreateNewCssPropertiesByElement(tag)
			}
		}
		channel <- cssWidget
	}
var wg sync.WaitGroup
*/

/*
	func SetCssProperties(currentWidget *widget.Widget) {
		var currentCssProperties *StyleEngine.StyleProperty
		for _, class := range currentWidget.Class {
			if currentCssProperties = StyleEngine.GetCssRulesByClass(class); currentCssProperties != nil {
				updateCssProperties(currentWidget.CssProperties, currentCssProperties)
			}
			if currentCssProperties = StyleEngine.GetRulesByElementAndClass(class, currentWidget.HtmlTag); currentCssProperties != nil {
				updateCssProperties(currentWidget.CssProperties, currentCssProperties)
			}
		}
		if currentCssProperties = StyleEngine.GetCssRulesByElement(currentWidget.HtmlTag); currentCssProperties != nil {
			updateCssProperties(currentWidget.CssProperties, currentCssProperties)
		}
		if currentCssProperties = StyleEngine.GetCssPropertiesByID(currentWidget.Id); currentCssProperties != nil {
			updateCssProperties(currentWidget.CssProperties, currentCssProperties)
		}
		if currentWidget.Style != "" {
			ParseCssFromInlineStyle(currentWidget.CssProperties, currentWidget.Style)
		}
	}

	func parserCssParameters(cssWidgetList []*StyleEngine.StyleProperty, cssText string) {
		varName, varValue, found := strings.Cut(cssText, ":")
		if found {
			index := utils.IndexFounder(StyleEngine.cssPropertiesNameList, varName, StyleEngine.cssPropertyCount)
			if index != -1 {
				function := StyleEngine.functionList[index]
				for _, properties := range cssWidgetList {
					if function != nil {
						function(properties, varValue)
					}
				}
			}
		}

}

func parseCssParameters(cssWidgetList []*StyleEngine.StyleProperty, cssText string) { //fix the Split function empty string problem

		cssProperties := strings.Split(cssText, ";")
		for _, property := range cssProperties {
			if len(property) > 0 {
				parserCssParameters(cssWidgetList, property)
			}
		}
	}
*/
/*
func getCssWidget(selector string, channel chan *StyleEngine.StyleProperty) {
	var cssWidget *StyleEngine.StyleProperty
	switch selector[0] {
	case '#':
		cssWidget = StyleEngine.GetCssPropertiesByID(selector[0:])
		if cssWidget == nil {
			cssWidget = StyleEngine.CreateNewCssPropertiesByID(selector[0:])
		}
	case '.':
		cssWidget = StyleEngine.GetCssPropertiesByClass(selector[1:])
		if cssWidget == nil {
			cssWidget = StyleEngine.CreateNewCssPropertiesByClass(selector[1:])
		}
	default:
		tag := HtmlParser.GetElementTag(selector[0:])
		cssWidget = StyleEngine.GetCssPropertiesByElement(tag)
		if cssWidget == nil {
			cssWidget = StyleEngine.CreateNewCssPropertiesByElement(tag)
		}
	}
	channel <- cssWidget
}

func getCssWidgetList(selectors string) (cssWidgetList []*StyleEngine.StyleProperty) {
	selectorList := strings.Split(selectors, ",")
	selectorCount := len(selectorList)
	channel := make(chan *StyleEngine.StyleProperty)
	for _, s := range selectorList {
		s = strings.TrimSpace(s)
		go getCssWidget(s, channel)
	}
	for selectorCount > 0 {
		cssWidgetList = append(cssWidgetList, <-channel)
		selectorCount--
	}
	return
}
*/
/*
	func ParseCssFromInlineStyle(properties *StyleEngine.StyleProperty, styleText string) {
		propertiesList := []*StyleEngine.StyleProperty{properties}
		parseCssParameters(propertiesList, styleText)
	}

	func ParseCssFromStyleTag(widget *widget.Widget) {
		cssTextWidget := widget.Children[0]
		styleWidget, ok := cssTextWidget.WidgetProperties.(tags.UntaggedText)
		if !ok {
			return
		}
		styleText := styleWidget.Value
		styleText = utils.RemoveCharsFromString(styleText)
		seek := 0
		index := 0
		index = strings.Index(styleText[seek:], "{")
		for index != -1 { //maybe go routine for every cssText
			index2 := strings.Index(styleText[seek:], "}")
			selectors := styleText[seek : seek+index]
			cssText := styleText[seek+index+1 : seek+index2]
			seek += index2 + 1
			cssWidgetList := getCssWidgetList(selectors)
			parseCssParameters(cssWidgetList, cssText)
			index = strings.Index(styleText[seek:], "{")
		}
		wg.Done() //this worker finished
	}
*/
/*
func SetInheritCssProperties(document *tags.Widget) {
	widgetList := []*tags.Widget{document}
	widgetIndexList := []int{0}
	currentIndex := 0
	for widgetIndexList[0] != document.ChildrenCount {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].ChildrenCount {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].ChildrenCount > 0 {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					widgetList = append(widgetList, widgetList[currentIndex].Children[widgetIndexList[currentIndex]])
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					StyleEngine.computeInheritCssProperties(currentWidget.CssProperties, currentWidget.Parent.CssProperties)
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].HtmlTag != HtmlParser.HTML_UNTAGGED_TEXT {
						currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
						StyleEngine.computeInheritCssProperties(currentWidget.CssProperties,
							currentWidget.Parent.CssProperties)
					}
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func initializeCssDocument(document *tags.Widget) {
	document.CssProperties = new(StyleEngine.StyleProperty)
	document.CssProperties.Display = enums.CSS_DISPLAY_TYPE_BLOCK
	document.CssProperties.Color = new(structs.ColorRGBA)
	document.CssProperties.Color.SetColorByRGB(0, 0, 0)
}

func ParseCssFromDocument(document *tags.Widget) {
	widgetList := []*tags.Widget{document}
	widgetIndexList := []int{0}
	initializeCssDocument(document)
	currentIndex := 0
	for widgetIndexList[0] != document.ChildrenCount {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].ChildrenCount {
			currentIndex--
			widgetIndexList = widgetIndexList[:len(widgetIndexList)-1]
			widgetList = widgetList[:len(widgetList)-1]
			widgetIndexList[currentIndex]++
		} else {
			if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].ChildrenCount > 0 {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					widgetList = append(widgetList, widgetList[currentIndex].Children[widgetIndexList[currentIndex]])
					widgetIndexList = append(widgetIndexList, 0)
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					if currentWidget.HtmlTag != HtmlParser.HTML_UNTAGGED_TEXT {
						SetCssProperties(currentWidget)
					}
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					if currentWidget.HtmlTag != HtmlParser.HTML_UNTAGGED_TEXT {
						SetCssProperties(currentWidget)
					}
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}
*/
/*
func CreateCssPropertiesFromStyleTag(styleElement StyleElement) {
	wg.Add(1)
	tree.CssStyleTagList = append(tree.CssStyleTagList, styleElement)
	go ParseCssFromStyleTag(styleElement)
}

func WaitCssScrapingOperations() {
	wg.Wait()
}
*/
