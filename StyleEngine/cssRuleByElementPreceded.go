package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementPreceded(tag, precededTag string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementPreceded, tag, precededTag, "")
	cssRuleListItem.Function = CssRuleListItem.IsElementPreceded
	return
}

func (receiver *CssRuleList) GetCssRulesByElementPreceded(element, precededElement string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementPreceded, element, precededElement, "")
}

func (receiver *StyleEngine) GetCssRulesByElementPreceded(element string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			result := sheet.cssRuleList.GenericSearch(sheet.cssRuleList.CssPropertiesByElementPreceded, element, "", "")
			if result != nil {
				ruleList = append(ruleList, result)
			}
		}
	}
	return
}
