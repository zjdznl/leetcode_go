package leetcode

import (
	"math/bits"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a, b := arr[i], arr[j]
		n, k := bits.OnesCount(uint(a)), bits.OnesCount(uint(b))
		if n != k {
			return n < k
		}
		return a < b
	})

	return arr
}

//
//func sortByBits(arr []int) []int {
//	sort.Slice(arr, func(i, j int) bool {
//		c1 := getOneBitCount(arr[i])
//		c2 := getOneBitCount(arr[j])
//		if c1 < c2 || (c1 == c2 && arr[i] < arr[j]) {
//			return true
//		}
//		return false
//	})
//	return arr
//}
//
//func getOneBitCount(i int) int {
//	c := 0
//	for ; i > 0; i = i & (i - 1) {
//		c++
//	}
//	return c
//}
