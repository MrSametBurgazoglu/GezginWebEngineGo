package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/StyleEngine/structs"
	"gezgin_web_engine/drawer/drawerBackend"
	"image"
)

type HtmlTagHeader struct {
	Widget
}

func (receiver *HtmlTagHeader) Draw(mainImage *image.RGBA) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, receiver.DrawProperties.Texture, receiver.DrawProperties)
	}
}

func (receiver *HtmlTagHeader) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForH1Tag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.HtmlElement = element
	widget.Initialize()
	widget.LayoutProperty.Display = "block"
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 24
	return widget
}

func SetWidgetPropertiesForH2Tag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.HtmlElement = element
	widget.Initialize()
	widget.LayoutProperty.Display = "block"
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 20

	return widget
}

func SetWidgetPropertiesForH3Tag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.HtmlElement = element
	widget.Initialize()
	widget.LayoutProperty.Display = "block"
	widget.StyleProperty = new(StyleEngine.StyleProperty)
	widget.StyleProperty.Initialize()
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 12
	return widget
}

func SetWidgetPropertiesForH4Tag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.HtmlElement = element
	widget.Initialize()
	widget.LayoutProperty.Display = "block"
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 10
	return widget
}

func SetWidgetPropertiesForH5Tag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.HtmlElement = element
	widget.Initialize()
	widget.LayoutProperty.Display = "block"
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 8
	return widget
}

func SetWidgetPropertiesForH6Tag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagHeader)
	widget.HtmlElement = element
	widget.Initialize()
	widget.LayoutProperty.Display = "block"
	widget.StyleProperty.Font = new(structs.Font)
	widget.StyleProperty.Font.FontSizeValue = 6
	return widget
}
