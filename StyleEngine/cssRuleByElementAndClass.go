package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementAndClass(tag, class string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementAndClassList, tag, class, "")
	cssRuleListItem.Function = DefaultValidator
	return
}

func (receiver *CssRuleList) GetRulesByElementAndClass(element, class string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementAndClassList, element, class, "")
}
