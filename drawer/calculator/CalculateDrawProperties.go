package calculator

import (
	"gezgin_web_engine/css_scraper/enums"
	"gezgin_web_engine/drawer/ScreenProperties"
	"gezgin_web_engine/drawer/structs"
	"gezgin_web_engine/html_scraper/htmlVariables"
)

func CalculateWidthOfWidget(widget any) int {
	switch w := widget.(type) {
	case structs.UntaggedTextDrawableWidget:
		return int(w.GetRect().W)
	case structs.DrawableWidget:
		if w.GetCssProperties().Display == enums.CSS_DISPLAY_TYPE_BLOCK {
			return ScreenProperties.WindowWidth
		}
	}
	return ScreenProperties.WindowWidth
}

func CalculateHeightOfWidget(widget structs.DrawableWidget) (totalHeight int) {
	if widget.GetHtmlTag() == htmlVariables.HTML_UNTAGGED_TEXT {
		return int(widget.GetRect().H)
	}
	for _, widgetInterface := range widget.GetChildren() {
		drawableWidget, ok := widgetInterface.(structs.DrawableWidget)
		if ok {
			totalHeight += int(drawableWidget.GetRect().H)
		}
	}
	return totalHeight
}

func CalculateXPosOfWidget(widget any) int32 {
	var parent structs.DrawableWidget
	switch w := widget.(type) {
	case structs.UntaggedTextDrawableWidget:
		return w.GetRect().W
	case structs.DrawableWidget:
		parent = w.GetParent().(structs.DrawableWidget)
		switch w.GetCssProperties().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			return parent.GetRect().X
		case enums.CSS_POSITION_TYPE_EMPTY:
			return parent.GetRect().X
		case enums.CSS_POSITION_TYPE_STATIC:
			return parent.GetRect().X
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			break
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			if w.GetCssProperties().Left != 0 {
				return parent.GetRect().X + int32(w.GetCssProperties().Left)
			} else if w.GetCssProperties().Right != 0 {
				return parent.GetRect().W - int32(w.GetCssProperties().Right)
			} else {
				return parent.GetRect().X
			}
		}
	}
	return 0
}

func CalculateYPosOfWidget(widget any) int32 {
	//TODO MAKE GET CHILD BY INDEX
	var beforeCurrentWidget structs.DrawableWidget
	var parent structs.DrawableWidget
	switch w := widget.(type) {
	case structs.UntaggedTextDrawableWidget:
		beforeCurrentWidget = w.GetParent().GetChild(w.GetChildrenIndex() - 1).(structs.DrawableWidget)
		if w.GetChildrenIndex() == 0 {
			return beforeCurrentWidget.GetRect().Y
		} else {
			return beforeCurrentWidget.GetRect().Y + beforeCurrentWidget.GetRect().H
		}
	case structs.DrawableWidget:
		parent = w.GetParent().(structs.DrawableWidget)
		beforeCurrentWidget = w.GetParent().GetChild(w.GetChildrenIndex() - 1).(structs.DrawableWidget)
		switch w.GetCssProperties().Position {
		case enums.CSS_POSITION_TYPE_STICKY:
			return beforeCurrentWidget.GetRect().X
		case enums.CSS_POSITION_TYPE_EMPTY:
			if w.GetChildrenIndex() > 0 && beforeCurrentWidget.GetHtmlTag() == htmlVariables.HTML_UNTAGGED_TEXT {
				return beforeCurrentWidget.GetRect().Y + beforeCurrentWidget.GetRect().H
			} else {
				return parent.GetRect().Y
			}
		case enums.CSS_POSITION_TYPE_STATIC:
			if w.GetChildrenIndex() > 0 {
				return beforeCurrentWidget.GetRect().Y + beforeCurrentWidget.GetRect().H
			} else {
				return parent.GetRect().Y + parent.GetRect().H
			}
		case enums.CSS_POSITION_TYPE_ABSOLUTE:
			break
		case enums.CSS_POSITION_TYPE_FIXED:
			break
		case enums.CSS_POSITION_TYPE_RELATIVE:
			break
		}
	}
	return 0
}
