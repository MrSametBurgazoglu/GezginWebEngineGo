package htmlParser

type HtmlElement struct {
	ChildrenCount int
	ChildrenIndex int
	HtmlTag       HtmlTags
	Attributes    map[string]string
	Children      []*HtmlElement
	Parent        *HtmlElement
	Text          string
}

func (receiver *HtmlElement) GetChildren(index int) *HtmlElement {
	return receiver.Children[index]
}

func (receiver *HtmlElement) GetText() string {
	return receiver.Text
}
