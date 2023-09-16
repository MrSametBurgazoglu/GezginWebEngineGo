package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssRulesByClass(class string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	if class == "cover-container" || class == "mx-auto" {
		println("hey")
	}
	cssRuleListItem = &CssRuleListItem.CssRuleListItem{Identifier1: class}
	cssRuleListItem.Initialize()
	cssRuleListItem.Function = DefaultValidator
	receiver.CssPropertiesByClassList = append(receiver.CssPropertiesByClassList, cssRuleListItem)
	return
}

func (receiver *CssRuleList) GetCssRulesByClass(class string) *CssRuleListItem.CssRuleListItem {
	for _, item := range receiver.CssPropertiesByClassList {
		if item.Identifier1 == class {
			return item
		}
	}
	return nil
}

func (receiver *StyleEngine) GetCssRulesByClass(class string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
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
