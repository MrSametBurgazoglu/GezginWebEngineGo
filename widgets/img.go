package widgets

import (
	"bytes"
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/drawer/drawerBackend"
	"image"
	"image/draw"
	"strconv"
	"time"
)

type HtmlTagImg struct {
	Widget
	isMap          bool
	alt            string
	sizes          string
	Src            string
	srcSet         string
	useMap         string
	longDesc       string
	height         int
	width          int
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
		receiver.height, _ = strconv.Atoi(variableValue)
	case "width":
		receiver.width, _ = strconv.Atoi(variableValue)
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
	draw.Draw(mainImage, *receiver.DrawProperties.Rect, receiver.DrawProperties.Texture, image.Point{X: 0, Y: 0}, draw.Src)
}

func (receiver *HtmlTagImg) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {
	for !resourceManager.CheckResource(receiver.Src) {
		time.Sleep(time.Millisecond)
	}
	resource, err := resourceManager.GetResource(receiver.Src)
	img, _, err := image.Decode(bytes.NewReader(resource.GetData())) //TODO PERFORMANCE UPDATE
	if err == nil {
		drawerBackend.GetImageTexture(
			&img,
			receiver.DrawProperties.Texture,
			receiver.DrawProperties.Rect,
		)
	}
}

func SetWidgetPropertiesForImgTag(element *HtmlParser.HtmlElement) WidgetInterface {
	widget := new(HtmlTagImg)
	widget.HtmlElement = element
	widget.Initialize()
	widget.Src = widget.HtmlElement.Attributes["src"]
	return widget
}
