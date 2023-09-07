package StyleEngine

import (
	"gezgin_web_engine/GlobalTypes"
	"gezgin_web_engine/StyleEngine/CssRuleListItem"
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/widget"
	"github.com/gammazero/workerpool"
	"runtime"
	"strings"
)

type CssSheet interface {
}

type CssParserResult interface {
	GetRuleCount() int
	GetRuleByIndex(int) GlobalTypes.CssRuleInterface
}

type HtmlElement interface {
	GetAttribute(string) string
	GetHtmlTag() int
}

type StyleSheet struct {
	external    bool
	cssRuleList *CssRuleList
}

type StyleEngine struct {
	WorkerPool        *workerpool.WorkerPool
	CssStyleSheetList []*StyleSheet
	Root              *StyleProperty.StyleProperty
}

func (receiver *StyleEngine) Initialize() {
	receiver.WorkerPool = workerpool.New(runtime.NumCPU() - 1)
	receiver.Root = new(StyleProperty.StyleProperty)
	receiver.Root.Initialize()
}

func (receiver *StyleEngine) InitializeRoot() {
	rules := receiver.GetAllCssRules(":root", nil, "root")
	receiver.ApplyRules(receiver.Root, rules)
}

func (receiver *StyleEngine) CreateCssSheet(external bool) (cssSheet *StyleSheet) {
	cssSheet = new(StyleSheet)
	cssSheet.external = external
	cssSheet.cssRuleList = new(CssRuleList)
	cssSheet.cssRuleList.Initialize()
	receiver.CssStyleSheetList = append(receiver.CssStyleSheetList, cssSheet)
	return
}

func (receiver *StyleEngine) CreateStyleRules(styleSheet *StyleSheet, cssRules CssParserResult) {
	ruleCount := cssRules.GetRuleCount()
	for index := 0; index < ruleCount; index++ {
		rule := cssRules.GetRuleByIndex(index)
		selectors := rule.GetSelectors()
		declarations := rule.GetDeclarations()
		cssRuleList := styleSheet.GetCssRuleListItems(selectors)
		for _, declaration := range declarations {
			property := declaration.GetProperty()
			value := declaration.GetValue()
			for _, item := range cssRuleList {
				item.Declarations[property] = value
			}
		}
	}
}

// GetCssRuleListItems /* make this as one goroutine*/
func (receiver *StyleSheet) GetCssRuleListItems(selectors []string) (cssRuleList []*CssRuleListItem.CssRuleListItem) {
	for _, s := range selectors {
		s = strings.TrimSpace(s)
		cssRuleList = append(cssRuleList, receiver.GetCssRuleItem(s))
	}
	return
}

func (receiver *StyleSheet) GetCssRuleItem(selector string) *CssRuleListItem.CssRuleListItem {
	var cssRuleList *CssRuleListItem.CssRuleListItem
	switch selector[0] {
	case '#':
		cssRuleList = receiver.cssRuleList.GetCssRulesByID(selector[0:])
		if cssRuleList == nil {
			cssRuleList = receiver.cssRuleList.CreateNewCssRulesByID(selector[0:])
		}
	case '.':
		secondDotIndex := strings.Index(selector[1:], ".")
		if secondDotIndex == -1 {
			cssRuleList = receiver.cssRuleList.GetCssRulesByClass(selector[1:])
			if cssRuleList == nil {
				cssRuleList = receiver.cssRuleList.CreateNewCssRulesByClass(selector[1:])
			}
		} else {
			if selector[secondDotIndex] == ' ' {
				firstClass := selector[1:secondDotIndex]
				secondClass := selector[secondDotIndex+2:]
				cssRuleList = receiver.cssRuleList.GetCssRulesByClassDescendant(firstClass, secondClass)
				if cssRuleList == nil {
					cssRuleList = receiver.cssRuleList.CreateNewCssRulesByClassDescendant(firstClass, secondClass)
				}
			} else {
				firstClass := selector[1:secondDotIndex]
				secondClass := selector[secondDotIndex:]
				cssRuleList = receiver.cssRuleList.GetCssRulesByClassBoth(firstClass, secondClass)
				if cssRuleList == nil {
					cssRuleList = receiver.cssRuleList.CreateNewCssRulesByClassBoth(firstClass, secondClass)
				}
			}
		}
	case ':':
		cssRuleList = receiver.cssRuleList.GetCssRulesByID(selector)
		if cssRuleList == nil {
			cssRuleList = receiver.cssRuleList.CreateNewCssRulesByID(selector)
		}
	default:
		tag := selector[0:]
		cssRuleList = receiver.cssRuleList.GetCssRulesByElement(tag)
		if cssRuleList == nil {
			cssRuleList = receiver.cssRuleList.CreateNewCssPropertiesByElement(tag)
		}
	}
	return cssRuleList
}

func (receiver *StyleEngine) GetAllCssRulesByElement(htmlName string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	rules := receiver.GetCssRulesByTag(htmlName, external)
	rules = append(rules, receiver.GetCssRulesByElementAndAttribute(htmlName, external)...)
	rules = append(rules, receiver.GetCssRulesByElementAndAttributeAndValue(htmlName, external)...)
	rules = append(rules, receiver.GetCssRulesByElementAndAttributeAndBeginsValue(htmlName, external)...)
	rules = append(rules, receiver.GetCssRulesByElementAndAttributeAndStartsValue(htmlName, external)...)
	rules = append(rules, receiver.GetCssRulesByElementAndAttributeAndContainsSubstringsValue(htmlName, external)...)
	rules = append(rules, receiver.GetCssRulesByElementAndAttributeAndEndsValue(htmlName, external)...)

	return
}

/*TODO PROCESSES HERE WITH STYLE PROPERTY DONE IN STYLE_ENGINE NOT STYLE PROPERTY*/
func (receiver *StyleEngine) GetAllCssRules(id string, classes []string, htmlName string) []*CssRuleListItem.CssRuleListItem {
	var rules []*CssRuleListItem.CssRuleListItem
	rules = append(rules, receiver.GetCssRulesByTag(htmlName, true)...)
	rules = append(rules, receiver.GetCssRulesByTag(htmlName, false)...)
	if classes != nil {
		for _, class := range classes {
			rules = append(rules, receiver.GetCssRulesByClass(class, true)...)
			rules = append(rules, receiver.GetCssRulesByClass(class, false)...)
		}
	}
	if id != "" {
		rules = append(rules, receiver.GetCssRulesByID(id, true)...)
		rules = append(rules, receiver.GetCssRulesByID(id, false)...)
	}
	for _, rule := range rules {
		println(rule.Function)
	}
	return rules
}

func (receiver *StyleEngine) ApplyRules(styleProperty *StyleProperty.StyleProperty, rules []*CssRuleListItem.CssRuleListItem) {
	for _, rule := range rules {
		receiver.ApplyRule(styleProperty, rule)
	}
}

/*TODO MAKE STYLE ENGINE ROOT TO HTML ELEMENT STYLE PROPERTY AND GIVE IT HERE FOR GLOBAL CSS VARIABLES*/
/*TODO MAKE STYLE PROPERTIES MAP FOR CSS VARIABLES AND GIVE HERE PARENT STYLE PROPERTY FOR APPLYING*/
func (receiver *StyleEngine) ApplyCssRules(currentWidget widget.WidgetInterface) {
	rules := receiver.GetAllCssRules(currentWidget.GetID(), currentWidget.GetClasses(), currentWidget.GetHtmlName())
	for _, rule := range rules {
		if rule.Function(currentWidget, rule) {
			receiver.ApplyRule(currentWidget.GetStyleProperty(), rule)
		}
	}
	//receiver.ApplyRules(currentWidget.GetStyleProperty(), rules)
	currentWidget.GetStyleProperty().ApplyInlineRules(currentWidget.GetStyleRules())
}

func (receiver *StyleEngine) ApplyRule(styleProperty *StyleProperty.StyleProperty, rule *CssRuleListItem.CssRuleListItem) {
	for property, value := range rule.Declarations {
		styleProperty.ApplyDeclaration(property, value)
	}
}
