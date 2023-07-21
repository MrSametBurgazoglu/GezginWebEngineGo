package ScreenProperties

var WindowHeight int
var WindowWidth int

func SetWindowSize(width int, height int) {
	WindowHeight = height
	WindowWidth = width
}

func GetWindowSize() (int, int) {
	return WindowHeight, WindowWidth
}
