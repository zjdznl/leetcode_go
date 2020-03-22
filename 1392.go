package leetcode

import "strings"

//236 ms 6MB
//TODO use rolling hash
func longestPrefix(s string) string {
	sLen := len(s)
	if sLen == 1 {
		return ""
	}

	maxIndex := 0
	for i := 1; i < sLen; i++ {
		if strings.HasPrefix(s, s[:i]) && strings.HasSuffix(s, s[:i]) {
			maxIndex = i
		}
	}

	return s[:maxIndex]
}
