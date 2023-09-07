package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByEveryElementAndAttributeAndContainsValue(attribute, value string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByEveryElementAndAttributeAndContainValue, "", attribute, value)
	cssRuleListItem.Function = CssRuleListItem.IsElementAndAttributeAndContainsValue
	return
}

func (receiver *CssRuleList) GetCssRulesByElementEveryAndAttributeAndContainsValue(attribute, value string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByEveryElementAndAttributeAndContainValue, "", attribute, value)
}
