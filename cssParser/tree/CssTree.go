package tree

import (
	"gezgin_web_engine/cssParser/structs"
	"gezgin_web_engine/htmlParser"
	"gezgin_web_engine/htmlParser/widget"
)

type CssPropertyListItem struct {
	Identifier1   string
	Identifier2   htmlParser.HtmlTags
	CssProperties *structs.CssProperties
}

type CssStyleSheets struct {
	CssPropertiesByIDList      []*CssPropertyListItem
	CssPropertiesByClassList   []*CssPropertyListItem
	CssPropertiesByElementList [htmlParser.HtmlTagCount]*structs.CssProperties
	CssStyleTagList            []*widget.Widget
	CssStyleLinkList           []string
}

func (receiver *CssStyleSheets) CreateNewCssPropertiesByID(id string) (newCssProperties *structs.CssProperties) {
	newCssProperties = new(structs.CssProperties)
	cssPropertyItem := &CssPropertyListItem{Identifier1: id, CssProperties: newCssProperties}
	receiver.CssPropertiesByIDList = append(receiver.CssPropertiesByIDList, cssPropertyItem)
	return
}

func (receiver *CssStyleSheets) CreateNewCssPropertiesByClass(id string) (newCssProperties *structs.CssProperties) {
	newCssProperties = new(structs.CssProperties)
	cssPropertyItem := &CssPropertyListItem{Identifier1: id, CssProperties: newCssProperties}
	receiver.CssPropertiesByClassList = append(receiver.CssPropertiesByClassList, cssPropertyItem)
	return
}

func (receiver *CssStyleSheets) CreateNewCssPropertiesByElement(tags htmlParser.HtmlTags) (newCssProperties *structs.CssProperties) {
	newCssProperties = new(structs.CssProperties)
	receiver.CssPropertiesByElementList[tags] = newCssProperties
	return
}

func (receiver *CssStyleSheets) CreateNewCssPropertiesByElementAndClass(id string) (newCssProperties *structs.CssProperties) {
	cssPropertyItem := &CssPropertyListItem{Identifier1: id, CssProperties: newCssProperties}
	receiver.CssPropertiesByIDList = append(receiver.CssPropertiesByIDList, cssPropertyItem)
	return
}

func (receiver *CssStyleSheets) GetCssPropertiesByID(id string) *structs.CssProperties {
	for _, item := range receiver.CssPropertiesByIDList {
		if item.Identifier1 == id {
			return item.CssProperties
		}
	}
	return nil
}

func (receiver *CssStyleSheets) GetCssPropertiesByClass(class string) *structs.CssProperties {
	for _, item := range receiver.CssPropertiesByClassList {
		if item.Identifier1 == class {
			return item.CssProperties
		}
	}
	return nil
}

func (receiver *CssStyleSheets) GetCssPropertiesByElement(element htmlParser.HtmlTags) *structs.CssProperties {
	return receiver.CssPropertiesByElementList[element]
}

func (receiver *CssStyleSheets) GetCssPropertiesByElementAndClass(class string, element htmlParser.HtmlTags) *structs.CssProperties {
	for _, item := range CssPropertiesByIDList {
		if item.Identifier1 == class && item.Identifier2 == element {
			return item.CssProperties
		}
	}
	return nil
}

// CssPropertiesByIDList make this one binary tree
var CssPropertiesByIDList []*CssPropertyListItem
var CssPropertiesByClassList []*CssPropertyListItem
var CssPropertiesByElementList [htmlParser.HtmlTagCount]*structs.CssProperties
var CssStyleTagList []*widget.Widget
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

func CreateNewCssPropertiesByElement(tags htmlParser.HtmlTags) (newCssProperties *structs.CssProperties) {
	newCssProperties = new(structs.CssProperties)
	CssPropertiesByElementList[tags] = newCssProperties
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
	for _, item := range CssPropertiesByClassList {
		if item.Identifier1 == class {
			return item.CssProperties
		}
	}
	return nil
}

func GetCssPropertiesByElement(element htmlParser.HtmlTags) *structs.CssProperties {
	return CssPropertiesByElementList[element]
}

func GetCssPropertiesByElementAndClass(class string, element htmlParser.HtmlTags) *structs.CssProperties {
	for _, item := range CssPropertiesByIDList {
		if item.Identifier1 == class && item.Identifier2 == element {
			return item.CssProperties
		}
	}
	return nil
}
