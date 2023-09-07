package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssRulesByClassDescendant(class1, class2 string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByClassDescendantList, class1, class2, "")
	cssRuleListItem.Function = CssRuleListItem.IsClassDescendant
	return
}

func (receiver *CssRuleList) GetCssRulesByClassDescendant(class1, class2 string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByClassDescendantList, class1, class2, "")
}
