package tags

type HtmlTagAudio struct {
	autoplay bool
	controls bool
	loop     bool
	muted    bool
	preload  PreLoadOptionType
	src      string
}
