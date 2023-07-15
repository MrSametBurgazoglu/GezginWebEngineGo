package Fonts

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	_ "golang.org/x/image/font"
	"image"
	"image/color"
	"log"
	"os"
)

type GezginFont struct {
	Font    *truetype.Font
	Size    float64
	Spacing float64
	Fg      image.Uniform
	Bg      image.Uniform
	Context *freetype.Context
}

var FontMap = make(map[int]*GezginFont)

func InitFont(size int) (*GezginFont, error) {

	// Read the font data.
	fontBytes, err := os.ReadFile("fonts/Sans.ttf")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	newFont := new(GezginFont)
	newFont.Size = float64(size)
	newFont.Font = f
	fg := image.Black
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(newFont.Font)
	c.SetFontSize(newFont.Size)
	c.SetSrc(fg)
	newFont.Context = c
	return newFont, nil
}

func GetFont(size int) *GezginFont {
	if font, ok := FontMap[size]; ok {
		return font
	} else {
		if font, err := InitFont(size); err != nil {
			panic(err)
		} else {
			FontMap[size] = font
			return font
		}
	}
	return nil
}

func DrawText(font *GezginFont, text []string, destination *image.RGBA) {
	// Initialize the context.
	var err error
	ruler := color.RGBA{R: 0xdd, G: 0xdd, B: 0xdd, A: 0xff}
	rgba := image.NewRGBA(image.Rect(0, 0, 640, 480))
	//draw.Draw(rgba, rgba.Bounds(), bg, image.Point{X: 0, Y: 0}, draw.Src)
	var context *freetype.Context = font.Context
	context.SetDst(destination)

	// Draw the guidelines.
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}

	// Draw the text.
	pt := freetype.Pt(10, 10+int(context.PointToFixed(font.Size)>>6))
	for _, s := range text {
		_, err = context.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += context.PointToFixed(font.Size * font.Spacing)
	}
}
