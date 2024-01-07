package main

import (
	"fmt"
	"gezgin_web_engine/web_engine"
	"os"
	"time"

	_ "embed"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

const appID = "com.github.diamondburned.gotk4-examples.gtk4.drawingarea"

func main() {
	app := gtk.NewApplication(appID, gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

var currentTab *web_engine.WebTab

// State describes the cursor state.
type State struct {
	X float64
	Y float64
}

func drawingFunction(area *gtk.DrawingArea, cr *cairo.Context, w, h int) {

	if currentTab.IsRendered() == false {
		startTime := time.Now()
		currentTab.RenderPage()
		currentTab.DrawPage()
		currentTab.SetRendered(true)
		fmt.Println("Drawing Total time taken ", time.Since(startTime).Milliseconds())
	}

	tab := currentTab
	imageS := tab.GetWebView()
	surface := cairo.CreateSurfaceFromImage(imageS)
	area.SetContentHeight(imageS.Rect.Max.Y)
	cr.SetSourceSurface(surface, 0, 0)
	cr.Paint()
}

func activate(app *gtk.Application) {
	var state State

	scrolledWindow := gtk.NewScrolledWindow()

	drawArea := gtk.NewDrawingArea()
	drawArea.SetVExpand(true)
	drawArea.SetDrawFunc(drawingFunction)

	motionCtrl := gtk.NewEventControllerMotion()
	motionCtrl.ConnectMotion(func(x, y float64) {
		state.X = x
		state.Y = y
		drawArea.QueueDraw()
	})
	drawArea.AddController(motionCtrl)
	/*
		gesture := gtk.NewGestureClick()
		gesture.Connect("pressed", func(count int, x, y float64, data ...any) {
			println("hey", x, y, count)
		})
		drawArea.AddController(gesture)

	*/

	scrolledWindow.SetChild(drawArea)

	window := gtk.NewApplicationWindow(app)
	window.SetTitle("drawingarea - gotk4 Example")
	window.SetChild(scrolledWindow)
	window.SetSizeRequest(1900, 1000)
	window.SetResizable(true)

	startTime := time.Now()
	web_engine.InitDrawer(1900, 1000)
	newTab := web_engine.NewTab()
	//newTab.OpenWebPageFromFile("exampleHtmlFiles/newExa.html")
	newTab.OpenWebPageFromWeb("https://getbootstrap.com/docs/5.0/examples/pricing/")
	currentTab = newTab
	fmt.Println("Total time taken ", time.Since(startTime).Milliseconds())

	window.Show()
}
