package calculator

import (
	"gezgin_web_engine/cssParser/enums"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/htmlParser"
	"gezgin_web_engine/htmlParser/widget"
)

func CalculateWidthOfWidget(widget *widget.Widget) int {
	if widget.HtmlTag == htmlParser.HTML_UNTAGGED_TEXT {
		return int(widget.DrawProperties.Rect.W)
	} else if widget.HtmlTag == htmlParser.HTML_IMG {
		return int(widget.DrawProperties.Rect.W)
	} else if widget.CssProperties != nil {
		if widget.CssProperties.Display == enums.CSS_DISPLAY_TYPE_BLOCK {
			if widget.CssProperties.Width != 0 {
				return int(widget.CssProperties.Width)
			} else if widget.Parent != nil {
				return int(widget.Parent.DrawProperties.Rect.W)
			}
		}
	}
	return ScreenProperties.WindowWidth
}

func CalculateHeightOfWidget(widget *widget.Widget) (totalHeight int) {
	if widget.HtmlTag == htmlParser.HTML_UNTAGGED_TEXT {
		return int(widget.DrawProperties.Rect.H)
	} else if widget.HtmlTag == htmlParser.HTML_IMG {
		return int(widget.DrawProperties.Rect.H)
	}
	for i := 0; i < widget.ChildrenCount; i++ {
		if widget.Children[i].Draw {
			totalHeight += int(widget.Children[i].DrawProperties.Rect.H)
		}
	}
	return totalHeight
}

func CalculateXPosOfWidget(currentWidget *widget.Widget) int32 {
	if currentWidget.CssProperties != nil {
		switch currentWidget.CssProperties.Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			return currentWidget.Parent.DrawProperties.Rect.X
		case enums.CSS_POSITION_TYPE_EMPTY:
			return currentWidget.Parent.DrawProperties.Rect.X
		case enums.CSS_POSITION_TYPE_STATIC:
			return currentWidget.Parent.DrawProperties.Rect.X
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			if currentWidget.CssProperties.Left != 0 {
				return currentWidget.Parent.DrawProperties.Rect.X + int32(currentWidget.CssProperties.Left)
			} else if currentWidget.CssProperties.Right != 0 {
				return currentWidget.Parent.DrawProperties.Rect.W - int32(currentWidget.CssProperties.Right)
			} else {
				return currentWidget.Parent.DrawProperties.Rect.X
			}
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			if currentWidget.CssProperties.Left != 0 {
				return currentWidget.Parent.DrawProperties.Rect.X + int32(currentWidget.CssProperties.Left)
			} else if currentWidget.CssProperties.Right != 0 {
				return currentWidget.Parent.DrawProperties.Rect.W - int32(currentWidget.CssProperties.Right)
			} else {
				return currentWidget.Parent.DrawProperties.Rect.X
			}
		}
	} else {
		return currentWidget.Parent.DrawProperties.Rect.X
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
			if currentWidget.ChildrenIndex > 0 && (currentWidget.Parent.Children[currentWidget.ChildrenIndex-1].Draw || currentWidget.Parent.Children[currentWidget.ChildrenIndex-1].HtmlTag == htmlParser.HTML_UNTAGGED_TEXT) {
				beforeCurrentWidget = currentWidget.Parent.Children[currentWidget.ChildrenIndex-1]
				marginTop := 0
				if currentWidget.CssProperties.Margin != nil {
					marginTop = currentWidget.CssProperties.Margin.MarginTop
				}
				return beforeCurrentWidget.DrawProperties.Rect.Y + beforeCurrentWidget.DrawProperties.Rect.H + int32(marginTop)
			} else {
				beforeCurrentWidget = currentWidget.Parent
				marginTop := 0
				if currentWidget.CssProperties.Margin != nil {
					marginTop = currentWidget.CssProperties.Margin.MarginTop
				}
				return beforeCurrentWidget.DrawProperties.Rect.Y + int32(marginTop)
			}
		case enums.CSS_POSITION_TYPE_STATIC:
			if currentWidget.ChildrenIndex > 0 {
				beforeCurrentWidget = currentWidget.Parent.Children[currentWidget.ChildrenIndex-1]
			} else {
				beforeCurrentWidget = currentWidget.Parent
			}
			marginTop := 0
			if currentWidget.CssProperties.Margin != nil {
				marginTop = currentWidget.CssProperties.Margin.MarginTop
			}
			return beforeCurrentWidget.DrawProperties.Rect.Y + beforeCurrentWidget.DrawProperties.Rect.H + int32(marginTop)
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			if currentWidget.CssProperties.Top != 0 {
				beforeCurrentWidget = currentWidget.Parent
				return beforeCurrentWidget.DrawProperties.Rect.Y + int32(currentWidget.CssProperties.Top)
			} else if currentWidget.CssProperties.Bottom != 0 {
				beforeCurrentWidget = currentWidget.Parent
				return beforeCurrentWidget.DrawProperties.Rect.Y + beforeCurrentWidget.DrawProperties.Rect.H - int32(currentWidget.CssProperties.Bottom)
			} else {
				beforeCurrentWidget = currentWidget.Parent.Children[currentWidget.ChildrenIndex-1]
				return beforeCurrentWidget.DrawProperties.Rect.Y + beforeCurrentWidget.DrawProperties.Rect.H
			}
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
