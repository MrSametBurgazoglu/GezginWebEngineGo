package StyleEngine

import "gezgin_web_engine/HtmlParser"

const HtmlTagCount = 105

type CssRuleListItem struct {
	Identifier1  string
	Identifier2  HtmlParser.HtmlTags
	declarations map[string]string
}

/*we are going to store rules here with identifiers but unique values*/

type CssRuleList struct {
	CssPropertiesByIDList      []*CssRuleListItem
	CssPropertiesByClassList   []*CssRuleListItem
	CssPropertiesByElementList [HtmlTagCount]CssRuleListItem //it can be only map
}

func (receiver *CssRuleList) CreateNewCssRulesByID(id string) (cssRuleListItem *CssRuleListItem) {
	cssRuleListItem = &CssRuleListItem{Identifier1: id}
	receiver.CssPropertiesByIDList = append(receiver.CssPropertiesByIDList, cssRuleListItem)
	return
}

func (receiver *CssRuleList) CreateNewCssRulesByClass(class string) (cssRuleListItem *CssRuleListItem) {
	cssRuleListItem = &CssRuleListItem{Identifier1: class}
	receiver.CssPropertiesByClassList = append(receiver.CssPropertiesByClassList, cssRuleListItem)
	return
}

func (receiver *CssRuleList) CreateNewCssPropertiesByElement(tags HtmlParser.HtmlTags) (cssRuleListItem *CssRuleListItem) {
	return &receiver.CssPropertiesByElementList[tags]
}

func (receiver *CssRuleList) CreateNewCssPropertiesByElementAndClass(id string) (cssRuleListItem *CssRuleListItem) {
	cssPropertyItem := &CssRuleListItem{Identifier1: id}
	receiver.CssPropertiesByIDList = append(receiver.CssPropertiesByIDList, cssPropertyItem)
	return
}

func (receiver *CssRuleList) GetCssRulesByID(id string) (cssRuleListItem *CssRuleListItem) {
	for _, item := range receiver.CssPropertiesByIDList {
		if item.Identifier1 == id {
			return item
		}
	}
	return nil
}

func (receiver *CssRuleList) GetCssRulesByClass(class string) *CssRuleListItem {
	for _, item := range receiver.CssPropertiesByClassList {
		if item.Identifier1 == class {
			return item
		}
	}
	return nil
}

func (receiver *CssRuleList) GetCssRulesByElement(element HtmlParser.HtmlTags) *CssRuleListItem {
	return &receiver.CssPropertiesByElementList[element]
}

func (receiver *CssRuleList) GetRulesByElementAndClass(class string, element HtmlParser.HtmlTags) *CssRuleListItem {
	for _, item := range receiver.CssPropertiesByIDList {
		if item.Identifier1 == class && item.Identifier2 == element {
			return item
		}
	}
	return nil
}
