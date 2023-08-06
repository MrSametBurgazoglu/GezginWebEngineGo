package utils

import "strings"

func MergeAttributes(values []string) (result []string) {
	lastStr := ""
	equalFound := false
	for i, value := range values[:len(values)-1] {
		if values[i+1][0] == '=' || strings.Contains(value, "=") {
			equalFound = true
		}
		if !equalFound {
			result = append(result, value)
		} else {
			lastStr += value
			if value[len(value)-1] == '"' && !strings.HasSuffix(value, "=\"") {
				result = append(result, lastStr)
				lastStr = ""
				equalFound = false
			}
		}
	}
	value := values[len(values)-1]
	if value[len(value)-1] == '"' {
		if !strings.Contains(value, "=") {
			lastStr += " " + value
		}
		result = append(result, lastStr)
	} else {
		result = append(result, value)
	}
	return
}
