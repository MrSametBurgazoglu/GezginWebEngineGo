package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByEveryElementAndAttributeAndContainsValue(attribute, value string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(&receiver.CssPropertiesByEveryElementAndAttributeAndContainValue, "", attribute, value)
	cssRuleListItem.Function = CssRuleListItem.IsElementAndAttributeAndContainsValue
	return
}

func (receiver *CssRuleList) GetCssRulesByEveryElementAndAttributeAndContainsValue(attribute, value string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch3(receiver.CssPropertiesByEveryElementAndAttributeAndContainValue, "", attribute, value)
}

func (receiver *StyleEngine) GetCssRulesByEveryElementAndAttributeAndContainsValue(external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			ruleList = append(ruleList, sheet.cssRuleList.CssPropertiesByEveryElementAndAttributeAndContainValue...)
		}
	}
	return
}
