package htmlVariables

import (
	"gezgin_web_engine/css_scraper/structs"
	"gezgin_web_engine/utils"
)

const htmlTagCount = 105

type HtmlTags uint8

const (
	HTML_DOCUMENT HtmlTags = iota
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
	tag                    HtmlTags
	widgetPropertyFunction func(widget *Widget) //it's unique to html element some of them doesn't have this function
	//void (*widget_draw_function) (struct widget*, SDL_Renderer*);//for drawing rendered object
	//void (*widget_render_function) (struct widget*, SDL_Renderer*);//render element
	endTag bool
	draw   bool
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
	{tag: HTML_A, draw: true},
	{tag: HTML_ABBR, draw: true},
	{tag: HTML_ADDRESS},
	{tag: HTML_AREA, endTag: true, draw: true},
	{tag: HTML_ARTICLE, draw: true},
	{tag: HTML_ASIDE, draw: true},
	{tag: HTML_AUDIO, draw: true},
	{tag: HTML_B, draw: true},
	{tag: HTML_BASE, endTag: true},
	{tag: HTML_BDI},
	{tag: HTML_BDO},
	{tag: HTML_BLOCKQUOTE, draw: true},
	{tag: HTML_BODY, draw: true},
	{tag: HTML_BR, endTag: true, draw: true},
	{tag: HTML_BUTTON, draw: true},
	{tag: HTML_CANVAS, draw: true},
	{tag: HTML_CAPTION, draw: true},
	{tag: HTML_CITE, draw: true},
	{tag: HTML_CODE, draw: true},
	{tag: HTML_COL, endTag: true, draw: true},
	{tag: HTML_COLGROUP, draw: true},
	{tag: HTML_DATA},
	{tag: HTML_DATALIST, draw: true},
	{tag: HTML_DD, draw: true},
	{tag: HTML_DEL, draw: true},
	{tag: HTML_DETAILS, draw: true},
	{tag: HTML_DFN, draw: true},
	{tag: HTML_DIALOG},
	{tag: HTML_DIV, draw: true},
	{tag: HTML_DL, draw: true},
	{tag: HTML_DT, draw: true},
	{tag: HTML_EM, draw: true},
	{tag: HTML_FIELDSET, draw: true},
	{tag: HTML_FIGCAPTION, draw: true},
	{tag: HTML_FIGURE, draw: true},
	{tag: HTML_FOOTER, draw: true},
	{tag: HTML_FORM, draw: true},
	{tag: HTML_H1, draw: true},
	{tag: HTML_H2, draw: true},
	{tag: HTML_H3, draw: true},
	{tag: HTML_H4, draw: true},
	{tag: HTML_H5, draw: true},
	{tag: HTML_H6, draw: true},
	{tag: HTML_HEAD},
	{tag: HTML_HEADER, draw: true},
	{tag: HTML_HR, endTag: true},
	{tag: HTML_HTML, draw: true},
	{tag: HTML_I, draw: true},
	{tag: HTML_IFRAME, draw: true},
	{tag: HTML_IMG, endTag: true, draw: true},
	{tag: HTML_INPUT, endTag: true, draw: true},
	{tag: HTML_INS, draw: true},
	{tag: HTML_KBD, draw: true},
	{tag: HTML_LABEL, draw: true},
	{tag: HTML_LEGEND, draw: true},
	{tag: HTML_LI, draw: true},
	{tag: HTML_LINK, endTag: true},
	{tag: HTML_MAIN, draw: true},
	{tag: HTML_MAP, draw: true},
	{tag: HTML_MARK, draw: true},
	{tag: HTML_META, endTag: true},
	{tag: HTML_METER},
	{tag: HTML_NAV, draw: true},
	{tag: HTML_OL, draw: true},
	{tag: HTML_OPTGROUP},
	{tag: HTML_OPTION},
	{tag: HTML_OUTPUT, draw: true},
	{tag: HTML_P, draw: true},
	{tag: HTML_PARAM, endTag: true},
	{tag: HTML_PICTURE},
	{tag: HTML_PRE},
	{tag: HTML_PROGRESS},
	{tag: HTML_Q, draw: true},
	{tag: HTML_S, draw: true},
	{tag: HTML_SAMP, draw: true},
	{tag: HTML_SCRIPT},
	{tag: HTML_SECTION, draw: true},
	{tag: HTML_SELECT},
	{tag: HTML_SMALL, draw: true},
	{tag: HTML_SOURCE, endTag: true},
	{tag: HTML_SPAN},
	{tag: HTML_STRONG, draw: true},
	{tag: HTML_STYLE},
	{tag: HTML_SUB, draw: true},
	{tag: HTML_SUMMARY, draw: true},
	{tag: HTML_SUP, draw: true},
	{tag: HTML_SVG},
	{tag: HTML_TABLE, draw: true},
	{tag: HTML_TBODY, draw: true},
	{tag: HTML_TD, draw: true},
	{tag: HTML_TEMPLATE},
	{tag: HTML_TEXTAREA},
	{tag: HTML_TFOOT, draw: true},
	{tag: HTML_TH, draw: true},
	{tag: HTML_THEAD, draw: true},
	{tag: HTML_TIME},
	{tag: HTML_TITLE},
	{tag: HTML_TR, draw: true},
	{tag: HTML_TRACK, endTag: true},
	{tag: HTML_U, draw: true},
	{tag: HTML_UL, draw: true},
	{tag: HTML_VAR, draw: true},
	{tag: HTML_VIDEO, draw: true},
	{tag: HTML_WBR},
}

func (htmlTag *HtmlTags) SetHtmlTag(tag string, widget *Widget) bool {
	index := utils.IndexFounder(htmlTagList, tag, htmlTagCount)
	if index != -1 {
		widget.HtmlTag = tagHtmlVariables[index].tag
		if tagHtmlVariables[index].widgetPropertyFunction != nil {
			tagHtmlVariables[index].widgetPropertyFunction(widget)
		}
		if tagHtmlVariables[index].draw {
			widget.CssProperties = new(structs.CssProperties)
			//set render and draw functions and draw properties
		}
		return tagHtmlVariables[index].endTag
	}
	return false
}

func (htmlTag *HtmlTags) getString() string {
	return htmlTagList[*htmlTag]
}
