package cssParser

import (
	"gezgin_web_engine/GlobalTypes"
	"strings"
)

type CssRule struct {
	cssStyleSheet       *CssStyleSheet
	cssSelectors        []*CssSelector
	cssDeclarationBlock *CssDeclarationBlock
}

func (receiver *CssRule) SetStyleSheet(styleSheet *CssStyleSheet) {
	receiver.cssStyleSheet = styleSheet
}

func (receiver *CssRule) SetCssSelectors(text string) {
	selectorList := strings.Split(text, ",")
	for _, s := range selectorList {
		newCssSelector := new(CssSelector)
		newCssSelector.SetSelector(s)
		receiver.cssSelectors = append(receiver.cssSelectors, newCssSelector)
	}
}

func (receiver *CssRule) GetSelectors() (list []string) {
	for _, selector := range receiver.cssSelectors {
		list = append(list, selector.selector)
	}
	return
}

func (receiver *CssRule) SetCssDeclarationBlock(text string) {
	cssDeclarations := strings.Split(text, ";")
	for _, declaration := range cssDeclarations {
		newCssDeclaration := new(CssDeclaration)
		newCssDeclaration.Set(declaration)
		receiver.cssDeclarationBlock.cssDeclarations = append(receiver.cssDeclarationBlock.cssDeclarations, newCssDeclaration)
	}
}

func (receiver *CssRule) GetDeclarations() (list []GlobalTypes.CssDeclaration) {
	for _, declaration := range receiver.cssDeclarationBlock.cssDeclarations {
		list = append(list, declaration)
	}
	return
}
