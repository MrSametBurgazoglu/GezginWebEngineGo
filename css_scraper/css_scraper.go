package css_scraper

import (
	"gezgin_web_engine/css_scraper/structs"
	"gezgin_web_engine/css_scraper/tree"
	"gezgin_web_engine/utils"
	"strings"
)

func scrapeCssParameters(cssText string) {
	varName, varValue, found := strings.Cut(cssText, "=")
	if found {
		index := utils.IndexFounder(cssPropertiesNameList, varName, cssPropertyCount)
		println(varName, varValue, index)
	}

}

func ScrapeCssFromInlineStyle(properties structs.CssProperties, styleText string) {
	if styleText != "" {
		parameters := strings.Split(styleText, ";")
		println(parameters)
	}
}

func scrapeCssFromStyleTag(styleText string) {
	strings.Split(styleText, ";")
}

func CreateCssPropertiesFromStyleTags() {
	for _, widget := range tree.CssStyleTagList {
		scrapeCssFromStyleTag(widget.StandardHtmlVariables.Style)
	}
}

func ExecuteCssScraper() {
	CreateCssPropertiesFromStyleTags()
}
