package CssRuleListItem

import "gezgin_web_engine/widget"

type CssRuleListItem struct {
	Identifier1  string
	Identifier2  string
	Identifier3  string
	Declarations map[string]string
	Function     func(widget widget.WidgetInterface, item *CssRuleListItem) bool //we use this Function for checking advanced css selectors
}

func (receiver *CssRuleListItem) Initialize() {
	receiver.Declarations = make(map[string]string)
}
