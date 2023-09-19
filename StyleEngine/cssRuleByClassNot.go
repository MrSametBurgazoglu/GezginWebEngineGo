package StyleEngine

import "gezgin_web_engine/StyleEngine/CssRuleListItem"

func (receiver *CssRuleList) CreateNewCssRulesByClassNot(class1, class2 string) (cssRuleListItem *CssRuleListItem.CssRuleListItem) {
	cssRuleListItem = receiver.GenericCreate(&receiver.CssPropertiesByClassNotList, class1, class2, "")
	cssRuleListItem.Function = CssRuleListItem.IsClassNot
	return
}

func (receiver *CssRuleList) GetCssRulesByClassNot(class1, class2 string) *CssRuleListItem.CssRuleListItem {
	return receiver.GenericSearch3(receiver.CssPropertiesByClassNotList, class1, class2, "")
}

/*THIS GETTER WRONG CAUSE WE DON'T KNOW THE SECOND CLASS*/
func (receiver *StyleEngine) GetCssRulesByClassNot(class1 string, external bool) (ruleList []*CssRuleListItem.CssRuleListItem) {
	for _, sheet := range receiver.CssStyleSheetList {
		if sheet.external == external {
			result := sheet.cssRuleList.GenericSearch1(sheet.cssRuleList.CssPropertiesByClassNotList, class1)
			if result != nil {
				ruleList = append(ruleList, result)
			}
		}
	}
	return
}
