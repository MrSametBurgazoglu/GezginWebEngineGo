package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByEveryElementAndAttribute(attribute string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByEveryElementAndAttribute, "", attribute, "")
	cssRuleListItem.Function = CssRuleListItem.IsAttributeExist
	return
}

func (receiver *CssRuleList) GetCssRulesByEveryElementAndAttribute(attribute string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementAndAttribute, "", attribute, "")
}

func (receiver *StyleEngine) GetCssRulesByEveryElementAndAttribute(external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			ruleList = append(ruleList, sheet.cssRuleList.CssPropertiesByEveryElementAndAttribute...)
		}
	}
	return
}
