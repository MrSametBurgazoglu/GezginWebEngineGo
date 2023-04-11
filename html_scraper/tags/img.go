package tags

import (
	"gezgin_web_engine/html_scraper/widget"
	"strconv"
)

type HtmlTagImg struct {
	isMap          bool
	alt            string
	sizes          string
	src            string
	srcSet         string
	useMap         string
	longDesc       string
	height         int
	width          int
	crossOrigin    CrossOriginType
	loading        LoadingType
	referrerPolicy ReferrerPolicyType
}

func (receiver *HtmlTagImg) ContextReaderFunc(context string) {
	//receiver
	if context == "ismap" {
		receiver.isMap = true
	}
}

func (receiver *HtmlTagImg) VarReaderFunc(variableName string, variableValue string) {
	//receiver
	switch variableName {
	case "alt":
		receiver.alt = variableValue
	case "crossorgin":
		receiver.crossOrigin.Set(variableValue)
	case "height":
		receiver.height, _ = strconv.Atoi(variableValue)
	case "width":
		receiver.width, _ = strconv.Atoi(variableValue)
	case "loading":
		receiver.loading.Set(variableValue)
	case "longdesc":
		receiver.longDesc = variableValue
	case "referrerpolicy":
		receiver.referrerPolicy.Set(variableValue)
	case "sizes":
		receiver.sizes = variableValue
	case "src":
		receiver.src = variableValue
	case "srcset":
		receiver.srcSet = variableValue
	case "usemap":
		receiver.useMap = variableValue
	}
}

func SetWidgetPropertiesForImgTag(widget *widget.Widget) {
	widget.WidgetProperties = new(HtmlTagImg)
	widget.HaveAttrAsVar = true
	widget.HaveAttrAsContext = true
}
