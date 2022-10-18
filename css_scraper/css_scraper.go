package css_scraper

import (
	"gezgin_web_engine/css_scraper/structs"
	"gezgin_web_engine/css_scraper/tree"
	"gezgin_web_engine/html_scraper"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"gezgin_web_engine/utils"
	"strings"
)

func setCssProperties(currentWidget *htmlVariables.Widget) {
	var currentCssProperties *structs.CssProperties
	for _, class := range currentWidget.StandardHtmlVariables.Class {
		if currentCssProperties = tree.GetCssPropertiesByClass(class); currentCssProperties != nil {
			updateCssProperties(currentWidget.CssProperties, currentCssProperties)
		}
		if currentCssProperties = tree.GetCssPropertiesByElementAndClass(class, currentWidget.HtmlTag); currentCssProperties != nil {
			updateCssProperties(currentWidget.CssProperties, currentCssProperties)
		}
	}
	/* if currentCssProperties = tree.GetCssPropertiesByElement(); currentCssProperties != nil {
		updateCssProperties(currentWidget.CssProperties, currentCssProperties)
	} TODO change tree.CssPropertiesByElement string to html_Tag(index)
	*/
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
		cssWidget = tree.GetCssPropertiesByClass(selector[0:])
		if cssWidget == nil {
			cssWidget = tree.CreateNewCssPropertiesByClass(selector[0:])
		}
	default:
		cssWidget = tree.GetCssPropertiesByElement(selector[0:])
		if cssWidget == nil {
			cssWidget = tree.CreateNewCssPropertiesByElement(selector[0:])
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
	if styleText != "" {
		propertiesList := []*structs.CssProperties{properties}
		scrapeCssProperties(propertiesList, styleText)
	}
}

func scrapeCssFromStyleTag(widget *htmlVariables.Widget) {
	cssTextWidget := widget.Children[0]
	styleWidget, ok := cssTextWidget.WidgetProperties.(html_scraper.UntaggedText)
	if !ok {
		return
	}
	styleText := styleWidget.Value
	seek := 0
	index := 0
	index = strings.Index(styleText[seek:], "{")
	for index != -1 {
		index2 := strings.Index(styleText[seek:], "}")
		selectors := strings.Trim(styleText[seek:index], " \n")
		cssText := strings.Trim(styleText[seek+index+1:seek+index2], " \n")
		seek += index2
		cssWidgetList := getCssWidgetList(selectors)
		scrapeCssProperties(cssWidgetList, cssText)
		for _, properties := range cssWidgetList {
			println(properties)
		}
		index = strings.Index(styleText[seek:], "{")
	}
}

func SetInheritCssProperties(document *htmlVariables.Widget) {
	widgetList := []*htmlVariables.Widget{document}
	widgetIndexList := []int{0}
	//initialize document
	currentIndex := 0
	widgetCount := 0
	for widgetIndexList[0] != document.ChildrenCount {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].ChildrenCount {
			currentIndex--
			widgetCount--
			widgetIndexList = widgetIndexList[:widgetCount]
			widgetIndexList[currentIndex]++

		} else {
			if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].ChildrenCount > 0 {
				widgetCount++
				widgetList = append(widgetList, &widgetList[currentIndex].Children[widgetIndexList[currentIndex]])
				widgetIndexList[widgetCount-1] = 0
				currentIndex++
				if widgetList[currentIndex].Draw {
					computeInheritCssProperties(widgetList[currentIndex].CssProperties, widgetList[currentIndex-1].CssProperties)

				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					computeInheritCssProperties(widgetList[currentIndex].Children[widgetIndexList[currentIndex]].CssProperties,
						widgetList[currentIndex].CssProperties)
				}
				widgetIndexList[currentIndex]++
			}
		}
	}
}

func ScrapeCssFromDocument(document *htmlVariables.Widget) {
	widgetList := []*htmlVariables.Widget{document}
	widgetIndexList := []int{0}
	//initialize document
	currentIndex := 0
	widgetCount := 0
	for widgetIndexList[0] != document.ChildrenCount {
		if widgetIndexList[currentIndex] == widgetList[currentIndex].ChildrenCount {
			currentIndex--
			widgetCount--
			widgetIndexList = widgetIndexList[:widgetCount]
			widgetIndexList[currentIndex]++

		} else {
			if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].ChildrenCount > 0 {
				widgetCount++
				widgetList = append(widgetList, &widgetList[currentIndex].Children[widgetIndexList[currentIndex]])
				widgetIndexList[widgetCount-1] = 0
				currentIndex++
				if widgetList[currentIndex].Draw {
					setCssProperties(widgetList[currentIndex])
				}
			} else {
				if widgetList[currentIndex].Children[widgetIndexList[currentIndex]].Draw {
					setCssProperties(&widgetList[currentIndex].Children[widgetIndexList[currentIndex]])
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
