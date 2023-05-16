package tags

import (
	"gezgin_web_engine/cssParser/structs"
	"gezgin_web_engine/htmlParser/widget"
)

func SetWidgetPropertiesForH1Tag(widget *widget.Widget) {
	widget.CssProperties.Font = new(structs.Font)
	widget.CssProperties.Font.FontSizeValue = 24
}

func SetWidgetPropertiesForH2Tag(widget *widget.Widget) {
	widget.CssProperties.Font = new(structs.Font)
	widget.CssProperties.Font.FontSizeValue = 20
}

func SetWidgetPropertiesForH3Tag(widget *widget.Widget) {
	widget.CssProperties.Font = new(structs.Font)
	widget.CssProperties.Font.FontSizeValue = 12
}

func SetWidgetPropertiesForH4Tag(widget *widget.Widget) {
	widget.CssProperties.Font = new(structs.Font)
	widget.CssProperties.Font.FontSizeValue = 10
}

func SetWidgetPropertiesForH5Tag(widget *widget.Widget) {
	widget.CssProperties.Font = new(structs.Font)
	widget.CssProperties.Font.FontSizeValue = 8
}

func SetWidgetPropertiesForH6Tag(widget *widget.Widget) {
	widget.CssProperties.Font = new(structs.Font)
	widget.CssProperties.Font.FontSizeValue = 6
}
