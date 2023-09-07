package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssPropertiesByElement(tag string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	receiver.CssPropertiesByElementList[tag] = new(CssRuleListItem.CssRuleListItem)
	receiver.CssPropertiesByElementList[tag].Initialize()
	receiver.CssPropertiesByElementList[tag].Function = DefaultValidator
	return receiver.CssPropertiesByElementList[tag]
}

func (receiver *CssRuleList) GetCssRulesByElement(element string) *CssRuleListItem.CssRuleListItem {
	return receiver.CssPropertiesByElementList[element]
}

func (receiver *StyleEngine) GetCssRulesByTag(htmlTag string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			rules := sheet.cssRuleList.GetCssRulesByElement(htmlTag)
			if rules != nil {
				ruleList = append(ruleList, rules)
			}
		}
	}
	return
}
