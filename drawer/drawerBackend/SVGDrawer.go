package drawerBackend

import (
	"gezgin_web_engine/LayoutProperty"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"image"
	"image/draw"
	"io"
)

func DrawSvg(dest image.Image, ioReader io.Reader, layoutProperty *LayoutProperty.LayoutProperty) {
	icon, err := oksvg.ReadReplacingCurrentColor(ioReader, "black", oksvg.IgnoreErrorMode)
	if err != nil {
		panic(err.Error())
	}
	icon.SetTarget(0, 0, float64(layoutProperty.ContentWidth), float64(layoutProperty.ContentHeight))
	rgba := image.NewRGBA(image.Rect(0, 0, layoutProperty.ContentWidth, layoutProperty.ContentHeight))
	draw.Draw(rgba, image.Rect(0, 0, layoutProperty.ContentWidth, layoutProperty.ContentHeight), image.White, image.Point{X: 0, Y: 0}, draw.Src)
	icon.Draw(rasterx.NewDasher(layoutProperty.ContentWidth, layoutProperty.ContentHeight, rasterx.NewScannerGV(layoutProperty.ContentWidth, layoutProperty.ContentHeight, rgba, rgba.Bounds())), 1)
	draw.Draw(dest.(draw.Image), image.Rect(0, 0, layoutProperty.ContentWidth, layoutProperty.ContentHeight), rgba, image.Point{X: 0, Y: 0}, draw.Over)
}
