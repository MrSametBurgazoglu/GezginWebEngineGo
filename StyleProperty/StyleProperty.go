package StyleProperty

import (
	"gezgin_web_engine/StyleProperty/enums"
	structs2 "gezgin_web_engine/StyleProperty/structs"
	"gezgin_web_engine/utils"
	"strings"
)

type StyleProperty struct {
	Parent       *StyleProperty
	Children     []*StyleProperty
	CssVariables map[string]string
	//this could be bit field
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
	FloatInherit              bool
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

	AccentColor        *structs2.ColorRGBA
	AlignContent       enums.CssAlignType
	AlignItems         enums.CssAlignType
	AlignSelf          enums.CssAlignType
	Animation          *structs2.Animation
	Background         *structs2.Background
	BackdropFilter     enums.CssFilterType
	Border             *structs2.Border
	BorderCollapseType enums.CssBorderCollapseType
	Color              *structs2.ColorRGBA
	Display            enums.CssDisplayType
	Font               *structs2.Font
	Float              enums.CssFloatType
	Margin             *structs2.Margin
	Padding            *structs2.Padding
	TextAlign          enums.CssTextAlignType
	FlexDirection      enums.CssFlexDirectionType
	Height             uint
	Width              uint
	MinHeight          uint
	MaxHeight          uint
	MinWidth           uint
	MaxWidth           uint
	Top                uint
	Bottom             uint
	Left               uint
	Right              uint
	HeightValueType    enums.CssPropertyValueType
	MinHeightValueType enums.CssPropertyValueType
	MaxHeightValueType enums.CssPropertyValueType
	WidthValueType     enums.CssPropertyValueType
	MinWidthValueType  enums.CssPropertyValueType
	MaxWidthValueType  enums.CssPropertyValueType
	Position           enums.CssPositionType
	TopValueType       enums.CssPropertyValueType
	BottomValueType    enums.CssPropertyValueType
	LeftValueType      enums.CssPropertyValueType
	RightValueType     enums.CssPropertyValueType
	Visibility         enums.CssVisibilityType
}

func (receiver *StyleProperty) Initialize() {
	receiver.CssVariables = make(map[string]string)
}

func (receiver *StyleProperty) ApplyInlineRules(m map[string]string) {
	for property, value := range m {
		receiver.ApplyDeclaration(property, value)
	}
}

func (receiver *StyleProperty) ApplyDeclaration(property string, value string) {
	property = strings.ReplaceAll(property, " ", "")
	value = strings.ReplaceAll(value, " ", "")
	if strings.HasPrefix(property, "--") {
		receiver.AddVariable(property, value)
	} else if strings.HasPrefix(value, "var(") {
		variable := receiver.GetVariable(value[4 : len(value)-1])
		value = variable
	}
	if strings.Contains(value, "!important") {
		if property == "display" {
			value = strings.ReplaceAll(value, "!important", "")
		}
		println(value)
	}
	index := utils.IndexFounder(cssPropertiesNameList, property, cssPropertyCount)
	if index != -1 {
		function := functionList[index]
		if function != nil {
			function(receiver, value)
		} else {
			println("CSS PROPERTY ", property, " NOT DEFINED")
		}
	}
}

func (receiver *StyleProperty) SetInheritStyleProperties(source *StyleProperty) {
	UpdateBackground(receiver, source)
	UpdateColor(receiver, source)
	UpdateText(receiver, source)
}

func (receiver *StyleProperty) AddVariable(key, value string) {
	receiver.CssVariables[key] = value
}

func (receiver *StyleProperty) GetVariable(key string) string {
	if strings.Contains(key, ",") {
		sepIndex := strings.Index(key, ",")
		keyValue := key[:sepIndex]
		initialValue := key[sepIndex+1:]
		value := receiver.CssVariables[keyValue]
		if value == "" {
			value = initialValue
		}
		return value
	}
	return receiver.CssVariables[key]
}

func (receiver *StyleProperty) InheritVariables(dest *StyleProperty) {
	for key, value := range receiver.CssVariables {
		dest.CssVariables[key] = value
	}
}
