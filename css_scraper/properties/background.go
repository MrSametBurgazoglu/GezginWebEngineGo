package properties

import (
	"gezgin_web_engine/css_scraper/enums"
	"gezgin_web_engine/css_scraper/structs"
	"gezgin_web_engine/utils"
)

const BackgroundBlendModeCount = 10
const BackgroundRepeatTypeCount = 6
const BACKGROUND_ORIGIN_TYPE_COUNT = 3
const BACKGROUND_CLIP_TYPE_COUNT = 3
const BACKGROUND_ATTACHMENT_TYPE_COUNT = 3

var backgroundBlendModeStrings = []string{
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

var backgroundRepeatStrings = []string{
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

func setBackgroundBlendMode(background *structs.Background, value string) {
	index := utils.IndexFounder(backgroundBlendModeStrings, value, BackgroundBlendModeCount)
	if index != -1 {
		background.BackgroundBlendModeType = enums.CssBackgroundBlendModeType(index)
	} else {
		background.BackgroundBlendModeType = enums.CSS_BACKGROUND_BLEND_MODE_NORMAL
	}
}

func setBackgroundRepeat(background *structs.Background, value string) {
	index := utils.IndexFounder(backgroundRepeatStrings, value, BackgroundRepeatTypeCount)
	if index != -1 {
		background.BackgroundRepeatType = enums.CssBackgroundRepeatType(index)
	} else {
		background.BackgroundRepeatType = enums.CSS_BACKGROUND_REPEAT_TYPE_NO_REPEAT
	}
}

func setBackgroundOrigin(background *structs.Background, value string) {
	index := utils.IndexFounder(background_origin_strings, value, BACKGROUND_ORIGIN_TYPE_COUNT)
	if index != -1 {
		background.BackgroundOriginType = enums.CssBackgroundOriginType(index)
	} else {
		background.BackgroundOriginType = enums.CSS_BACKGROUND_ORIGIN_PADDING_BOX
	}
}

func setBackgroundClip(background *structs.Background, value string) {
	index := utils.IndexFounder(background_origin_strings, value, BACKGROUND_CLIP_TYPE_COUNT)
	if index != -1 {
		background.BackgroundClipType = enums.CssBackgroundClipType(index)
	} else {
		background.BackgroundClipType = enums.CSS_BACKGROUND_CLIP_BORDER_BOX
	}
}

func setBackgroundAttachment(background *structs.Background, value string) {
	index := utils.IndexFounder(background_attachment_strings, value, BACKGROUND_ATTACHMENT_TYPE_COUNT)
	if index != -1 {
		background.BackgroundAttachmentType = enums.CssBackgroundAttachmentType(index)
	} else {
		background.BackgroundAttachmentType = enums.CSS_BACKGROUND_ATTACHMENT_SCROLL
	}
}

func setBackgroundColor(background *structs.Background, value string) {
	background.BackgroundColor = new(structs.ColorRGBA)
	background.BackgroundColor.SetColor(value)
}

/*
func setBackgroundImageColor(background *structs.Background, values []string) {
	for i, value := range values {
		background.BackgroundImageColorList[i] =
	}
}
*/

func BackgroundColorPropertySetValue(properties *structs.CssProperties, value string) {
	if value == "inherit" {
		if !properties.BackgroundInherit {
			if properties.Background == nil {
				properties.Background = new(structs.Background)
			}
			properties.Background.BackgroundColorInherit = true
		}
	} else {
		if properties.Background == nil {
			properties.Background = new(structs.Background)
		}
		if properties.BackgroundInherit {
			properties.Background.BackgroundSizeInherit = true
			properties.Background.BackgroundPositionInherit = true
			properties.Background.BackgroundOriginInherit = true
			properties.Background.BackgroundImageInherit = true
			properties.Background.BackgroundClipInherit = true
			properties.Background.BackgroundAttachmentInherit = true
			properties.Background.BackgroundRepeatInherit = true
		}
		properties.Background.BackgroundColorInherit = false
		if properties.Background.BackgroundColor == nil {
			properties.Background.BackgroundColor = new(structs.ColorRGBA)
		}
		if value == "initial" {
			properties.Background.BackgroundColor.SetColorByRGBA(0, 0, 0, 0)
		} else {
			setBackgroundColor(properties.Background, value)
		}
	}
}

func UpdateBackground(properties *structs.CssProperties, source *structs.CssProperties) {
	if source.BackgroundInherit {
		properties.BackgroundInherit = true
		properties.Background = nil
	} else if source.Background != nil {
		properties.Background = source.Background
	}
}
