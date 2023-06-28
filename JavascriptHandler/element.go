package JavascriptHandler

import (
	"gezgin_web_engine/widgets"
	v8 "rogchap.com/v8go"
)

var Element *v8.ObjectTemplate

func (receiver *JavascriptEngine) searchElementById(elementId string) widgets.WidgetInterface {
	widgetList := []widgets.WidgetInterface{receiver.DocumentWidget}
	length := len(widgetList)
	keepGo := true
	for keepGo {
		keepGo = false
		for _, w := range widgetList {
			if w.GetChildrenCount() > 0 {
				for _, child := range w.GetChildren() {
					widgetList = append(widgetList, child)
					keepGo = true
				}
			}
			if w.GetID() == elementId {
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

func (receiver *JavascriptEngine) setAttributeByID(elementId, attribute, value string) {
	element := receiver.searchElementById(elementId)
	if element == nil {
		return
	}
	/*
		switch attribute {
		case "style":
			element.Style = value
			CssParser.SetCssProperties(element)
			globalDocument.Rendered = false
		}
	*/
}

func CreateElementTemplate(iso *v8.Isolate) {
	Element = v8.NewObjectTemplate(iso)
	SetAttribute := v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		/*
			filterType, _ := info.This().Get("filter_type")
			filterValue, _ := info.This().Get("filter_value")
			attribute := info.Args()[0]
			value := info.Args()[1]

			if filterType.String() == "id" {
				setAttributeByID(filterValue.String(), attribute.String(), value.String())
			}
		*/
		return nil
	})
	Element.Set("setAttribute", SetAttribute)
}
