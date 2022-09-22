package tags

type HtmlTagForm struct {
	acceptCharset string
	action        string
	name          string
	autoComplete  bool
	noValidate    bool
	formEnc       FormEncType
	formMethod    FormMethodType
	formRelType   FormRelType
	formTarget    FormTargetType
}
