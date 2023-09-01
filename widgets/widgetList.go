package widgets

import "gezgin_web_engine/HtmlParser"

type TaskManagerInterface interface {
	HandleWebImgResource(string)
	HandleWebLinkStyleSheet(string)
	SetHtmlElement(widgetInterface WidgetInterface)
	SetBodyElement(widgetInterface WidgetInterface)
}

var WidgetFunctions = []func(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface{
	nil,
	SetWidgetPropertiesForATag,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForBodyTag,
	nil,
	SetWidgetPropertiesForButtonTag,
	SetWidgetPropertiesForCanvasTag,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForDetailsTag,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForDivTag,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForH1Tag,
	SetWidgetPropertiesForH2Tag,
	SetWidgetPropertiesForH3Tag,
	SetWidgetPropertiesForH4Tag,
	SetWidgetPropertiesForH5Tag,
	SetWidgetPropertiesForH6Tag,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForHtmlTag,
	nil,
	nil,
	SetWidgetPropertiesForImgTag,
	nil,
	SetWidgetPropertiesForLabelTag,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForPTag,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForUntaggedText,
	SetCustomWidgetProperties,
}
