package CssParser

import "strings"

func IsMediaRuleCorrect(rule string) bool {
	if strings.HasPrefix(rule, "prefers-reduced-motion") {
		return false
	}
	return false
}
