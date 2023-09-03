package StyleEngine

import (
	"gezgin_web_engine/StyleEngine/enums"
	"gezgin_web_engine/StyleEngine/structs"
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

	AccentColor        *structs.ColorRGBA
	AlignContent       enums.CssAlignType
	AlignItems         enums.CssAlignType
	AlignSelf          enums.CssAlignType
	Animation          *structs.Animation
	Background         *structs.Background
	BackdropFilter     enums.CssFilterType
	Border             *structs.Border
	BorderCollapseType enums.CssBorderCollapseType
	Color              *structs.ColorRGBA
	Display            enums.CssDisplayType
	Font               *structs.Font
	Margin             *structs.Margin
	Padding            *structs.Padding
	TextAlign          enums.CssTextAlignType
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

/*TODO MAKE STYLE ENGINE ROOT TO HTML ELEMENT STYLE PROPERTY AND GIVE IT HERE FOR GLOBAL CSS VARIABLES*/
/*TODO MAKE STYLE PROPERTIES MAP FOR CSS VARIABLES AND GIVE HERE PARENT STYLE PROPERTY FOR APPLYING*/
func (receiver *StyleProperty) ApplyCssRules(styleEngine *StyleEngine, id string, classes []string, htmlName string, styleMap map[string]string) {
	externalTagRules := styleEngine.GetCssRulesByTag(htmlName, true)
	for _, rule := range externalTagRules {
		receiver.ApplyRules(rule)
	}
	internalTagRules := styleEngine.GetCssRulesByTag(htmlName, false)
	for _, rule := range internalTagRules {
		receiver.ApplyRules(rule)
	}
	if classes != nil {
		for _, class := range classes {
			externalClassRules := styleEngine.GetCssRulesByClass(class, true)
			for _, rule := range externalClassRules {
				receiver.ApplyRules(rule)
			}
			internalClassRules := styleEngine.GetCssRulesByClass(class, false)
			for _, rule := range internalClassRules {
				receiver.ApplyRules(rule)
			}
		}
	}
	if id != "" {
		externalIDRules := styleEngine.GetCssRulesByID(id, true)
		for _, rule := range externalIDRules {
			receiver.ApplyRules(rule)
		}
		internalIDRules := styleEngine.GetCssRulesByID(id, false)
		for _, rule := range internalIDRules {
			receiver.ApplyRules(rule)
		}
	}
	receiver.ApplyInlineRules(styleMap)
}

func (receiver *StyleProperty) ApplyRules(rule *CssRuleListItem) {
	for property, value := range rule.declarations {
		receiver.ApplyDeclaration(property, value)
	}
}

func (receiver *StyleProperty) ApplyInlineRules(m map[string]string) {
	for property, value := range m {
		receiver.ApplyDeclaration(property, value)
	}
}

func (receiver *StyleProperty) ApplyDeclaration(property string, value string) {
	if strings.HasPrefix(property, "--") {
		receiver.AddVariable(property, value)
	} else if strings.HasPrefix(value, "var(") {
		variable := receiver.GetVariable(value[4 : len(value)-1])
		value = variable
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
