package tags

type HtmlTagVideo struct {
	autoplay bool
	controls bool
	loop     bool
	muted    bool
	poster   string
	src      string
	height   int
	width    int
	preload  PreLoadOptionType
}
