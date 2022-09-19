package tags

import "fmt"

type HtmlTagA struct {
	download string
	href     string
	hrefLang string
}

func (*HtmlTagA) choose_variable_for_a_tag(variableName string, variableValue string) {
	fmt.Println("heyyo")
	fmt.Println("wow")
}
