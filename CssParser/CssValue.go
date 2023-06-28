package CssParser

type CssValue struct {
	value string
}

func (receiver *CssValue) Set(text string) {
	receiver.value = text
}
