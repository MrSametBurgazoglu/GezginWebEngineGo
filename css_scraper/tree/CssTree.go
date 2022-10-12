package tree

import (
	"gezgin_web_engine/css_scraper/structs"
	"gezgin_web_engine/html_scraper/htmlVariables"
)

type CssPropertyListItem struct {
	Identifier1   string
	Identifier2   htmlVariables.HtmlTags
	CssProperties *structs.CssProperties
}

// CssPropertiesByIDList make this one binary tree
var CssPropertiesByIDList []*CssPropertyListItem
var CssPropertiesByClassList []*CssPropertyListItem
var CssPropertiesByElementList []*CssPropertyListItem
var CssStyleTagList []*htmlVariables.Widget
var CssStyleLinkList []string

func (receiver CssPropertyListItem) CreateNewCssPropertiesByID(id string) (newCssProperties *structs.CssProperties) {
	cssPropertyItem := &CssPropertyListItem{Identifier1: id, CssProperties: newCssProperties}
	CssPropertiesByIDList = append(CssPropertiesByIDList, cssPropertyItem)
	return
}

func (receiver CssPropertyListItem) CreateNewCssPropertiesByClass(id string) (newCssProperties *structs.CssProperties) {
	cssPropertyItem := &CssPropertyListItem{Identifier1: id, CssProperties: newCssProperties}
	CssPropertiesByClassList = append(CssPropertiesByClassList, cssPropertyItem)
	return
}

func (receiver CssPropertyListItem) CreateNewCssPropertiesByElement(id string) (newCssProperties *structs.CssProperties) {
	cssPropertyItem := &CssPropertyListItem{Identifier1: id, CssProperties: newCssProperties}
	CssPropertiesByElementList = append(CssPropertiesByElementList, cssPropertyItem)
	return
}

func (receiver CssPropertyListItem) CreateNewCssPropertiesByElementAndClass(id string) (newCssProperties *structs.CssProperties) {
	cssPropertyItem := &CssPropertyListItem{Identifier1: id, CssProperties: newCssProperties}
	CssPropertiesByIDList = append(CssPropertiesByIDList, cssPropertyItem)
	return
}

func (receiver CssPropertyListItem) GetCssPropertiesByID(id string) *structs.CssProperties {
	for _, item := range CssPropertiesByIDList {
		if item.Identifier1 == id {
			return item.CssProperties
		}
	}
	return nil
}

func (receiver CssPropertyListItem) GetCssPropertiesByClass(class string) *structs.CssProperties {
	for _, item := range CssPropertiesByIDList {
		if item.Identifier1 == class {
			return item.CssProperties
		}
	}
	return nil
}

func (receiver CssPropertyListItem) GetCssPropertiesByElement(element htmlVariables.HtmlTags) *structs.CssProperties {
	for _, item := range CssPropertiesByIDList {
		if item.Identifier2 == element {
			return item.CssProperties
		}
	}
	return nil
}

func (receiver CssPropertyListItem) GetCssPropertiesByElementAndClass(class string, element htmlVariables.HtmlTags) *structs.CssProperties {
	for _, item := range CssPropertiesByIDList {
		if item.Identifier1 == class && item.Identifier2 == element {
			return item.CssProperties
		}
	}
	return nil
}
