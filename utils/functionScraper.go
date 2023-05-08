package utils

import "strings"

func ParseFunction(value string) (functionName string, functionParameters []string, ok bool) {
	cleanString := strings.TrimSpace(value)
	index1 := strings.Index(cleanString, "(")
	if index1 != -1 {
		index2 := strings.Index(cleanString, ")")
		if index2 != -1 && index2 > index1 {
			functionName = cleanString[:index1]
			functionParameters = strings.Split(cleanString[index1:index2], ",")
			ok = true
			return
		}
	}
	return
}
