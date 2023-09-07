package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssRulesByClassBoth(class1, class2 string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByClassBothList, class1, class2, "")
	cssRuleListItem.Function = CssRuleListItem.IsBothClass
	return
}

func (receiver *CssRuleList) GetCssRulesByClassBoth(class1, class2 string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByClassBothList, class1, class2, "")
}

func (receiver *StyleEngine) GetCssRulesByClassBoth(class1 string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			ruleList = append(ruleList, sheet.cssRuleList.GenericSearch(sheet.cssRuleList.CssPropertiesByClassBothList, class1, "", ""))
		}
	}
	return
}
