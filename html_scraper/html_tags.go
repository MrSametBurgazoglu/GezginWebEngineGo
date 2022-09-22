package html_scraper

import (
	"gezgin_web_engine/widget"
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
	HTML_NOSCRIPT
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
	widgetPropertyFunction func(widget widget.Widget) //it's unique to html element some of them doesn't have this function
	//void (*widget_draw_function) (struct widget*, SDL_Renderer*);//for drawing rendered object
	//void (*widget_render_function) (struct widget*, SDL_Renderer*);//render element
	endTag bool
	draw   bool
}

var htmlTagList = [105]string{
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

func (htmlTag *HtmlTags) setHtmlTag(tag string) {

}

func (htmlTag *HtmlTags) getString() {

}
