package ScreenProperties

var WindowHeight int
var WindowWidth int

func SetWindowSize(height int, width int) {
	WindowHeight = height
	WindowWidth = width
}

func GetWindowSize() (int, int) {
	return WindowHeight, WindowWidth
}
