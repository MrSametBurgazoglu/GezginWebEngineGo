package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssRulesByClassDescendantAndFirst(class1, class2, class3 string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(&receiver.CssPropertiesByClassDescendantAndFirstClassList, class2, class1, class3)
	cssRuleListItem.Function = CssRuleListItem.IsClassDescendantAndFirst
	return
}

func (receiver *CssRuleList) GetCssRulesByClassDescendantAndFirst(class1, class2, class3 string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch3(receiver.CssPropertiesByClassDescendantAndFirstClassList, class2, class1, class3)
}

func (receiver *StyleEngine) GetCssRulesByClassDescendantAndFirst(class1 string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			result := sheet.cssRuleList.GenericSearch1(sheet.cssRuleList.CssPropertiesByClassDescendantAndFirstClassList, class1)
			if result != nil {
				ruleList = append(ruleList, result)
			}
		}
	}
	return
}
