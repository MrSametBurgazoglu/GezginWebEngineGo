package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"

	_ "embed"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
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

// State describes the cursor state.
type State struct {
	X float64
	Y float64
}

func drawingFunction(area *gtk.DrawingArea, cr *cairo.Context, w, h int) {
	println("drawing")
	imageSurface := image.NewRGBA(image.Rect(0, 0, w, h))
	mygreen := color.RGBA{G: 100, A: 255} //  R, G, B, Alpha

	// backfill entire background surface with color mygreen
	draw.Draw(imageSurface, imageSurface.Bounds(), &image.Uniform{C: mygreen}, image.Point{Y: 0, X: 0}, draw.Src)
	surface := cairo.CreateSurfaceFromImage(imageSurface)
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
	window.SetDefaultSize(640, 480)
	window.Show()
}

func loadPNG(data []byte) (*gdkpixbuf.Pixbuf, error) {
	l, err := gdkpixbuf.NewPixbufLoaderWithType("png")
	if err != nil {
		return nil, fmt.Errorf("NewLoaderWithType png: %w", err)
	}
	defer l.Close()

	if err := l.Write(data); err != nil {
		return nil, fmt.Errorf("PixbufLoader.Write: %w", err)
	}

	if err := l.Close(); err != nil {
		return nil, fmt.Errorf("PixbufLoader.Close: %w", err)
	}

	return l.Pixbuf(), nil
}
