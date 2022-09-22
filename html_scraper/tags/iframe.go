package tags

type HtmlTagIframe struct {
	allow               string
	allowFullScreen     string
	allowPaymentRequest string
	name                string
	src                 string
	srcDoc              string
	width               int
	height              int
	loading             LoadingType
	referrerPolicy      ReferrerPolicyType
	sandbox             SandboxAllowType
}
