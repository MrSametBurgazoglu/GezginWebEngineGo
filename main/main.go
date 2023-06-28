package main

import (
	"flag"
	"fmt"
	"gezgin_web_engine/eventSystem"
	"gezgin_web_engine/web_engine"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
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

	startTime := time.Now()
	newTab := web_engine.NewTab()
	newTab.OpenWebPageFromFile("exampleHtmlFiles/newExa.html")
	//web_engine.OpenWebEngine("exampleHtmlFiles/newExa.html")
	fmt.Println("Total time taken ", time.Since(startTime).Milliseconds())

	web_engine.InitDrawer(700, 1300)
	window, renderer, err := sdl.CreateWindowAndRenderer(1300, 700, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	running := true
	//newTab.RenderPage(renderer)
	//newTab.DrawPage(renderer)
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			default:
				eventSystem.TakeInputFromSdl(event)
			}
		}
		if newTab.IsRendered() == false {
			renderer.SetDrawColor(250, 250, 250, 0)
			renderer.Clear()
			newTab.RenderPage(renderer)
			newTab.DrawPage(renderer)
			renderer.Present()
			newTab.SetRendered(true)
		}
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
