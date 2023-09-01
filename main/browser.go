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

	//web_engine.InitDrawer(w, h)
	if currentTab.IsRendered() == false {
		startTime := time.Now()
		currentTab.RenderPage()
		currentTab.DrawPage()
		currentTab.SetRendered(true)
		fmt.Println("Drawing Total time taken ", time.Since(startTime).Milliseconds())
	}

	tab := currentTab
	//file, err := os.Open("exampleHtmlFiles/browser-diagram.png")
	//img, err2 := png.Decode(file)
	imageS := tab.GetWebView()
	//if err == nil && err2 == nil {
	//	draw.Draw(currentTab.GetWebView(), currentTab.GetWebView().Bounds(), img, image.Point{X: 0, Y: 0}, draw.Src)
	//}

	surface := cairo.CreateSurfaceFromImage(imageS)
	cr.SetSourceSurface(surface, 0, 0)
	cr.Paint()
}

func activate(app *gtk.Application) {
	var state State

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

	window := gtk.NewApplicationWindow(app)
	window.SetTitle("drawingarea - gotk4 Example")
	window.SetChild(drawArea)
	window.SetDefaultSize(1300, 700)

	startTime := time.Now()
	web_engine.InitDrawer(1300, 700)
	newTab := web_engine.NewTab()
	//newTab.OpenWebPageFromFile("exampleHtmlFiles/newExa.html")
	newTab.OpenWebPageFromWeb("https://www.youtube.com")
	currentTab = newTab
	fmt.Println("Total time taken ", time.Since(startTime).Milliseconds())

	window.Show()
}
