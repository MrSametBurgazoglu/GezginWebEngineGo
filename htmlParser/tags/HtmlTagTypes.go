package tags

import (
	"gezgin_web_engine/utils"
)

type InputType uint8
type ButtonType uint8
type FormEncType uint8
type FormMethodType uint8
type FormTargetType uint8
type ShapeType uint8
type PreLoadOptionType uint8
type FormRelType uint8
type LoadingType uint8
type ReferrerPolicyType uint8
type SandboxAllowType uint8
type CrossOriginType uint8
type HttpEquivType uint8
type MetaNameType uint8
type OlType uint8
type TextAreaWrap uint8
type TableScope uint8
type TrackKind uint8
type HtmlDirectionType uint8

const (
	INPUT_TYPE_BUTTON InputType = iota
	INPUT_TYPE_CHECKBOX
	INPUT_TYPE_COLOR
	INPUT_TYPE_DATE
	INPUT_TYPE_DATETIME_LOCAL
	INPUT_TYPE_EMAIL
	INPUT_TYPE_FILE
	INPUT_TYPE_HIDDEN
	INPUT_TYPE_IMAGE
	INPUT_TYPE_MONTH
	INPUT_TYPE_NUMBER
	INPUT_TYPE_PASSWORD
	INPUT_TYPE_RADIO
	INPUT_TYPE_RANGE
	INPUT_TYPE_RESET
	INPUT_TYPE_SEARCH
	INPUT_TYPE_SUBMIT
	INPUT_TYPE_TEL
	INPUT_TYPE_TEXT
	INPUT_TYPE_TIME
	INPUT_TYPE_URL
	INPUT_TYPE_WEEK
)

const (
	BUTTON_TYPE_BUTTON ButtonType = iota
	BUTTON_TYPE_RESET
	BUTTON_TYPE_SUBMIT
)

const (
	FORM_ENC_TYPE_TEXT FormEncType = iota
	FORM_ENC_TYPE_MULTIPART
	FORM_ENC_TYPE_APPLICATION
)

const (
	FORM_METHOD_GET FormMethodType = iota
	FORM_METHOD_POST
)

const (
	FORM_TARGET_BLANK FormTargetType = iota
	FORM_TARGET_SELF
	FORM_TARGET_PARENT
	FORM_TARGET_TOP
	FORM_TARGET_CUSTOM
)

const (
	SHAPE_DEFAULT ShapeType = iota
	SHAPE_RECT
	SHAPE_CIRCLE
	SHAPE_POLY
)

const (
	PRELOAD_AUTO PreLoadOptionType = iota
	PRELOAD_METADATA
	PRELOAD_NONE
)

const (
	FORM_REL_ALTERNATE FormRelType = iota
	FORM_REL_AUTHOR
	FORM_REL_BOOKMARK
	FORM_REL_DNS_PREFETCH
	FORM_REL_EXTERNAL
	FORM_REL_HELP
	FORM_REL_ICON
	FORM_REL_LICENSE
	FORM_REL_NEXT
	FORM_REL_NOFOLLOW
	FORM_REL_NO_OPENER
	FORM_REL_NO_REFERRER
	FORM_REL_OPENER
	FORM_REL_PINGBACK
	FORM_REL_PRECONNECT
	FORM_REL_PREFETCH
	FORM_REL_PRELOAD
	FORM_REL_PRERENDER
	FORM_REL_PREV
	FORM_REL_SEARCH
	FORM_REL_STYLESHEET
	FORM_REL_TAG
)

const (
	LOADING_TYPE_EAGER LoadingType = iota
	LOADING_TYPE_LAZY
)

const (
	POLICY_NO_REFERRER ReferrerPolicyType = iota
	POLICY_NO_REFERRER_WHEN_DOWNGRADE
	POLICY_ORIGIN
	POLICY_ORIGIN_WHEN_CROSS_ORIGIN
	POLICY_SAME_ORIGIN
	POLICY_STRICT_ORIGIN
	POLICY_STRICT_ORIGIN_WHEN_CROSS_ORIGIN
	POLICY_UNSAFE_URL
)

const (
	SANDBOX_ALLOW_NONE SandboxAllowType = iota
	SANDBOX_ALLOW_FORMS
	SANDBOX_ALLOW_MODALS
	SANDBOX_ALLOW_ORIENTATION_LOCK
	SANDBOX_ALLOW_POINTER_LOCK
	SANDBOX_ALLOW_POPUPS
	SANDBOX_ALLOW_POPUPS_TO_ESCAPE_SANDBOX
	SANDBOX_ALLOW_PRESENTATION
	SANDBOX_ALLOW_SAME_ORIGIN
	SANDBOX_ALLOW_SCRIPTS
	SANDBOX_ALLOW_TOP_NAVIGATION
	SANDBOX_ALLOW_TOP_NAVIGATION_BY_USER_ACTIVATION
)

const (
	CROSS_ORIGIN_ANONYMOUS CrossOriginType = iota
	CROSS_ORIGIN_USE_CREDENTIALS
)

const (
	HTTP_EQUIV_CONTENT_SECURITY_POLICY HttpEquivType = iota
	HTTP_EQUIV_CONTENT_TYPE
	HTTP_EQUIV_DEFAULT_STYLE
	HTTP_EQUIV_REFRESH
)

const (
	META_NAME_APPLICATION_NAME MetaNameType = iota
	META_NAME_AUTHOR
	META_NAME_DESCRIPTION
	META_NAME_GENERATOR
	META_NAME_KEYWORDS
	META_NAME_VIEWPORT
)

const (
	OL_TYPE_1 OlType = iota
	OL_TYPE_a
	OL_TYPE_A
	OL_TYPE_i
	OL_TYPE_I
)

const (
	TEXTAREA_WRAP_HARD TextAreaWrap = iota
	TEXTAREA_WRAP_SOFT
)

const (
	TABLE_SCOPE_COL TableScope = iota
	TABLE_SCOPE_ROW
	TABLE_SCOPE_COLGROUP
	TABLE_SCOPE_ROWGROUP
)

const (
	TRACK_KIND_CAPTIONS TrackKind = iota
	TRACK_KIND_CHAPTERS
	TRACK_KIND_DESCRIPTIONS
	TRACK_KIND_METADATA
	TRACK_KIND_SUBTITLES
)

const (
	HTML_DIRECTION_LEFT HtmlDirectionType = iota
	HTML_DIRECTION_RIGHT
	HTML_DIRECTION_UP
	HTML_DIRECTION_DOWN
)

const referrerPolicyStringCount = 6
const formRelStringCount = 15

var referrerPolicyStrings = []string{
	"no-referrer",
	"no-referrer-when-downgrade",
	"origin",
	"same-origin",
	"origin-when-cross-origin",
	"strict-origin-when-cross-origin",
	"unsafe-url",
}

var formRelStrings = []string{
	"alternate",
	"author",
	"dns-prefetch",
	"help",
	"icon",
	"license",
	"next",
	"pingback",
	"preconnect",
	"prefetch",
	"preload",
	"prerender",
	"prev",
	"search",
	"stylesheet",
}

func (r *ReferrerPolicyType) Set(value string) {
	index := utils.IndexFounder(referrerPolicyStrings, value, referrerPolicyStringCount)
	*r = ReferrerPolicyType(index)
}
func (r *FormRelType) Set(value string) {
	index := utils.IndexFounder(formRelStrings, value, formRelStringCount)
	*r = FormRelType(index)
}

func (c *CrossOriginType) Set(value string) {
	if value == "anonymous" {
		*c = CROSS_ORIGIN_ANONYMOUS
	} else if value == "use-credentials" {
		*c = CROSS_ORIGIN_USE_CREDENTIALS
	}
}

func (t *FormMethodType) Set(value string) {
	switch value {
	case "get":
		*t = FORM_METHOD_GET
	case "post":
		*t = FORM_METHOD_POST
	}
}

func (e *FormEncType) Set(value string) {
	switch value {
	case "text/plain":
		*e = FORM_ENC_TYPE_TEXT
	case "multipart/form-data":
		*e = FORM_ENC_TYPE_MULTIPART
	case "application/x-www-form-urlencoded":
		*e = FORM_ENC_TYPE_APPLICATION
	}
}

func (t *FormTargetType) Set(value string) {
	switch value {
	case "_blank":
		*t = FORM_TARGET_BLANK
	case "_self":
		*t = FORM_TARGET_SELF
	case "_parent":
		*t = FORM_TARGET_PARENT
	case "_top":
		*t = FORM_TARGET_TOP
	default:
		*t = FORM_TARGET_CUSTOM
	}
}

func (b *ButtonType) Set(value string) {
	switch value {
	case "button":
		*b = BUTTON_TYPE_BUTTON
	case "reset":
		*b = BUTTON_TYPE_RESET
	case "submit":
		*b = BUTTON_TYPE_SUBMIT
	}
}

func (l *LoadingType) Set(value string) {
	if value == "eager" {
		*l = LOADING_TYPE_EAGER
	} else if value == "lazy" {
		*l = LOADING_TYPE_LAZY
	}
}

func (t *TextAreaWrap) Set(value string) {
	if value == "hard" {
		*t = TEXTAREA_WRAP_HARD
	} else if value == "soft" {
		*t = TEXTAREA_WRAP_SOFT
	}
}