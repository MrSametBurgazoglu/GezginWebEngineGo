package main

import (
	"gezgin_web_engine/web_engine"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	var err error
	var font *ttf.Font

	if err = ttf.Init(); err != nil {
		panic(err)
	}
	defer ttf.Quit()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	if font, err = ttf.OpenFont("fonts/Sans.ttf", 14); err != nil {
		panic(err)
	}
	defer font.Close()

	web_engine.OpenWebEngine("exampleHtmlFiles/example.html")
	web_engine.InitDrawer()

	window, renderer, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		renderer.SetDrawColor(250, 250, 250, 0)
		renderer.Clear()
		web_engine.RenderPage(renderer)
		web_engine.DrawPage(renderer)
		renderer.Present()
	}
}
