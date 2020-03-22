package leetcode

func numTimesAllBlue(light []int) int {
	res, maxRightIndex := 0, 0
	for i, v := range light {
		maxRightIndex = max(maxRightIndex, v-1)
		if i == maxRightIndex {
			res += 1
		}
	}
	return res
}
func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

//
//func numTimesAllBlue(light []int) int {
//	permutationSum := 0
//	satisfySum := 0
//	allBlueCount := 0
//	for i, v := range light {
//		permutationSum += i + 1
//		satisfySum += v
//		if permutationSum == satisfySum {
//			allBlueCount += 1
//		}
//	}
//	return allBlueCount
//}
//

//
//func numTimesAllBlue(light []int) int {
//	allBlue := 0
//	lightFlag := make([]int, len(light)) // 第几个灯泡之前是否全都变亮
//
//	curMaxIndex := 0
//
//	for i, v := range light {
//		curMaxIndex = max(curMaxIndex, i)
//		lightFlag[v-1] = 1
//		if sumIntSlice(lightFlag[:curMaxIndex+1]) == i+1 {
//			allBlue += 1
//		}
//	}
//
//	return allBlue
//}
//
//func max(i1, i2 int) int {
//	if i1 > i2 {
//		return i1
//	}
//	return i2
//}
//
//func sumIntSlice(s []int) int {
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//	return sum
//}
