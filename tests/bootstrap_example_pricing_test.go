package tests

import (
	"gezgin_web_engine/LayoutProperty"
	"gezgin_web_engine/web_engine"
	"github.com/stretchr/testify/assert"
	"testing"
)

var pricingExpectedBody = LayoutProperty.LayoutProperty{
	Position:  "BODY",
	XPosition: 0, YPosition: 0, Width: 1468, Height: 1195,
	ContentXPosition: 0, ContentYPosition: 0, ContentWidth: 1468, ContentHeight: 1195,
	MarginTop: 0, MarginBottom: 0, MarginLeft: 0, MarginRight: 0,
	PaddingTop: 0, PaddingBottom: 0, PaddingLeft: 0, PaddingRight: 0,
	Children: []*LayoutProperty.LayoutProperty{&pricingExpectedSvg, &pricingExpectedDiv1},
}

var pricingExpectedSvg = LayoutProperty.LayoutProperty{
	Position:  "SVG",
	XPosition: 0, YPosition: 0, Width: 1468, Height: 0, //todo actually it should be zero
	ContentXPosition: 0, ContentYPosition: 0, ContentWidth: 1468, ContentHeight: 0,
	MarginTop: 0, MarginBottom: 0, MarginLeft: 0, MarginRight: 0,
	PaddingTop: 0, PaddingBottom: 0, PaddingLeft: 0, PaddingRight: 0,
}

var pricingExpectedDiv1 = LayoutProperty.LayoutProperty{
	Position:  "DIV1",
	XPosition: 0, YPosition: 0, Width: 1468, Height: 1195,
	ContentXPosition: 266, ContentYPosition: 16, ContentWidth: 936, ContentHeight: 1600,
	MarginTop: 0, MarginBottom: 0, MarginLeft: 256, MarginRight: 256,
	PaddingTop: 16, PaddingBottom: 16, PaddingLeft: 12, PaddingRight: 12,
}

func testTree(t *testing.T, expected, actual *LayoutProperty.LayoutProperty) {
	testLayout(t, expected, actual)
	for i, child := range expected.Children {
		testTree(t, child, actual.Children[i])
	}

}

func testLayout(t *testing.T, expected, actual *LayoutProperty.LayoutProperty) {
	println(expected.Position)
	assert.Equal(t, expected.XPosition, actual.XPosition)
	assert.Equal(t, expected.YPosition, actual.YPosition)
	assert.Equal(t, expected.Width, actual.Width)
	assert.Equal(t, expected.Height, actual.Height)
	assert.Equal(t, expected.ContentXPosition, actual.ContentXPosition)
	assert.Equal(t, expected.ContentYPosition, actual.ContentYPosition)
	assert.Equal(t, expected.ContentWidth, actual.ContentWidth)
	assert.Equal(t, expected.ContentHeight, actual.ContentHeight)
	assert.Equal(t, expected.MarginTop, actual.MarginTop)
	assert.Equal(t, expected.MarginBottom, actual.MarginBottom)
	assert.Equal(t, expected.MarginLeft, actual.MarginLeft)
	assert.Equal(t, expected.PaddingTop, actual.PaddingTop)
	assert.Equal(t, expected.PaddingBottom, actual.PaddingBottom)
	assert.Equal(t, expected.PaddingLeft, actual.PaddingLeft)
	assert.Equal(t, expected.PaddingRight, actual.PaddingRight)
}

func TestPricingExample(t *testing.T) {
	web_engine.InitDrawer(1468, 1000)
	newTab := web_engine.NewTab()
	newTab.OpenWebPageFromWeb("https://getbootstrap.com/docs/5.0/examples/pricing/")
	newTab.RenderPage()
	newTab.DrawPage()
	taskManager := newTab.GetTaskManager()
	documentLayout := taskManager.DocumentWidget.GetLayout()
	testTree(t, &pricingExpectedBody, documentLayout)
}
