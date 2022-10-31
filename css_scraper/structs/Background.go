package structs

import "gezgin_web_engine/css_scraper/enums"

type CssBackgroundSize struct {
	backgroundSizeType1  enums.CssBackgroundSizeType
	backgroundSizeType2  enums.CssBackgroundSizeType
	backgroundSizeValue1 int
	backgroundSizeValue2 int
}

type BackgroundImageColor struct {
	Color   ColorRGBA
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

	BackgroundColor          *ColorRGBA
	ImageList                []string
	BackgroundImageColorList [3]*BackgroundImageColor
	BackgroundType           enums.CssBackgroundType
	BackgroundPositionType   enums.CssBackgroundPositionType
	BackgroundRepeatType     enums.CssBackgroundRepeatType
	BackgroundSize           *CssBackgroundSize
	BackgroundAttachmentType enums.CssBackgroundAttachmentType
	BackgroundBlendModeType  enums.CssBackgroundBlendModeType
	BackgroundClipType       enums.CssBackgroundClipType
	BackgroundOriginType     enums.CssBackgroundOriginType
}
