package widgets

import (
	"gezgin_web_engine/StyleEngine/structs"
)

type HtmlTagHeader struct {
	*Widget
}

func (receiver *HtmlTagHeader) Draw() {

}

func (receiver *HtmlTagHeader) Render() {

}

func SetWidgetPropertiesForH1Tag() WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 24
	return widget
}

func SetWidgetPropertiesForH2Tag() WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 20
	return widget
}

func SetWidgetPropertiesForH3Tag() WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 12
	return widget
}

func SetWidgetPropertiesForH4Tag() WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 10
	return widget
}

func SetWidgetPropertiesForH5Tag() WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 8
	return widget
}

func SetWidgetPropertiesForH6Tag() WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 6
	return widget
}
