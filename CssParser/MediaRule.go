package CssParser

import (
	"fmt"
	"gezgin_web_engine/drawer/ScreenProperties"
	"strconv"
	"strings"
)

func IsMediaRuleCorrect(rule string) bool {
	currentWidth, currentHeight := ScreenProperties.GetWindowSize()
	if strings.HasPrefix(rule, "prefers-reduced-motion") {
		return false
	} else if strings.Contains(rule, "min-width") {
		parameterStart := strings.Index(rule, "(")
		parameterEnd := strings.Index(rule, ")")
		parameter := rule[parameterStart:parameterEnd]
		splittedValue := strings.Split(parameter, ":")
		value := splittedValue[1]
		var intValue int
		_, err := fmt.Sscanf(value, "%dpx", &intValue)
		if err != nil {
			return false
		}
		if intValue <= currentWidth {
			return true
		}
	} else if strings.Contains(rule, "min-height") {
		parameterStart := strings.Index(rule, "(")
		parameterEnd := strings.Index(rule, ")")
		parameter := rule[parameterStart:parameterEnd]
		splittedValue := strings.Split(parameter, ":")
		value := splittedValue[1]
		intValue, _ := strconv.Atoi(value)
		if intValue < currentHeight {
			return true
		}
	}
	return false
}
