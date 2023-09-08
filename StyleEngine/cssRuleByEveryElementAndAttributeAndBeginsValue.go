package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByEveryElementAndAttributeAndBeginsValue(attribute, value string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByEveryElementAndAttributeAndBeginsValue, "", attribute, value)
	cssRuleListItem.Function = CssRuleListItem.IsElementAndAttributeAndBeginsValue
	return
}

func (receiver *CssRuleList) GetCssRulesByEveryElementAndAttributeAndBeginsValue(attribute, value string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByEveryElementAndAttributeAndBeginsValue, "", attribute, value)
}

func (receiver *StyleEngine) GetCssRulesByEveryElementAndAttributeAndBeginsValue(external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			ruleList = append(ruleList, sheet.cssRuleList.CssPropertiesByEveryElementAndAttributeAndBeginsValue...)
		}
	}
	return
}
