package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementAndAttributeAndContainsSubstringsValue(tag, attribute, value string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementAndAttributeAndContainsSubstringValue, tag, attribute, value)
	cssRuleListItem.Function = CssRuleListItem.IsElementAndAttributeAndContainsSubstringValue
	return
}

func (receiver *CssRuleList) GetCssRulesByElementAndAttributeAndContainsSubstringsValue(element, attribute, value string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementAndAttributeAndContainsSubstringValue, element, attribute, value)
}

func (receiver *StyleEngine) GetCssRulesByElementAndAttributeAndContainsSubstringsValue(element string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			result := sheet.cssRuleList.GenericSearch(sheet.cssRuleList.CssPropertiesByElementAndAttributeAndContainsSubstringValue, element, "", "")
			if result != nil {
				ruleList = append(ruleList, result)
			}
		}
	}
	return
}
