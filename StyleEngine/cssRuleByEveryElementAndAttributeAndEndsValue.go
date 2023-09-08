package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByEveryElementAndAttributeAndEndsValue(attribute, value string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(&receiver.CssPropertiesByEveryElementAndAttributeAndEndsValue, "", attribute, value)
	cssRuleListItem.Function = CssRuleListItem.IsElementAndAttributeAndEndsValue
	return
}

func (receiver *CssRuleList) GetCssRulesByEveryElementAndAttributeAndEndsValue(attribute, value string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch3(receiver.CssPropertiesByEveryElementAndAttributeAndEndsValue, "", attribute, value)
}

func (receiver *StyleEngine) GetCssRulesByEveryElementAndAttributeAndEndsValue(external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			ruleList = append(ruleList, sheet.cssRuleList.CssPropertiesByEveryElementAndAttributeAndEndsValue...)
		}
	}
	return
}
