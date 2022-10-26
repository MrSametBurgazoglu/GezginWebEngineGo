package properties

import (
	"gezgin_web_engine/css_scraper/enums"
	"gezgin_web_engine/css_scraper/structs"
	"gezgin_web_engine/utils"
)

type CssBackgroundSize struct {
	backgroundSizeType1  enums.CssBackgroundSizeType
	backgroundSizeType2  enums.CssBackgroundSizeType
	backgroundSizeValue1 int
	backgroundSizeValue2 int
}

type BackgroundImageColor struct {
	Color   structs.ColorRGBA
	Percent uint8
}

type Background struct {
	BackgroundImageInherit      bool
	BackgroundPositionInherit   bool
	BackgroundColorInherit      bool
	BackgroundRepeatInherit     bool
	BackgroundOriginInherit     bool
	BackgroundClipInherit       bool
	BackgroundAttachmentInherit bool
	BackgroundSizeInherit       bool

	BackgroundColor          *structs.ColorRGBA
	ImageList                []string
	BackgroundImageColorList []*BackgroundImageColor
	BackgroundType           enums.CssBackgroundType
	BackgroundPositionType   enums.CssBackgroundPositionType
	BackgroundSize           *CssBackgroundSize
	BackgroundAttachment     enums.CssBackgroundAttachmentType
	BackgroundBlendMode      enums.CssBackgroundBlendModeType
	BackgroundClip           enums.CssBackgroundClipType
	BackgroundOrigin         enums.CssBackgroundOriginType
}

const BACKGROUND_BLEND_MODE_COUNT = 10
const BACKGROUND_REPEAT_TYPE_COUNT = 6
const BACKGROUND_ORIGIN_TYPE_COUNT = 3
const BACKGROUND_CLIP_TYPE_COUNT = 3
const BACKGROUND_ATTACHMENT_TYPE_COUNT = 3

var background_blend_mode_strings = []string{
	"color",
	"color-dodge",
	"darken",
	"lighten",
	"luminosity",
	"multiply",
	"normal",
	"overlay",
	"saturation",
	"screen",
}

var background_repeat_strings = []string{
	"no-repeat",
	"repeat",
	"repeat-x",
	"repeat-y",
	"round",
	"space",
}

var background_origin_strings = []string{
	"border-box",
	"content-box",
	"padding-box",
}

var background_attachment_strings = []string{
	"fixed",
	"local",
	"scroll",
}

func setBackgroundBlendMode(background Background, value string) {
	index := utils.IndexFounder(background_blend_mode_strings, value, BACKGROUND_BLEND_MODE_COUNT)
	if index != -1 {
		background.BackgroundBlendMode = enums.CssBackgroundBlendModeType(index)
	} else {
		background.BackgroundBlendMode = enums.CSS_BACKGROUND_BLEND_MODE_NORMAL
	}
}

func setBackgroundColor(background Background, value string) {
	background.BackgroundColor = new(structs.ColorRGBA)
	background.BackgroundColor.SetColor(value)
}
