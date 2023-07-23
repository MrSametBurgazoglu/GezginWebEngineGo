package Fonts

import (
	"gezgin_web_engine/StyleEngine/structs"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	_ "golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
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
	newFont.Spacing = 1.5
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

func DrawText(font *GezginFont, text []string, destination *image.RGBA, fontColor *structs.ColorRGBA) (int32, int32) {
	// Initialize the context.
	//var err error
	ruler := color.RGBA{R: 0xdd, G: 0xdd, B: 0xdd, A: 0xff}
	rgba := image.NewRGBA(image.Rect(0, 0, 1500, 800))
	alpha, red, green, blue := fontColor.GetColorByRGBA()
	fg := image.NewUniform(color.RGBA{R: red, G: green, B: blue, A: alpha})
	bg := image.NewUniform(color.Transparent)
	draw.Draw(rgba, rgba.Bounds(), fg, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(destination, rgba.Bounds(), bg, image.Point{X: 0, Y: 0}, draw.Src)
	var context = font.Context
	context.SetDst(destination)
	context.SetSrc(fg)
	context.SetClip(destination.Bounds())
	//context.SetDPI(20)

	// Draw the guidelines.
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}

	// Draw the text.
	pt := freetype.Pt(10, 10+int(context.PointToFixed(font.Size)>>6))
	var maxPt = fixed.Int26_6(0)
	//font.Font.HMetric()
	for _, s := range text {
		p, err := context.DrawString(s, pt)
		if p.X > maxPt {
			maxPt = p.X
		}
		if err != nil {
			log.Println(err)
			return 0, 0
		}
		pt.Y += context.PointToFixed(font.Size * font.Spacing)
	}
	pt.Y -= context.PointToFixed(font.Size)
	const shift, mask = 6, 1<<6 - 1
	return int32(pt.Y >> shift), int32(maxPt >> shift)
}
