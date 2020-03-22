package leetcode

//
//func createTargetArrayWithCopy(nums []int, index []int) []int {
//	lenNums := len(nums)
//	indexOffset := make([]int, lenNums)
//
//	numCount := make(map[int]int)
//
//	for i := lenNums - 1; i >= 0; i-- {
//		indexI := index[i]
//		minRigth := 0
//		for j := 0; j<=indexI;	j++  {
//			minRigth += numCount[j]
//		}
//		indexOffset[i] = minRigth
//		numCount[indexI]++
//	}
//
//	res := make([]int, lenNums)
//	for i, index := range index {
//		res[index+indexOffset[i]] = nums[i]
//	}
//	return res
//}

func createTargetArrayWithCopy(nums []int, index []int) []int {
	res := make([]int, len(index))
	for i := range res {
		res[i] = -1
	}

	for i, indexI := range index {
		if res[indexI] != -1 {
			copy(res[indexI+1:], res[indexI:])
		}
		res[indexI] = nums[i]
	}
	return res
}

func createTargetArrayWithAssignment(nums []int, index []int) []int {
	l := len(nums)
	target := make([]int, l)
	for i := 0; i < l; i++ {
		target[i] = -1
	}
	for i := 0; i < l; i++ {
		if target[index[i]] == -1 {
			target[index[i]] = nums[i]
			continue
		}
		for j := l - 1; j >= index[i] && j >= 1; j-- {
			target[j] = target[j-1]
		}
		target[index[i]] = nums[i]
	}
	return target
}
