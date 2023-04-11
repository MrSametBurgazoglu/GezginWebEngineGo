package tags

import "gezgin_web_engine/html_scraper/widget"

type HtmlTagLink struct {
	href           string
	hrefLang       string
	media          string
	sizes          string
	title          string
	mediaType      string
	crossOrigin    CrossOriginType
	referrerPolicy ReferrerPolicyType
	formRel        FormRelType
}

func (receiver *HtmlTagLink) VarReaderFunc(variableName string, variableValue string) {
	//receiver
	switch variableName {
	case "crossorgin":
		receiver.crossOrigin.Set(variableValue)
	case "href":
		receiver.href = variableValue
	case "hreflang":
		receiver.hrefLang = variableValue
	case "media":
		receiver.media = variableValue
	case "referrerpolicy":
		receiver.referrerPolicy.Set(variableValue)
	case "rel":
		receiver.formRel.Set(variableValue)
	case "sizes":
		receiver.sizes = variableValue
	case "title":
		receiver.title = variableValue
	case "type":
		receiver.mediaType = variableValue
	}
}

func SetWidgetPropertiesForLinkTag(widget *widget.Widget) {
	widget.WidgetProperties = new(HtmlTagLink)
	widget.HaveAttrAsVar = true
}
