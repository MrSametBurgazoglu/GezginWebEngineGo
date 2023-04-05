package utils

import "strings"

func RemoveCharsFromString(str string) string {
	charsToRemove := []string{" ", "\n", "\t"}
	for _, char := range charsToRemove {
		str = strings.ReplaceAll(str, char, "")
	}
	return str
}
