package StyleProperty

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/utils"
)

const AlignStringCount = 8

var alignStrings = []string{
	"baseline",
	"center",
	"flex-end",
	"flex-start",
	"space-around",
	"space-between",
	"space-evenly",
	"stretch",
}

func setAlignContent(cssProperties *StyleProperty, value string) {
	index := utils.IndexFounder(alignStrings, value, AlignStringCount)
	if index != -1 {
		cssProperties.AlignContent = enums.CssAlignType(index)
	} else {
		cssProperties.AlignContent = enums.CSS_ALIGN_STRETCH
	}
}

func setAlignItems(cssProperties *StyleProperty, value string) {
	index := utils.IndexFounder(alignStrings, value, AlignStringCount)
	if index != -1 {
		cssProperties.AlignItems = enums.CssAlignType(index)
	} else {
		cssProperties.AlignItems = enums.CSS_ALIGN_STRETCH
	}
}

func setAlignSelf(cssProperties *StyleProperty, value string) {
	index := utils.IndexFounder(alignStrings, value, AlignStringCount)
	if index != -1 {
		cssProperties.AlignSelf = enums.CssAlignType(index)
	} else {
		cssProperties.AlignSelf = enums.CSS_ALIGN_STRETCH
	}
}

func AlignContentPropertySetValue(cssProperties *StyleProperty, value string) {
	if value == "inherit" {
		cssProperties.AlignContentInherit = true
	} else if value == "initial" {
		cssProperties.AlignContent = enums.CSS_ALIGN_STRETCH
	} else {
		setAlignContent(cssProperties, value)
	}
}

func AlignItemsPropertySetValue(cssProperties *StyleProperty, value string) {
	if value == "inherit" {
		cssProperties.AlignItemsInherit = true
	} else if value == "initial" {
		cssProperties.AlignItems = enums.CSS_ALIGN_STRETCH
	} else {
		setAlignItems(cssProperties, value)
	}
}

func AlignSelfPropertySetValue(cssProperties *StyleProperty, value string) {
	if value == "inherit" {
		cssProperties.AlignSelfInherit = true
	} else if value == "initial" {
		cssProperties.AlignSelf = enums.CSS_ALIGN_STRETCH
	} else {
		setAlignSelf(cssProperties, value)
	}
}

func ComputeInheritAlign(properties *StyleProperty, source *StyleProperty) {
	if properties.AlignContentInherit {
		properties.AlignContent = source.AlignContent
	}
	if properties.AlignItemsInherit {
		properties.AlignItems = source.AlignItems
	}
	if properties.AlignSelfInherit {
		properties.AlignSelf = source.AlignSelf
	}
}

func updateAlign(properties *StyleProperty, source *StyleProperty) {
	if source.AlignContentInherit {
		properties.AlignContentInherit = true
	} else if source.AlignContent != enums.CSS_ALIGN_EMPTY {
		properties.AlignContent = source.AlignContent
	}
	if source.AlignItemsInherit {
		properties.AlignItemsInherit = true
	} else if source.AlignItems != enums.CSS_ALIGN_EMPTY {
		properties.AlignItems = source.AlignItems
	}
	if source.AlignSelfInherit {
		properties.AlignSelfInherit = true
	} else if source.AlignSelf != enums.CSS_ALIGN_EMPTY {
		properties.AlignSelf = source.AlignSelf
	}
}
