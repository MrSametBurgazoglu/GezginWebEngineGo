package StyleEngine

import "gezgin_web_engine/widget"

const HtmlTagCount = 105

/*
.class	.intro	Selects all elements with class="intro"
WE ALREADY CAN DO THIS BY SEARCHING WITH CLASS
.class1.class2	.name1.name2	Selects all elements with both name1 and name2 set within its class attribute
WE CAN DO THIS BY SEARCHING CLASS BY FIRST ONE AND AFTER FIND IT WE CAN LOOK TO OTHER CLASS IS EXIST
.class1 .class2	.name1 .name2	Selects all elements with name2 that is a descendant of an element with name1
WE CAN DO THIS BY SEARCHING NAME 2 AND LOOK PARENTS FOR NAME 1
#id	#firstname	Selects the element with id="firstname"
WE CAN ALREADY DO THIS BY SEARCH WITH ID
*	*	Selects all elements
element	p	Selects all <p> elements
WE CAN ALREADY DO THIS BY SEARCH WITH TAG
element.class	p.intro	Selects all <p> elements with class="intro"
WE CAN ALREADY DO THIS BY SEARCH WITH ID AND CLASS TOGETHER
element,element	div, p	Selects all <div> elements and all <p> elements
WE CAN ALREADY DO THIS BY SEARCH WITH TAG
element element	div p	Selects all <p> elements inside <div> elements
WE CAN DO THIS BY SEARCHING ELEMENT P THEN LOOKING FOR PARENTS
element>element	div > p	Selects all <p> elements where the parent is a <div> element
WE CAN DO THIS BY SEARCHING ELEMENT P THEN LOOK FOR PARENT
element+element	div + p	Selects the first <p> element that is placed immediately after <div> elements
WE CAN DO THIS BY SEARCHING ELEMENT P THEN LOOK FOR PARENTS CHILDREN
element1~element2	p ~ ul	Selects every <ul> element that is preceded by a <p> element
WE CAN DO THIS BY SEARCHING ELEMENT P THEN LOOK FOR PARENTS CHILDREN
[attribute]	[target]	Selects all elements with a target attribute
WE CAN DO THIS ATTRIBUTE FILTERS BY SEARCHING ELEMENT AND LOOK ATTRIBUTE BY FUNCTION
[attribute=value]	[target="_blank"]	Selects all elements with target="_blank"
[attribute~=value]	[title~="flower"]	Selects all elements with a title attribute containing the word "flower"
[attribute|=value]	[lang|="en"]	Selects all elements with a lang attribute value equal to "en" or starting with "en-"
[attribute^=value]	a[href^="https"]	Selects every <a> element whose href attribute value begins with "https"
[attribute$=value]	a[href$=".pdf"]	Selects every <a> element whose href attribute value ends with ".pdf"
[attribute*=value]	a[href*="w3schools"]	Selects every <a> element whose href attribute value contains the substring "w3schools"
:active	a:active	Selects the active link

//START this ones is damn strange.
::after	p::after	Insert something after the content of each <p> element
::before	p::before	Insert something before the content of each <p> element
//END
WE CAN DO THIS FILTERS BY SEARCHING ELEMENT AND LOOK ATTRIBUTE BY FUNCTION
:checked	input:checked	Selects every checked <input> element
:default	input:default	Selects the default <input> element
:disabled	input:disabled	Selects every disabled <input> element
:empty	p:empty	Selects every <p> element that has no children (including text nodes)
:enabled	input:enabled	Selects every enabled <input> element
:first-child	p:first-child	Selects every <p> element that is the first child of its parent
::first-letter	p::first-letter	Selects the first letter of every <p> element
::first-line	p::first-line	Selects the first line of every <p> element
:first-of-type	p:first-of-type	Selects every <p> element that is the first <p> element of its parent
:focus	input:focus	Selects the input element which has focus
:fullscreen	:fullscreen	Selects the element that is in full-screen mode
:hover	a:hover	Selects links on mouse over
:in-range	input:in-range	Selects input elements with a value within a specified range
:indeterminate	input:indeterminate	Selects input elements that are in an indeterminate state
:invalid	input:invalid	Selects all input elements with an invalid value
:lang(language)	p:lang(it)	Selects every <p> element with a lang attribute equal to "it" (Italian)
:last-child	p:last-child	Selects every <p> element that is the last child of its parent
:last-of-type	p:last-of-type	Selects every <p> element that is the last <p> element of its parent
:link	a:link	Selects all unvisited links
::marker	::marker	Selects the markers of list items
:not(selector)	:not(p)	Selects every element that is not a <p> element
:nth-child(n)	p:nth-child(2)	Selects every <p> element that is the second child of its parent
:nth-last-child(n)	p:nth-last-child(2)	Selects every <p> element that is the second child of its parent, counting from the last child
:nth-last-of-type(n)	p:nth-last-of-type(2)	Selects every <p> element that is the second <p> element of its parent, counting from the last child
:nth-of-type(n)	p:nth-of-type(2)	Selects every <p> element that is the second <p> element of its parent
:only-of-type	p:only-of-type	Selects every <p> element that is the only <p> element of its parent
:only-child	p:only-child	Selects every <p> element that is the only child of its parent
:optional	input:optional	Selects input elements with no "required" attribute
:out-of-range	input:out-of-range	Selects input elements with a value outside a specified range
::placeholder	input::placeholder	Selects input elements with the "placeholder" attribute specified
:read-only	input:read-only	Selects input elements with the "readonly" attribute specified
:read-write	input:read-write	Selects input elements with the "readonly" attribute NOT specified
:required	input:required	Selects input elements with the "required" attribute specified
:root	:root	Selects the document's root element
::selection	::selection	Selects the portion of an element that is selected by a user
:target	#news:target	Selects the current active #news element (clicked on a URL containing that anchor name)
:valid	input:valid	Selects all input elements with a valid value
:visited	a:visited	Selects all visited links
*/

type CSS_RULE_TYPE uint16

const (
	CSS_RULE_TYPE_TAG_ONLY = iota
	CSS_RULE_TYPE_CLASS_ONLY
	CSS_RULE_TYPE_TAG_AND_CLASS
	CSS_RULE_TYPE_CLASS_BOTH
	CSS_RULE_TYPE_CLASS_DESCENDANT
	CSS_RULE_TYPE_TAG_DESCENDANT
	CSS_RULE_TYPE_TAG_DESCENDANT_FIRST
)

type CssRuleListItem struct {
	Identifier1  string
	Identifier2  string
	Declarations map[string]string
	function     func(widget *widget.Widget, item *CssRuleListItem) //we use this function for checking advanced css selectors
}

func (receiver *CssRuleListItem) Initialize() {
	receiver.Declarations = make(map[string]string)
}

type CssRuleList struct {
	CssPropertiesByIDList      []*CssRuleListItem
	CssPropertiesByClassList   []*CssRuleListItem
	CssPropertiesByElementList map[string]*CssRuleListItem //it can be only map
}

func (receiver *CssRuleList) Initialize() {
	receiver.CssPropertiesByElementList = make(map[string]*CssRuleListItem)
}

func (receiver *CssRuleList) CreateNewCssRulesByID(id string) (cssRuleListItem *CssRuleListItem) {
	cssRuleListItem = &CssRuleListItem{Identifier1: id}
	cssRuleListItem.Initialize()
	receiver.CssPropertiesByIDList = append(receiver.CssPropertiesByIDList, cssRuleListItem)
	return
}

func (receiver *CssRuleList) CreateNewCssRulesByClass(class string) (cssRuleListItem *CssRuleListItem) {
	cssRuleListItem = &CssRuleListItem{Identifier1: class}
	cssRuleListItem.Initialize()
	receiver.CssPropertiesByClassList = append(receiver.CssPropertiesByClassList, cssRuleListItem)
	return
}

func (receiver *CssRuleList) CreateNewCssRulesByClassDescendant(class1, class2 string) (cssRuleListItem *CssRuleListItem) {
	cssRuleListItem = &CssRuleListItem{Identifier1: class1, Identifier2: class2}
	cssRuleListItem.Initialize()
	receiver.CssPropertiesByClassList = append(receiver.CssPropertiesByClassList, cssRuleListItem)
	return
}

func (receiver *CssRuleList) CreateNewCssRulesByClassBoth(class1, class2 string) (cssRuleListItem *CssRuleListItem) {
	cssRuleListItem = &CssRuleListItem{Identifier1: class1, Identifier2: class2}
	cssRuleListItem.Initialize()
	receiver.CssPropertiesByClassList = append(receiver.CssPropertiesByClassList, cssRuleListItem)
	return
}

func (receiver *CssRuleList) CreateNewCssPropertiesByElement(tag string) (cssRuleListItem *CssRuleListItem) {
	receiver.CssPropertiesByElementList[tag] = new(CssRuleListItem)
	receiver.CssPropertiesByElementList[tag].Initialize()
	return receiver.CssPropertiesByElementList[tag]
}

func (receiver *CssRuleList) CreateNewCssPropertiesByElementAndClass(id string) (cssRuleListItem *CssRuleListItem) {
	cssRuleListItem = &CssRuleListItem{Identifier1: id}
	cssRuleListItem.Initialize()
	receiver.CssPropertiesByIDList = append(receiver.CssPropertiesByIDList, cssRuleListItem)
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

func (receiver *CssRuleList) GetCssRulesByClassDescendant(class1, class2 string) *CssRuleListItem {
	for _, item := range receiver.CssPropertiesByClassList {
		if item.Identifier1 == class1 && item.Identifier2 == class2 {
			return item
		}
	}
	return nil
}

func (receiver *CssRuleList) GetCssRulesByClassBoth(class1, class2 string) *CssRuleListItem {
	for _, item := range receiver.CssPropertiesByClassList {
		if item.Identifier1 == class1 && item.Identifier2 == class2 {
			return item
		}
	}
	return nil
}

func (receiver *CssRuleList) GetCssRulesByElement(element string) *CssRuleListItem {
	return receiver.CssPropertiesByElementList[element]
}

func (receiver *CssRuleList) GetRulesByElementAndClass(class, element string) *CssRuleListItem {
	for _, item := range receiver.CssPropertiesByIDList {
		if item.Identifier1 == class && item.Identifier2 == element {
			return item
		}
	}
	return nil
}
