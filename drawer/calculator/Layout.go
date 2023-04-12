package calculator

import (
	"gezgin_web_engine/css_scraper/enums"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"gezgin_web_engine/html_scraper/widget"
)

// Calculate position and size from top and bottom with go routines simultaneously
func CalculateLayoutWidthOfWidget(widget *widget.Widget) int {
	if widget.CssProperties != nil {
		if widget.CssProperties.Display == enums.CSS_DISPLAY_TYPE_BLOCK {
			marginLeft := CalculateMarginLeftValue(widget)
			paddingLeft := CalculatePaddingLeftValue(widget)
			marginRight := CalculateMarginRightValue(widget)
			paddingRight := CalculatePaddingRightValue(widget)
			return marginLeft + paddingLeft + int(widget.DrawProperties.ContentRect.W) + paddingRight + marginRight
		}
	} else if widget.HtmlTag == htmlVariables.HTML_UNTAGGED_TEXT {
		return int(widget.DrawProperties.ContentRect.W)
	}
	return ScreenProperties.WindowWidth
}

func CalculateLayoutHeightOfWidget(widget *widget.Widget) int {
	if widget.CssProperties != nil {
		if widget.CssProperties.Display == enums.CSS_DISPLAY_TYPE_BLOCK {
			marginTop := CalculateMarginTopValue(widget)
			paddingTop := CalculatePaddingTopValue(widget)
			paddingBottom := CalculatePaddingBottomValue(widget)
			marginBottom := CalculateMarginBottomValue(widget)
			return marginTop + paddingTop + int(widget.DrawProperties.ContentRect.H) + paddingBottom + marginBottom
		}
	} else if widget.HtmlTag == htmlVariables.HTML_UNTAGGED_TEXT {
		return int(widget.DrawProperties.ContentRect.H)
	}
	return ScreenProperties.WindowWidth
}

func CalculateXPositionOfWidget(widget *widget.Widget) int32 {
	if widget.CssProperties != nil {
		a := widget.CssProperties.Position
		println(a, enums.CSS_POSITION_TYPE_RELATIVE)
		switch widget.CssProperties.Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			return widget.Parent.DrawProperties.Rect.X
		case enums.CSS_POSITION_TYPE_EMPTY:
			return widget.Parent.DrawProperties.Rect.X
		case enums.CSS_POSITION_TYPE_STATIC:
			return widget.Parent.DrawProperties.Rect.X
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			break
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			if widget.CssProperties.Left != 0 {
				return widget.Parent.DrawProperties.Rect.X + int32(widget.CssProperties.Left)
			} else if widget.CssProperties.Right != 0 {
				return widget.Parent.DrawProperties.Rect.W - int32(widget.CssProperties.Right)
			} else {
				return widget.Parent.DrawProperties.Rect.X
			}
		}
	} else {
		return widget.Parent.DrawProperties.Rect.X
	}
	return 0
}

func CalculateYPositionOfWidget(widget *widget.Widget) int {
	if widget.CssProperties != nil {
		if widget.CssProperties.Display == enums.CSS_DISPLAY_TYPE_BLOCK {
			marginTop := CalculateMarginTopValue(widget)
			paddingTop := CalculatePaddingTopValue(widget)
			paddingBottom := CalculatePaddingBottomValue(widget)
			marginBottom := CalculateMarginBottomValue(widget)
			return marginTop + paddingTop + int(widget.DrawProperties.ContentRect.H) + paddingBottom + marginBottom
		}
	} else if widget.HtmlTag == htmlVariables.HTML_UNTAGGED_TEXT {
		return int(widget.DrawProperties.ContentRect.H)
	}
	return ScreenProperties.WindowWidth
}
