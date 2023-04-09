package calculator

import (
	"gezgin_web_engine/css_scraper/enums"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"gezgin_web_engine/html_scraper/widget"
)

func CalculateContentWidthOfWidget(widget *widget.Widget) int {
	if widget.CssProperties != nil {
		if widget.CssProperties.Display == enums.CSS_DISPLAY_TYPE_BLOCK {
			childMaxWidth := 0
			for _, child := range widget.Children {
				if int(child.DrawProperties.LayoutRect.W) > childMaxWidth {
					childMaxWidth = int(child.DrawProperties.LayoutRect.W)
				}
			}
			return childMaxWidth
		}
	} else if widget.HtmlTag == htmlVariables.HTML_UNTAGGED_TEXT {
		return int(widget.DrawProperties.Rect.W)
	}
	return ScreenProperties.WindowWidth
}

func CalculateContentHeightOfWidget(widget *widget.Widget) (totalHeight int) {
	if widget.HtmlTag == htmlVariables.HTML_UNTAGGED_TEXT {
		return int(widget.DrawProperties.Rect.H)
	}
	for i := 0; i < widget.ChildrenCount; i++ {
		if widget.Children[i].Draw {
			totalHeight += int(widget.Children[i].DrawProperties.Rect.H)
		}
	}
	return totalHeight
}
