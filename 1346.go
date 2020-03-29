package leetcode

func checkIfExist(arr []int) bool {
	checkList := make(map[int]*struct{}, len(arr))
	for _, i := range arr {
		if checkList[2*i] != nil || (i&1 == 0 && checkList[i/2] != nil) {
			return true
		}
		checkList[i] = &struct{}{}
	}

	return false
}
