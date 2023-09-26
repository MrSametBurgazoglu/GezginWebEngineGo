package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementAndAttributeAndValue(tag, attribute, value string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(&receiver.CssPropertiesByElementAndAttributeAndValue, tag, attribute, value)
	cssRuleListItem.Function = CssRuleListItem.IsElementAndAttributeAndValue
	return
}

func (receiver *CssRuleList) GetCssRulesByElementAndAttributeAndValue(element, attribute, value string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch3(receiver.CssPropertiesByElementAndAttributeAndValue, element, attribute, value)
}

func (receiver *StyleEngine) GetCssRulesByElementAndAttributeAndValue(element string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			result := sheet.cssRuleList.GenericSearch3(sheet.cssRuleList.CssPropertiesByElementAndAttributeAndValue, element, "", "")
			if result != nil {
				ruleList = append(ruleList, result)
			}
		}
	}
	return
}