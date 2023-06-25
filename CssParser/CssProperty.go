package CssParser

type CssProperty struct {
	property string
}

func (receiver *CssProperty) Set(text string) {
	receiver.property = text
}
