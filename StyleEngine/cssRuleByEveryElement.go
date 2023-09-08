package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) GetCssRulesByEveryElement() *CssRuleListItem.CssRuleListItem {
	return &receiver.CssPropertiesByEveryElement
}

func (receiver *StyleEngine) GetCssRulesByEveryElement(external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			ruleList = append(ruleList, &sheet.cssRuleList.CssPropertiesByEveryElement)
		}
	}
	return
}
