package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) GetCssRulesByEveryElement() *CssRuleListItem.CssRuleListItem {
	return &receiver.CssPropertiesByEveryElement
}
