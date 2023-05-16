package tags

import (
	"gezgin_web_engine/cssParser/structs"
	"gezgin_web_engine/htmlParser/widget"
)

func SetWidgetPropertiesForPTag(widget *widget.Widget) {
	widget.CssProperties.Font = new(structs.Font)
	widget.CssProperties.Font.FontSizeValue = 14
	widget.CssProperties.Margin = new(structs.Margin)
	widget.CssProperties.Margin.MarginTop = 10
}
