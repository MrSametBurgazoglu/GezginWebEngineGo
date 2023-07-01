package eventSystem

import (
	"github.com/veandco/go-sdl2/sdl"
)

var EventMap = make(map[string][]InputWidget)

type InputWidget interface {
	GetRect() *sdl.Rect
	GetOnclick() string
}

func SetInput(event string, currentWidget InputWidget) {
	EventMap[event] = append(EventMap[event], currentWidget)
}

func TakeInput(event sdl.Event) {
	switch currentEvent := (event).(type) {
	case *sdl.MouseButtonEvent:
		widgetList := EventMap["onclick"]
		for _, w := range widgetList {
			rect := *(w.GetRect())
			if rect.X < currentEvent.X &&
				currentEvent.X < rect.X+rect.W &&
				rect.Y < currentEvent.Y &&
				currentEvent.Y < rect.Y+rect.H {
				println(w.GetOnclick())
			}
		}
		println(currentEvent.X)
		println(currentEvent.Y)
		println("Mouse input")
		break
	}
}
