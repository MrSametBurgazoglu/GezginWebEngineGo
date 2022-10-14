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
var CssPropertiesByElementList = make(map[string]*structs.CssProperties)
var CssStyleTagList []*htmlVariables.Widget
var CssStyleLinkList []string

func CreateNewCssPropertiesByID(id string) (newCssProperties *structs.CssProperties) {
	newCssProperties = new(structs.CssProperties)
	cssPropertyItem := &CssPropertyListItem{Identifier1: id, CssProperties: newCssProperties}
	CssPropertiesByIDList = append(CssPropertiesByIDList, cssPropertyItem)
	return
}

func CreateNewCssPropertiesByClass(id string) (newCssProperties *structs.CssProperties) {
	newCssProperties = new(structs.CssProperties)
	cssPropertyItem := &CssPropertyListItem{Identifier1: id, CssProperties: newCssProperties}
	CssPropertiesByClassList = append(CssPropertiesByClassList, cssPropertyItem)
	return
}

func CreateNewCssPropertiesByElement(element string) (newCssProperties *structs.CssProperties) {
	newCssProperties = new(structs.CssProperties)
	CssPropertiesByElementList[element] = newCssProperties
	return
}

func CreateNewCssPropertiesByElementAndClass(id string) (newCssProperties *structs.CssProperties) {
	cssPropertyItem := &CssPropertyListItem{Identifier1: id, CssProperties: newCssProperties}
	CssPropertiesByIDList = append(CssPropertiesByIDList, cssPropertyItem)
	return
}

func GetCssPropertiesByID(id string) *structs.CssProperties {
	for _, item := range CssPropertiesByIDList {
		if item.Identifier1 == id {
			return item.CssProperties
		}
	}
	return nil
}

func GetCssPropertiesByClass(class string) *structs.CssProperties {
	for _, item := range CssPropertiesByIDList {
		if item.Identifier1 == class {
			return item.CssProperties
		}
	}
	return nil
}

func GetCssPropertiesByElement(element string) *structs.CssProperties {
	value, ok := CssPropertiesByElementList[element]
	if ok {
		return value
	}
	return nil
}

func GetCssPropertiesByElementAndClass(class string, element htmlVariables.HtmlTags) *structs.CssProperties {
	for _, item := range CssPropertiesByIDList {
		if item.Identifier1 == class && item.Identifier2 == element {
			return item.CssProperties
		}
	}
	return nil
}
