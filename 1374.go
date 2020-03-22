package leetcode

import (
	"fmt"
	"strings"
)

func generateTheString(n int) string {
	if n&1 == 1 {
		return strings.Repeat("a", n)
	}
	return fmt.Sprintf("%s%s", strings.Repeat("a", 1), strings.Repeat("b", n-1))
}

// n is a positive number
func getMaxDividedBy2Counter(n int) int {
	maxCount := 0
	for {
		if n&1 == 1 {
			break
		}
		maxCount += 1
		n = n >> 1
	}
	return maxCount
}
