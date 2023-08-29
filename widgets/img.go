package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/drawer/drawerBackend"
	"image"
	"image/draw"
	"image/png"
	"os"
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
		println("waiting for resource")
	}
	//resource, err := resourceManager.GetResource(receiver.Src)
	//img, format, err2 := image.Decode(bytes.NewReader(resource.GetData())) //TODO PERFORMANCE UPDATE
	file, err := os.Open("exampleHtmlFiles/browser-diagram.png")
	img, err2 := png.Decode(file)
	//if err == nil && err2 == nil {
	//	draw.Draw(mainImage, mainImage.Bounds(), img, image.Point{X: 0, Y: 0}, draw.Src)
	//}
	if img.Bounds().Size() != receiver.DrawProperties.Texture.Bounds().Size() {
		receiver.DrawProperties.Texture = image.NewRGBA(image.Rect(0, 0, img.Bounds().Size().X, img.Bounds().Size().Y))
	}
	if err == nil && err2 == nil {
		drawerBackend.GetImageTexture(
			&img,
			receiver.DrawProperties.Texture,
			receiver.LayoutProperty,
		)
	}
}

func SetWidgetPropertiesForImgTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) WidgetInterface {
	widget := new(HtmlTagImg)
	widget.HtmlElement = element
	widget.Initialize()
	widget.Src = widget.HtmlElement.Attributes["src"]
	return widget
}
