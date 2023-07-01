package Fonts

import (
	"github.com/golang/freetype"
	"github.com/veandco/go-sdl2/ttf"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
)

var FontMap = make(map[int]*ttf.Font)

func InitFont() {
	fontBytes, err := os.ReadFile("fonts/Sans.ttf")
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	// Initialize the context.
	fg, bg := image.Black, image.White
	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	rgba := image.NewRGBA(image.Rect(0, 0, 640, 480))
	draw.Draw(rgba, rgba.Bounds(), bg, image.Point{X: 0, Y: 0}, draw.Src)
	c := freetype.NewContext()
	size := 12.0
	spacing := 5.0
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)

	// Draw the guidelines.
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}

	// Draw the text.
	text := "hello world"
	pt := freetype.Pt(10, 10+int(c.PointToFixed(size)>>6))
	for _, s := range text {
		_, err = c.DrawString(string(s), pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(size * spacing)
	}
}

func GetFont(size int) *ttf.Font {
	if font, ok := FontMap[size]; ok {
		return font
	} else {
		if font, err := ttf.OpenFont("fonts/Sans.ttf", size); err != nil {
			panic(err)
		} else {
			FontMap[size] = font
			return font
		}
	}
	return nil
}
