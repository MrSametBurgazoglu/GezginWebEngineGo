package widgets

import (
	"gezgin_web_engine/HtmlParser"
	"gezgin_web_engine/ResourceManager"
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/StyleProperty/structs"
	"gezgin_web_engine/drawer/drawerBackend"
	"gezgin_web_engine/widget"
	"image"
	"image/color"
	"image/draw"
	"strconv"
)

type HtmlTagSVG struct {
	widget.Widget
	Width  int
	Height int
	//Children []widget.WidgetInterface
}

func (receiver *HtmlTagSVG) Draw(mainImage *image.RGBA) {
	svgImage := image.NewRGBA(image.Rect(0, 0, receiver.Width, receiver.Height))
	draw.Draw(svgImage, image.Rect(0, 0, receiver.Width, receiver.Height), image.White, image.Point{X: 0, Y: 0}, draw.Src)

	for _, child := range receiver.Children {
		switch child.GetHtmlName() {
		case "circle":
			xAttr := child.GetAttributes()["cx"]
			yAttr := child.GetAttributes()["cy"]
			radiusAttr := child.GetAttributes()["r"]
			strokeAttr := child.GetAttributes()["stroke"]
			fillAttr := child.GetAttributes()["fill"]
			x, _ := strconv.Atoi(xAttr)
			y, _ := strconv.Atoi(yAttr)
			r, _ := strconv.Atoi(radiusAttr)
			stroke := structs.ColorRGBA{}
			stroke.SetColor(strokeAttr)
			alpha, red, green, blue := stroke.GetColorByRGBA()
			fill := structs.ColorRGBA{}
			fill.SetColor(fillAttr)
			falpha, fred, fgreen, fblue := fill.GetColorByRGBA()
			println(x, y, r, "heyyo")
			if fillAttr != "" {
				drawerBackend.DrawFilledCircle(svgImage, x, y, r, color.RGBA{R: fred, G: fgreen, B: fblue, A: falpha})
			}
			drawerBackend.DrawCircle(svgImage, x, y, r, color.RGBA{R: red, G: green, B: blue, A: alpha})
		}
		println("svg elements", child.GetHtmlName())
	}
	draw.Draw(mainImage, svgImage.Bounds(), svgImage, image.Point{X: receiver.LayoutProperty.ContentXPosition, Y: receiver.LayoutProperty.ContentYPosition}, draw.Src)
}

func (receiver *HtmlTagSVG) Render(mainImage *image.RGBA, resourceManager *ResourceManager.ResourceManager) {

}

func SetWidgetPropertiesForSVGTag(element *HtmlParser.HtmlElement, taskManager TaskManagerInterface) widget.WidgetInterface {
	widget := new(HtmlTagSVG)
	widget.HtmlElement = element
	widget.Initialize()
	wAttr := element.Attributes["width"]
	hAttr := element.Attributes["height"]
	w, _ := strconv.Atoi(wAttr)
	h, _ := strconv.Atoi(hAttr)
	widget.Width = w
	widget.Height = h
	widget.StyleProperty.Display = enums.CSS_DISPLAY_TYPE_INLINE
	return widget
}
