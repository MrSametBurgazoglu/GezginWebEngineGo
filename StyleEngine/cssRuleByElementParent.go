package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementParent(tag, parentTag string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(&receiver.CssPropertiesByElementParent, tag, parentTag, "")
	cssRuleListItem.Function = CssRuleListItem.IsElementDescendant
	return
}

func (receiver *CssRuleList) GetCssRulesByElementParent(element, parentElement string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch3(receiver.CssPropertiesByElementParent, element, parentElement, "")
}

func (receiver *StyleEngine) GetCssRulesByElementParent(element string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			result := sheet.cssRuleList.GenericSearch3(sheet.cssRuleList.CssPropertiesByElementParent, element, "", "")
			if result != nil {
				ruleList = append(ruleList, result)
			}
		}
	}
	return
}
