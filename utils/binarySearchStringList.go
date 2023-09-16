package utils

import "strings"

/*TODO BUGFIX LAST STRING CAN'T BE FOUND*/
func IndexFounder(stringList []string, string string, listCount int) int {
	var m, result int
	l := 0
	r := listCount - 1

	for l <= r {
		m = l + (r-l)/2
		result = strings.Compare(stringList[m], string)
		if result == 0 {
			return m
		} else if result < 0 {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}
