package tags

type HtmlTagImg struct {
	isMap          bool
	alt            string
	sizes          string
	src            string
	srcSet         string
	useMap         string
	longDesc       string
	height         int
	width          int
	crossOrigin    CrossOriginType
	loading        LoadingType
	referrerPolicy ReferrerPolicyType
}
