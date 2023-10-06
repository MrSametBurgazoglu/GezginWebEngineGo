package widgets

import (
	"bytes"
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleProperty"
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/widget"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"strconv"
	"strings"
	"time"
)

type HtmlTagImg struct {
	widget.Widget
	isMap          bool
	alt            string
	sizes          string
	Src            string
	srcSet         string
	useMap         string
	longDesc       string
	height         uint
	width          uint
	crossOrigin    CrossOriginType
	loading        LoadingType
	referrerPolicy ReferrerPolicyType
}

func (receiver *HtmlTagImg) ContextReaderFunc(context string) {
	//receiver
	if context == "ismap" {
		receiver.isMap = true
	}
}

func (receiver *HtmlTagImg) VarReaderFunc(variableName string, variableValue string) {
	//receiver
	switch variableName {
	case "alt":
		receiver.alt = variableValue
	case "crossorgin":
		receiver.crossOrigin.Set(variableValue)
	case "height":
		value, _ := strconv.Atoi(variableValue)
		receiver.height = uint(value)
	case "width":
		value, _ := strconv.Atoi(variableValue)
		receiver.width = uint(value)
	case "loading":
		receiver.loading.Set(variableValue)
	case "longdesc":
		receiver.longDesc = variableValue
	case "referrerpolicy":
		receiver.referrerPolicy.Set(variableValue)
	case "sizes":
		receiver.sizes = variableValue
	case "src":
		receiver.Src = variableValue[1 : len(variableValue)-1]
	case "srcset":
		receiver.srcSet = variableValue
	case "usemap":
		receiver.useMap = variableValue
	}
}

func (receiver *HtmlTagImg) Draw(mainImage *image.RGBA) {
	//file, err := os.Open("exampleHtmlFiles/browser-diagram.png")
	//img, err2 := png.Decode(file)
	//if err == nil && err2 == nil {
	draw.Draw(mainImage,
		image.Rect(receiver.LayoutProperty.XPosition,
			receiver.LayoutProperty.YPosition,
			receiver.LayoutProperty.Width+receiver.LayoutProperty.XPosition,
			receiver.LayoutProperty.Height+receiver.LayoutProperty.YPosition),
		receiver.DrawProperties.Texture,
		image.Point{X: 0, Y: 0},
		draw.Over)
	//}
}

func (receiver *HtmlTagImg) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {
	for !resourceManager.CheckResource(receiver.Src) {
		time.Sleep(time.Millisecond)
	}
	resource, _ := resourceManager.GetResource(receiver.Src)
	img, _, err2 := image.Decode(bytes.NewReader(resource.GetData())) //TODO PERFORMANCE UPDATE
	var imgTemp *image.RGBA
	if err2 != nil {
		imgTemp = image.NewRGBA(image.Rect(0, 0, receiver.LayoutProperty.ContentWidth, receiver.LayoutProperty.ContentHeight))
		if strings.HasSuffix(receiver.Src, ".svg") {
			drawerBackend.DrawSvg(imgTemp, bytes.NewReader(resource.GetData()), receiver.LayoutProperty)
		}
	} else {
		imgTemp = image.NewRGBA(img.Bounds())
		if receiver.LayoutProperty.ContentWidth != 0 && receiver.LayoutProperty.Height != 0 {
			img = resize.Resize(uint(receiver.LayoutProperty.ContentWidth), uint(receiver.LayoutProperty.ContentHeight), img, resize.Lanczos2)
		}
		draw.Draw(imgTemp, img.Bounds(), img, image.Point{X: 0, Y: 0}, draw.Src)
	}

	if imgTemp.Bounds().Size() != receiver.DrawProperties.Texture.Bounds().Size() {
		receiver.DrawProperties.Texture = image.NewRGBA(imgTemp.Bounds())
	}
	var imgText image.Image = imgTemp
	drawerBackend.GetImageTexture(
		&imgText,
		receiver.DrawProperties.Texture,
		receiver.LayoutProperty,
	)

}

func SetWidgetPropertiesForImgTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) widget.WidgetInterface {
	widget := new(HtmlTagImg)
	widget.HtmlElement = element
	widget.Initialize()
	widget.Src = widget.HtmlElement.Attributes["src"]
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
	widget.LayoutProperty.ContentWidth = widget.LayoutProperty.GetWidthFromStyleProperty()
	widget.LayoutProperty.ContentHeight = widget.LayoutProperty.GetPresetHeight()
	return widget
}
