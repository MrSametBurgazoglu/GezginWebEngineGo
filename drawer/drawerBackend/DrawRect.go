package drawerBackend

import (
	"gezgin_web_engine/drawer/ScreenProperties"
	"github.com/go-gl/gl/v2.1/gl"
)

func DrawRect(rect Rect, red, green, blue float32) {
	gl.Color3f(red, green, blue)
	x, y := rect.GetPosition()
	width, height := rect.GetSize()
	gl.Rectf(
		getRealValue(ScreenProperties.WindowWidth, x),
		getRealValue(ScreenProperties.WindowHeight, y),
		getRealValue(ScreenProperties.WindowWidth, x+width),
		getRealValue(ScreenProperties.WindowHeight, y+height),
	)
}
