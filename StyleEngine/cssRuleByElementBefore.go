package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementBefore(tag, beforeTag string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementBefore, tag, beforeTag, "")
	cssRuleListItem.Function = CssRuleListItem.IsElementBefore
	return
}

func (receiver *CssRuleList) GetCssRulesByElementBefore(element, beforeElement string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementBefore, element, beforeElement, "")
}

func (receiver *StyleEngine) GetCssRulesByElementBefore(element string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			ruleList = append(ruleList, sheet.cssRuleList.GenericSearch(sheet.cssRuleList.CssPropertiesByElementBefore, element, "", ""))
		}
	}
	return
}
