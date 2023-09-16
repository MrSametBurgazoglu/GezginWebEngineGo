package structs

import (
	"gezgin_web_engine/StyleProperty/enums"
)

type BorderLineWidth struct {
	BorderLineType enums.CssBorderlineType
	Value          int
}

type BorderWidth struct {
	BorderWidthTopInherit    bool
	BorderWidthLeftInherit   bool
	BorderWidthRightInherit  bool
	BorderWidthBottomInherit bool

	TopBorder    *BorderLineWidth
	LeftBorder   *BorderLineWidth
	RightBorder  *BorderLineWidth
	BottomBorder *BorderLineWidth
}

type BorderStyle struct {
	BorderStyleTopInherit    bool
	BorderStyleLeftInherit   bool
	BorderStyleRightInherit  bool
	BorderStyleBottomInherit bool

	BorderStyleTypeTop    enums.CssBorderStyleType
	BorderStyleTypeBottom enums.CssBorderStyleType
	BorderStyleTypeLeft   enums.CssBorderStyleType
	BorderStyleTypeRight  enums.CssBorderStyleType
}

type BorderColor struct {
	BorderColorTopInherit    bool
	BorderColorLeftInherit   bool
	BorderColorRightInherit  bool
	BorderColorBottomInherit bool

	TopBorderColor    *ColorRGBA
	BottomBorderColor *ColorRGBA
	RightBorderColor  *ColorRGBA
	LeftBorderColor   *ColorRGBA
}

type BorderRadius struct {
	BorderTopLeftRadiusInherit     bool
	BorderTopRightRadiusInherit    bool
	BorderBottomLeftRadiusInherit  bool
	BorderBottomRightRadiusInherit bool

	BorderTopLeftRadiusValue     int
	BorderTopRightRadiusValue    int
	BorderBottomLeftRadiusValue  int
	BorderBottomRightRadiusValue int

	BorderTopLeftRadiusType     enums.CssBorderRadiusType
	BorderTopRightRadiusType    enums.CssBorderRadiusType
	BorderBottomLeftRadiusType  enums.CssBorderRadiusType
	BorderBottomRightRadiusType enums.CssBorderRadiusType
}

type BorderSpacing struct {
	HSpacing int
	VSpacing int
}

type Border struct {
	BorderWidthInherit bool
	BorderStyleInherit bool
	BorderColorInherit bool
	BorderWidth        *BorderWidth
	BorderStyle        *BorderStyle
	BorderColor        *BorderColor
}
