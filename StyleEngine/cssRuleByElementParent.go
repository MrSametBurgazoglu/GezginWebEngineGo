package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementParent(tag, parentTag string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementParent, tag, parentTag, "")
	cssRuleListItem.Function = CssRuleListItem.IsElementDescendant
	return
}

func (receiver *CssRuleList) GetCssRulesByElementParent(element, parentElement string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementParent, element, parentElement, "")
}
