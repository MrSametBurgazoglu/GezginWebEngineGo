package structs

import "gezgin_web_engine/css_scraper/enums"

type CssProperties struct {
	AccentColorInherit        bool
	AlignContentInherit       bool
	AlignItemsInherit         bool
	AlignSelfInherit          bool
	AnimationInherit          bool
	BackdropfilterInherit     bool
	BackfaceVisibilityInherit bool
	BackgroundInherit         bool
	BorderInherit             bool
	BorderCollapseInherit     bool
	BorderImageInherit        bool
	BorderRadiusInherit       bool
	BorderSpacingInherit      bool
	ColorInherit              bool
	FontInherit               bool
	TextAlignInherit          bool
	TextAlignLastInherit      bool
	TextDecorationInherit     bool
	TextIndentInherit         bool
	TextJustifyInherit        bool
	TextOverflowInherit       bool
	TextShadowInherit         bool
	TextTransformInherit      bool
	ColumnCountInherit        bool
	ColumnFillInherit         bool
	ColumnGapInherit          bool
	ColumnRuleInherit         bool
	ColumnSpanInherit         bool
	ColumnWidthInherit        bool
	FlexInherit               bool
	FlexFlowInherit           bool
	PositionInherit           bool
	WidthInherit              bool
	MinWidthInherit           bool
	MaxWidthInherit           bool
	HeightInherit             bool
	MinHeightInherit          bool
	MaxHeightInherit          bool
	TopInherit                bool
	LeftInherit               bool
	RightInherit              bool
	BottomInherit             bool
	MarginInherit             bool
	OutlineInherit            bool
	OverflowInherit           bool
	PaddingInherit            bool
	VisibilityInherit         bool
	OpacityInherit            bool
	ResizeInherit             bool

	Background *Background
	Color      *ColorRGBA
	Display    enums.CssDisplayType
	Position   enums.CssPositionType
	Top        uint
	Bottom     uint
	Left       uint
	Right      uint8
}

func (receiver *CssProperties) GetCssProperties() *CssProperties {
	return receiver
}
