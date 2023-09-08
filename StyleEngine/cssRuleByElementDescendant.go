package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementDescendant(tag, descendantTag string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementDescendant, tag, descendantTag, "")
	cssRuleListItem.Function = CssRuleListItem.IsElementDescendant
	return
}

func (receiver *CssRuleList) GetCssRulesByElementDescendant(element, descendantElement string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementDescendant, element, descendantElement, "")
}

func (receiver *StyleEngine) GetCssRulesByElementDescendant(element string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			result := sheet.cssRuleList.GenericSearch(sheet.cssRuleList.CssPropertiesByElementDescendant, element, "", "")
			if result != nil {
				ruleList = append(ruleList, result)
			}
		}
	}
	return
}
