package StyleEngine

import (
	"gezgin_web_engine/GlobalTypes"
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
	Root              *StyleProperty
}

func (receiver *StyleEngine) Initialize() {
	receiver.WorkerPool = workerpool.New(runtime.NumCPU() - 1)
	receiver.Root = new(StyleProperty)
	receiver.Root.Initialize()
}

func (receiver *StyleEngine) InitializeRoot() {
	receiver.Root.ApplyCssRules(receiver, ":root", nil, "root", map[string]string{})
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
				item.declarations[property] = value
			}
		}
	}
}

func (receiver *StyleEngine) GetCssRulesByClass(class string, external bool) (ruleList []*CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			rules := sheet.cssRuleList.GetCssRulesByClass(class)
			if rules != nil {
				ruleList = append(ruleList, rules)
			}
		}
	}
	return
}

func (receiver *StyleEngine) GetCssRulesByTag(htmlTag string, external bool) (ruleList []*CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			rules := sheet.cssRuleList.GetCssRulesByElement(htmlTag)
			if rules != nil {
				ruleList = append(ruleList, rules)
			}
		}
	}
	return
}

func (receiver *StyleEngine) GetCssRulesByID(id string, external bool) (ruleList []*CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			rules := sheet.cssRuleList.GetCssRulesByID(id)
			if rules != nil {
				ruleList = append(ruleList, rules)
			}
		}
	}
	return
}

// GetCssRuleListItems /* make this as one goroutine*/
func (receiver *StyleSheet) GetCssRuleListItems(selectors []string) (cssRuleList []*CssRuleListItem) {
	for _, s := range selectors {
		s = strings.TrimSpace(s)
		cssRuleList = append(cssRuleList, receiver.GetCssRuleItem(s))
	}
	return
}

func (receiver *StyleSheet) GetCssRuleItem(selector string) *CssRuleListItem {
	var cssRuleList *CssRuleListItem
	switch selector[0] {
	case '#':
		cssRuleList = receiver.cssRuleList.GetCssRulesByID(selector[0:])
		if cssRuleList == nil {
			cssRuleList = receiver.cssRuleList.CreateNewCssRulesByID(selector[0:])
		}
	case '.':
		secondDotIndex := strings.Index(selector[0:], ".")
		if secondDotIndex == -1 {
			cssRuleList = receiver.cssRuleList.GetCssRulesByClass(selector[1:])
			if cssRuleList == nil {
				cssRuleList = receiver.cssRuleList.CreateNewCssRulesByClass(selector[1:])
			}
		} else {
			if selector[secondDotIndex] == ' ' {
				firstClass := selector[1:secondDotIndex]
				secondClass := selector[secondDotIndex+1:]
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
