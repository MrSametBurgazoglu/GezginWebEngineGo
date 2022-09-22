package tags

type HtmlTagInput struct {
	autoComplete   bool
	autoFocus      bool
	checked        bool
	disabled       bool
	multiple       bool
	readonly       bool
	required       bool
	formNovalidate bool
	accept         string
	alt            string
	dirname        string
	form           string
	formAction     string
	list           string
	max            string
	min            string
	name           string
	pattern        string
	placeHolder    string
	src            string
	value          string
	width          int
	height         int
	maxLength      int
	minLength      int
	size           int
	step           int
	formEncType    FormEncType
	formMethod     FormMethodType
	formTarget     FormTargetType
	inputType      InputType
}
