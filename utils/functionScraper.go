package utils

import "strings"

func ScrapeFunction(value string) {
	cleanString := strings.TrimSpace(value)
	index1 := strings.Index(cleanString, "(")
	index2 := strings.Index(cleanString, ")")
	if index1 != -1 && index2 != -1 && index2 > index1 {
		println("its okey")
	}
}
