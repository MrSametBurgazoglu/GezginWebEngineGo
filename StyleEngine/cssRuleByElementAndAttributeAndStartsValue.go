package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementAndAttributeAndStartsValue(tag, attribute, value string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementAndAttributeAndStartsValue, tag, attribute, value)
	cssRuleListItem.Function = CssRuleListItem.IsElementAndAttributeAndStartsValue
	return
}

func (receiver *CssRuleList) GetCssRulesByElementAndAttributeAndStartsValue(element, attribute, value string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementAndAttributeAndStartsValue, element, attribute, value)
}

func (receiver *StyleEngine) GetCssRulesByElementAndAttributeAndStartsValue(element string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			result := sheet.cssRuleList.GenericSearch(sheet.cssRuleList.CssPropertiesByElementAndAttributeAndStartsValue, element, "", "")
			if result != nil {
				ruleList = append(ruleList, result)
			}
		}
	}
	return
}
