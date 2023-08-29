package structs

import (
	"gezgin_web_engine/drawer/Fonts"
	"image"
)

type DrawProperties struct {
	Font              *Fonts.GezginFont
	Rect              *image.Rectangle
	Texture           *image.RGBA
	BackgroundTexture *image.RGBA
}

func (receiver *DrawProperties) Initialize() {
	receiver.Rect = new(image.Rectangle)
	receiver.Texture = new(image.RGBA)
	receiver.BackgroundTexture = new(image.RGBA)
}
