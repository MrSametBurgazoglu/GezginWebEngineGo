package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementAndAttribute(tag, attribute string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementAndAttribute, tag, attribute, "")
	cssRuleListItem.Function = CssRuleListItem.IsAttributeExist
	return
}

func (receiver *CssRuleList) GetCssRulesByElementAndAttribute(element, attribute string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementAndAttribute, element, attribute, "")
}

func (receiver *StyleEngine) GetCssRulesByElementAndAttribute(element string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			result := sheet.cssRuleList.GenericSearch(sheet.cssRuleList.CssPropertiesByElementAndAttribute, element, "", "")
			if result != nil {
				ruleList = append(ruleList, result)
			}
		}
	}
	return
}
