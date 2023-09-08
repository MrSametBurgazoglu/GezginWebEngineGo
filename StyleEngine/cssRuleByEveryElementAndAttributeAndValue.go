package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByEveryElementAndAttributeAndValue(attribute, value string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByEveryElementAndAttributeAndValue, "", attribute, value)
	cssRuleListItem.Function = CssRuleListItem.IsElementAndAttributeAndValue
	return
}

func (receiver *CssRuleList) GetCssRulesByEveryElementAndAttributeAndValue(attribute, value string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByEveryElementAndAttributeAndValue, "", attribute, value)
}

func (receiver *StyleEngine) GetCssRulesByEveryElementAndAttributeAndValue(external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			ruleList = append(ruleList, sheet.cssRuleList.CssPropertiesByEveryElementAndAttributeAndValue...)
		}
	}
	return
}
