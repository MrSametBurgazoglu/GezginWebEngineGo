package calculator

import (
	"gezgin_web_engine/css_scraper/enums"
	"gezgin_web_engine/html_scraper/widget"
)

func CalculatePaddingLeftValue(widget *widget.Widget) int {
	if widget.CssProperties.Padding == nil {
		return 0
	}
	switch widget.CssProperties.Padding.PaddingLeftValueType {
	case enums.CSS_PROPERTY_VALUE_TYPE_LENGTH:
		return widget.CssProperties.Padding.PaddingLeft
	case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
		return int(widget.Parent.DrawProperties.ContentRect.W) * widget.CssProperties.Padding.PaddingLeft
	default:
		return 0
	}
}

func CalculatePaddingRightValue(widget *widget.Widget) int {
	if widget.CssProperties.Padding == nil {
		return 0
	}
	switch widget.CssProperties.Padding.PaddingRightValueType {
	case enums.CSS_PROPERTY_VALUE_TYPE_LENGTH:
		return widget.CssProperties.Padding.PaddingRight
	case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
		return int(widget.Parent.DrawProperties.ContentRect.W) * widget.CssProperties.Padding.PaddingRight
	default:
		return 0
	}
}

func CalculatePaddingTopValue(widget *widget.Widget) int {
	if widget.CssProperties.Padding == nil {
		return 0
	}
	switch widget.CssProperties.Padding.PaddingTopValueType {
	case enums.CSS_PROPERTY_VALUE_TYPE_LENGTH:
		return widget.CssProperties.Padding.PaddingTop
	case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
		return int(widget.Parent.DrawProperties.ContentRect.W) * widget.CssProperties.Padding.PaddingTop
	default:
		return 0
	}
}

func CalculatePaddingBottomValue(widget *widget.Widget) int {
	if widget.CssProperties.Padding == nil {
		return 0
	}
	switch widget.CssProperties.Padding.PaddingBottomValueType {
	case enums.CSS_PROPERTY_VALUE_TYPE_LENGTH:
		return widget.CssProperties.Padding.PaddingBottom
	case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
		return int(widget.Parent.DrawProperties.ContentRect.W) * widget.CssProperties.Padding.PaddingBottom
	default:
		return 0
	}
}
