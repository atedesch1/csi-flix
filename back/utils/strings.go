package utils

import "strings"

func SplitAndTrim(str string, separator string) []string {
	var arr []string
	if str != "" {
		arr = strings.Split(str, separator)
		for i, item := range arr {
			arr[i] = strings.TrimSpace(item)
		}
	} else {
		arr = []string{}
	}

	return arr
}
