package tags

type HtmlTagButton struct {
	autoFocus      bool
	disabled       bool
	formNovalidate bool
	form           string
	formAction     string
	name           string
	value          string
	formEncType    FormEncType
	formMethod     FormMethodType
	formTarget     FormTargetType
	buttonType     ButtonType
}
