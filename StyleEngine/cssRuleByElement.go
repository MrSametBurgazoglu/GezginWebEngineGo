package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElement(tag string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	receiver.CssPropertiesByElementList[tag] = new(CssRuleListItem.CssRuleListItem)
	receiver.CssPropertiesByElementList[tag].Initialize()
	receiver.CssPropertiesByElementList[tag].Function = DefaultValidator
	return receiver.CssPropertiesByElementList[tag]
}

func (receiver *CssRuleList) GetCssRulesByElement(element string) *CssRuleListItem.CssRuleListItem {
	return receiver.CssPropertiesByElementList[element]
}
