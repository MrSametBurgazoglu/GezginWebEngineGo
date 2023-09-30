package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/widget"
	"image"
	"image/draw"
	"strconv"
	"strings"
)

type HtmlTagSVG struct {
	widget.Widget
	Width  int
	Height int
	//Children []widget.WidgetInterface
}

func (s *HtmlTagSVG) InitRect(child widget.WidgetInterface) {

}

func (s *HtmlTagSVG) InitChildren() {
	for _, child := range s.Children {
		switch child.GetHtmlName() {
		case "rect":
			s.InitRect(child)
		}
	}
}

func (receiver *HtmlTagSVG) Draw(mainImage *image.RGBA) {
	svgImage := image.NewRGBA(image.Rect(0, 0, receiver.LayoutProperty.ContentWidth, receiver.LayoutProperty.ContentHeight))
	svgFileText := "<?xml version=\"1.0\" encoding=\"iso-8859-1\"?> <svg"
	for key, value := range receiver.GetAttributes() {
		if strings.HasSuffix(value, "%") {
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				valueInt = 100
			}
			if key == "width" {
				value = strconv.Itoa(receiver.LayoutProperty.ContentWidth * valueInt / 100)
			} else if key == "height" {
				value = strconv.Itoa(receiver.LayoutProperty.ContentHeight * valueInt / 100)
			}
		}
		svgFileText += " "
		if value == "" {
			svgFileText += key
		} else {
			svgFileText += key + "=\"" + value + "\""
		}
	}
	svgFileText += ">"
	svgFileText += receiver.Children[0].(*UntaggedText).Value
	svgFileText += "</svg>"
	if strings.Contains(svgFileText, "<text") {
		start := strings.Index(svgFileText, "<text")
		end := strings.Index(svgFileText, "</text")
		svgFileText = svgFileText[:start] + svgFileText[end+7:]
	}
	ioReader := strings.NewReader(svgFileText)
	drawerBackend.DrawSvg(svgImage, ioReader, receiver.LayoutProperty)
	draw.Draw(mainImage, image.Rect(receiver.LayoutProperty.ContentXPosition, receiver.LayoutProperty.ContentYPosition, receiver.LayoutProperty.ContentXPosition+receiver.LayoutProperty.ContentWidth, receiver.LayoutProperty.ContentYPosition+receiver.LayoutProperty.ContentHeight), svgImage, image.Point{X: 0, Y: 0}, draw.Over)
}

func (receiver *HtmlTagSVG) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForSVGTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) widget.WidgetInterface {
	widget := new(HtmlTagSVG)
	widget.HtmlElement = element
	widget.Initialize()
	widget.IsNotDrawChildren = true
	wAttr := element.Attributes["width"]
	hAttr := element.Attributes["height"]
	if wAttr != "" {
		if !strings.HasSuffix(wAttr, "%") {
			wAttr += "px"
		}
		StyleProperty.WidthPropertySetValue(widget.StyleProperty, wAttr)
	}
	if hAttr != "" {
		if !strings.HasSuffix(hAttr, "%") {
			hAttr += "px"
		}
		StyleProperty.HeightPropertySetValue(widget.StyleProperty, hAttr)
	}
	widget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_BLOCK
	return widget
}
