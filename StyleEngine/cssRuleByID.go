package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssRulesByID(id string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = &CssRuleListItem.CssRuleListItem{Identifier1: id}
	cssRuleListItem.Initialize()
	cssRuleListItem.Function = DefaultValidator
	receiver.CssPropertiesByIDList = append(receiver.CssPropertiesByIDList, cssRuleListItem)
	return
}

func (receiver *CssRuleList) GetCssRulesByID(id string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	for _, item := range receiver.CssPropertiesByIDList {
		if item.Identifier1 == id {
			return item
		}
	}
	return nil
}

func (receiver *StyleEngine) GetCssRulesByID(id string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			rules := sheet.cssRuleList.GetCssRulesByID(id)
			if rules != nil {
				ruleList = append(ruleList, rules)
			}
		}
	}
	return
}