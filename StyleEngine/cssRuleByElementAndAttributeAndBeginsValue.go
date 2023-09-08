package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementAndAttributeAndBeginsValue(tag, attribute, value string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementAndAttributeAndBeginsValue, tag, attribute, value)
	cssRuleListItem.Function = CssRuleListItem.IsElementAndAttributeAndBeginsValue
	return
}

func (receiver *CssRuleList) GetCssRulesByElementAndAttributeAndBeginsValue(element, attribute, value string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementAndAttributeAndBeginsValue, element, attribute, value)
}

func (receiver *StyleEngine) GetCssRulesByElementAndAttributeAndBeginsValue(element string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			result := sheet.cssRuleList.GenericSearch(sheet.cssRuleList.CssPropertiesByElementAndAttributeAndBeginsValue, element, "", "")
			if result != nil {
				ruleList = append(ruleList, result)
			}
		}
	}
	return
}
