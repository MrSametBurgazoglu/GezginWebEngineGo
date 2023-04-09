package calculator

import (
	"gezgin_web_engine/css_scraper/enums"
	"gezgin_web_engine/html_scraper/widget"
)

func CalculateMarginLeftValue(widget *widget.Widget) int {
	if widget.CssProperties.Margin == nil {
		return 0
	}
	switch widget.CssProperties.Margin.MarginLeftValueType {
	case enums.CSS_PROPERTY_VALUE_TYPE_LENGTH:
		return widget.CssProperties.Margin.MarginLeft
	case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
		return int(widget.Parent.DrawProperties.ContentRect.W) * widget.CssProperties.Margin.MarginLeft
	default:
		return 0
	}
}

func CalculateMarginRightValue(widget *widget.Widget) int {
	if widget.CssProperties.Margin == nil {
		return 0
	}
	switch widget.CssProperties.Margin.MarginRightValueType {
	case enums.CSS_PROPERTY_VALUE_TYPE_LENGTH:
		return widget.CssProperties.Margin.MarginRight
	case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
		return int(widget.Parent.DrawProperties.ContentRect.W) * widget.CssProperties.Margin.MarginRight
	default:
		return 0
	}
}

func CalculateMarginTopValue(widget *widget.Widget) int {
	if widget.CssProperties.Margin == nil {
		return 0
	}
	switch widget.CssProperties.Margin.MarginTopValueType {
	case enums.CSS_PROPERTY_VALUE_TYPE_LENGTH:
		return widget.CssProperties.Margin.MarginTop
	case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
		return int(widget.Parent.DrawProperties.ContentRect.W) * widget.CssProperties.Margin.MarginTop
	default:
		return 0
	}
}

func CalculateMarginBottomValue(widget *widget.Widget) int {
	if widget.CssProperties.Margin == nil {
		return 0
	}
	switch widget.CssProperties.Margin.MarginBottomValueType {
	case enums.CSS_PROPERTY_VALUE_TYPE_LENGTH:
		return widget.CssProperties.Margin.MarginBottom
	case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
		return int(widget.Parent.DrawProperties.ContentRect.W) * widget.CssProperties.Margin.MarginBottom
	default:
		return 0
	}
}
