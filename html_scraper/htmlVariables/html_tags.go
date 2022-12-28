package htmlVariables

import (
	"gezgin_web_engine/html_scraper/HtmlElementWidget"
	"gezgin_web_engine/html_scraper/HtmlTags"
	"gezgin_web_engine/html_scraper/tags"
	"gezgin_web_engine/utils"
)

const HtmlTagCount = 105

//type HtmlTags uint8

const (
	HTML_DOCUMENT HtmlTags.HtmlTags = iota
	HTML_DOCTYPE
	HTML_A
	HTML_ABBR
	HTML_ADDRESS
	HTML_AREA
	HTML_ARTICLE
	HTML_ASIDE
	HTML_AUDIO
	HTML_B
	HTML_BASE
	HTML_BDI
	HTML_BDO
	HTML_BLOCKQUOTE
	HTML_BODY
	HTML_BR
	HTML_BUTTON
	HTML_CANVAS
	HTML_CAPTION
	HTML_CITE
	HTML_CODE
	HTML_COL
	HTML_COLGROUP
	HTML_DATA
	HTML_DATALIST
	HTML_DD
	HTML_DEL
	HTML_DETAILS
	HTML_DFN
	HTML_DIALOG
	HTML_DIV
	HTML_DL
	HTML_DT
	HTML_EM
	HTML_FIELDSET
	HTML_FIGCAPTION
	HTML_FIGURE
	HTML_FOOTER
	HTML_FORM
	HTML_H1
	HTML_H2
	HTML_H3
	HTML_H4
	HTML_H5
	HTML_H6
	HTML_HEAD
	HTML_HEADER
	HTML_HR
	HTML_HTML
	HTML_I
	HTML_IFRAME
	HTML_IMG
	HTML_INPUT
	HTML_INS
	HTML_KBD
	HTML_LABEL
	HTML_LEGEND
	HTML_LI
	HTML_LINK
	HTML_MAIN
	HTML_MAP
	HTML_MARK
	HTML_META
	HTML_METER
	HTML_NAV
	HTML_OL
	HTML_OPTGROUP
	HTML_OPTION
	HTML_OUTPUT
	HTML_P
	HTML_PARAM
	HTML_PICTURE
	HTML_PRE
	HTML_PROGRESS
	HTML_Q
	HTML_S
	HTML_SAMP
	HTML_SCRIPT
	HTML_SECTION
	HTML_SELECT
	HTML_SMALL
	HTML_SOURCE
	HTML_SPAN
	HTML_STRONG
	HTML_STYLE
	HTML_SUB
	HTML_SUMMARY
	HTML_SUP
	HTML_SVG
	HTML_TABLE
	HTML_TBODY
	HTML_TD
	HTML_TEMPLATE
	HTML_TEXTAREA
	HTML_TFOOT
	HTML_TH
	HTML_THEAD
	HTML_TIME
	HTML_TITLE
	HTML_TR
	HTML_TRACK
	HTML_U
	HTML_UL
	HTML_VAR
	HTML_VIDEO
	HTML_WBR
	HTML_UNTAGGED_TEXT
)

type HtmlTagVariables struct {
	tag                    HtmlTags.HtmlTags
	widgetPropertyFunction func() HtmlElementWidget.HtmlElementWidgetInterface //it's unique to html element some of them doesn't have this function
	endTag                 bool
}

var htmlTagList = []string{
	"!DOCTYPE",
	"a",
	"abbr",
	"address",
	"area",
	"article",
	"aside",
	"audio",
	"b",
	"base",
	"bdi",
	"bdo",
	"blockquote",
	"body",
	"br",
	"button",
	"canvas",
	"caption",
	"cite",
	"code",
	"col",
	"colgroup",
	"data",
	"datalist",
	"dd",
	"del",
	"details",
	"dfn",
	"dialog",
	"div",
	"dl",
	"dt",
	"em",
	"fieldset",
	"figcaption",
	"figure",
	"footer",
	"form",
	"h1",
	"h2",
	"h3",
	"h4",
	"h5",
	"h6",
	"head",
	"header",
	"hr",
	"html",
	"i",
	"iframe",
	"img",
	"input",
	"ins",
	"kbd",
	"label",
	"legend",
	"li",
	"link",
	"main",
	"map",
	"mark",
	"meta",
	"meter",
	"nav",
	"ol",
	"optgroup",
	"option",
	"output",
	"p",
	"param",
	"picture",
	"pre",
	"progress",
	"q",
	"s",
	"samp",
	"script",
	"section",
	"select",
	"small",
	"source",
	"span",
	"strong",
	"style",
	"sub",
	"summary",
	"sup",
	"svg",
	"table",
	"tbody",
	"td",
	"template",
	"textarea",
	"tfoot",
	"th",
	"thead",
	"time",
	"title",
	"tr",
	"track",
	"u",
	"ul",
	"var",
	"video",
	"wbr",
}

var tagHtmlVariables = []HtmlTagVariables{
	{tag: HTML_DOCTYPE, endTag: true},
	{tag: HTML_A, widgetPropertyFunction: tags.SetWidgetPropertiesForATag},
	{tag: HTML_ABBR, widgetPropertyFunction: tags.SetWidgetPropertiesForAbbrTag},
	{tag: HTML_ADDRESS},
	{tag: HTML_AREA, endTag: true},
	{tag: HTML_ARTICLE},
	{tag: HTML_ASIDE},
	{tag: HTML_AUDIO},
	{tag: HTML_B},
	{tag: HTML_BASE, endTag: true},
	{tag: HTML_BDI},
	{tag: HTML_BDO},
	{tag: HTML_BLOCKQUOTE},
	{tag: HTML_BODY, widgetPropertyFunction: tags.SetWidgetPropertiesForBodyTag},
	{tag: HTML_BR, endTag: true},
	{tag: HTML_BUTTON},
	{tag: HTML_CANVAS},
	{tag: HTML_CAPTION},
	{tag: HTML_CITE},
	{tag: HTML_CODE},
	{tag: HTML_COL, endTag: true},
	{tag: HTML_COLGROUP},
	{tag: HTML_DATA},
	{tag: HTML_DATALIST},
	{tag: HTML_DD},
	{tag: HTML_DEL},
	{tag: HTML_DETAILS},
	{tag: HTML_DFN},
	{tag: HTML_DIALOG},
	{tag: HTML_DIV, widgetPropertyFunction: tags.SetWidgetPropertiesForDivTag},
	{tag: HTML_DL},
	{tag: HTML_DT},
	{tag: HTML_EM},
	{tag: HTML_FIELDSET},
	{tag: HTML_FIGCAPTION},
	{tag: HTML_FIGURE},
	{tag: HTML_FOOTER},
	{tag: HTML_FORM},
	{tag: HTML_H1},
	{tag: HTML_H2},
	{tag: HTML_H3},
	{tag: HTML_H4},
	{tag: HTML_H5},
	{tag: HTML_H6},
	{tag: HTML_HEAD},
	{tag: HTML_HEADER},
	{tag: HTML_HR, endTag: true},
	{tag: HTML_HTML, widgetPropertyFunction: tags.SetWidgetPropertiesForHtmlTag},
	{tag: HTML_I},
	{tag: HTML_IFRAME},
	{tag: HTML_IMG, endTag: true},
	{tag: HTML_INPUT, endTag: true},
	{tag: HTML_INS},
	{tag: HTML_KBD},
	{tag: HTML_LABEL},
	{tag: HTML_LEGEND},
	{tag: HTML_LI},
	{tag: HTML_LINK, endTag: true},
	{tag: HTML_MAIN},
	{tag: HTML_MAP},
	{tag: HTML_MARK},
	{tag: HTML_META, endTag: true},
	{tag: HTML_METER},
	{tag: HTML_NAV},
	{tag: HTML_OL},
	{tag: HTML_OPTGROUP},
	{tag: HTML_OPTION},
	{tag: HTML_OUTPUT},
	{tag: HTML_P},
	{tag: HTML_PARAM, endTag: true},
	{tag: HTML_PICTURE},
	{tag: HTML_PRE},
	{tag: HTML_PROGRESS},
	{tag: HTML_Q},
	{tag: HTML_S},
	{tag: HTML_SAMP},
	{tag: HTML_SCRIPT},
	{tag: HTML_SECTION},
	{tag: HTML_SELECT},
	{tag: HTML_SMALL},
	{tag: HTML_SOURCE, endTag: true},
	{tag: HTML_SPAN},
	{tag: HTML_STRONG},
	{tag: HTML_STYLE},
	{tag: HTML_SUB},
	{tag: HTML_SUMMARY},
	{tag: HTML_SUP},
	{tag: HTML_SVG},
	{tag: HTML_TABLE},
	{tag: HTML_TBODY},
	{tag: HTML_TD},
	{tag: HTML_TEMPLATE},
	{tag: HTML_TEXTAREA},
	{tag: HTML_TFOOT},
	{tag: HTML_TH},
	{tag: HTML_THEAD},
	{tag: HTML_TIME},
	{tag: HTML_TITLE},
	{tag: HTML_TR},
	{tag: HTML_TRACK, endTag: true},
	{tag: HTML_U},
	{tag: HTML_UL},
	{tag: HTML_VAR},
	{tag: HTML_VIDEO},
	{tag: HTML_WBR},
}

func GetElementIndex(tag string) int {
	return utils.IndexFounder(htmlTagList, tag, HtmlTagCount)
}

func SetHtmlTag(tag string, widget HtmlElementWidget.HtmlElementWidgetInterface) bool {
	index := utils.IndexFounder(htmlTagList, tag, HtmlTagCount)
	if index != -1 {
		widget = tagHtmlVariables[index].widgetPropertyFunction()
		widget.SetHtmlTag(tagHtmlVariables[index].tag)
		return tagHtmlVariables[index].endTag
	}
	return false
}

func getString(htmlTag *HtmlTags.HtmlTags) string {
	return htmlTagList[*htmlTag]
}
