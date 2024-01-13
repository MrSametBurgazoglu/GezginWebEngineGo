package tests

import (
	"gezgin_web_engine/web_engine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPricingExample(t *testing.T) {
	web_engine.InitDrawer(1468, 1000)
	newTab := web_engine.NewTab()
	newTab.OpenWebPageFromWeb("https://getbootstrap.com/docs/5.0/examples/pricing/")
	newTab.RenderPage()
	newTab.DrawPage()
	taskManager := newTab.GetTaskManager()
	documentLayout := taskManager.DocumentWidget.GetLayout()
	documentLayout.XPosition = 0
	documentLayout.YPosition = 0
	assert.Equal(t, 0, documentLayout.XPosition)
	assert.Equal(t, 0, documentLayout.YPosition)
	assert.Equal(t, 1468, documentLayout.Width)
	assert.Equal(t, 1195, documentLayout.Height)
	div1 := documentLayout.Children[1]
	assert.Equal(t, 266, div1.XPosition)
	assert.Equal(t, 0, div1.YPosition)
	assert.Equal(t, 936, div1.ContentWidth)
	assert.Equal(t, 1195, div1.ContentHeight)
	header := div1.Children[0]
	println(documentLayout.Children[1].PaddingTop)
	assert.Equal(t, 266, header.XPosition)
	assert.Equal(t, 16, header.YPosition)
	assert.Equal(t, 936, header.ContentWidth)
	assert.Equal(t, 198, header.ContentHeight)
}
