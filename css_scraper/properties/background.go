package properties

import "gezgin_web_engine/css_scraper/structs"

type CssBackgroundSize struct {
	backgroundSizeType1
}

type BackgroundImageColor struct {
	Color   structs.ColorRGBA
	Percent uint8
}

type Background struct {
	backgroundImageInherit      bool
	backgroundPositionInherit   bool
	backgroundColorInherit      bool
	backgroundRepeatInherit     bool
	backgroundOriginInherit     bool
	backgroundClipInherit       bool
	backgroundAttachmentInherit bool
	backgroundSizeInherit       bool

	backgroundColor *structs.ColorRGBA
	imageList       []string
}
