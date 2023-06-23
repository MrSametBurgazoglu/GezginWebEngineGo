package StyleEngine

import "gezgin_web_engine/GlobalTypes"

type StyleRule struct {
	selector    string
	declaration map[string]string
}

func (receiver *StyleRule) SetSelector(newSelector string) {
	receiver.selector = newSelector
}

func (receiver *StyleRule) AddDeclaration(rule GlobalTypes.CssDeclaration) {
	receiver.declaration[rule.GetProperty()] = rule.GetValue()
}
