package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementPreceded(tag, precededTag string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementPreceded, tag, precededTag, "")
	cssRuleListItem.Function = CssRuleListItem.IsElementPreceded
	return
}

func (receiver *CssRuleList) GetCssRulesByElementPreceded(element, precededElement string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementPreceded, element, precededElement, "")
}
