package css_scraper

import (
	"gezgin_web_engine/css_scraper/structs"
	"gezgin_web_engine/css_scraper/tree"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"gezgin_web_engine/html_scraper/tags"
	"gezgin_web_engine/html_scraper/widget"
	"gezgin_web_engine/utils"
	"strings"
)

func SetCssProperties(currentWidget *widget.Widget) {
	//if currentWidget.HtmlTag == htmlVariables.HTML_UNTAGGED_TEXT {
	//	return // untagged text don't have class or element css properties
	//}
	var currentCssProperties *structs.CssProperties
	for _, class := range currentWidget.StandardHtmlVariables.Class {
		if currentCssProperties = tree.GetCssPropertiesByClass(class); currentCssProperties != nil {
			updateCssProperties(currentWidget.CssProperties, currentCssProperties)
		}
		if currentCssProperties = tree.GetCssPropertiesByElementAndClass(class, currentWidget.HtmlTag); currentCssProperties != nil {
			updateCssProperties(currentWidget.CssProperties, currentCssProperties)
		}
	}
	if currentCssProperties = tree.GetCssPropertiesByElement(currentWidget.HtmlTag); currentCssProperties != nil {
		updateCssProperties(currentWidget.CssProperties, currentCssProperties)
	}
	if currentCssProperties = tree.GetCssPropertiesByID(currentWidget.StandardHtmlVariables.Id); currentCssProperties != nil {
		updateCssProperties(currentWidget.CssProperties, currentCssProperties)
	}
	if currentWidget.StandardHtmlVariables.Style != "" {
		ScrapeCssFromInlineStyle(currentWidget.CssProperties, currentWidget.StandardHtmlVariables.Style)
	}
}

func scrapeCssParameters(cssWidgetList []*structs.CssProperties, cssText string) {
	varName, varValue, found := strings.Cut(cssText, ":")
	if found {
		index := utils.IndexFounder(cssPropertiesNameList, varName, cssPropertyCount)
		if index != -1 {
			function := functionList[index]
			for _, properties := range cssWidgetList {
				if function != nil {
					function(properties, varValue)
				}
			}
		}
	}

}

func scrapeCssProperties(cssWidgetList []*structs.CssProperties, cssText string) { //fix the Split function empty string problem
	cssProperties := strings.Split(cssText, ";")
	for _, property := range cssProperties {
		if len(property) > 0 {
			scrapeCssParameters(cssWidgetList, property)
		}
	}
}

func getCssWidget(selector string, channel chan *structs.CssProperties) {
	var cssWidget *structs.CssProperties
	switch selector[0] {
	case '#':
		cssWidget = tree.GetCssPropertiesByID(selector[0:])
		if cssWidget == nil {
			cssWidget = tree.CreateNewCssPropertiesByID(selector[0:])
		}
	case '.':
		cssWidget = tree.GetCssPropertiesByClass(selector[1:])
		if cssWidget == nil {
			cssWidget = tree.CreateNewCssPropertiesByClass(selector[1:])
		}
	default:
		//find index by string to element
		tag := htmlVariables.GetElementTag(selector[0:])
		cssWidget = tree.GetCssPropertiesByElement(tag)
		if cssWidget == nil {
			cssWidget = tree.CreateNewCssPropertiesByElement(tag)
		}
	}
	channel <- cssWidget
}

func getCssWidgetList(selectors string) (cssWidgetList []*structs.CssProperties) {
	selectorList := strings.Split(selectors, ",")
	selectorCount := len(selectorList)
	channel := make(chan *structs.CssProperties)
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

func ScrapeCssFromInlineStyle(properties *structs.CssProperties, styleText string) {
	propertiesList := []*structs.CssProperties{properties}
	scrapeCssProperties(propertiesList, styleText)
}

func scrapeCssFromStyleTag(widget *widget.Widget) {
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
	for index != -1 {
		index2 := strings.Index(styleText[seek:], "}")
		selectors := styleText[seek : seek+index]
		cssText := styleText[seek+index+1 : seek+index2]
		seek += index2 + 1
		cssWidgetList := getCssWidgetList(selectors)
		scrapeCssProperties(cssWidgetList, cssText)
		index = strings.Index(styleText[seek:], "{")
	}
}

func SetInheritCssProperties(document *widget.Widget) {
	widgetList := []*widget.Widget{document}
	widgetIndexList := []int{0}
	//initialize document
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
					computeInheritCssProperties(currentWidget.CssProperties, currentWidget.Parent.CssProperties)
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].HtmlTag != htmlVariables.HTML_UNTAGGED_TEXT {
						currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
						computeInheritCssProperties(currentWidget.CssProperties,
							currentWidget.Parent.CssProperties)
					}
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func initializeCssDocument(document *widget.Widget) {
	document.CssProperties = new(structs.CssProperties)
	document.CssProperties.Color = new(structs.ColorRGBA)
	document.CssProperties.Color.SetColorByRGB(0, 0, 0)
}

func ScrapeCssFromDocument(document *widget.Widget) {
	widgetList := []*widget.Widget{document}
	widgetIndexList := []int{0}
	//initialize document
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
					if currentWidget.HtmlTag != htmlVariables.HTML_UNTAGGED_TEXT {
						SetCssProperties(currentWidget)
					}
					currentIndex++
				} else {
					widgetIndexList[currentIndex]++
				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					currentWidget := widgetList[currentIndex].Children[widgetIndexList[currentIndex]]
					if currentWidget.HtmlTag != htmlVariables.HTML_UNTAGGED_TEXT {
						SetCssProperties(currentWidget)
					}
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func CreateCssPropertiesFromStyleTags() {
	for _, widget := range tree.CssStyleTagList {
		scrapeCssFromStyleTag(widget) //maybe we can call this function as goroutine
	}
}

func CreateCssPropertiesFromStyleFiles() {
	for _, s := range tree.CssStyleLinkList {
		println("file name:", s)
	}
}

func ExecuteCssScraper() {
	CreateCssPropertiesFromStyleTags()
	CreateCssPropertiesFromStyleFiles()
}
