package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementAndAttribute(tag, attribute string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementAndAttribute, tag, attribute, "")
	cssRuleListItem.Function = CssRuleListItem.IsElementAndAttribute
	return
}

func (receiver *CssRuleList) GetCssRulesByElementAndAttribute(element, attribute string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementAndAttribute, element, attribute, "")
}

func (receiver *StyleEngine) GetCssRulesByElementAndAttribute(element string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			ruleList = append(ruleList, sheet.cssRuleList.GenericSearch(sheet.cssRuleList.CssPropertiesByElementAndAttribute, element, "", ""))
		}
	}
	return
}
