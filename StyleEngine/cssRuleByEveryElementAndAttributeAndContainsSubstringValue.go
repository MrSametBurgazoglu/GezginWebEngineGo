package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByEveryElementAndAttributeAndContainsSubstringsValue(attribute, value string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByEveryElementAndAttributeAndContainsSubstringValue, "", attribute, value)
	cssRuleListItem.Function = CssRuleListItem.IsElementAndAttributeAndContainsSubstringValue
	return
}

func (receiver *CssRuleList) GetCssRulesByEveryElementAndAttributeAndContainsSubstringsValue(attribute, value string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByEveryElementAndAttributeAndContainsSubstringValue, "", attribute, value)
}

func (receiver *StyleEngine) GetCssRulesByEveryElementAndAttributeAndContainsSubstringsValue(external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			ruleList = append(ruleList, sheet.cssRuleList.CssPropertiesByEveryElementAndAttributeAndContainsSubstringValue...)
		}
	}
	return
}
