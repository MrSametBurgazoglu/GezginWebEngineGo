package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/StyleEngine/enums"
	"gezgin_web_engine/drawer/ScreenProperties"
)

func CalculateWidthOfWidget(widget WidgetInterface) int {
	if HtmlParser.HtmlTags(widget.GetHtmlTag()) == HtmlParser.HTML_UNTAGGED_TEXT {
		return int(widget.GetDrawProperties().W)
	} else if HtmlParser.HtmlTags(widget.GetHtmlTag()) == HtmlParser.HTML_IMG {
		return int(widget.GetDrawProperties().W)
	} else if widget.GetStyleProperty() != nil {
		if widget.GetStyleProperty().Display == enums.CSS_DISPLAY_TYPE_BLOCK {
			if widget.GetStyleProperty().Width != 0 {
				return int(widget.GetStyleProperty().Width)
			} else if widget.GetParent() != nil {
				return int(widget.GetParent().GetDrawProperties().W)
			}
		}
	}
	return ScreenProperties.WindowWidth
}

func CalculateHeightOfWidget(widget WidgetInterface) (totalHeight int) {
	if HtmlParser.HtmlTags(widget.GetHtmlTag()) == HtmlParser.HTML_UNTAGGED_TEXT {
		return int(widget.GetDrawProperties().H)
	} else if HtmlParser.HtmlTags(widget.GetHtmlTag()) == HtmlParser.HTML_IMG {
		return int(widget.GetDrawProperties().H)
	}
	for i := 0; i < widget.GetChildrenCount(); i++ {
		totalHeight += int(widget.GetChildrenByIndex(i).GetDrawProperties().H)
	}
	return totalHeight
}

func CalculateXPosOfWidget(currentWidget WidgetInterface) int32 {
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			return currentWidget.GetParent().GetDrawProperties().X
		case enums.CSS_POSITION_TYPE_EMPTY:
			return currentWidget.GetParent().GetDrawProperties().X
		case enums.CSS_POSITION_TYPE_STATIC:
			return currentWidget.GetParent().GetDrawProperties().X
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			if currentWidget.GetStyleProperty().Left != 0 {
				return currentWidget.GetParent().GetDrawProperties().X + int32(currentWidget.GetStyleProperty().Left)
			} else if currentWidget.GetStyleProperty().Right != 0 {
				return currentWidget.GetParent().GetDrawProperties().W - int32(currentWidget.GetStyleProperty().Right)
			} else {
				return currentWidget.GetParent().GetDrawProperties().X
			}
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			if currentWidget.GetStyleProperty().Left != 0 {
				return currentWidget.GetParent().GetDrawProperties().X + int32(currentWidget.GetStyleProperty().Left)
			} else if currentWidget.GetStyleProperty().Right != 0 {
				return currentWidget.GetParent().GetDrawProperties().W - int32(currentWidget.GetStyleProperty().Right)
			} else {
				return currentWidget.GetParent().GetDrawProperties().X
			}
		}
	} else {
		return currentWidget.GetParent().GetDrawProperties().X
	}
	return 0
}

func CalculateYPosOfWidget(currentWidget WidgetInterface) int32 {
	var beforeCurrentWidget WidgetInterface
	if currentWidget.GetStyleProperty() != nil {
		switch currentWidget.GetStyleProperty().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			return currentWidget.GetParent().GetDrawProperties().X
		case enums.CSS_POSITION_TYPE_EMPTY:
			if currentWidget.GetChildrenIndex() > 0 {
				beforeCurrentWidget = currentWidget.GetParent().GetChildrenByIndex(currentWidget.GetChildrenIndex() - 1)
				marginTop := 0
				if currentWidget.GetStyleProperty().Margin != nil {
					marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
				}
				return beforeCurrentWidget.GetDrawProperties().Y + beforeCurrentWidget.GetDrawProperties().H + int32(marginTop)
			} else {
				beforeCurrentWidget = currentWidget.GetParent()
				marginTop := 0
				if currentWidget.GetStyleProperty().Margin != nil {
					marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
				}
				return beforeCurrentWidget.GetDrawProperties().Y + int32(marginTop)
			}
		case enums.CSS_POSITION_TYPE_STATIC:
			if currentWidget.GetChildrenIndex() > 0 {
				beforeCurrentWidget = currentWidget.GetParent().GetChildrenByIndex(currentWidget.GetChildrenIndex() - 1)
			} else {
				beforeCurrentWidget = currentWidget.GetParent()
			}
			marginTop := 0
			if currentWidget.GetStyleProperty().Margin != nil {
				marginTop = currentWidget.GetStyleProperty().Margin.MarginTop
			}
			return beforeCurrentWidget.GetDrawProperties().Y + beforeCurrentWidget.GetDrawProperties().H + int32(marginTop)
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			if currentWidget.GetStyleProperty().Top != 0 {
				beforeCurrentWidget = currentWidget.GetParent()
				return beforeCurrentWidget.GetDrawProperties().Y + int32(currentWidget.GetStyleProperty().Top)
			} else if currentWidget.GetStyleProperty().Bottom != 0 {
				beforeCurrentWidget = currentWidget.GetParent()
				return beforeCurrentWidget.GetDrawProperties().Y + beforeCurrentWidget.GetDrawProperties().H - int32(currentWidget.GetStyleProperty().Bottom)
			} else {
				beforeCurrentWidget = currentWidget.GetParent().GetChildrenByIndex(currentWidget.GetChildrenIndex() - 1)
				return beforeCurrentWidget.GetDrawProperties().Y + beforeCurrentWidget.GetDrawProperties().H
			}
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			if currentWidget.GetChildrenIndex() > 0 {
				beforeCurrentWidget = currentWidget.GetParent().GetChildrenByIndex(currentWidget.GetChildrenIndex() - 1)
				return beforeCurrentWidget.GetDrawProperties().Y + beforeCurrentWidget.GetDrawProperties().H + int32(currentWidget.GetStyleProperty().Top)
			} else {
				beforeCurrentWidget = currentWidget.GetParent()
				return beforeCurrentWidget.GetDrawProperties().Y + int32(currentWidget.GetStyleProperty().Top)
			}
		}
	} else {
		beforeCurrentWidget = currentWidget.GetParent()
		if currentWidget.GetChildrenIndex() == 0 {
			return beforeCurrentWidget.GetDrawProperties().Y
		} else {
			return beforeCurrentWidget.GetChildrenByIndex(currentWidget.GetChildrenIndex()-1).GetDrawProperties().Y + beforeCurrentWidget.GetChildrenByIndex(currentWidget.GetChildrenIndex()-1).GetDrawProperties().H
		}
	}
	return 0
}
