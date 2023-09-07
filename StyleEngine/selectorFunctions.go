package StyleEngine

import (
	"gezgin_web_engine/widget"
	"strings"
)

func IsClassDescendant(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	for currentWidget.GetParent() != nil {
		for _, class := range currentWidget.GetClasses() {
			if item.Identifier2 == class {
				return true
			}
		}
	}
	return false
}

func IsBothClass(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	for _, class := range currentWidget.GetClasses() {
		if item.Identifier2 == class {
			return true
		}
	}
	return false
}

func IsElementDescendant(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	for currentWidget.GetParent() != nil {
		if item.Identifier2 == currentWidget.GetHtmlName() {
			return true
		}
	}
	return false
}

func IsElementBefore(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	if currentWidget.GetChildrenIndex() == 0 {
		return false
	} else {
		if currentWidget.GetParent().GetChildrenByIndex(currentWidget.GetChildrenIndex()-1).GetHtmlName() == item.Identifier2 {
			return true
		}
		return false
	}
}

func IsElementPreceded(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	if currentWidget.GetChildrenIndex() == currentWidget.GetParent().GetChildrenCount()-1 {
		return false
	} else {
		if currentWidget.GetParent().GetChildrenByIndex(currentWidget.GetChildrenIndex()+1).GetHtmlName() == item.Identifier2 {
			return true
		}
		return false
	}
}

func IsElementAndAttribute(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	for attribute, _ := range currentWidget.GetAttributes() {
		if item.Identifier2 == attribute {
			return true
		}
	}
	return false
}

func IsElementAndAttributeAndValue(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	for attribute, value := range currentWidget.GetAttributes() {
		if item.Identifier2 == attribute && value == item.Identifier3 {
			return true
		}
	}
	return false
}

func IsElementAndAttributeAndContainsValue(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	for attribute, value := range currentWidget.GetAttributes() {
		if item.Identifier2 == attribute && strings.Contains(value, item.Identifier3) {
			return true
		}
	}
	return false
}
