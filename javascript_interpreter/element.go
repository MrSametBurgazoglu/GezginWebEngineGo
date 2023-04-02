package javascript_interpreter

import (
	"gezgin_web_engine/css_scraper"
	"gezgin_web_engine/html_scraper/widget"
	v8 "rogchap.com/v8go"
)

var Element *v8.ObjectTemplate

func searchElementById(elementId string) *widget.Widget {
	widgetList := []*widget.Widget{globalDocument}
	length := len(widgetList)
	keepGo := true
	for keepGo {
		keepGo = false
		for _, w := range widgetList {
			if w.ChildrenCount > 0 {
				for _, child := range w.Children {
					widgetList = append(widgetList, child)
					keepGo = true
				}
			}
			if w.StandardHtmlVariables.Id == elementId {
				return w
			}
		}
		if keepGo {
			widgetList = widgetList[length:]
			length = len(widgetList)
		}
	}
	return nil
}

func setAttributeByID(elementId, attribute, value string) {
	element := searchElementById(elementId)
	if element == nil {
		return
	}
	switch attribute {
	case "style":
		println("set style attribute", value)
		element.StandardHtmlVariables.Style = value
		css_scraper.SetCssProperties(element)
		println(element.CssProperties.Background.BackgroundColor)
	}
}

func CreateElementTemplate(iso *v8.Isolate) {
	Element = v8.NewObjectTemplate(iso)
	SetAttribute := v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		filterType, _ := info.This().Get("filter_type")
		filterValue, _ := info.This().Get("filter_value")
		attribute := info.Args()[0]
		value := info.Args()[1]
		if filterType.String() == "id" {
			setAttributeByID(filterValue.String(), attribute.String(), value.String())
		}
		return nil
	})
	Element.Set("setAttribute", SetAttribute)
}
