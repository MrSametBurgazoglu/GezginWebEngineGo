package cssParser

type CssSelector struct {
	selector string
}

func (receiver *CssSelector) SetSelector(text string) {
	receiver.selector = text
}

func (receiver *CssSelector) GetSelector() string {
	return receiver.selector
}
