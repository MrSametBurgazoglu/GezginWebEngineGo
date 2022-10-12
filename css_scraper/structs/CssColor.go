package structs

type CssColor struct {
	alpha uint8
	red   uint8
	green uint8
	blue  uint8
}

func (receiver *CssColor) setByName(colorName string) {
	if colorName == "red" {
		receiver.alpha = 255
		receiver.red = 255
		receiver.green = 0
		receiver.blue = 0
	}
}
