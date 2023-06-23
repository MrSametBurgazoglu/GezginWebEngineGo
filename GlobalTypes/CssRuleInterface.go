package GlobalTypes

type CssRuleInterface interface {
	GetSelectors() []string
	GetDeclarations() []CssDeclaration
}
