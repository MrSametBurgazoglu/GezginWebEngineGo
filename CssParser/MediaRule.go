package CssParser

import (
	"strconv"
	"strings"
)

func IsMediaRuleCorrect(rule string) bool {
	if strings.HasPrefix(rule, "prefers-reduced-motion") {
		return false
	} else if strings.HasPrefix(rule, "min-width") {
		splittedValue := strings.Split(rule, ":")
		value := splittedValue[1]
		intValue, _ := strconv.Atoi(value)
		if intValue == 992 {
			return true
		}
	}
	return false
}
