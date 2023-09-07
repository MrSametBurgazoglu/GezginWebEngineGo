package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementDescendant(tag, descendantTag string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementDescendant, tag, descendantTag, "")
	cssRuleListItem.Function = CssRuleListItem.IsElementDescendant
	return
}

func (receiver *CssRuleList) GetCssRulesByElementDescendant(element, descendantElement string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementDescendant, element, descendantElement, "")
}
