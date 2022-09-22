package tags

type HtmlTagScript struct {
	async          bool
	defer_         bool
	noModule       bool
	integrity      string
	src            string
	type_          string
	crossOrigin    CrossOriginType
	referrerPolicy ReferrerPolicyType
}
