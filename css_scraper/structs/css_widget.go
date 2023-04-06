package structs

import "gezgin_web_engine/css_scraper/enums"

type CssProperties struct {
	AccentColorInherit        bool
	AlignContentInherit       bool
	AlignItemsInherit         bool
	AlignSelfInherit          bool
	AnimationInherit          bool
	BackdropFilterInherit     bool
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

	AccentColor        *ColorRGBA
	AlignContent       enums.CssAlignType
	AlignItems         enums.CssAlignType
	AlignSelf          enums.CssAlignType
	Animation          *Animation
	Background         *Background
	BackdropFilter     enums.CssFilterType
	Color              *ColorRGBA
	Display            enums.CssDisplayType
	Margin             *Margin
	Padding            *Padding
	Height             uint
	MinHeight          uint
	MaxHeight          uint
	HeightValueType    enums.CssPropertyValueType
	MinHeightValueType enums.CssPropertyValueType
	MaxHeightValueType enums.CssPropertyValueType
	Width              uint
	MinWidth           uint
	MaxWidth           uint
	WidthValueType     enums.CssPropertyValueType
	MinWidthValueType  enums.CssPropertyValueType
	MaxWidthValueType  enums.CssPropertyValueType
	Position           enums.CssPositionType
	Top                uint
	TopValueType       enums.CssPropertyValueType
	Bottom             uint
	BottomValueType    enums.CssPropertyValueType
	Left               uint
	LeftValueType      enums.CssPropertyValueType
	Right              uint
	RightValueType     enums.CssPropertyValueType
	Visibility         enums.CssVisibilityType
}
