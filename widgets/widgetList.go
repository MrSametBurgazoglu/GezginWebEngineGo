package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/widget"
)

type TaskManagerInterface interface {
	HandleWebImgResource(string)
	HandleWebLinkStyleSheet(string)
	SetHtmlElement(widgetInterface widget.WidgetInterface)
	SetBodyElement(widgetInterface widget.WidgetInterface)
}

var WidgetFunctions = []func(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) widget.WidgetInterface{
	nil,
	SetWidgetPropertiesForATag,
	nil,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForAsideTag,
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
	SetWidgetPropertiesForCodeTag,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForDetailsTag,
	nil,
	nil,
	SetWidgetPropertiesForDivTag,
	nil,
	nil,
	SetWidgetPropertiesForEMTag,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForFooterTag,
	SetWidgetPropertiesForFormTag,
	SetWidgetPropertiesForH1Tag,
	SetWidgetPropertiesForH2Tag,
	SetWidgetPropertiesForH3Tag,
	SetWidgetPropertiesForH4Tag,
	SetWidgetPropertiesForH5Tag,
	SetWidgetPropertiesForH6Tag,
	nil,
	SetWidgetPropertiesForHeaderTag,
	SetWidgetPropertiesForHRTag,
	SetWidgetPropertiesForHtmlTag,
	nil,
	nil,
	SetWidgetPropertiesForImgTag,
	SetWidgetPropertiesForInputTag,
	SetWidgetPropertiesForLabelTag,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForLITag,
	nil,
	SetWidgetPropertiesForMainTag,
	nil,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForNavTag,
	nil,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForPTag,
	nil,
	nil,
	SetWidgetPropertiesForPRETag,
	nil,
	nil,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForSectionTag,
	nil,
	SetWidgetPropertiesForSmallTag,
	nil,
	SetWidgetPropertiesForSpanTag,
	SetWidgetPropertiesForStrongTag,
	nil,
	nil,
	SetWidgetPropertiesForSummaryTag,
	nil,
	SetWidgetPropertiesForSVGTag,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForTitleTag,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForULTag,
	nil,
	nil,
	nil,
	SetWidgetPropertiesForUntaggedText,
	SetCustomWidgetProperties,
}
