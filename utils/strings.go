package utils

import "strings"

func LeftPadWith(s string, c string, totalLength int) string {
	if totalLength <= len(s) {
		return s
	}

	return strings.Repeat(c, totalLength-len(s)) + s
}
