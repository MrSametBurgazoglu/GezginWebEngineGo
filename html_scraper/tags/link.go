package tags

type HtmlTagLink struct {
	href           string
	hrefLang       string
	media          string
	sizes          string
	title          string
	mediaType      string
	crossOrigin    CrossOriginType
	referrerPolicy ReferrerPolicyType
	formRel        FormRelType
}
