package CssParser

import (
	"gezgin_web_engine/GlobalTypes"
	"gezgin_web_engine/utils"
	"strings"
	"sync"
)

type CssParser struct {
	wg sync.WaitGroup
}

type Result struct {
	ruleCount          int
	CssStyleSheetRules []GlobalTypes.CssRuleInterface
}

func (receiver *Result) GetRuleCount() int {
	return receiver.ruleCount
}

func (receiver *Result) GetRuleByIndex(index int) GlobalTypes.CssRuleInterface {
	return receiver.CssStyleSheetRules[index]
}

type StyleElement interface {
}

func (receiver *CssParser) ParseCssFromStyleTag(styleElement StyleElement, styleText string) (result *Result) {
	result = new(Result)
	newCssStyleSheet := new(CssStyleSheet)
	styleText = utils.RemoveCharsFromString(styleText)
	seek := 0
	index := 0
	index = strings.Index(styleText[seek:], "{")
	for index != -1 { //maybe go routine for every cssText
		newCssRule := new(CssRule)
		index2 := strings.Index(styleText[seek:], "}")
		selectors := styleText[seek : seek+index]
		cssText := styleText[seek+index+1 : seek+index2]
		seek += index2 + 1
		newCssRule.SetStyleSheet(newCssStyleSheet)
		newCssRule.SetCssSelectors(selectors)
		newCssRule.SetCssDeclarationBlock(cssText)
		result.CssStyleSheetRules = append(result.CssStyleSheetRules, newCssRule)
		index = strings.Index(styleText[seek:], "{")
	}
	return
}

func ParseCssFromInlineStyle(cssText string) (m map[string]string) {
	m = make(map[string]string)
	declarations := strings.Split(cssText, ";")
	for _, declaration := range declarations {
		list := strings.Split(declaration, ":")
		m[list[0]] = list[1]
	}
	return
}
