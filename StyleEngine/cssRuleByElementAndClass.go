package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElementAndClass(tag, class string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(receiver.CssPropertiesByElementAndClassList, tag, class, "")
	cssRuleListItem.Function = DefaultValidator
	return
}

func (receiver *CssRuleList) GetRulesByElementAndClass(element, class string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch(receiver.CssPropertiesByElementAndClassList, element, class, "")
}

func (receiver *StyleEngine) GetCssRulesByElementAndClass(element string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			ruleList = append(ruleList, sheet.cssRuleList.GenericSearch(sheet.cssRuleList.CssPropertiesByElementAndClassList, element, "", ""))
		}
	}
	return
}
