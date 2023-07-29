package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/drawer/drawerBackend"
	"image"
)

type HtmlTagHtml struct {
	Widget
}

func (receiver *HtmlTagHtml) Draw(mainImage *image.RGBA) {
	if receiver.GetStyleProperty().Background != nil {
		alpha, red, green, blue := receiver.StyleProperty.Background.BackgroundColor.GetColorByRGBA()
		drawerBackend.DrawBackground(red, green, blue, alpha, receiver.DrawProperties.Texture, receiver.DrawProperties)
	}
}

func (receiver *HtmlTagHtml) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForHtmlTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagHtml)
	widget.Initialize()
	return widget
}
