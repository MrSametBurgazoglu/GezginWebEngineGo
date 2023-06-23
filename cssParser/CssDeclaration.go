package cssParser

import "strings"

type CssDeclaration struct {
	cssProperty *CssProperty
	cssValue    *CssValue
}

func (receiver *CssDeclaration) Set(text string) {
	property, value, found := strings.Cut(text, ":")
	if found {
		receiver.cssProperty.Set(property)
		receiver.cssValue.Set(value)
	}
}

func (receiver *CssDeclaration) GetProperty() string {
	return receiver.cssProperty.property
}

func (receiver *CssDeclaration) GetValue() string {
	return receiver.cssValue.value
}
