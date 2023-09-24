package LayoutEngine

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/widget"
	"math"
)

func JustifyContent(parentSize, childrenTotalSize, childrenCount, CType int) (int, int) {
	if parentSize == childrenTotalSize {
		return 0, 0
	}
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
func SetFlexContainerRowChildrenWidthNoWrap(currentWidget widget.WidgetInterface) {
	parentWidth := currentWidget.GetLayout().ContentWidth
	totalChildrenWidth, childrenWidths := GetTotalWidthOfChildren(currentWidget)
	if totalChildrenWidth > parentWidth {
		childWidth := parentWidth / len(childrenWidths)
		for _, widgetInterface := range currentWidget.GetLayout().Children {
			widgetInterface.Width = childWidth
			widgetInterface.ContentWidth = childWidth
		}
	} else {
		for i, widgetInterface := range currentWidget.GetLayout().Children {
			widgetInterface.Width = childrenWidths[i]
			widgetInterface.ContentWidth = childrenWidths[i]
		}
	}
}

func SetFlexContainerRowChildrenPositionNoWrap(currentWidget widget.WidgetInterface) {
	parentWidth := currentWidget.GetLayout().ContentWidth
	totalChildrenWidth, childrenWidths := GetTotalWidthOfChildren(currentWidget)

	startPos, spaceBetweenItems := JustifyContent(parentWidth, totalChildrenWidth, len(childrenWidths), 1)
	currentPos := startPos
	for _, widgetInterface := range currentWidget.GetLayout().Children {
		widgetInterface.XPosition = currentPos
		widgetInterface.ContentXPosition = currentPos
		currentPos += spaceBetweenItems
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

func GetTotalWidthOfWidgets(widgets []widget.WidgetInterface) int {
	totalWidth := 0
	for _, widgetInterface := range widgets {
		totalWidth += widgetInterface.GetLayout().Width
	}
	return totalWidth
}

func SetFlexContainerRowChildrenPositionWrap(currentWidget widget.WidgetInterface) {
	parentWidth := currentWidget.GetLayout().ContentWidth
	currentWidth := 0
	var currentSubContainers [][]widget.WidgetInterface
	var currentChildren []widget.WidgetInterface

	for _, widgetInterface := range currentWidget.GetChildren() {
		currentChildren = append(currentChildren, widgetInterface)
		currentWidth += widgetInterface.GetLayout().Width
		if currentWidth > parentWidth {
			copyChildren := currentChildren
			currentSubContainers = append(currentSubContainers, copyChildren)
			currentChildren = nil
			currentWidth = 0
		}
	}
	parentHeight := currentWidget.GetLayout().ContentHeight
	for i, container := range currentSubContainers {
		totalWidthOfChildren := GetTotalWidthOfWidgets(container)
		startPos, spaceBetweenItems := JustifyContent(parentWidth, totalWidthOfChildren, len(container), 1)
		currentPos := startPos
		for _, widgetInterface := range container {
			widgetInterface.GetLayout().XPosition = currentPos
			widgetInterface.GetLayout().ContentXPosition = currentPos
			currentPos += spaceBetweenItems
		}

		containerHeight := parentHeight / len(currentSubContainers)
		currentYPosition := containerHeight * i
		for _, widgetInterface := range container {
			yPosition := AlignItems(containerHeight, widgetInterface.GetLayout().Height, 0)
			widgetInterface.GetLayout().YPosition = currentYPosition + yPosition
			widgetInterface.GetLayout().ContentYPosition = currentYPosition + yPosition
		}
	}
}

func SetFlexContainerRowChildrenPosition(currentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty().FlexWrap == enums.CSS_FLEX_WRAP_NOWRAP {
		SetFlexContainerRowChildrenPositionNoWrap(currentWidget)
	} else {
		SetFlexContainerRowChildrenPositionWrap(currentWidget)
	}
}

func SetFlexContainerColumnChildrenPositionNoWrap(currentWidget widget.WidgetInterface) {
	parentHeight := currentWidget.GetLayout().ContentHeight
	totalChildrenHeight := GetTotalChildrenHeight(currentWidget)

	startPos, spaceBetweenItems := JustifyContent(parentHeight, totalChildrenHeight, len(currentWidget.GetLayout().Children), 1)
	currentPos := startPos
	for _, widgetInterface := range currentWidget.GetLayout().Children {
		widgetInterface.YPosition = currentPos
		widgetInterface.ContentYPosition = currentPos
		currentPos += spaceBetweenItems
	}

	for _, widgetInterface := range currentWidget.GetChildren() {
		widgetInterface.GetLayout().XPosition = 0
		widgetInterface.GetLayout().ContentXPosition = 0
	}

	parentWidth := currentWidget.GetLayout().ContentWidth

	for _, widgetInterface := range currentWidget.GetChildren() {
		xPosition := AlignItems(parentWidth, widgetInterface.GetLayout().Width, 0)
		widgetInterface.GetLayout().XPosition += xPosition
		widgetInterface.GetLayout().ContentXPosition += xPosition
	}

}

func GetTotalHeightOfWidgets(widgets []widget.WidgetInterface) int {
	totalHeight := 0
	for _, widgetInterface := range widgets {
		totalHeight += widgetInterface.GetLayout().Height
	}
	return totalHeight
}

func SetFlexContainerColumnChildrenPositionWrap(currentWidget widget.WidgetInterface) {
	parentHeight := currentWidget.GetLayout().ContentHeight
	currentHeight := 0
	var currentSubContainers [][]widget.WidgetInterface
	var currentChildren []widget.WidgetInterface

	for _, widgetInterface := range currentWidget.GetChildren() {
		currentChildren = append(currentChildren, widgetInterface)
		currentHeight += widgetInterface.GetLayout().Height
		if currentHeight > parentHeight {
			copyChildren := currentChildren
			currentSubContainers = append(currentSubContainers, copyChildren)
			currentChildren = nil
			currentHeight = 0
		}
	}
	parentWidth := currentWidget.GetLayout().ContentWidth
	for i, container := range currentSubContainers {
		totalHeightOfChildren := GetTotalHeightOfWidgets(container)
		startPos, spaceBetweenItems := JustifyContent(parentHeight, totalHeightOfChildren, len(container), 1)
		currentPos := startPos
		for _, widgetInterface := range container {
			widgetInterface.GetLayout().YPosition = currentPos
			widgetInterface.GetLayout().ContentYPosition = currentPos
			currentPos += spaceBetweenItems
		}

		containerWidth := parentWidth / len(currentSubContainers)
		currentXPosition := containerWidth * i
		for _, widgetInterface := range container {
			xPosition := AlignItems(containerWidth, widgetInterface.GetLayout().Width, 0)
			widgetInterface.GetLayout().XPosition = currentXPosition + xPosition
			widgetInterface.GetLayout().ContentXPosition = currentXPosition + xPosition
		}
	}
}

func SetFlexContainerColumnChildrenPosition(currentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty().FlexWrap == enums.CSS_FLEX_WRAP_NOWRAP {
		SetFlexContainerColumnChildrenPositionNoWrap(currentWidget)
	} else {
		SetFlexContainerColumnChildrenPositionWrap(currentWidget)
	}
}

func SetFlexContainerChildrenPosition(currentWidget widget.WidgetInterface) {
	if currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_EMPTY || currentWidget.GetStyleProperty().FlexDirection == enums.CSS_FLEX_DIRECTION_ROW {
		SetFlexContainerRowChildrenPosition(currentWidget)
	} else {
		SetFlexContainerColumnChildrenPosition(currentWidget)
	}
}

func SetFlexContainerRowChildrenWidthWrap(currentWidget widget.WidgetInterface) {
	for i, child := range currentWidget.GetChildren() {
		width := LookForWidth(child.GetLayout())
		if currentWidget.GetStyleProperty().Children != nil && currentWidget.GetStyleProperty().Children[i].Width != 0 {
			width = currentWidget.GetLayout().Children[i].GetWidthFromStyleProperty()
		}
		child.GetLayout().Width = width
		child.GetLayout().ContentWidth = width
	}
}

func SetFlexRowContainerChildrenWidth(currentWidget widget.WidgetInterface) {
	switch currentWidget.GetStyleProperty().FlexWrap {
	case enums.CSS_FLEX_WRAP_NOWRAP:
		SetFlexContainerRowChildrenWidthNoWrap(currentWidget)
	case enums.CSS_FLEX_WRAP_WRAP:
		SetFlexContainerRowChildrenWidthWrap(currentWidget)
	case enums.CSS_FLEX_WRAP_EMPTY:
		SetFlexContainerRowChildrenWidthNoWrap(currentWidget)
	}
}

func SetFlexColumnContainerChildrenWidthNoWrap(currentWidget widget.WidgetInterface) {
	for i, child := range currentWidget.GetChildren() {
		width := currentWidget.GetLayout().Width
		if currentWidget.GetStyleProperty().Children[i].Width != 0 {
			width = currentWidget.GetLayout().Children[i].GetWidthFromStyleProperty()
		}
		child.GetLayout().Width = width
		child.GetLayout().ContentWidth = width
	}
}

func SetFlexColumnContainerChildrenWidthWrap(currentWidget widget.WidgetInterface) {
	parentHeight := currentWidget.GetLayout().ContentHeight
	parentWidth := currentWidget.GetLayout().ContentWidth
	totalChildrenHeight := GetTotalChildrenHeight(currentWidget)
	if totalChildrenHeight > parentHeight {
		childWidth := parentWidth / int(math.Ceil(float64(parentHeight)/float64(totalChildrenHeight)))
		for i, widgetInterface := range currentWidget.GetLayout().Children {
			width := childWidth
			if currentWidget.GetStyleProperty().Children[i].Width != 0 {
				width = currentWidget.GetLayout().Children[i].GetWidthFromStyleProperty()
			}
			widgetInterface.Width = width
			widgetInterface.ContentWidth = width
		}
	} else {
		for i, child := range currentWidget.GetChildren() {
			width := currentWidget.GetLayout().Width
			if currentWidget.GetStyleProperty().Children[i].Width != 0 {
				width = currentWidget.GetLayout().Children[i].GetWidthFromStyleProperty()
			}
			child.GetLayout().Width = width
			child.GetLayout().ContentWidth = width
		}
	}
}

/*TOOO IF WRAP THEN ITS MEANS WE DON'T FIT THEM INTO ONE LINE WE USE THEIR WIDTHS*/
func SetFlexColumnContainerChildrenWidth(currentWidget widget.WidgetInterface) {
	switch currentWidget.GetStyleProperty().FlexWrap {
	case enums.CSS_FLEX_WRAP_NOWRAP:
		SetFlexColumnContainerChildrenWidthNoWrap(currentWidget)
	case enums.CSS_FLEX_WRAP_WRAP:
		SetFlexColumnContainerChildrenWidthWrap(currentWidget)
	case enums.CSS_FLEX_WRAP_EMPTY:
		SetFlexColumnContainerChildrenWidthNoWrap(currentWidget)
	}
}
