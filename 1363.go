package leetcode

import (
	"bytes"
	"fmt"
	"strings"
)

func largestMultipleOfThree(digits []int) string {
	countMap := make(map[int]int, 9)
	remainder := 0
	for _, d := range digits {
		countMap[d] += 1
		remainder = (remainder + d) % 3
	}
	list1 := []int{1, 4, 7} // %3 == 1
	list2 := []int{2, 5, 8} // %3 == 2
	if remainder == 0 {
		return buildRes(countMap)
	} else {
		if remainder == 2 {
			tmpList := list1
			list1 = list2
			list2 = tmpList
		}
		for _, v := range list1 {
			if countMap[v] > 0 {
				countMap[v] -= 1
				return buildRes(countMap)
			}
		}
		for _, v := range list2 {
			for _, vv := range list2 {
				if v == vv {
					if countMap[v] > 1 {
						countMap[v] -= 2
						return buildRes(countMap)
					}
				} else {
					if countMap[v] > 0 && countMap[vv] > 0 {
						countMap[v] -= 1
						countMap[vv] -= 1
						return buildRes(countMap)
					}
				}
			}

		}
	}
	panic("no possible")

}

func buildRes(countMap map[int]int) string {
	if len(countMap) == 0 {
		return ""
	}

	var res bytes.Buffer
	for i := 9; i >= 0; i -= 1 {
		res.WriteString(strings.Repeat(fmt.Sprint(i), countMap[i]))
	}
	resStr := res.String()
	if strings.HasPrefix(resStr, "0") {
		return "0"
	}
	return strings.TrimLeft(resStr, "0")
}
