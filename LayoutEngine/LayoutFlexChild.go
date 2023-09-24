package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
)

func JustifyContent(parentSize, childrenTotalSize, childrenCount, CType int) (int, int) {
	switch CType {
	case 0: //flex-start
		return 0, 0
	case 1: //center
		return (parentSize - childrenTotalSize) / 2, 0
	case 2: //flex-end
		return parentSize - childrenTotalSize, 0
	case 3: //space-between
		return 0, (parentSize - childrenTotalSize) / childrenCount
	case 4: //space-around
		space := (parentSize - childrenTotalSize) / (childrenCount * 2)
		return space, space * 2
	case 5:
		space := (parentSize - childrenTotalSize) / (childrenCount + 1)
		return space, space
	default:
		return 0, 0
	}
}

func AlignItems(parentSize, childrenSize, CType int) int {
	switch CType {
	case 0: //flex-start
		return 0
	case 1: //center
		return (parentSize - childrenSize) / 2
	case 2: //flex-end
		return parentSize - childrenSize
	default:
		return 0
	}
}

/*TODO Positionları set position'a al. width kısmında sadece elementlerin width'lerini set et*/
/*TODO height kısmında hem kendisini hem de alt elementlerin height'lerini set et*/
func SetFLexContainerWidthRowNoWrap(currentWidget widget.WidgetInterface) {
	parentWidth := currentWidget.GetLayout().ContentWidth
	totalChildrenWidth, childrenWidths := GetTotalWidthOfChildren(currentWidget)
	if totalChildrenWidth > parentWidth {
		childWidth := parentWidth / len(childrenWidths)
		currentPos := 0
		for _, widgetInterface := range currentWidget.GetLayout().Children {
			widgetInterface.Width = childWidth
			widgetInterface.ContentWidth = childWidth
			widgetInterface.XPosition = currentPos
			widgetInterface.ContentXPosition = currentPos
			currentPos += childWidth
		}
	} else {
		/*
			startPos, spaceBetweenItems := JustifyContent(parentWidth, totalChildrenWidth, len(childrenWidths), 1)
			currentPos := startPos
		*/
		for i, widgetInterface := range currentWidget.GetLayout().Children {
			widgetInterface.Width = childrenWidths[i]
			widgetInterface.ContentWidth = childrenWidths[i]
		}
	}
	/*
		for _, widgetInterface := range currentWidget.GetChildren() {
			widgetInterface.GetLayout().YPosition = 0
			widgetInterface.GetLayout().ContentYPosition = 0
		}

		preSetHeight := currentWidget.GetLayout().GetPresetHeight()
		if preSetHeight > 0 {
			totalChildrenHeight := LookForHeight(currentWidget.GetLayout())
			if preSetHeight > totalChildrenHeight {
				for _, widgetInterface := range currentWidget.GetChildren() {
					childHeight := widgetInterface.GetLayout().GetPresetHeight()
					if childHeight == 0 {
						childHeight = LookForHeight(widgetInterface.GetLayout())
					}
					yPosition := AlignItems(preSetHeight, childHeight, 0)
					widgetInterface.GetLayout().YPosition += yPosition
					widgetInterface.GetLayout().ContentYPosition += yPosition
				}
			}
		}

	*/
}

func SetFlexContainerRowChildrenPositionNoWrap(currentWidget widget.WidgetInterface) {
	parentWidth := currentWidget.GetLayout().ContentWidth
	totalChildrenWidth, childrenWidths := GetTotalWidthOfChildren(currentWidget)
	if totalChildrenWidth > parentWidth {
		childWidth := parentWidth / len(childrenWidths)
		currentPos := 0
		for _, widgetInterface := range currentWidget.GetLayout().Children {
			widgetInterface.XPosition = currentPos
			widgetInterface.ContentXPosition = currentPos
			currentPos += childWidth
		}
	} else {
		startPos, spaceBetweenItems := JustifyContent(parentWidth, totalChildrenWidth, len(childrenWidths), 1)
		currentPos := startPos
		for _, widgetInterface := range currentWidget.GetLayout().Children {
			widgetInterface.XPosition = currentPos
			widgetInterface.ContentXPosition = currentPos
			currentPos += spaceBetweenItems
		}
	}

	for _, widgetInterface := range currentWidget.GetChildren() {
		widgetInterface.GetLayout().YPosition = 0
		widgetInterface.GetLayout().ContentYPosition = 0
	}

	preSetHeight := currentWidget.GetLayout().GetPresetHeight()
	if preSetHeight > 0 {
		totalChildrenHeight := LookForHeight(currentWidget.GetLayout())
		if preSetHeight > totalChildrenHeight {
			for _, widgetInterface := range currentWidget.GetChildren() {
				childHeight := widgetInterface.GetLayout().GetPresetHeight()
				if childHeight == 0 {
					childHeight = LookForHeight(widgetInterface.GetLayout())
				}
				yPosition := AlignItems(preSetHeight, childHeight, 0)
				widgetInterface.GetLayout().YPosition += yPosition
				widgetInterface.GetLayout().ContentYPosition += yPosition
			}
		}
	}

}

func SetFlexContainerRowChildrenPositionWrap(currentWidget widget.WidgetInterface) {

}

func SetFlexContainerRowChildrenPosition(currentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty().FlexWrap == enums.CSS_FLEX_WRAP_NOWRAP {
		SetFlexContainerRowChildrenPositionNoWrap(currentWidget)
	} else {
		SetFlexContainerRowChildrenPositionWrap(currentWidget)
	}
}

func SetFlexContainerColumnChildrenPosition(currentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty().FlexWrap == enums.CSS_FLEX_WRAP_NOWRAP {
		SetFlexContainerRowChildrenPositionNoWrap(currentWidget)
	} else {
		SetFlexContainerRowChildrenPositionWrap(currentWidget)
	}
}

func SetFlexContainerChildrenPosition(currentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_EMPTY || currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_ROW {
		SetFlexContainerRowChildrenPosition(currentWidget)
	} else {
		SetFlexContainerColumnChildrenPosition(currentWidget)
	}
}

func SetFLexContainerWidthRowWrap(currentWidget widget.WidgetInterface) {
	for i, child := range currentWidget.GetChildren() {
		width := LookForWidth(child.GetLayout())
		if currentWidget.GetStyleProperty().Children != nil && currentWidget.GetStyleProperty().Children[i].Width != 0 {
			width = currentWidget.GetLayout().Children[i].GetWidthFromStyleProperty()
		}
		child.GetLayout().Width = width
		child.GetLayout().ContentWidth = width
	}
}

func SetFlexRowContainerChildrenSizeAndPosition(currentWidget widget.WidgetInterface) {
	switch currentWidget.GetStyleProperty().FlexWrap {
	case enums.CSS_FLEX_WRAP_NOWRAP:
		SetFLexContainerWidthRowNoWrap(currentWidget)
	case enums.CSS_FLEX_WRAP_WRAP:
		SetFLexContainerWidthRowWrap(currentWidget)
	case enums.CSS_FLEX_WRAP_EMPTY:
		SetFLexContainerWidthRowNoWrap(currentWidget)
	}
	/*
		for i, child := range currentWidget.GetChildren() {
			width := LookForWidth(child.GetLayout())
			if currentWidget.GetStyleProperty().Children != nil && currentWidget.GetStyleProperty().Children[i].Width != 0 {
				switch currentWidget.GetStyleProperty().Children[i].WidthValueType {
				case enums.CSS_PROPERTY_VALUE_TYPE_PIXEL:
					width = int(currentWidget.GetStyleProperty().Children[i].Width)
				case enums.CSS_PROPERTY_VALUE_TYPE_PERCENTAGE:
					width = int(float64(currentWidget.GetLayout().Width) * (float64(currentWidget.GetStyleProperty().Children[i].Width) / 100.0))
				}
				if currentWidget.GetStyleProperty().Children[i].MaxWidth != 0 && width > int(currentWidget.GetStyleProperty().Children[i].MaxWidth) {
					width = int(currentWidget.GetStyleProperty().Children[i].MaxWidth)
				}
				if width < int(currentWidget.GetStyleProperty().Children[i].MinWidth) {
					width = int(currentWidget.GetStyleProperty().Children[i].MinWidth)
				}
			}
			child.GetLayout().Width = width
			child.GetLayout().ContentWidth = width
			totalWidth += width
		}

	*/
}

/*TOOO IF WRAP THEN ITS MEANS WE DON'T FIT THEM INTO ONE LINE WE USE THEIR WIDTHS*/
func SetFlexColumnContainerChildrenSizeAndPosition(currentWidget widget.WidgetInterface) {
	for i, child := range currentWidget.GetChildren() {
		width := currentWidget.GetLayout().Width
		if currentWidget.GetStyleProperty().Children[i].Width != 0 {
			width = currentWidget.GetLayout().Children[i].GetWidthFromStyleProperty()
		}
		child.GetLayout().Width = width
		child.GetLayout().ContentWidth = width
	}
}
