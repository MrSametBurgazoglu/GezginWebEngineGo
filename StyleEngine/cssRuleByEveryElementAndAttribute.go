package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByEveryElementAndAttribute(attribute string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByEveryElementAndAttribute, "", attribute, "")
	cssRuleListItem.Function = CssRuleListItem.IsElementAndAttribute
	return
}

func (receiver *CssRuleList) GetCssRulesByEveryElementAndAttribute(attribute string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementAndAttribute, "", attribute, "")
}
