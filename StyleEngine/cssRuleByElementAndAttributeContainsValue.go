package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementAndAttributeAndContainsValue(tag, attribute, value string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementAndAttributeAndContainValue, tag, attribute, value)
	cssRuleListItem.Function = CssRuleListItem.IsElementAndAttributeAndContainsValue
	return
}

func (receiver *CssRuleList) GetCssRulesByElementAndAttributeAndContainsValue(element, attribute, value string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementAndAttributeAndContainValue, element, attribute, value)
}

func (receiver *StyleEngine) GetCssRulesByElementAndAttributeAndContainsValue(element string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			ruleList = append(ruleList, sheet.cssRuleList.GenericSearch(sheet.cssRuleList.CssPropertiesByElementAndAttributeAndContainValue, element, "", ""))
		}
	}
	return
}
