package calculator

import (
	"gezgin_web_engine/css_scraper/enums"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/html_scraper/htmlVariables"
	"gezgin_web_engine/html_scraper/widget"
)

func CalculateWidthOfWidget(widget *widget.Widget) int {
	if widget.CssProperties != nil {
		if widget.CssProperties.Display == enums.CSS_DISPLAY_TYPE_BLOCK {
			return ScreenProperties.WindowWidth
		}
	} else if widget.HtmlTag == htmlVariables.HTML_UNTAGGED_TEXT {
		return int(widget.DrawProperties.Rect.W)
	}
	return ScreenProperties.WindowWidth
}

func CalculateHeightOfWidget(widget *widget.Widget) (totalHeight int) {
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

func CalculateXPosOfWidget(widget *widget.Widget) int32 {
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

func CalculateYPosOfWidget(currentWidget *widget.Widget) int32 {
	var beforeCurrentWidget *widget.Widget
	if currentWidget.CssProperties != nil {
		switch currentWidget.CssProperties.Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			return currentWidget.Parent.DrawProperties.Rect.X
		case enums.CSS_POSITION_TYPE_EMPTY:
			if currentWidget.ChildrenIndex > 0 && (currentWidget.Parent.Children[currentWidget.ChildrenIndex-1].Draw || currentWidget.Parent.Children[currentWidget.ChildrenIndex-1].HtmlTag == htmlVariables.HTML_UNTAGGED_TEXT) {
				beforeCurrentWidget = currentWidget.Parent.Children[currentWidget.ChildrenIndex-1]
				return beforeCurrentWidget.DrawProperties.Rect.Y + beforeCurrentWidget.DrawProperties.Rect.H
			} else {
				beforeCurrentWidget = currentWidget.Parent
				return beforeCurrentWidget.DrawProperties.Rect.Y
			}
		case enums.CSS_POSITION_TYPE_STATIC:
			if currentWidget.ChildrenIndex > 0 {
				beforeCurrentWidget = currentWidget.Parent.Children[currentWidget.ChildrenIndex-1]
			} else {
				beforeCurrentWidget = currentWidget.Parent
			}
			return beforeCurrentWidget.DrawProperties.Rect.Y + beforeCurrentWidget.DrawProperties.Rect.H
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			break
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			if currentWidget.ChildrenIndex > 0 {
				beforeCurrentWidget = currentWidget.Parent.Children[currentWidget.ChildrenIndex-1]
				return beforeCurrentWidget.DrawProperties.Rect.Y + beforeCurrentWidget.DrawProperties.Rect.H + int32(currentWidget.CssProperties.Top)
			} else {
				beforeCurrentWidget = currentWidget.Parent
				return beforeCurrentWidget.DrawProperties.Rect.Y + int32(currentWidget.CssProperties.Top)
			}
		}
	} else {
		beforeCurrentWidget = currentWidget.Parent
		if currentWidget.ChildrenIndex == 0 {
			return beforeCurrentWidget.DrawProperties.Rect.Y
		} else {
			return beforeCurrentWidget.Children[currentWidget.ChildrenIndex-1].DrawProperties.Rect.Y + beforeCurrentWidget.Children[currentWidget.ChildrenIndex-1].DrawProperties.Rect.H
		}
	}
	return 0
}
