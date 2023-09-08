package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementAndClass(tag, class string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(&receiver.CssPropertiesByElementAndClassList, tag, class, "")
	cssRuleListItem.Function = CssRuleListItem.IsClassExist
	return
}

func (receiver *CssRuleList) GetRulesByElementAndClass(element, class string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch3(receiver.CssPropertiesByElementAndClassList, element, class, "")
}

func (receiver *StyleEngine) GetCssRulesByElementAndClass(element string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			result := sheet.cssRuleList.GenericSearch3(sheet.cssRuleList.CssPropertiesByElementAndClassList, element, "", "")
			if result != nil {
				ruleList = append(ruleList, result)
			}
		}
	}
	return
}
