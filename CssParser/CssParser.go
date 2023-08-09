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

/*SUPPORT MEDIA QUERIES*/ //we should rewrite css parser like a programming language
func (receiver *CssParser) ParseCssFromStyleTag(styleElement StyleElement, styleText string) (result *Result) {
	result = new(Result)
	newCssStyleSheet := new(CssStyleSheet)
	styleText = utils.RemoveCharsFromString(styleText)
	seek := 0
	index := 0
	commentStart := strings.Index(styleText[seek:], "/*")
	index = strings.Index(styleText[seek:], "{")
	if commentStart < index {
		commentEnd := strings.LastIndex(styleText[seek:index], "*/")
		seek += commentEnd + 2
		index -= commentEnd + 2
	}
	frontRule := false
	for index != -1 { //maybe go routine for every cssText
		if styleText[seek] == '@' {
			frontRule = true
			firstSeek := seek
			println("frontRule started", seek)
			index = strings.Index(styleText[seek:], "{")
			seek += index + 2
			index = strings.Index(styleText[seek:], "{")
			println("frontRule ended", seek+index)
			println("frontRule", styleText[firstSeek:seek+index])
			if strings.HasPrefix(styleText[firstSeek:], "@media") {
				if !IsMediaRuleCorrect(styleText[firstSeek : seek+index]) {
					endOfAllRule := strings.Index(styleText[seek+index:], "}}")
					println("end of rule", endOfAllRule, styleText[seek+index+endOfAllRule:seek+index+endOfAllRule+3])
					seek = seek + index + endOfAllRule + 2
					frontRule = false
					index = strings.Index(styleText[seek:], "{")
				}
				println("type media")
			}
			continue
		} else if strings.HasPrefix(styleText[seek:], "/*") {
			commentEnd := strings.LastIndex(styleText[seek:], "*/")
			seek += commentEnd + 2
		}
		newCssRule := new(CssRule)
		index2 := strings.Index(styleText[seek:], "}")
		if index2 == -1 {
			return
		}
		selectors := styleText[seek : seek+index]
		println(selectors)
		cssText := styleText[seek+index+1 : seek+index2]
		seek += index2 + 1
		newCssRule.cssDeclarationBlock = new(CssDeclarationBlock)
		newCssRule.SetStyleSheet(newCssStyleSheet)
		newCssRule.SetCssSelectors(selectors)
		newCssRule.SetCssDeclarationBlock(cssText)
		result.CssStyleSheetRules = append(result.CssStyleSheetRules, newCssRule)
		result.ruleCount += 1
		println(string(styleText[seek]))
		if frontRule && styleText[seek] == '}' {
			frontRule = false
			seek += 1
		}
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
