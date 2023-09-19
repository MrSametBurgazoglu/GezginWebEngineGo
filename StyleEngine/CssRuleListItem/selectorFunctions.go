package CssRuleListItem

import (
	"gezgin_web_engine/widget"
	"strings"
)

func IsClassExist(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	for currentWidget.GetParent() != nil {
		for _, class := range currentWidget.GetClasses() {
			if item.Identifier2 == class {
				return true
			}
		}
	}
	return false
}

func IsClassDescendant(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	for currentWidget.GetParent() != nil {
		for _, class := range currentWidget.GetClasses() {
			if item.Identifier2 == class {
				return true
			}
		}
		currentWidget = currentWidget.GetParent()
	}
	return false
}

func IsClassNot(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	for _, class := range currentWidget.GetClasses() {
		if item.Identifier2 == class {
			return false
		}
	}
	return true
}

func IsClassDescendantAndFirst(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	if currentWidget.GetChildrenIndex() == 0 || currentWidget.GetParent() == nil {
		return false
	}
	beforeCurrentWidget := currentWidget.GetParent().GetChildrenByIndex(currentWidget.GetChildrenIndex() - 1) /*THIS CHECK PROBABLY MUST BE FIRST ONE CAUSE ITS MORE EASY AND UNIQUE*/
	for _, class2 := range beforeCurrentWidget.GetClasses() {
		if item.Identifier3 != class2 {
			continue
		}
		currentParent := currentWidget.GetParent()
		for currentParent != nil {
			for _, class := range currentParent.GetClasses() {
				if item.Identifier2 == class {
					return true
				}
			}
			currentParent = currentParent.GetParent()
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

func IsAttributeExist(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
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

func IsElementAndAttributeAndBeginsValue(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	for attribute, value := range currentWidget.GetAttributes() {
		if item.Identifier2 == attribute && item.Identifier3 == value || strings.HasPrefix(value, item.Identifier3+"-") {
			return true
		}
	}
	return false
}

func IsElementAndAttributeAndStartsValue(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	for attribute, value := range currentWidget.GetAttributes() {
		if item.Identifier2 == attribute && strings.HasPrefix(value, item.Identifier3) {
			return true
		}
	}
	return false
}

func IsElementAndAttributeAndEndsValue(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	for attribute, value := range currentWidget.GetAttributes() {
		if item.Identifier2 == attribute && strings.HasSuffix(value, item.Identifier3) {
			return true
		}
	}
	return false
}

func IsElementAndAttributeAndContainsSubstringValue(currentWidget widget.WidgetInterface, item *CssRuleListItem) bool {
	for attribute, value := range currentWidget.GetAttributes() {
		if item.Identifier2 == attribute {
			for _, s := range strings.Fields(value) {
				if s == item.Identifier3 {
					return true
				}
			}
		}
	}
	return false
}
